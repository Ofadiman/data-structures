package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type TreeNode struct {
	ID       uuid.UUID   `json:"id"`
	Parent   *TreeNode   `json:"-"`
	Value    int         `json:"value"`
	Children []*TreeNode `json:"children,omitempty"`
}

func NewTreeNode(value int, parent *TreeNode) *TreeNode {
	return &TreeNode{
		ID:       uuid.New(),
		Parent:   parent,
		Value:    value,
		Children: []*TreeNode{},
	}
}

type Tree struct {
	Root *TreeNode
}

func NewTree(root *TreeNode) *Tree {
	return &Tree{
		Root: root,
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

		if node.ID == id {
			return node, nil
		}

		for _, child := range node.Children {
			foundNode, err := search(child)
			if err == nil {
				return foundNode, nil
			}
		}

		return nil, errors.New("node not found")
	}

	return search(r.Root)
}

func (r *Tree) Insert(parentID string, child *TreeNode) error {
	parent, err := r.FindNodeByID(parentID)
	if err != nil {
		return err
	}

	parent.Children = append(parent.Children, child)

	return nil
}

func (r *Tree) Delete(nodeID string) error {
	node, err := r.FindNodeByID(nodeID)
	if err != nil {
		return err
	}
	fmt.Printf("found node %#v\n", node)

	isRoot := node.Parent == nil
	if isRoot {
		r.Root = nil
		return nil
	}

	for i, child := range node.Parent.Children {
		if child.ID.String() == nodeID {
			node.Parent.Children = append(node.Parent.Children[:i], node.Parent.Children[i+1:]...)
			break
		}
	}

	return nil
}

func (r *Tree) ForEachNodeDepthFirst(callback func(node *TreeNode)) {
	r.depthFirstTraversal(r.Root, callback)
}

func (r *Tree) depthFirstTraversal(node *TreeNode, callback func(node *TreeNode)) {
	if node == nil {
		return
	}

	callback(node)

	for _, child := range node.Children {
		r.depthFirstTraversal(child, callback)
	}
}

func (r *Tree) ForEachNodeBreadthFirst(callback func(node *TreeNode)) {
	if r.Root == nil {
		return
	}

	queue := []*TreeNode{r.Root}

	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]

		callback(currentNode)

		for _, child := range currentNode.Children {
			queue = append(queue, child)
		}
	}
}

func (r *Tree) Serialize() (string, error) {
	data, err := json.Marshal(r.Root)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (r *Tree) Deserialize(data string) error {
	if r == nil {
		return errors.New("cannot deserialize into a nil tree")
	}

	if data == "null" {
		r.Root = nil
		return nil
	}

	var root TreeNode
	err := json.Unmarshal([]byte(data), &root)
	if err != nil {
		return err
	}

	queue := []*TreeNode{&root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, child := range node.Children {
			child.Parent = node
			queue = append(queue, child)
		}
	}

	r.Root = &root

	return nil
}
