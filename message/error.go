package message

import "github.com/tychoish/grip/level"

type errorMessage struct {
	Err error          `json:"error" bson:"error" yaml:"error"`
	P   level.Priority `bson:"priority" json:"priority" yaml:"priority"`
}

// NewErrorMessage takes an error object and returns a Composer
// instance that only renders a loggable message when the error is
// non-nil.
func NewErrorMessage(p level.Priority, err error) Composer {
	return &errorMessage{
		Err: err,
		P:   p,
	}
}

func (e *errorMessage) Resolve() string {
	if e.Err == nil {
		return ""
	}
	return e.Err.Error()
}

func (e *errorMessage) Loggable() bool {
	return e.Err != nil
}

func (e *errorMessage) Raw() interface{} {
	return e
}

func (e *errorMessage) Priority() level.Priority {
	return e.P
}
