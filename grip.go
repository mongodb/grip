package grip

import (
	"errors"
	"fmt"
	"strings"
)

func Catch(err error) {
	if err != nil {
		fmt.Println("ERROR:", err)
	}
}

type multiCatcher struct {
	errs  []string
	count int
}

func NewCatcher() *multiCatcher {
	return &multiCatcher{}
}

func (self *multiCatcher) Add(err error) {
	if err != nil {
		self.errs = append(self.errs, err.Error())
		self.count += 1
	}
}

func (self *multiCatcher) Resolve() (err error) {
	if self.count == 0 {
		err = nil
	} else if self.count == 1 {
		err = errors.New(self.errs[0])
	} else {
		err = errors.New(strings.Join(self.errs, ", "))
	}

	self = NewCatcher()
	return err
}
