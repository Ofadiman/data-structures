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

type Tree struct {
	root *Node
}

func NewTree() *Tree {
	return &Tree{
		root: nil,
	}
}

func (r *Tree) Insert(node *Node) {
	if r.root == nil {
		r.root = node
	} else {
		r.root.Insert(node)
	}
}

func (r *Tree) Print(printer io.Writer, node *Node, padding int, prefix string) {
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
