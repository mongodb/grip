package grip

import (
	"errors"
	"strings"
)

// Provides an interface to collect and coalesse error messages within
// a function or other sequence of operations. Used to implement a kind
// of "continue on error"-style operations

type MultiCatcher struct {
	errs  []string
	count int
}

func NewCatcher() *MultiCatcher {
	return &MultiCatcher{}
}

func (self *MultiCatcher) Add(err error) {
	if err != nil {
		self.errs = append(self.errs, err.Error())
		self.count += 1
	}
}

func (self *MultiCatcher) String() string {
	return strings.Join(self.errs, ", ")
}

func (self *MultiCatcher) Resolve() (err error) {
	if self.count == 0 {
		err = nil
	} else if self.count == 1 {
		err = errors.New(self.errs[0])
	} else {
		err = errors.New(self.String())
	}

	self = NewCatcher()
	return err
}
