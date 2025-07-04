package goin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackPushAndPop(t *testing.T) {
	stack := &Stack[int]{}
	
	// Test empty stack
	assert.True(t, stack.IsEmpty())
	
	// Test push and pop
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	
	assert.False(t, stack.IsEmpty())
	
	// Test pop in LIFO order
	assert.Equal(t, 3, stack.Pop())
	assert.Equal(t, 2, stack.Pop())
	assert.Equal(t, 1, stack.Pop())
	
	assert.True(t, stack.IsEmpty())
}

func TestStackPeek(t *testing.T) {
	stack := &Stack[string]{}
	
	// Test peek on empty stack
	val, ok := stack.Peek()
	assert.False(t, ok)
	assert.Equal(t, "", val)
	
	// Test peek with elements
	stack.Push("hello")
	stack.Push("world")
	
	val, ok = stack.Peek()
	assert.True(t, ok)
	assert.Equal(t, "world", val)
	
	// Ensure peek doesn't modify stack
	assert.Equal(t, "world", stack.Pop())
	assert.Equal(t, "hello", stack.Pop())
}

func TestStackToSlice(t *testing.T) {
	stack := &Stack[int]{}
	
	// Test empty stack
	slice := stack.ToSlice()
	assert.Empty(t, slice)
	
	// Test with elements
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	
	slice = stack.ToSlice()
	expected := []int{1, 2, 3}
	assert.Equal(t, expected, slice)
}

func TestStackPopEmptyStack(t *testing.T) {
	stack := &Stack[int]{}
	
	// Pop from empty stack should return zero value
	val := stack.Pop()
	assert.Equal(t, 0, val)
}