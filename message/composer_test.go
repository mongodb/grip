package message

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tychoish/grip/level"
)

const testMsg = "hello"

func TestMessageComposerImplementations(t *testing.T) {
	assert := assert.New(t)
	// map objects to output
	cases := map[Composer]string{
		NewString(testMsg):                                        testMsg,
		NewDefaultMessage(level.Error, testMsg):                   testMsg,
		NewBytes([]byte(testMsg)):                                 testMsg,
		NewBytesMessage(level.Error, []byte(testMsg)):             testMsg,
		NewError(errors.New(testMsg)):                             testMsg,
		NewErrorMessage(level.Error, errors.New(testMsg)):         testMsg,
		NewErrorWrap(errors.New(testMsg), ""):                     testMsg,
		NewErrorWrapMessage(level.Error, errors.New(testMsg), ""): testMsg,
		MakeFieldsMessage(testMsg, Fields{}):                      fmt.Sprintf("[msg='%s']", testMsg),
	}

	for msg, output := range cases {
		assert.NotNil(msg)
		assert.NotEmpty(output)
		assert.Implements((*Composer)(nil), msg)
		assert.Equal(msg.String(), output)

		if msg.Priority() != level.Invalid {
			assert.Equal(msg.Priority(), level.Error)
		}
	}
}
