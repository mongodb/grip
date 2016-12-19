package message

import (
	"fmt"

	"github.com/tychoish/grip/level"
)

type formatMessenger struct {
	base     string
	args     []interface{}
	priority level.Priority
}

// NewFormatedMessage takes arguments as fmt.Sprintf(), and returns
// an object that only runs the format operation as part of the
// Resolve() method.
func NewFormatedMessage(p level.Priority, base string, args ...interface{}) Composer {
	return &formatMessenger{
		base:     base,
		args:     args,
		priority: p,
	}
}

func (f *formatMessenger) Resolve() string {
	return fmt.Sprintf(f.base, f.args...)
}

func (f *formatMessenger) Loggable() bool {
	return f.base != ""
}

func (f *formatMessenger) Raw() interface{} {
	return &struct {
		Message  string `json:"message" bson:"message" yaml:"message"`
		Loggable bool   `json:"loggable" bson:"loggable" yaml:"loggable"`
		Priority int    `bson:"priority" json:"priority" yaml:"priority"`
		Level    string `bson:"level" json:"level" yaml:"level"`
	}{
		Message:  f.Resolve(),
		Loggable: f.Loggable(),
		Priority: int(f.priority),
		Level:    f.priority.String(),
	}
}

func (f *formatMessenger) Priority() level.Priority {
	return f.priority
}
