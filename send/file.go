package send

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tychoish/grip/level"
)

type fileLogger struct {
	name string

	defaultLevel   level.Priority
	thresholdLevel level.Priority

	options  map[string]string
	template string

	logger  *log.Logger
	fileObj *os.File
}

func NewFileLogger(name, filePath string, thresholdLevel, defaultLevel level.Priority) (Sender, error) {
	l := &fileLogger{
		name:     name,
		options:  make(map[string]string),
		template: "[p=%d]: %s\n",
	}

	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		l, _ := NewNativeLogger(name, thresholdLevel, defaultLevel)
		return l, fmt.Errorf("error opening logging file, %s, falling back to stdOut logging", err.Error())
	}
	l.fileObj = f
	l.createLogger()

	err = l.SetDefaultLevel(defaultLevel)
	if err != nil {
		return l, err
	}

	err = l.SetThresholdLevel(thresholdLevel)
	if err != nil {
		return l, err
	}

	return l, nil
}

func (f *fileLogger) createLogger() {
	f.logger = log.New(f.fileObj, strings.Join([]string{"[", f.name, "] "}, ""), log.LstdFlags)
}

func (f *fileLogger) Close() {
	f.fileObj.Close()
}

func (f *fileLogger) Send(p level.Priority, m string) {
	f.logger.Printf(f.template, int(p), m)
}

func (f *fileLogger) Name() string {
	return f.name
}

func (f *fileLogger) SetName(name string) {
	f.name = name
	f.createLogger()
}

func (f *fileLogger) AddOption(key, value string) {
	f.options[key] = value
}

func (f *fileLogger) GetDefaultLevel() level.Priority {
	return f.defaultLevel
}

func (f *fileLogger) GetThresholdLevel() level.Priority {
	return f.thresholdLevel
}

func (f *fileLogger) SetDefaultLevel(p level.Priority) error {
	if level.IsValidPriority(p) {
		f.defaultLevel = p
		return nil
	} else {
		return fmt.Errorf("%s (%d) is not a valid priority value (0-6)", p, int(p))
	}
}

func (f *fileLogger) SetThresholdLevel(p level.Priority) error {
	if level.IsValidPriority(p) {
		f.thresholdLevel = p
		return nil
	} else {
		return fmt.Errorf("%s (%d) is not a valid priority value (0-6)", p, int(p))
	}
}
