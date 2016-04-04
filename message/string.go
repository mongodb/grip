package message

type stringMessage struct {
	content string
}

func NewDefaultMessage(message string) *stringMessage {
	return &stringMessage{message}
}

func (s *stringMessage) Resolve() string {
	return s.content
}

func (s *stringMessage) Loggable() bool {
	if s.content != "" {
		return true
	}
	return false
}
