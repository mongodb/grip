package message

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func funcA() string {
	return funcB()
}

func funcB() string {
	return funcC()
}

func funcC() string {
	return NewStack(0, "").String()
}

// Don't add any code above this line unless you modify the line numbers in
// TestPrintStack.

func TestPrintStack(t *testing.T) {
	assert := assert.New(t)
	stack := funcA()
	assert.Contains(stack, `message/stack_test.go:18 (funcC)`)
	assert.Contains(stack, `message/stack_test.go:14 (funcB)`)
	assert.Contains(stack, `message/stack_test.go:10 (funcA)`)
}
