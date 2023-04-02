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

func TestStack_Peek(t *testing.T) {
	t.Run("should return error when stack is empty", func(t *testing.T) {
		stack := NewStack[int]()

		top, err := stack.Peek()

		assert.Nil(t, top)
		assert.Error(t, err)
	})

	t.Run("should return last item pushed to the stack", func(t *testing.T) {
		stack := NewStack[int]()

		stack.Push(5)

		size := len(stack.items)
		top, err := stack.Peek()

		assert.Equal(t, 1, size)
		assert.Equal(t, 5, *top)
		assert.Nil(t, err)
	})
}

func TestStack_IsEmpty(t *testing.T) {
	t.Run("should return true when stack is empty", func(t *testing.T) {
		stack := NewStack[int]()

		isEmpty := stack.IsEmpty()

		assert.True(t, isEmpty)
	})

	t.Run("should return false when stack is not empty", func(t *testing.T) {
		stack := NewStack[int]()

		stack.Push(5)

		size := len(stack.items)
		top, err := stack.Peek()

		assert.Equal(t, 1, size)
		assert.Equal(t, 5, *top)
		assert.Nil(t, err)

		isEmpty := stack.IsEmpty()

		assert.False(t, isEmpty)
	})
}

func TestStack_Size(t *testing.T) {
	t.Run("should return stack size when stack is empty", func(t *testing.T) {
		stack := NewStack[int]()

		size := stack.Size()

		assert.Equal(t, 0, size)
	})

	t.Run("should return stack size when stack has elements", func(t *testing.T) {
		stack := NewStack[int]()

		stack.Push(5)

		size := stack.Size()

		assert.Equal(t, 1, size)
	})
}

func TestStack_Empty(t *testing.T) {
	t.Run("should do nothing when stack is already empty", func(t *testing.T) {
		stack := NewStack[int]()

		size := stack.Size()
		top, err := stack.Peek()

		assert.Equal(t, 0, size)
		assert.Nil(t, top)
		assert.Error(t, err)

		stack.Empty()

		size = stack.Size()
		top, err = stack.Peek()

		assert.Equal(t, 0, size)
		assert.Nil(t, top)
		assert.Error(t, err)
	})

	t.Run("should empty stack", func(t *testing.T) {
		stack := NewStack[int]()

		stack.Push(5)
		stack.Push(10)
		stack.Push(15)

		size := stack.Size()
		top, err := stack.Peek()

		assert.Equal(t, 3, size)
		assert.Equal(t, 15, *top)
		assert.Nil(t, err)

		stack.Empty()

		size = stack.Size()
		top, err = stack.Peek()

		assert.Equal(t, 0, size)
		assert.Nil(t, top)
		assert.Error(t, err)
	})
}
