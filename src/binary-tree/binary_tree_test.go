package binary_tree

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestBinaryTree_Insert(t *testing.T) {
	t.Run("should initialize empty binary tree", func(t *testing.T) {
		tree := NewTree(&strings.Builder{})

		assert.Nil(t, tree.root)
	})

	t.Run("should add root node", func(t *testing.T) {
		root := NewNode(5)
		tree := NewTree(&strings.Builder{})

		tree.Insert(root)

		assert.Equal(t, tree.root, root)
	})

	t.Run("should add root node and a few nodes", func(t *testing.T) {
		root := NewNode(5)
		left := NewNode(0)
		right := NewNode(1)
		tree := NewTree(&strings.Builder{})

		tree.Insert(root)
		tree.Insert(left)
		tree.Insert(right)

		assert.Equal(t, tree.root, root)
		assert.Equal(t, tree.root.Left, left)
		assert.Equal(t, tree.root.Right, right)
	})
}
