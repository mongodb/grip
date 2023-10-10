package send

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/google/go-github/v53/github"
)

type githubClientMock struct {
	failSend       bool
	httpStatusCode int
	numSent        int

	lastRepo string
}

func (g *githubClientMock) Init(_ string, _ int, _ time.Duration) {}

func (g *githubClientMock) Create(_ context.Context, _ string, _ string, _ *github.IssueRequest) (*github.Issue, *github.Response, error) {
	if g.failSend {
		return nil, nil, errors.New("failed to create issue")
	}

	g.numSent++
	return nil, g.createResponse(), nil
}
func (g *githubClientMock) CreateComment(_ context.Context, _ string, _ string, _ int, _ *github.IssueComment) (*github.IssueComment, *github.Response, error) {
	if g.failSend {
		return nil, nil, errors.New("failed to create comment")
	}

	g.numSent++
	return nil, g.createResponse(), nil
}

func (g *githubClientMock) CreateStatus(_ context.Context, repo, owner, ref string, _ *github.RepoStatus) (*github.RepoStatus, *github.Response, error) {
	if g.failSend {
		return nil, nil, errors.New("failed to create status")
	}

	g.numSent++
	g.lastRepo = fmt.Sprintf("%s/%s@%s", repo, owner, ref)
	return nil, g.createResponse(), nil
}

func (g *githubClientMock) createResponse() *github.Response {
	statusCode := g.httpStatusCode
	if statusCode <= 0 {
		statusCode = http.StatusOK
	}

	return &github.Response{
		Response: &http.Response{
			StatusCode: statusCode,
			Body:       io.NopCloser(strings.NewReader("body")),
		},
	}
}
