package send

import (
	"errors"

	"github.com/slack-go/slack"
)

// implements the slackClient interface for use in tests.
type slackClientMock struct {
	failAuthTest       bool
	failedGettingUser  bool
	failSendingMessage bool
	numSent            int
	lastTarget         string
	lastMsgOptions     *[]slack.MsgOption
}

func (c *slackClientMock) Create(_ string) {}
func (c *slackClientMock) AuthTest() (*slack.AuthTestResponse, error) {
	if c.failAuthTest {
		return nil, errors.New("mock failed auth test")
	}
	return nil, nil
}

func (c *slackClientMock) PostMessage(channelID string, options ...slack.MsgOption) (string, string, error) {
	if c.failSendingMessage {
		return "", "", errors.New("mock failed auth test")
	}

	c.numSent++
	c.lastTarget = channelID
	c.lastMsgOptions = &options
	return "", "", nil
}
func (c *slackClientMock) GetUserByEmail(email string) (*slack.User, error) {
	if c.failedGettingUser {
		return nil, errors.New("mock failed getting user")
	}
	return nil, nil
}
