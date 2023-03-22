package binary_tree

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryTree_Insert(t *testing.T) {
	t.Run("should initialize empty binary tree", func(t *testing.T) {
		tree := NewTree()

		assert.Nil(t, tree.root)
	})

	t.Run("should add root node", func(t *testing.T) {
		root := NewNode(5)
		tree := NewTree()

		tree.Insert(root)

		assert.Equal(t, tree.root, root)
	})

	t.Run("should add root node and a few nodes", func(t *testing.T) {
		root := NewNode(5)
		n1 := NewNode(0)
		n2 := NewNode(1)
		tree := NewTree()

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

		root := NewNode(10)
		n1 := NewNode(11)
		n2 := NewNode(9)
		n3 := NewNode(100)
		n4 := NewNode(99)
		n5 := NewNode(200)
		tree := NewTree()

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
