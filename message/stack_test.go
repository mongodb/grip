package message

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintStack(t *testing.T) {
	assert := assert.New(t)
	stack := funcA()
	// Depending on whether this is run inside of a GOPATH, it may produce
	// different results. If grip is a subdirectory of the GOPATH, it'll include
	// the full path to it (e.g. github.com/mongodb/grip/message/stack_test.go).
	// However, if it's run outside of the GOPATH, it only includes the path
	// relative to the go module root directory.
	assert.Contains(stack, `message/stack_test.go:26 (funcC)`)
	assert.Contains(stack, `message/stack_test.go:22 (funcB)`)
	assert.Contains(stack, `message/stack_test.go:18 (funcA)`)
}

func funcA() string {
	return funcB()
}

func funcB() string {
	return funcC()
}

func funcC() string {
	return NewStack(0, "").String()
}

// don't add any code above this line unless you modify the line numbers in TestPrintStack
