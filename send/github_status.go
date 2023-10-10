package send

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/go-github/v53/github"
	"github.com/mongodb/grip/message"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type githubStatusMessageLogger struct {
	opts *GithubOptions
	ref  string

	gh githubClient
	*Base
}

func (s *githubStatusMessageLogger) Send(m message.Composer) {
	if s.Level().ShouldLog(m) {
		var status *github.RepoStatus
		owner := ""
		repo := ""
		ref := ""

		switch v := m.Raw().(type) {
		case *message.GithubStatus:
			status = githubStatusMessagePayloadToRepoStatus(v)
			if v != nil {
				owner = v.Owner
				repo = v.Repo
				ref = v.Ref
			}
		case message.GithubStatus:
			status = githubStatusMessagePayloadToRepoStatus(&v)
			owner = v.Owner
			repo = v.Repo
			ref = v.Ref

		case *message.Fields:
			status = s.githubMessageFieldsToStatus(v)
			owner, repo, ref = githubMessageFieldsToRepo(v)
		case message.Fields:
			status = s.githubMessageFieldsToStatus(&v)
			owner, repo, ref = githubMessageFieldsToRepo(&v)
		}
		if len(owner) == 0 {
			owner = s.opts.Account
		}
		if len(repo) == 0 {
			owner = s.opts.Repo
		}
		if len(ref) == 0 {
			owner = s.ref
		}
		if status == nil {
			s.ErrorHandler()(errors.New("composer cannot be converted to GitHub status"), m)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		ctx, span := tracer.Start(ctx, "CreateStatus", trace.WithAttributes(
			attribute.String(githubEndpointAttribute, "CreateStatus"),
			attribute.String(githubOwnerAttribute, owner),
			attribute.String(githubRepoAttribute, repo),
			attribute.String(githubRefAttribute, ref),
		))
		defer span.End()

		if _, resp, err := s.gh.CreateStatus(ctx, owner, repo, ref, status); err != nil {
			s.ErrorHandler()(errors.Wrap(err, "sending GitHub create status request"), m)

			span.RecordError(err)
			span.SetStatus(codes.Error, "sending status")
		} else if err = handleHTTPResponseError(resp.Response); err != nil {
			s.ErrorHandler()(errors.Wrap(err, "creating GitHub status"), m)

			span.RecordError(err)
			span.SetStatus(codes.Error, "sending status")
		}
	}
}

func (s *githubStatusMessageLogger) Flush(_ context.Context) error { return nil }

// NewGithubStatusLogger returns a Sender to send payloads to the Github Status
// API. Statuses will be attached to the given ref.
func NewGithubStatusLogger(name string, opts *GithubOptions, ref string) (Sender, error) {
	opts.populate()
	s := &githubStatusMessageLogger{
		Base: NewBase(name),
		gh:   &githubClientImpl{},
		ref:  ref,
	}

	s.gh.Init(opts.Token, opts.MaxAttempts, opts.MinDelay)

	fallback := log.New(os.Stdout, "", log.LstdFlags)
	if err := s.SetErrorHandler(ErrorHandlerFromLogger(fallback)); err != nil {
		return nil, err
	}

	if err := s.SetFormatter(MakePlainFormatter()); err != nil {
		return nil, err
	}

	s.reset = func() {
		fallback.SetPrefix(fmt.Sprintf("[%s] [%s/%s] ", s.Name(), opts.Account, opts.Repo))
	}

	s.SetName(name)

	return s, nil
}

func (s *githubStatusMessageLogger) githubMessageFieldsToStatus(m *message.Fields) *github.RepoStatus {
	if m == nil {
		return nil
	}

	state, ok := getStringPtrFromField((*m)["state"])
	if !ok {
		return nil
	}
	context, ok := getStringPtrFromField((*m)["context"])
	if !ok {
		return nil
	}
	URL, ok := getStringPtrFromField((*m)["URL"])
	if !ok {
		return nil
	}
	var description *string
	if description != nil {
		description, ok = getStringPtrFromField((*m)["description"])
		if description != nil && len(*description) == 0 {
			description = nil
		}
		if !ok {
			return nil
		}
	}

	status := &github.RepoStatus{
		State:       state,
		Context:     context,
		TargetURL:   URL,
		Description: description,
	}

	return status
}

func getStringPtrFromField(i interface{}) (*string, bool) {
	if ret, ok := i.(string); ok {
		return &ret, true
	}
	if ret, ok := i.(*string); ok {
		return ret, ok
	}

	return nil, false
}
func githubStatusMessagePayloadToRepoStatus(c *message.GithubStatus) *github.RepoStatus {
	if c == nil {
		return nil
	}

	s := &github.RepoStatus{
		Context: github.String(c.Context),
		State:   github.String(string(c.State)),
	}
	if len(c.URL) > 0 {
		s.TargetURL = github.String(c.URL)
	}
	if len(c.Description) > 0 {
		s.Description = github.String(c.Description)
	}

	return s
}

func githubMessageFieldsToRepo(m *message.Fields) (string, string, string) {
	if m == nil {
		return "", "", ""
	}

	owner, ok := getStringPtrFromField((*m)["owner"])
	if !ok {
		owner = github.String("")
	}
	repo, ok := getStringPtrFromField((*m)["repo"])
	if !ok {
		repo = github.String("")
	}
	ref, ok := getStringPtrFromField((*m)["ref"])
	if !ok {
		ref = github.String("")
	}

	return *owner, *repo, *ref
}
