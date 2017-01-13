package send

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/tychoish/grip/level"
	"github.com/tychoish/grip/message"
)

type callSiteLogger struct {
	depth  int
	logger *log.Logger
	*base
}

func NewCallSiteConsoleLogger(name string, depth int, l LevelInfo) (Sender, error) {
	s := MakeCallSiteConsoleLogger(depth)

	if err := s.SetLevel(l); err != nil {
		return nil, err
	}

	s.SetName(name)

	return s, nil
}

func MakeCallSiteConsoleLogger(depth int) Sender {
	s := &callSiteLogger{
		depth: depth,
		base:  newBase(""),
	}

	s.level = LevelInfo{level.Trace, level.Trace}

	s.reset = func() {
		s.logger = log.New(os.Stdout, strings.Join([]string{"[", s.Name(), "] "}, ""), log.LstdFlags)
	}

	return s
}

func NewCallSiteFileLogger(name, fileName string, depth int, l LevelInfo) (Sender, error) {
	s, err := MakeCallSiteFileLogger(fileName, depth)
	if err != nil {
		return nil, err
	}

	if err := s.SetLevel(l); err != nil {
		return nil, err
	}

	s.SetName(name)

	return s, nil
}

func MakeCallSiteFileLogger(fileName string, depth int) (Sender, error) {
	s := &callSiteLogger{
		depth: depth,
		base:  newBase(""),
	}

	s.level = LevelInfo{level.Trace, level.Trace}

	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("error opening logging file, %s", err.Error())
	}

	s.reset = func() {
		s.logger = log.New(f, strings.Join([]string{"[", s.Name(), "] "}, ""), log.LstdFlags)
	}

	s.closer = func() error {
		return f.Close()
	}

	return s, nil
}

func (s *callSiteLogger) Type() SenderType { return CallSite }
func (s *callSiteLogger) Send(m message.Composer) {
	if s.level.ShouldLog(m) {
		file, line := callerInfo(s.depth)
		s.logger.Printf("[p=%s] [%s:%d]: %s", m.Priority(), file, line, m.Resolve())
	}
}

func callerInfo(depth int) (string, int) {
	// increase depth to account for callerInfo itself.
	depth++

	// get caller info.
	_, file, line, _ := runtime.Caller(depth)

	// get the directory and filename
	dir, fileName := filepath.Split(file)
	file = filepath.Join(filepath.Base(dir), fileName)

	return file, line
}
