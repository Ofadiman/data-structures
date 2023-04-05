package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBinarySearchTree(t *testing.T) {
	t.Run("should create an empty tree", func(t *testing.T) {
		tree := NewBinarySearchTree()

		assert.Nil(t, tree.Root)
	})

	t.Run("should create a tree with initial values", func(t *testing.T) {
		tree := NewBinarySearchTree(8, 3, 10)

		assert.NotNil(t, tree.Root)
		assert.Equal(t, 8, tree.Root.Value)
		assert.Equal(t, 3, tree.Root.Left.Value)
		assert.Equal(t, 10, tree.Root.Right.Value)
	})
}

func TestBinarySearchTree_Insert(t *testing.T) {
	t.Run("should insert values maintaining the BST property", func(t *testing.T) {
		tree := NewBinarySearchTree()
		tree.Insert(8)
		tree.Insert(3)
		tree.Insert(10)
		tree.Insert(1)
		tree.Insert(6)
		tree.Insert(14)
		tree.Insert(4)
		tree.Insert(7)
		tree.Insert(13)

		assert.NotNil(t, tree.Root)
		assert.Equal(t, 8, tree.Root.Value)
		assert.Equal(t, 3, tree.Root.Left.Value)
		assert.Equal(t, 10, tree.Root.Right.Value)

		assert.Equal(t, 1, tree.Root.Left.Left.Value)
		assert.Equal(t, 6, tree.Root.Left.Right.Value)

		assert.Equal(t, 14, tree.Root.Right.Right.Value)
		assert.Equal(t, 4, tree.Root.Left.Right.Left.Value)
		assert.Equal(t, 7, tree.Root.Left.Right.Right.Value)
		assert.Equal(t, 13, tree.Root.Right.Right.Left.Value)
	})
}
