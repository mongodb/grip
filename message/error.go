package message

type errorMessage struct {
	Err error `json:"error" bson:"error" yaml:"error"`
}

func NewErrorMessage(err error) *errorMessage {
	return &errorMessage{err}
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
