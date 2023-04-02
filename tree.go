package main

import "github.com/google/uuid"

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
