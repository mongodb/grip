package message

import (
	"testing"

	"github.com/google/go-github/github"
	"github.com/mongodb/grip/level"
	"github.com/stretchr/testify/assert"
)

func TestGithubStatus(t *testing.T) {
	assert := assert.New(t)

	c, err := NewGithubStatus(level.Info, "example", GithubStatePending, "https://example.com/hi", "description")
	assert.NoError(err)
	assert.NotNil(c)

	raw, ok := c.Raw().(*github.RepoStatus)
	assert.True(ok)

	assert.NotPanics(func() {
		assert.Equal("example", *raw.Context)
		assert.Equal(GithubStatePending, *raw.State)
		assert.Equal("https://example.com/hi", *raw.URL)
		assert.Equal("description", *raw.Description)
	})

	assert.Equal("example pending: description (https://example.com/hi)", c.String())
}
