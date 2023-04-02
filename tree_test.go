package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTreeNode(t *testing.T) {
	t.Run("should create node", func(t *testing.T) {
		node := NewTreeNode(5)

		assert.Equal(t, 5, node.value)
		assert.Empty(t, node.children)
		assert.IsType(t, "string", node.id.String())
	})
}

func TestNewTree(t *testing.T) {
	t.Run("should create tree", func(t *testing.T) {
		root := NewTreeNode(5)
		tree := NewTree(root)

		assert.Equal(t, root, tree.root)
	})
}
