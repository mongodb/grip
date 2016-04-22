package grip

import (
	"errors"
	"strings"
)

// Provides an interface to collect and coalesse error messages within
// a function or other sequence of operations. Used to implement a kind
// of "continue on error"-style operations

type MultiCatcher struct {
	errs []string
}

func NewCatcher() *MultiCatcher {
	return &MultiCatcher{}
}

func (c *MultiCatcher) Add(err error) {
	if err != nil {
		c.errs = append(c.errs, err.Error())
	}
}

func (c *MultiCatcher) Len() int {
	return len(c.errs)
}

func (c *MultiCatcher) HasErrors() bool {
	return c.Len() > 0
}

func (c *MultiCatcher) String() string {
	return strings.Join(c.errs, ", ")
}

func (c *MultiCatcher) Resolve() (err error) {
	if c.Len() == 0 {
		err = nil
	} else if c.Len() == 1 {
		err = errors.New(c.errs[0])
	} else {
		err = errors.New(c.String())
	}

	c = NewCatcher()
	return err
}
