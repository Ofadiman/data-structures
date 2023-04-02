package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryTree_Insert(t *testing.T) {
	t.Run("should initialize empty binary tree", func(t *testing.T) {
		tree := NewBinaryTree()

		assert.Nil(t, tree.root)
	})

	t.Run("should add root node", func(t *testing.T) {
		root := NewBinaryTreeNode(5)
		tree := NewBinaryTree()

		tree.Insert(root)

		assert.Equal(t, tree.root, root)
	})

	t.Run("should add root node and a few nodes", func(t *testing.T) {
		root := NewBinaryTreeNode(5)
		n1 := NewBinaryTreeNode(0)
		n2 := NewBinaryTreeNode(1)
		tree := NewBinaryTree()

		tree.Insert(root)
		tree.Insert(n1)
		tree.Insert(n2)

		assert.Equal(t, tree.root, root)
		assert.Equal(t, tree.root.Right, n1)
		assert.Equal(t, tree.root.Right.Left, n2)
	})
}

func TestBinaryTree_ForEach(t *testing.T) {
	t.Run("should print binary tree with nodes", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		root := NewBinaryTreeNode(10)
		n1 := NewBinaryTreeNode(11)
		n2 := NewBinaryTreeNode(9)
		n3 := NewBinaryTreeNode(100)
		n4 := NewBinaryTreeNode(99)
		n5 := NewBinaryTreeNode(200)
		tree := NewBinaryTree()

		tree.Insert(root)
		tree.Insert(n1)
		tree.Insert(n2)
		tree.Insert(n3)
		tree.Insert(n4)
		tree.Insert(n5)

		tree.Print(buffer, tree.root, 0, "Root")

		assert.Equal(t, buffer.String(), "Root: 10\n  Left: 11\n    Left: 100\n      Left: 200\n      Right: 99\n  Right: 9\n")
	})
}
