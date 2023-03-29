package main

import (
	"errors"
	"fmt"
)

func NewSinglyLinkedListNode(value int, next *SinglyLinkedListNode) *SinglyLinkedListNode {
	return &SinglyLinkedListNode{
		value: value,
		next:  next,
	}
}

type SinglyLinkedListNode struct {
	value int
	next  *SinglyLinkedListNode
}

func (r *SinglyLinkedListNode) Value() int {
	return r.value
}

func (r *SinglyLinkedListNode) Next() *SinglyLinkedListNode {
	return r.next
}

func NewSinglyLinkedList(values ...int) *SinglyLinkedList {
	list := &SinglyLinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}

	if len(values) > 0 {
		for i := 0; i < len(values); i++ {
			node := NewSinglyLinkedListNode(values[i], nil)

			if list.Head() == nil {
				list.head = node
				list.tail = node
			} else {
				list.tail.next = node
				list.tail = node
			}
		}

		list.length = len(values)
	}

	return list
}

type SinglyLinkedList struct {
	head   *SinglyLinkedListNode
	tail   *SinglyLinkedListNode
	length int
}

func (r *SinglyLinkedList) Size() int {
	return r.length
}

func (r *SinglyLinkedList) Head() *SinglyLinkedListNode {
	return r.head
}

func (r *SinglyLinkedList) Tail() *SinglyLinkedListNode {
	return r.tail
}

func (r *SinglyLinkedList) ForEachNode(callback func(node *SinglyLinkedListNode)) {
	if r.IsEmpty() {
		return
	}

	node := r.Head()
	for {
		if node == nil {
			break
		}

		callback(node)
		node = node.next
	}
}

func (r *SinglyLinkedList) IsEmpty() bool {
	return r.Head() == nil
}

func (r *SinglyLinkedList) Append(node *SinglyLinkedListNode) {
	if r.IsEmpty() {
		r.head = node
		r.tail = node
	} else {
		r.tail.next = node
		r.tail = node
	}
}

func (r *SinglyLinkedList) Prepend(node *SinglyLinkedListNode) {
	if r.IsEmpty() {
		r.head = node
		r.tail = node
	} else {
		node.next = r.head
		r.head = node
	}
}

func (r *SinglyLinkedList) InsertAfter(value int, newNode *SinglyLinkedListNode) error {
	if r.IsEmpty() {
		return errors.New("list is empty")
	}

	didInsert := false
	r.ForEachNode(func(node *SinglyLinkedListNode) {
		if node.Value() == value {
			didInsert = true

			if node.Next() == nil {
				r.tail = newNode
			}

			temp := node.Next()
			node.next = newNode
			newNode.next = temp
		}
	})

	if didInsert == false {
		return errors.New(fmt.Sprintf("could not find element with value %v in list", value))
	}

	return nil
}

func (r *SinglyLinkedList) InsertBefore(value int, newNode *SinglyLinkedListNode) error {
	if r.IsEmpty() {
		return errors.New("list is empty")
	}

	didInsert := false
	var current *SinglyLinkedListNode
	var prev *SinglyLinkedListNode

	for i := 0; i < r.length; i++ {
		if current == nil && prev == nil {
			current = r.Head()
		} else {
			prev = current
			current = current.Next()
		}

		if didInsert == true {
			break
		}

		if i == 0 && current.Value() == value {
			didInsert = true
			r.head = newNode
			newNode.next = current
		}
		if i != 0 && current.Value() == value {
			didInsert = true
			newNode.next = current
			prev.next = newNode
		}
	}

	if didInsert == false {
		return errors.New(fmt.Sprintf("could not find node with %v value in the list", value))
	}

	return nil
}
