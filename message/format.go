package message

import (
	"fmt"

	"github.com/tychoish/grip/level"
)

type formatMessenger struct {
	base string
	args []interface{}
	Base `bson:"metadata" json:"metadata" yaml:"metadata"`
}

// NewFormatedMessage takes arguments as fmt.Sprintf(), and returns
// an object that only runs the format operation as part of the
// Resolve() method.
func NewFormatedMessage(p level.Priority, base string, args ...interface{}) Composer {
	m := &formatMessenger{
		base: base,
		args: args,
	}
	m.SetPriority(p)

	return m
}

func NewFormated(base string, args ...interface{}) Composer {
	return &formatMessenger{
		base: base,
		args: args,
	}
}

func (f *formatMessenger) Resolve() string {
	return fmt.Sprintf(f.base, f.args...)
}

func (f *formatMessenger) Loggable() bool {
	return f.base != ""
}

func (f *formatMessenger) Raw() interface{} {
	p := f.Priority()

	return &struct {
		Message  string `json:"message" bson:"message" yaml:"message"`
		Loggable bool   `json:"loggable" bson:"loggable" yaml:"loggable"`
		Priority int    `bson:"priority" json:"priority" yaml:"priority"`
		Level    string `bson:"level" json:"level" yaml:"level"`
	}{
		Message:  f.Resolve(),
		Loggable: f.Loggable(),
		Priority: int(p),
		Level:    p.String(),
	}
}
