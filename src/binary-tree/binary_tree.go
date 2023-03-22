package binary_tree

import (
	"fmt"
	"io"
)

type Node struct {
	Left  *Node
	Right *Node
	Value int
}

func NewNode(value int) *Node {
	return &Node{
		Left:  nil,
		Right: nil,
		Value: value,
	}
}

func (r *Node) Insert(node *Node) {
	if r.Left != nil && r.Right != nil {
		if r.Left.Value > r.Right.Value {
			r.Right.Insert(node)
		} else {
			r.Left.Insert(node)
		}
	}

	if r.Left == nil {
		fmt.Println("inserting node at left")
		r.Left = node
	} else {
		fmt.Println("inserting node at right")
		r.Right = node
	}
}

type Tree struct {
	root    *Node
	printer io.Writer
}

func NewTree(printer io.Writer) *Tree {
	return &Tree{
		root:    nil,
		printer: printer,
	}
}

func (r *Tree) Insert(node *Node) {
	if r.root == nil {
		r.root = node
	} else {
		r.root.Insert(node)
	}
}
