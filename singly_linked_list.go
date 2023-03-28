package main

import "fmt"

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
	if r.Head() == nil {
		return
	}

	node := r.Head()
	for {
		if node == nil {
			break
		}

		fmt.Printf("calling callback with node: {value: %v}\n", node.value)
		callback(node)
		node = node.next
	}
}
