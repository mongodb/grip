package send

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync/atomic"
	"testing"
	"time"

	"github.com/evergreen-ci/utility"
	"github.com/google/go-github/v79/github"
	"github.com/mongodb/grip/level"
	"github.com/mongodb/grip/message"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGithubOptionsPopulateDefaultsToBadGateway(t *testing.T) {
	opts := &GithubOptions{}
	require.NoError(t, opts.populate())

	assert.Equal(t, numGithubAttempts, opts.MaxAttempts)
	assert.Equal(t, []int{http.StatusBadGateway}, opts.RetryableHTTPStatusCodes)
}

func TestGithubOptionsRejectsNonErrorRetryStatuses(t *testing.T) {
	for _, test := range []struct {
		name       string
		statusCode int
	}{
		{name: "zero", statusCode: 0},
		{name: "informational", statusCode: http.StatusContinue},
		{name: "success", statusCode: http.StatusCreated},
		{name: "redirect", statusCode: http.StatusPermanentRedirect},
		{name: "above valid range", statusCode: 600},
	} {
		t.Run(test.name, func(t *testing.T) {
			opts := &GithubOptions{RetryableHTTPStatusCodes: []int{test.statusCode}}
			err := opts.populate()

			require.Error(t, err)
			assert.Contains(t, err.Error(), "must be between 400 and 599")
		})
	}
}

func TestGithubOptionsAllowsHTTPErrorRetryStatuses(t *testing.T) {
	opts := &GithubOptions{RetryableHTTPStatusCodes: []int{
		http.StatusBadRequest,
		http.StatusTooManyRequests,
		http.StatusInternalServerError,
		599,
	}}

	require.NoError(t, opts.populate())
}

func TestGitHubSenderConstructorsRejectNonErrorRetryStatuses(t *testing.T) {
	constructors := map[string]func(*GithubOptions) (Sender, error){
		"issues": func(opts *GithubOptions) (Sender, error) {
			return NewGithubIssuesLogger("issues", opts)
		},
		"comments": func(opts *GithubOptions) (Sender, error) {
			return NewGithubCommentLogger("comments", 1, opts)
		},
		"statuses": func(opts *GithubOptions) (Sender, error) {
			return NewGithubStatusLogger("statuses", opts, "ref")
		},
	}

	for name, constructor := range constructors {
		t.Run(name, func(t *testing.T) {
			sender, err := constructor(&GithubOptions{RetryableHTTPStatusCodes: []int{http.StatusCreated}})

			require.Error(t, err)
			assert.Nil(t, sender)
			assert.Contains(t, err.Error(), "invalid GitHub options")
		})
	}
}

func TestGitHubShouldRetryConfiguredStatusesWithinAttemptLimit(t *testing.T) {
	ctx, attempts := withGitHubAttemptTracker(t.Context())
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://api.github.com", nil)
	require.NoError(t, err)
	retry := githubShouldRetry(3, []int{
		http.StatusInternalServerError,
		http.StatusBadGateway,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout,
	})

	assert.True(t, retry(0, req, &http.Response{StatusCode: http.StatusInternalServerError}, nil))
	assert.Equal(t, 1, attempts.count)
	assert.True(t, retry(1, req, &http.Response{StatusCode: http.StatusServiceUnavailable}, nil))
	assert.Equal(t, 2, attempts.count)
	assert.False(t, retry(2, req, &http.Response{StatusCode: http.StatusGatewayTimeout}, nil))
	assert.Equal(t, 3, attempts.count)
	assert.False(t, retry(0, req, &http.Response{StatusCode: http.StatusNotImplemented}, nil))
}

func TestGitHubStatusRetriesConfiguredHTTPFailures(t *testing.T) {
	for _, test := range []struct {
		name           string
		statuses       []int
		expectedStatus int
		expectedError  bool
	}{
		{
			name:     "eventually succeeds",
			statuses: []int{http.StatusInternalServerError, http.StatusServiceUnavailable, http.StatusCreated},
		},
		{
			name:           "exhausts retries",
			statuses:       []int{http.StatusInternalServerError, http.StatusServiceUnavailable, http.StatusGatewayTimeout},
			expectedStatus: http.StatusGatewayTimeout,
			expectedError:  true,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			var requestCount atomic.Int32
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				body, err := io.ReadAll(r.Body)
				if err != nil {
					t.Errorf("reading request body: %s", err)
				} else if len(body) == 0 {
					t.Error("request body is empty")
				}

				attempt := int(requestCount.Add(1))
				if attempt > len(test.statuses) {
					t.Errorf("received unexpected request attempt %d", attempt)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(test.statuses[attempt-1])
				_, _ = w.Write([]byte(`{}`))
			}))
			defer server.Close()

			retryableStatuses := []int{
				http.StatusInternalServerError,
				http.StatusServiceUnavailable,
				http.StatusGatewayTimeout,
			}
			httpClient := utility.SetupOauth2CustomHTTPRetryableClient(
				"token",
				githubShouldRetry(len(test.statuses), retryableStatuses),
				func(int, *http.Request, *http.Response, error) time.Duration { return 0 },
				server.Client(),
			)
			githubClient := github.NewClient(httpClient)
			baseURL, err := url.Parse(server.URL + "/")
			require.NoError(t, err)
			githubClient.BaseURL = baseURL

			sender := &githubStatusMessageLogger{
				Base: NewBase("github-status"),
				opts: &GithubOptions{RetryableHTTPStatusCodes: retryableStatuses},
				gh: &githubClientImpl{
					IssuesService: githubClient.Issues,
					repos:         githubClient.Repositories,
				},
			}
			require.NoError(t, sender.SetLevel(LevelInfo{Default: level.Info, Threshold: level.Trace}))

			err = sender.SendWithError(t.Context(), message.NewGithubStatusMessageWithRepo(level.Info, message.GithubStatus{
				Owner:   "evergreen-ci",
				Repo:    "evergreen",
				Ref:     "abc123",
				Context: "evergreen/code-health",
				State:   message.GithubStateSuccess,
			}))
			assert.Equal(t, int32(len(test.statuses)), requestCount.Load())

			if !test.expectedError {
				require.NoError(t, err)
				return
			}
			require.Error(t, err)
			var sendErr *GitHubSendError
			require.True(t, errors.As(err, &sendErr))
			assert.Equal(t, test.expectedStatus, sendErr.StatusCode)
			assert.Equal(t, len(test.statuses), sendErr.Attempts)
			assert.True(t, sendErr.Retryable)
		})
	}
}
