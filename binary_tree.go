package main

import (
	"fmt"
	"io"
)

type BinaryTreeNode struct {
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
	Value int
}

func NewBinaryTreeNode(value int) *BinaryTreeNode {
	return &BinaryTreeNode{
		Left:  nil,
		Right: nil,
		Value: value,
	}
}

func (r *BinaryTreeNode) Insert(node *BinaryTreeNode) {
	if r.Value > node.Value {
		if r.Right == nil {
			r.Right = node
		} else {
			r.Right.Insert(node)
		}
	} else {
		if r.Left == nil {
			r.Left = node
		} else {
			r.Left.Insert(node)
		}
	}
}

type BinaryTree struct {
	root *BinaryTreeNode
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{
		root: nil,
	}
}

func (r *BinaryTree) Insert(node *BinaryTreeNode) {
	if r.root == nil {
		r.root = node
	} else {
		r.root.Insert(node)
	}
}

func (r *BinaryTree) Print(printer io.Writer, node *BinaryTreeNode, padding int, prefix string) {
	if node == nil {
		return
	}

	for i := 0; i < padding; i++ {
		fmt.Fprint(printer, " ")
	}

	fmt.Fprintf(printer, "%v: %v\n", prefix, node.Value)

	r.Print(printer, node.Left, padding+2, "Left")
	r.Print(printer, node.Right, padding+2, "Right")
}
