package grip

import "fmt"

type lineMessenger struct {
	lines []interface{}
}

func NewDefaultMessage(args ...interface{}) *lineMessenger {
	return &lineMessenger{args}
}

func (l *lineMessenger) Resolve() string {
	return fmt.Sprintln(l.lines)
}

type formatMessenger struct {
	base string
	args []interface{}
}

func NewFormatedMessage(base string, args ...interface{}) *formatMessenger {
	return &formatMessenger{base, args}
}

func (f *formatMessenger) Resolve() string {
	return fmt.Sprintf(f.base, f.args...)
}
