package message

func NewErrorMessage(err error) *errorMessage {
	return &errorMessage{err}
}
func (e *errorMessage) Resolve() string {
	if e.err == nil {
		return ""
	}
	return e.err.Error()
}

func (e *errorMessage) Loggable() bool {
	if e.err == nil {
		return false
	}
	return true

}
