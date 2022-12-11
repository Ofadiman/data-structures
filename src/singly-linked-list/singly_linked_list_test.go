package singly_linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
