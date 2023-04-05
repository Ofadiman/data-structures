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

func (r *BinarySearchTree) Search(value int) *BinarySearchTreeNode {
	return searchNode(r.Root, value)
}

func searchNode(currentNode *BinarySearchTreeNode, value int) *BinarySearchTreeNode {
	if currentNode == nil || currentNode.Value == value {
		return currentNode
	}

	if value < currentNode.Value {
		return searchNode(currentNode.Left, value)
	} else {
		return searchNode(currentNode.Right, value)
	}
}

func (r *BinarySearchTree) Delete(value int) {
	r.Root = deleteNode(r.Root, value)
}

func deleteNode(currentNode *BinarySearchTreeNode, value int) *BinarySearchTreeNode {
	if currentNode == nil {
		return nil
	}

	if value < currentNode.Value {
		currentNode.Left = deleteNode(currentNode.Left, value)
	} else if value > currentNode.Value {
		currentNode.Right = deleteNode(currentNode.Right, value)
	} else {
		// Case 1: Node with only one child or no children
		if currentNode.Left == nil {
			return currentNode.Right
		} else if currentNode.Right == nil {
			return currentNode.Left
		}

		// Case 2: Node with two children
		currentNode.Value = minValue(currentNode.Right)
		currentNode.Right = deleteNode(currentNode.Right, currentNode.Value)
	}
	return currentNode
}

func minValue(node *BinarySearchTreeNode) int {
	minValue := node.Value
	for node.Left != nil {
		node = node.Left
		minValue = node.Value
	}
	return minValue
}
