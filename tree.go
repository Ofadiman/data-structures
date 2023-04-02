package main

import (
	"errors"
	"github.com/google/uuid"
)

type TreeNode struct {
	id       uuid.UUID
	value    int
	children []*TreeNode
}

func NewTreeNode(value int) *TreeNode {
	return &TreeNode{
		id:       uuid.New(),
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
