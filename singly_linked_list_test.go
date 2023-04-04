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

func TestSinglyLinkedList_Prepend(t *testing.T) {
	t.Run("should prepend to an empty list", func(t *testing.T) {
		list := NewSinglyLinkedList()

		list.Prepend(1)

		assert.Equal(t, 1, list.Head.Value)
		assert.Equal(t, 1, list.Tail.Value)
		assert.Nil(t, list.Head.Next)
	})

	t.Run("should prepend to a non-empty list", func(t *testing.T) {
		list := NewSinglyLinkedList(2)

		list.Prepend(1)

		assert.Equal(t, 1, list.Head.Value)
		assert.Equal(t, 2, list.Tail.Value)
		assert.Equal(t, list.Head.Next, list.Tail)
	})
}

func TestSinglyLinkedList_Insert(t *testing.T) {
	t.Run("should insert at the beginning of the list", func(t *testing.T) {
		list := NewSinglyLinkedList(1)

		list.Insert(0, 0)

		assert.Equal(t, 0, list.Head.Value)
		assert.Equal(t, 1, list.Tail.Value)
		assert.Equal(t, list.Head.Next, list.Tail)
	})

	t.Run("should insert at the end of the list", func(t *testing.T) {
		list := NewSinglyLinkedList(1, 2)

		list.Insert(3, 2)

		assert.Equal(t, 1, list.Head.Value)
		assert.Equal(t, 3, list.Tail.Value)
		assert.Equal(t, list.Head.Next.Next, list.Tail)
	})

	t.Run("should insert in the middle of the list", func(t *testing.T) {
		list := NewSinglyLinkedList(1, 3)

		list.Insert(2, 1)

		assert.Equal(t, 1, list.Head.Value)
		assert.Equal(t, 2, list.Head.Next.Value)
		assert.Equal(t, 3, list.Tail.Value)
	})

	t.Run("should insert at the end when position is greater than list length", func(t *testing.T) {
		list := NewSinglyLinkedList(1, 2)

		list.Insert(3, 4)

		assert.Equal(t, 1, list.Head.Value)
		assert.Equal(t, 2, list.Head.Next.Value)
		assert.Equal(t, 3, list.Tail.Value)
		assert.Equal(t, list.Head.Next.Next, list.Tail)
	})
}

func TestSinglyLinkedList_Remove(t *testing.T) {
	t.Run("should remove the head when the value matches", func(t *testing.T) {
		list := NewSinglyLinkedList(1, 2, 3)

		list.Remove(1)

		assert.Equal(t, 2, list.Head.Value)
		assert.Equal(t, 3, list.Head.Next.Value)
	})

	t.Run("should remove the tail when the value matches", func(t *testing.T) {
		list := NewSinglyLinkedList(1, 2, 3)

		list.Remove(3)

		assert.Equal(t, 1, list.Head.Value)
		assert.Equal(t, 2, list.Head.Next.Value)
		assert.Nil(t, list.Head.Next.Next)
	})

	t.Run("should remove a middle node when the value matches", func(t *testing.T) {
		list := NewSinglyLinkedList(1, 2, 3)

		list.Remove(2)

		assert.Equal(t, 1, list.Head.Value)
		assert.Equal(t, 3, list.Head.Next.Value)
	})

	t.Run("should do nothing when the value does not exist", func(t *testing.T) {
		list := NewSinglyLinkedList(1, 2, 3)

		list.Remove(4)

		assert.Equal(t, 1, list.Head.Value)
		assert.Equal(t, 2, list.Head.Next.Value)
		assert.Equal(t, 3, list.Head.Next.Next.Value)
	})
}

func TestSinglyLinkedList_RemoveAt(t *testing.T) {
	t.Run("should return an error when the list is empty", func(t *testing.T) {
		list := NewSinglyLinkedList()

		err := list.RemoveAt(0)

		assert.Error(t, err)
	})

	t.Run("should return an error when the index is negative", func(t *testing.T) {
		list := NewSinglyLinkedList(1, 2, 3)

		err := list.RemoveAt(-1)

		assert.Error(t, err)
	})

	t.Run("should remove the head when the index is 0", func(t *testing.T) {
		list := NewSinglyLinkedList(1, 2, 3)

		err := list.RemoveAt(0)

		assert.Nil(t, err)
		assert.Equal(t, 2, list.Head.Value)
		assert.Equal(t, 3, list.Head.Next.Value)
	})

	t.Run("should remove a middle node when the index is valid", func(t *testing.T) {
		list := NewSinglyLinkedList(1, 2, 3)

		err := list.RemoveAt(1)

		assert.Nil(t, err)
		assert.Equal(t, 1, list.Head.Value)
		assert.Equal(t, 3, list.Head.Next.Value)
	})

	t.Run("should remove the tail when the index is valid", func(t *testing.T) {
		list := NewSinglyLinkedList(1, 2, 3)

		err := list.RemoveAt(2)

		assert.Nil(t, err)
		assert.Equal(t, 1, list.Head.Value)
		assert.Equal(t, 2, list.Head.Next.Value)
		assert.Nil(t, list.Head.Next.Next)
	})

	t.Run("should return an error when the index is greater than the list length", func(t *testing.T) {
		list := NewSinglyLinkedList(1, 2, 3)

		err := list.RemoveAt(3)

		assert.Error(t, err)
	})
}
