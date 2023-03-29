package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSinglyLinkedList(t *testing.T) {
	t.Run("should initialize empty list", func(t *testing.T) {
		list := NewSinglyLinkedList()

		assert.Nil(t, list.Head())
		assert.Nil(t, list.Tail())
	})

	t.Run("should initialize singly linked list with values", func(t *testing.T) {
		list := NewSinglyLinkedList(1, 2, 3)

		head := list.Head()
		next := head.Next()
		assert.Equal(t, head.Value(), 1)
		assert.Equal(t, next.Value(), 2)
		assert.Equal(t, list.Head().Next().Next().Value(), 3)
	})
}

func TestSinglyLinkedList_Size(t *testing.T) {
	t.Run("should return 0 if list is empty", func(t *testing.T) {
		list := NewSinglyLinkedList()

		assert.Equal(t, list.Size(), 0)
	})

	t.Run("should return list size", func(t *testing.T) {
		list := NewSinglyLinkedList(1, 2, 3)

		assert.Equal(t, list.Size(), 3)
	})
}

func TestSinglyLinkedList_Head(t *testing.T) {
	t.Run("should return nil if list is empty", func(t *testing.T) {
		list := NewSinglyLinkedList()

		assert.Nil(t, list.Head())
	})

	t.Run("should return list head", func(t *testing.T) {
		list := NewSinglyLinkedList(1, 2, 3)

		assert.Equal(t, list.Head().value, 1)
	})

}

func TestSinglyLinkedList_Tail(t *testing.T) {
	t.Run("should return nil if list is empty", func(t *testing.T) {
		list := NewSinglyLinkedList()

		assert.Nil(t, list.Tail())
	})

	t.Run("should return list tail", func(t *testing.T) {
		list := NewSinglyLinkedList(1, 2, 3)

		assert.Equal(t, list.Tail().value, 3)
	})
}

func TestSinglyLinkedList_ForEachNode(t *testing.T) {
	t.Run("should do nothing if list is empty", func(t *testing.T) {
		var elements []int
		callback := func(node *SinglyLinkedListNode) {
			elements = append(elements, node.value)
		}

		list := NewSinglyLinkedList()

		list.ForEachNode(callback)

		assert.Len(t, elements, 0)
	})

	t.Run("should call callback function for each list node", func(t *testing.T) {
		var elements []int
		callback := func(node *SinglyLinkedListNode) {
			elements = append(elements, node.value)
		}

		list := NewSinglyLinkedList(1, 2, 3)

		list.ForEachNode(callback)

		assert.Equal(t, elements, []int{1, 2, 3})
	})
}

func TestSinglyLinkedList_IsEmpty(t *testing.T) {
	t.Run("should return true if list is empty", func(t *testing.T) {
		list := NewSinglyLinkedList()

		assert.Equal(t, true, list.IsEmpty())
	})

	t.Run("should return false if list is not empty", func(t *testing.T) {
		list := NewSinglyLinkedList(1)

		assert.Equal(t, false, list.IsEmpty())
	})
}

func TestSinglyLinkedList_Append(t *testing.T) {
	t.Run("should create head and tail if list is empty", func(t *testing.T) {
		list := NewSinglyLinkedList()

		node := NewSinglyLinkedListNode(5, nil)
		list.Append(node)

		assert.Equal(t, list.Head(), node)
		assert.Equal(t, list.Tail(), node)
	})

	t.Run("should append new node to list", func(t *testing.T) {
		list := NewSinglyLinkedList(1)

		node := NewSinglyLinkedListNode(5, nil)
		list.Append(node)

		assert.Equal(t, list.Tail(), node)
		assert.Equal(t, list.Head().Next(), node)
	})
}

func TestSinglyLinkedList_Prepend(t *testing.T) {
	t.Run("should create head and tail if list is empty", func(t *testing.T) {
		list := NewSinglyLinkedList()

		node := NewSinglyLinkedListNode(1, nil)
		list.Append(node)

		assert.Equal(t, list.Head(), node)
		assert.Equal(t, list.Tail(), node)
	})

	t.Run("should prepend node to list", func(t *testing.T) {
		list := NewSinglyLinkedList(1)

		initialHead := list.Head()
		node := NewSinglyLinkedListNode(2, nil)
		list.Prepend(node)

		assert.Equal(t, list.Head(), node)
		assert.Equal(t, list.Tail(), initialHead)
		assert.Equal(t, list.Head().Next(), initialHead)
	})
}

func TestSinglyLinkedList_InsertAfter(t *testing.T) {
	t.Run("should return error if list is empty", func(t *testing.T) {
		list := NewSinglyLinkedList()

		node := NewSinglyLinkedListNode(1, nil)
		err := list.InsertAfter(5, node)

		assert.Error(t, err)
	})

	t.Run("should return error if node with passed value does not exist in list", func(t *testing.T) {
		list := NewSinglyLinkedList(1)

		node := NewSinglyLinkedListNode(5, nil)
		err := list.InsertAfter(3, node)

		assert.Error(t, err)
	})

	t.Run("should append node to the end of the list", func(t *testing.T) {
		list := NewSinglyLinkedList(1)

		node := NewSinglyLinkedListNode(5, nil)
		err := list.InsertAfter(1, node)

		assert.Nil(t, err)
		assert.Equal(t, list.Tail(), node)
		assert.Equal(t, list.Head().Next(), node)
	})

	t.Run("should insert node after selected node", func(t *testing.T) {
		list := NewSinglyLinkedList(1, 2)

		node := NewSinglyLinkedListNode(5, nil)
		err := list.InsertAfter(1, node)

		assert.Nil(t, err)
		assert.Equal(t, list.Head().Next(), node)
		assert.Equal(t, node.Next().Value(), 2)
	})
}

func TestSinglyLinkedList_InsertBefore(t *testing.T) {
	t.Run("should return error when list is empty", func(t *testing.T) {
		list := NewSinglyLinkedList()

		node := NewSinglyLinkedListNode(3, nil)
		err := list.InsertBefore(5, node)

		assert.Error(t, err)
	})

	t.Run("should return error when node with passed value does not exist", func(t *testing.T) {
		list := NewSinglyLinkedList(1)

		node := NewSinglyLinkedListNode(3, nil)
		err := list.InsertBefore(5, node)

		assert.Error(t, err)
	})

	t.Run("should insert node before expected node and set a new head", func(t *testing.T) {
		list := NewSinglyLinkedList(1)
		initialHead := list.Head()

		node := NewSinglyLinkedListNode(3, nil)
		err := list.InsertBefore(1, node)

		assert.Nil(t, err)
		assert.Equal(t, list.Head(), node)
		assert.Equal(t, node.Next(), initialHead)
	})

	t.Run("should insert node before expected node", func(t *testing.T) {
		list := NewSinglyLinkedList(1, 2)

		node := NewSinglyLinkedListNode(3, nil)
		err := list.InsertBefore(2, node)

		assert.Nil(t, err)
		assert.Equal(t, list.Head().Next(), node)
		assert.Equal(t, node.Next(), list.Tail())
	})
}
