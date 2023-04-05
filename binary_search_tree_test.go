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

func TestBinarySearchTree_Search(t *testing.T) {
	t.Run("should return nil for an empty tree", func(t *testing.T) {
		tree := NewBinarySearchTree()

		assert.Nil(t, tree.Search(5))
	})

	t.Run("should return the node with the searched value if present", func(t *testing.T) {
		tree := NewBinarySearchTree(8, 3, 10, 1, 6, 14, 4, 7, 13)

		node := tree.Search(6)
		assert.NotNil(t, node)
		assert.Equal(t, 6, node.Value)
	})

	t.Run("should return nil if the searched value is not present", func(t *testing.T) {
		tree := NewBinarySearchTree(8, 3, 10, 1, 6, 14, 4, 7, 13)

		assert.Nil(t, tree.Search(5))
	})
}

func TestBinarySearchTree_Delete(t *testing.T) {
	t.Run("should not panic when deleting from an empty tree", func(t *testing.T) {
		tree := NewBinarySearchTree()

		assert.NotPanics(t, func() { tree.Delete(5) })
	})

	t.Run("should delete a leaf node", func(t *testing.T) {
		tree := NewBinarySearchTree(8, 3, 10, 1, 6, 14, 4, 7, 13)

		tree.Delete(1)
		assert.Nil(t, tree.Search(1))
	})

	t.Run("should delete a node with one child", func(t *testing.T) {
		tree := NewBinarySearchTree(8, 3, 10, 1, 6, 14, 4, 7, 13)

		tree.Delete(14)
		assert.Nil(t, tree.Search(14))
		assert.Equal(t, 13, tree.Root.Right.Right.Value)
	})

	t.Run("should delete a node with two children", func(t *testing.T) {
		tree := NewBinarySearchTree(8, 3, 10, 1, 6, 14, 4, 7, 13)

		tree.Delete(3)
		assert.Nil(t, tree.Search(3))
		assert.Equal(t, 4, tree.Root.Left.Value)
	})

	t.Run("should delete the root node", func(t *testing.T) {
		tree := NewBinarySearchTree(8, 3, 10, 1, 6, 14, 4, 7, 13)

		tree.Delete(8)
		assert.Nil(t, tree.Search(8))
		assert.Equal(t, 10, tree.Root.Value)
	})
}

func TestBinarySearchTree_InOrderTraversalWithCallback(t *testing.T) {
	tree := NewBinarySearchTree(8, 3, 10, 1, 6, 14, 4, 7, 13)
	expectedOrder := []int{1, 3, 4, 6, 7, 8, 10, 13, 14}
	orderIndex := 0

	tree.InOrderTraversal(func(node *BinarySearchTreeNode) {
		assert.Equal(t, expectedOrder[orderIndex], node.Value)
		orderIndex++
	})
}
