package main

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTreeNode(t *testing.T) {
	t.Run("should create node", func(t *testing.T) {
		node := NewTreeNode(5, nil)

		assert.Equal(t, 5, node.Value)
		assert.Empty(t, node.Children)
		assert.IsType(t, "string", node.ID.String())
	})
}

func TestNewTree(t *testing.T) {
	t.Run("should create tree", func(t *testing.T) {
		root := NewTreeNode(5, nil)
		tree := NewTree(root)

		assert.Equal(t, root, tree.Root)
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

		node, err := tree.FindNodeByID(root.ID.String())

		assert.Equal(t, root, node)
		assert.Nil(t, err)
	})

	t.Run("should return node when node with passed id exists", func(t *testing.T) {
		root := NewTreeNode(5, nil)
		child := NewTreeNode(10, root)
		root.Children = append(root.Children, child)
		tree := NewTree(root)

		node, err := tree.FindNodeByID(child.ID.String())

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

		err := tree.Insert(root.ID.String(), child)

		assert.Nil(t, err)
		assert.Equal(t, []*TreeNode{child}, root.Children)
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

		err := tree.Delete(root.ID.String())

		assert.Nil(t, err)
		assert.Nil(t, tree.Root)
	})

	t.Run("should delete regular node by id", func(t *testing.T) {
		root := NewTreeNode(5, nil)
		child := NewTreeNode(10, root)
		root.Children = append(root.Children, child)
		tree := NewTree(root)

		err := tree.Delete(child.ID.String())

		assert.Nil(t, err)
		assert.Equal(t, []*TreeNode{}, root.Children)
	})
}

func TestTree_ForEachNodeDepthFirst(t *testing.T) {
	t.Run("should call callback for each node in depth-first order", func(t *testing.T) {
		root := NewTreeNode(1, nil)
		child1 := NewTreeNode(2, root)
		child2 := NewTreeNode(3, root)
		child3 := NewTreeNode(4, child1)

		root.Children = append(root.Children, child1, child2)
		child1.Children = append(child1.Children, child3)

		tree := NewTree(root)

		var visited []*TreeNode
		tree.ForEachNodeDepthFirst(func(node *TreeNode) {
			visited = append(visited, node)
		})

		expected := []*TreeNode{root, child1, child3, child2}
		assert.Equal(t, expected, visited)
	})

	t.Run("should not call callback for an empty tree", func(t *testing.T) {
		tree := NewTree(nil)

		var visited []*TreeNode
		tree.ForEachNodeDepthFirst(func(node *TreeNode) {
			visited = append(visited, node)
		})

		assert.Empty(t, visited)
	})

	t.Run("should call callback only for the root node when the tree has only one node", func(t *testing.T) {
		root := NewTreeNode(1, nil)
		tree := NewTree(root)

		var visited []*TreeNode
		tree.ForEachNodeDepthFirst(func(node *TreeNode) {
			visited = append(visited, node)
		})

		expected := []*TreeNode{root}
		assert.Equal(t, expected, visited)
	})
}

func TestTree_ForEachNodeBreadthFirst(t *testing.T) {
	t.Run("should call callback for each node in breadth-first order", func(t *testing.T) {
		root := NewTreeNode(1, nil)
		child1 := NewTreeNode(2, root)
		child2 := NewTreeNode(3, root)
		child3 := NewTreeNode(4, child1)

		root.Children = append(root.Children, child1, child2)
		child1.Children = append(child1.Children, child3)

		tree := NewTree(root)

		var visited []*TreeNode
		tree.ForEachNodeBreadthFirst(func(node *TreeNode) {
			visited = append(visited, node)
		})

		expected := []*TreeNode{root, child1, child2, child3}
		assert.Equal(t, expected, visited)
	})

	t.Run("should not call callback for an empty tree", func(t *testing.T) {
		tree := NewTree(nil)

		var visited []*TreeNode
		tree.ForEachNodeBreadthFirst(func(node *TreeNode) {
			visited = append(visited, node)
		})

		assert.Empty(t, visited)
	})

	t.Run("should call callback only for the root node when the tree has only one node", func(t *testing.T) {
		root := NewTreeNode(1, nil)
		tree := NewTree(root)

		var visited []*TreeNode
		tree.ForEachNodeBreadthFirst(func(node *TreeNode) {
			visited = append(visited, node)
		})

		expected := []*TreeNode{root}
		assert.Equal(t, expected, visited)
	})
}

func TestTree_Serialize(t *testing.T) {
	t.Run("should serialize an empty tree", func(t *testing.T) {
		tree := NewTree(nil)

		serialized, err := tree.Serialize()

		assert.Nil(t, err)
		assert.Equal(t, "null", serialized)
	})

	t.Run("should serialize a tree with only a root node", func(t *testing.T) {
		root := NewTreeNode(1, nil)
		tree := NewTree(root)

		serialized, err := tree.Serialize()

		expected := `{"id":"` + root.ID.String() + `","value":1}`
		assert.Nil(t, err)
		assert.Equal(t, expected, serialized)
	})

	t.Run("should serialize a tree with multiple nodes", func(t *testing.T) {
		root := NewTreeNode(1, nil)
		child1 := NewTreeNode(2, root)
		child2 := NewTreeNode(3, root)
		child3 := NewTreeNode(4, child1)

		root.Children = append(root.Children, child1, child2)
		child1.Children = append(child1.Children, child3)

		tree := NewTree(root)

		serialized, err := tree.Serialize()

		expected := `{"id":"` + root.ID.String() + `","value":1,"children":[{"id":"` + child1.ID.String() + `","value":2,"children":[{"id":"` + child3.ID.String() + `","value":4}]},{"id":"` + child2.ID.String() + `","value":3}]}`
		assert.Nil(t, err)
		assert.Equal(t, expected, serialized)
	})
}

func TestTree_Deserialize(t *testing.T) {
	t.Run("should return an error when deserializing into a nil tree", func(t *testing.T) {
		var tree *Tree

		err := tree.Deserialize(`{"id":"` + uuid.New().String() + `","value":1}`)

		assert.Error(t, err)
	})

	t.Run("should deserialize an empty tree", func(t *testing.T) {
		tree := NewTree(nil)

		err := tree.Deserialize("null")

		assert.Nil(t, err)
		assert.Nil(t, tree.Root)
	})

	t.Run("should deserialize a tree with only a root node", func(t *testing.T) {
		tree := NewTree(nil)

		data := `{"id":"` + uuid.New().String() + `","value":1}`
		err := tree.Deserialize(data)

		assert.Nil(t, err)
		assert.NotNil(t, tree.Root)
		assert.Equal(t, 1, tree.Root.Value)
	})

	t.Run("should deserialize a tree with multiple nodes", func(t *testing.T) {
		tree := NewTree(nil)

		data := `{"id":"` + uuid.New().String() + `","value":1,"children":[{"id":"` + uuid.New().String() + `","value":2,"children":[{"id":"` + uuid.New().String() + `","value":4}]},{"id":"` + uuid.New().String() + `","value":3}]}`
		err := tree.Deserialize(data)

		assert.Nil(t, err)

		root := tree.Root
		child1 := root.Children[0]
		child2 := root.Children[1]
		child3 := child1.Children[0]

		assert.NotNil(t, root)
		assert.Equal(t, 1, root.Value)
		assert.Equal(t, 2, child1.Value)
		assert.Equal(t, 3, child2.Value)
		assert.Equal(t, 4, child3.Value)
	})
}

func TestTree_Height(t *testing.T) {
	t.Run("should return 0 for empty tree", func(t *testing.T) {
		tree := NewTree(nil)

		assert.Equal(t, 0, tree.Height())
	})

	t.Run("should return 0 for tree with only root", func(t *testing.T) {
		root := NewTreeNode(1, nil)
		tree := NewTree(root)

		assert.Equal(t, 0, tree.Height())
	})

	t.Run("should return height for balanced tree", func(t *testing.T) {
		root := NewTreeNode(1, nil)
		child1 := NewTreeNode(2, root)
		child2 := NewTreeNode(3, root)
		root.Children = append(root.Children, child1, child2)

		child1_1 := NewTreeNode(4, child1)
		child1.Children = append(child1.Children, child1_1)

		child2_1 := NewTreeNode(5, child2)
		child2.Children = append(child2.Children, child2_1)

		tree := NewTree(root)

		assert.Equal(t, 2, tree.Height())
	})

	t.Run("should return height for unbalanced tree", func(t *testing.T) {
		root := NewTreeNode(1, nil)
		child1 := NewTreeNode(2, root)
		child2 := NewTreeNode(3, root)
		root.Children = append(root.Children, child1, child2)

		child1_1 := NewTreeNode(4, child1)
		child1.Children = append(child1.Children, child1_1)

		child1_1_1 := NewTreeNode(6, child1_1)
		child1_1.Children = append(child1_1.Children, child1_1_1)

		tree := NewTree(root)

		assert.Equal(t, 3, tree.Height())
	})
}
