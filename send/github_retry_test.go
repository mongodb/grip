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
	opts.populate()

	assert.Equal(t, numGithubAttempts, opts.MaxAttempts)
	assert.Equal(t, []int{http.StatusBadGateway}, opts.RetryableHTTPStatusCodes)
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
