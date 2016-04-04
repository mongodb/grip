package message

import "fmt"

type formatMessenger struct {
	base string
	args []interface{}
}

// NewFormatedMessage() takes arguments as fmt.Sprintf(), and returns
// an object that only runs the format operation as part of the
// Resolve() method.  Use in combination with Compose[*] logging
// methods.
func NewFormatedMessage(base string, args ...interface{}) *formatMessenger {
	return &formatMessenger{base, args}
}

func (f *formatMessenger) Resolve() string {
	return fmt.Sprintf(f.base, f.args...)
}

func (f *formatMessenger) Loggable() bool {
	if f.base == "" {
		return false
	}
	return true
}

type errorMessage struct {
	err error
}
