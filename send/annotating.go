package send

import "github.com/mongodb/grip/message"

type annotatingSender struct {
	Sender
	annotations map[string]interface{}
}

// NewAnnotatingSender adds the annotations defined in the annotations
// map to every argument.
//
// Calling code should assume that the sender owns the annotations map
// and it should not attempt to modify that data after calling the
// sender constructor.
//
// While you can wrap an existing sender with the annotator, changes
// to the annotating sender (e.g. level, formater, error handler) will
// propagate to the embedded sender.
func NewAnnotatingSender(s Sender, annotations map[string]interface{}) Sender {
	return &annotatingSender{
		Sender:      s,
		annotations: annotations,
	}
}

func (s *annotatingSender) Send(m message.Composer) {
	if !s.Sender.Level().ShouldLog(m) {
		return
	}

	for k, v := range s.annotations {
		m.Annotate(k, v)
	}
	s.Sender.Send(m)
}
