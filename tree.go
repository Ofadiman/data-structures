package main

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type TreeNode struct {
	id       uuid.UUID
	parent   *TreeNode
	value    int
	children []*TreeNode
}

func NewTreeNode(value int, parent *TreeNode) *TreeNode {
	return &TreeNode{
		id:       uuid.New(),
		parent:   parent,
		value:    value,
		children: []*TreeNode{},
	}
}

type Tree struct {
	root *TreeNode
}

func NewTree(root *TreeNode) *Tree {
	return &Tree{
		root: root,
	}
}

func (r *Tree) FindNodeByID(nodeID string) (*TreeNode, error) {
	id, err := uuid.Parse(nodeID)
	if err != nil {
		return nil, errors.New("invalid node ID")
	}

	var search func(node *TreeNode) (*TreeNode, error)
	search = func(node *TreeNode) (*TreeNode, error) {
		if node == nil {
			return nil, errors.New("node not found")
		}

		if node.id == id {
			return node, nil
		}

		for _, child := range node.children {
			foundNode, err := search(child)
			if err == nil {
				return foundNode, nil
			}
		}

		return nil, errors.New("node not found")
	}

	return search(r.root)
}

func (r *Tree) Insert(parentID string, child *TreeNode) error {
	parent, err := r.FindNodeByID(parentID)
	if err != nil {
		return err
	}

	parent.children = append(parent.children, child)

	return nil
}

func (r *Tree) Delete(nodeID string) error {
	node, err := r.FindNodeByID(nodeID)
	if err != nil {
		return err
	}
	fmt.Printf("found node %#v\n", node)

	isRoot := node.parent == nil
	if isRoot {
		r.root = nil
		return nil
	}

	for i, child := range node.parent.children {
		if child.id.String() == nodeID {
			node.parent.children = append(node.parent.children[:i], node.parent.children[i+1:]...)
			break
		}
	}

	return nil
}

func (r *Tree) ForEachNodeDepthFirst(callback func(node *TreeNode)) {
	r.depthFirstTraversal(r.root, callback)
}

func (r *Tree) depthFirstTraversal(node *TreeNode, callback func(node *TreeNode)) {
	if node == nil {
		return
	}

	callback(node)

	for _, child := range node.children {
		r.depthFirstTraversal(child, callback)
	}
}

func (r *Tree) ForEachNodeBreadthFirst(callback func(node *TreeNode)) {
	if r.root == nil {
		return
	}

	queue := []*TreeNode{r.root}

	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]

		callback(currentNode)

		for _, child := range currentNode.children {
			queue = append(queue, child)
		}
	}
}
