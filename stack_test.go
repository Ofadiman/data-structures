package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStack(t *testing.T) {
	t.Run("should initialize empty stack", func(t *testing.T) {
		stack := NewStack[int]()

		assert.Equal(t, 0, len(stack.items))
	})
}

func TestStack_Push(t *testing.T) {
	t.Run("should push item to stack", func(t *testing.T) {
		stack := NewStack[int]()

		stack.Push(5)

		top := stack.items[0]
		size := len(stack.items)

		assert.Equal(t, 5, *top)
		assert.Equal(t, 1, size)

		stack.Push(10)

		top = stack.items[1]
		size = len(stack.items)

		assert.Equal(t, 10, *top)
		assert.Equal(t, 2, size)
	})
}

func TestStack_Pop(t *testing.T) {
	t.Run("should return error when stack is empty", func(t *testing.T) {
		stack := NewStack[int]()

		item, err := stack.Pop()

		assert.Nil(t, item)
		assert.Error(t, err)
	})

	t.Run("should pop item from stack", func(t *testing.T) {
		stack := NewStack[int]()

		stack.Push(5)

		top := stack.items[0]
		size := len(stack.items)

		assert.Equal(t, 1, size)
		assert.Equal(t, 5, *top)

		item, err := stack.Pop()
		size = len(stack.items)

		assert.Equal(t, 0, size)
		assert.Equal(t, 5, *item)
		assert.Nil(t, err)
	})
}
