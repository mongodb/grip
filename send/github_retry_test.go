package send

import (
	"errors"
	"net/http"
	"testing"

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

func TestGitHubStatusSendWithErrorReturnsDeliveryDetails(t *testing.T) {
	sender := &githubStatusMessageLogger{
		Base: NewBase("github-status"),
		opts: &GithubOptions{RetryableHTTPStatusCodes: []int{http.StatusInternalServerError}},
		gh:   &githubClientMock{httpStatusCode: http.StatusInternalServerError},
	}
	require.NoError(t, sender.SetLevel(LevelInfo{Default: level.Info, Threshold: level.Trace}))

	err := sender.SendWithError(t.Context(), message.NewGithubStatusMessageWithRepo(level.Info, message.GithubStatus{
		Owner:   "evergreen-ci",
		Repo:    "evergreen",
		Ref:     "abc123",
		Context: "evergreen/code-health",
		State:   message.GithubStateSuccess,
	}))
	require.Error(t, err)

	var sendErr *GitHubSendError
	require.True(t, errors.As(err, &sendErr))
	assert.Equal(t, http.StatusInternalServerError, sendErr.StatusCode)
	assert.Equal(t, 1, sendErr.Attempts)
	assert.True(t, sendErr.Retryable)
}
