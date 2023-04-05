package main

type BinarySearchTreeNode struct {
	Value int
	Left  *BinarySearchTreeNode
	Right *BinarySearchTreeNode
}

func NewBinarySearchTreeNode(value int) *BinarySearchTreeNode {
	return &BinarySearchTreeNode{Value: value}
}

type BinarySearchTree struct {
	Root *BinarySearchTreeNode
}

func NewBinarySearchTree(values ...int) *BinarySearchTree {
	bst := &BinarySearchTree{}
	for _, value := range values {
		bst.Insert(value)
	}
	return bst
}

func (r *BinarySearchTree) Insert(value int) {
	newNode := NewBinarySearchTreeNode(value)

	if r.Root == nil {
		r.Root = newNode
	} else {
		insertNode(r.Root, newNode)
	}
}

func insertNode(currentNode, newNode *BinarySearchTreeNode) {
	if newNode.Value < currentNode.Value {
		if currentNode.Left == nil {
			currentNode.Left = newNode
		} else {
			insertNode(currentNode.Left, newNode)
		}
	} else {
		if currentNode.Right == nil {
			currentNode.Right = newNode
		} else {
			insertNode(currentNode.Right, newNode)
		}
	}
}
