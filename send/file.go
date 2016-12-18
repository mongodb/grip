package send

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/tychoish/grip/level"
	"github.com/tychoish/grip/message"
)

type fileLogger struct {
	name  string
	level LevelInfo

	template string

	logger  *log.Logger
	fileObj *os.File

	sync.RWMutex
}

// NewFileLogger creates a Sender implementation that writes log
// output to a file. Returns an error but falls back to a standard
// output logger if there's problems with the file. Internally using
// the go standard library logging system.
func NewFileLogger(name, filePath string, thresholdLevel, defaultLevel level.Priority) (Sender, error) {
	l := &fileLogger{
		name:     name,
		template: "[p=%d]: %s\n",
	}

	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		l, _ := NewNativeLogger(name, thresholdLevel, defaultLevel)
		return l, fmt.Errorf("error opening logging file, %s, falling back to stdOut logging", err.Error())
	}
	l.fileObj = f
	l.createLogger()

	level := LevelInfo{defaultLevel, thresholdLevel}
	if !level.Valid() {
		return nil, fmt.Errorf("level configuration is invalid: %+v", level)
	}
	l.level = level

	return l, nil
}

func (f *fileLogger) Close()           { _ = f.fileObj.Close() }
func (f *fileLogger) Type() SenderType { return File }
func (f *fileLogger) Name() string     { return f.name }

func (f *fileLogger) createLogger() {
	f.logger = log.New(f.fileObj, strings.Join([]string{"[", f.name, "] "}, ""), log.LstdFlags)
}

func (f *fileLogger) Send(p level.Priority, m message.Composer) {
	if !GetMessageInfo(f.level, p, m).ShouldLog() {
		return
	}

	f.logger.Printf(f.template, int(p), m.Resolve())
}

func (f *fileLogger) SetName(name string) {
	f.Lock()
	defer f.Unlock()

	f.name = name
	f.createLogger()
}

func (f *fileLogger) SetLevel(l LevelInfo) error {
	if !l.Valid() {
		return fmt.Errorf("level settings are not valid: %+v", l)
	}

	f.Lock()
	defer f.Unlock()

	f.level = l

	return nil
}

func (f *fileLogger) Level() LevelInfo {
	f.RLock()
	defer f.RUnlock()

	return f.level
}
