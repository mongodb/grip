package message

import "github.com/tychoish/grip/level"

type errorMessage struct {
	Err  error `json:"error" bson:"error" yaml:"error"`
	Base `bson:"metadata" json:"metadata" yaml:"metadata"`
}

// NewErrorMessage takes an error object and returns a Composer
// instance that only renders a loggable message when the error is
// non-nil.
func NewErrorMessage(p level.Priority, err error) Composer {
	m := &errorMessage{
		Err: err,
	}

	m.SetPriority(p)
	return m
}

func NewError(err error) Composer {
	return &errorMessage{Err: err}
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
