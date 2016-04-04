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

func (s *stringMessage) Raw() interface{} {
	return &struct {
		Message  string `json:"message" bson:"message" yaml:"message"`
		Loggable bool   `json:"loggable" bson:"loggable" yaml:"loggable"`
	}{
		Message:  s.Resolve(),
		Loggable: s.Loggable(),
	}
}
