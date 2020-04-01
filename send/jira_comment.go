package send

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mongodb/grip/level"
	"github.com/mongodb/grip/message"
)

type jiraCommentJournal struct {
	issueID string
	opts    *JiraOptions
	*Base
}

// MakeJiraCommentLogger is the same as NewJiraCommentLogger but uses a warning
// level of Trace
func MakeJiraCommentLogger(ctx context.Context, id string, opts *JiraOptions) (Sender, error) {
	return NewJiraCommentLogger(ctx, id, opts, LevelInfo{level.Trace, level.Trace})
}

// NewJiraCommentLogger constructs a Sender that creates issues to jira, given
// options defined in a JiraOptions struct. id parameter is the ID of the issue.
// ctx is used as the request context in the OAuth HTTP client
func NewJiraCommentLogger(ctx context.Context, id string, opts *JiraOptions, l LevelInfo) (Sender, error) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	j := &jiraCommentJournal{
		opts:    opts,
		issueID: id,
		Base:    NewBase(id),
	}

	if err := j.opts.client.CreateClient(opts.HTTPClient, opts.BaseURL); err != nil {
		return nil, err
	}

	authOpts := jiraAuthOpts{
		username:           opts.BasicAuthOpts.Username,
		password:           opts.BasicAuthOpts.Password,
		addBasicAuthHeader: opts.BasicAuthOpts.UseBasicAuth,
		accessToken:        opts.Oauth1Opts.AccessToken,
		tokenSecret:        opts.Oauth1Opts.TokenSecret,
		privateKey:         opts.Oauth1Opts.PrivateKey,
		consumerKey:        opts.Oauth1Opts.ConsumerKey,
	}
	if err := j.opts.client.Authenticate(ctx, authOpts); err != nil {
		return nil, fmt.Errorf("jira authentication error: %v", err)
	}

	if err := j.SetLevel(l); err != nil {
		return nil, err
	}

	fallback := log.New(os.Stdout, "", log.LstdFlags)
	if err := j.SetErrorHandler(ErrorHandlerFromLogger(fallback)); err != nil {
		return nil, err
	}

	j.SetName(id)
	j.reset = func() {
		fallback.SetPrefix(fmt.Sprintf("[%s] ", j.Name()))
	}

	return j, nil
}

// Send post issues via jiraCommentJournal with information in the message.Composer
func (j *jiraCommentJournal) Send(m message.Composer) {
	if j.Level().ShouldLog(m) {
		issue := j.issueID
		if c, ok := m.Raw().(*message.JIRAComment); ok {
			issue = c.IssueID
		}
		if err := j.opts.client.PostComment(issue, m.String()); err != nil {
			j.ErrorHandler()(err, m)
		}
	}
}

func (j *jiraCommentJournal) Flush(_ context.Context) error { return nil }
