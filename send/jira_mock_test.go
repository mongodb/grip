package send

import (
	"errors"

	jira "github.com/andygrunwald/go-jira"
)

type jiraClientMock struct {
	failCreate bool
	failAuth   bool
	failSend   bool
	numSent    int
}

func (j *jiraClientMock) CreateClient(_ string) error {
	if j.failCreate {
		return errors.New("mock failed to create client")
	}
	return nil
}

func (j *jiraClientMock) Authenticate(_ string, _ string) (bool, error) {
	if j.failAuth {
		return true, errors.New("mock failed authentication")
	}
	return false, nil
}

func (j *jiraClientMock) PostIssue(_ *jira.IssueFields) error {
	if j.failSend {
		return errors.New("mock failed sending")
	}

	j.numSent++

	return nil
}
