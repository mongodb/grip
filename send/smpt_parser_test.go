package send

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmailParser(t *testing.T) {
	assert := assert.New(t)

	assert.Implements((*emailAddressParser)(nil), &emailAddressParserImpl{})
}
