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

func TestTree_FindNodeByID(t *testing.T) {
	t.Run("should return error when node with passed id does not exist", func(t *testing.T) {
		root := NewTreeNode(5)
		tree := NewTree(root)

		node, err := tree.FindNodeByID("56c40ac1-ef6a-430d-9d16-f0a17cd9c142")

		assert.Nil(t, node)
		assert.Error(t, err)
	})

	t.Run("should return root node", func(t *testing.T) {
		root := NewTreeNode(5)
		tree := NewTree(root)

		node, err := tree.FindNodeByID(root.id.String())

		assert.Equal(t, root, node)
		assert.Nil(t, err)
	})

	t.Run("should return node when node with passed id exists", func(t *testing.T) {
		root := NewTreeNode(5)
		child := NewTreeNode(10)
		root.children = append(root.children, child)
		tree := NewTree(root)

		node, err := tree.FindNodeByID(child.id.String())

		assert.Equal(t, child, node)
		assert.Nil(t, err)
	})
}
