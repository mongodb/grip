package message

import (
	"testing"

	"github.com/mongodb/grip/level"
	"github.com/stretchr/testify/assert"
)

func TestSlackMessage(t *testing.T) {
	assert := assert.New(t) //nolint

	m := NewSlackMessage(level.Notice, "abc", "i'm awesome", nil)
	assert.True(m.Loggable())
	assert.Equal("abc: i'm awesome", m.String())
	_, ok := m.Raw().(*Slack)
	assert.True(ok)

	m = NewSlackMessage(level.Notice, "", "i'm awesome", nil)
	assert.True(m.Loggable())

	m = NewSlackMessage(level.Notice, "abc", "", nil)
	assert.False(m.Loggable())
}
