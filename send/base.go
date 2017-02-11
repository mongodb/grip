package send

import (
	"errors"
	"fmt"
	"sync"

	"github.com/tychoish/grip/message"
)

type Base struct {
	name       string
	level      LevelInfo
	reset      func()
	closer     func() error
	errHandler ErrorHandler
	mutex      sync.RWMutex
}

func NewBase(n string) *Base {
	return &Base{
		name:       n,
		reset:      func() {},
		closer:     func() error { return nil },
		errHandler: func(error, message.Composer) {},
	}
}

func (b *Base) Close() error { return b.closer() }

func (b *Base) Name() string {
	b.mutex.RLock()
	defer b.mutex.RUnlock()

	return b.name
}

func (b *Base) SetName(name string) {
	b.mutex.Lock()
	b.name = name
	b.mutex.Unlock()

	b.reset()
}

func (b *Base) SetErrorHandler(eh ErrorHandler) error {
	if eh == nil {
		return errors.New("error handler must be non-nil")
	}

	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.errHandler = eh

	return nil
}

func (b *Base) SetLevel(l LevelInfo) error {
	if !l.Valid() {
		return fmt.Errorf("level settings are not valid: %+v", l)
	}

	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.level = l

	return nil
}

func (b *Base) Level() LevelInfo {
	b.mutex.RLock()
	defer b.mutex.RUnlock()

	return b.level
}
