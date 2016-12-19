package send

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

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
func NewFileLogger(name, filePath string, l LevelInfo) (Sender, error) {
	s := &fileLogger{
		name:     name,
		template: "[p=%s]: %s\n",
	}

	if err := s.SetLevel(l); err != nil {
		return nil, err
	}

	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("error opening logging file, %s, falling back to stdOut logging", err.Error())
	}
	s.fileObj = f
	s.createLogger()

	return s, nil
}

func (f *fileLogger) Close()           { _ = f.fileObj.Close() }
func (f *fileLogger) Type() SenderType { return File }
func (f *fileLogger) Name() string     { return f.name }

func (f *fileLogger) createLogger() {
	f.logger = log.New(f.fileObj, strings.Join([]string{"[", f.name, "] "}, ""), log.LstdFlags)
}

func (f *fileLogger) Send(m message.Composer) {
	if !f.level.ShouldLog(m) {
		return
	}

	f.logger.Printf(f.template, m.Priority(), m.Resolve())
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
