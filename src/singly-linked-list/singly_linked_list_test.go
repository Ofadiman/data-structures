package singly_linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSinglyLinkedList(t *testing.T) {
	t.Run("should create singly linked list without any arguments", func(t *testing.T) {
		list := NewSinglyLinkedList()

		assert.Nil(t, list.head)
		assert.Nil(t, list.tail)
		assert.Equal(t, list.length, 0)
	})

	t.Run("should create singly linked list with nodes", func(t *testing.T) {
		list := NewSinglyLinkedList(1, 2, 3)

		assert.Equal(t, list.head.value, 1)
		assert.Equal(t, list.tail.value, 3)
		assert.Equal(t, list.length, 3)
	})
}
