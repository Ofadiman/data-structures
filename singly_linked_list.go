package main

import (
	"encoding/json"
	"errors"
)

type SinglyLinkedListNode struct {
	Value int
	Next  *SinglyLinkedListNode
}

type SinglyLinkedList struct {
	Head *SinglyLinkedListNode
	Tail *SinglyLinkedListNode
}

func NewSinglyLinkedList(values ...int) *SinglyLinkedList {
	list := &SinglyLinkedList{}

	for _, value := range values {
		list.Append(value)
	}

	return list
}

func (r *SinglyLinkedList) Append(value int) {
	newNode := &SinglyLinkedListNode{Value: value, Next: nil}

	if r.Head == nil {
		r.Head = newNode
		r.Tail = newNode
	} else {
		r.Tail.Next = newNode
		r.Tail = newNode
	}
}

func (r *SinglyLinkedList) Prepend(value int) {
	newNode := &SinglyLinkedListNode{
		Value: value,
		Next:  r.Head,
	}

	if r.Head == nil {
		r.Head = newNode
		r.Tail = newNode
	} else {
		r.Head = newNode
	}
}

func (r *SinglyLinkedList) Insert(value int, position int) {
	newNode := &SinglyLinkedListNode{
		Value: value,
	}

	if position <= 0 || r.Head == nil {
		r.Prepend(value)
		return
	}

	current := r.Head
	for i := 1; i < position && current.Next != nil; i++ {
		current = current.Next
	}

	newNode.Next = current.Next
	current.Next = newNode

	if newNode.Next == nil {
		r.Tail = newNode
	}
}

func (r *SinglyLinkedList) Remove(value int) {
	if r.Head == nil {
		return
	}

	if r.Head.Value == value {
		r.Head = r.Head.Next
		if r.Head == nil {
			r.Tail = nil
		}
		return
	}

	current := r.Head
	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
			if current.Next == nil {
				r.Tail = current
			}
			return
		}
		current = current.Next
	}
}

func (r *SinglyLinkedList) RemoveAt(index int) error {
	if r.Head == nil || index < 0 {
		return errors.New("invalid index")
	}

	if index == 0 {
		r.Head = r.Head.Next
		if r.Head == nil {
			r.Tail = nil
		}
		return nil
	}

	current := r.Head
	for i := 1; i < index && current.Next != nil; i++ {
		current = current.Next
	}

	if current.Next == nil {
		return errors.New("invalid index")
	}

	current.Next = current.Next.Next
	if current.Next == nil {
		r.Tail = current
	}
	return nil
}

func (r *SinglyLinkedList) Find(value int) *SinglyLinkedListNode {
	current := r.Head
	for current != nil {
		if current.Value == value {
			return current
		}
		current = current.Next
	}
	return nil
}

func (r *SinglyLinkedList) IndexOf(value int) int {
	current := r.Head
	index := 0
	for current != nil {
		if current.Value == value {
			return index
		}
		index++
		current = current.Next
	}
	return -1
}

func (r *SinglyLinkedList) Reverse() {
	var prev, current, next *SinglyLinkedListNode
	current = r.Head
	r.Tail = current

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	r.Head = prev
}

type JsonNode struct {
	Value int `json:"value"`
}

func (r *SinglyLinkedList) Serialize() (string, error) {
	nodes := make([]JsonNode, 0)

	current := r.Head
	for current != nil {
		nodes = append(nodes, JsonNode{Value: current.Value})
		current = current.Next
	}

	jsonData, err := json.Marshal(nodes)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
