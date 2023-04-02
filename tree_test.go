package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTreeNode(t *testing.T) {
	t.Run("should create node", func(t *testing.T) {
		node := NewTreeNode(5, nil)

		assert.Equal(t, 5, node.value)
		assert.Empty(t, node.children)
		assert.IsType(t, "string", node.id.String())
	})
}

func TestNewTree(t *testing.T) {
	t.Run("should create tree", func(t *testing.T) {
		root := NewTreeNode(5, nil)
		tree := NewTree(root)

		assert.Equal(t, root, tree.root)
	})
}

func TestTree_FindNodeByID(t *testing.T) {
	t.Run("should return error when node with passed id does not exist", func(t *testing.T) {
		root := NewTreeNode(5, nil)
		tree := NewTree(root)

		node, err := tree.FindNodeByID("56c40ac1-ef6a-430d-9d16-f0a17cd9c142")

		assert.Nil(t, node)
		assert.Error(t, err)
	})

	t.Run("should return root node", func(t *testing.T) {
		root := NewTreeNode(5, nil)
		tree := NewTree(root)

		node, err := tree.FindNodeByID(root.id.String())

		assert.Equal(t, root, node)
		assert.Nil(t, err)
	})

	t.Run("should return node when node with passed id exists", func(t *testing.T) {
		root := NewTreeNode(5, nil)
		child := NewTreeNode(10, root)
		root.children = append(root.children, child)
		tree := NewTree(root)

		node, err := tree.FindNodeByID(child.id.String())

		assert.Equal(t, child, node)
		assert.Nil(t, err)
	})
}

func TestTree_Insert(t *testing.T) {
	t.Run("should return error when node with passed id does not exist", func(t *testing.T) {
		root := NewTreeNode(5, nil)
		child := NewTreeNode(10, root)
		tree := NewTree(root)

		err := tree.Insert("d4363b15-b700-4ed0-94a2-9b12ac28ae46", child)

		assert.Error(t, err)
	})

	t.Run("should insert node when node with passed id exists", func(t *testing.T) {
		root := NewTreeNode(5, nil)
		child := NewTreeNode(10, root)
		tree := NewTree(root)

		err := tree.Insert(root.id.String(), child)

		assert.Nil(t, err)
		assert.Equal(t, []*TreeNode{child}, root.children)
	})
}

func TestTree_Delete(t *testing.T) {
	t.Run("should return error when node with passed id does not exist", func(t *testing.T) {
		root := NewTreeNode(5, nil)
		tree := NewTree(root)

		err := tree.Delete("853e3c2c-b053-4677-9449-b6d073694268")

		assert.Error(t, err)
	})

	t.Run("should delete root node by id", func(t *testing.T) {
		root := NewTreeNode(5, nil)
		tree := NewTree(root)

		err := tree.Delete(root.id.String())

		assert.Nil(t, err)
		assert.Nil(t, tree.root)
	})

	t.Run("should delete regular node by id", func(t *testing.T) {
		root := NewTreeNode(5, nil)
		child := NewTreeNode(10, root)
		root.children = append(root.children, child)
		tree := NewTree(root)

		err := tree.Delete(child.id.String())

		assert.Nil(t, err)
		assert.Equal(t, []*TreeNode{}, root.children)
	})
}
