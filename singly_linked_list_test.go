package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSinglyLinkedList(t *testing.T) {
	t.Run("should initialize empty list", func(t *testing.T) {
		list := NewSinglyLinkedList()

		assert.Nil(t, list.Head)
		assert.Nil(t, list.Tail)
	})

	t.Run("should initialize singly linked list with values", func(t *testing.T) {
		list := NewSinglyLinkedList(1, 2, 3)

		head := list.Head
		next := head.Next
		assert.Equal(t, head.Value, 1)
		assert.Equal(t, next.Value, 2)
		assert.Equal(t, list.Head.Next.Next.Value, 3)
	})
}

func TestSinglyLinkedList_Append(t *testing.T) {
	t.Run("should append item to an empty list", func(t *testing.T) {
		list := NewSinglyLinkedList()

		list.Append(5)

		head := list.Head
		tail := list.Tail

		assert.Equal(t, 5, head.Value)
		assert.Equal(t, head, tail)
	})

	t.Run("should append multiple items to the list", func(t *testing.T) {
		list := NewSinglyLinkedList()

		list.Append(5)
		list.Append(10)

		head := list.Head
		tail := list.Tail
		second := head.Next

		assert.Equal(t, 5, head.Value)
		assert.Equal(t, 10, second.Value)
		assert.Equal(t, 10, tail.Value)
		assert.Nil(t, tail.Next)
	})
}
