package send

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mongodb/grip/level"
	"github.com/mongodb/grip/message"
	"github.com/pkg/errors"
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
		personalAccessToken: opts.PersonalAccessTokenOpts.Token,
	}
	if err := j.opts.client.Authenticate(ctx, authOpts); err != nil {
		return nil, errors.Wrap(err, "authenticating")
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

		j.ErrorHandler()(j.opts.client.PostComment(issue, m.String()), m)
	}
}

func (j *jiraCommentJournal) Flush(_ context.Context) error { return nil }
