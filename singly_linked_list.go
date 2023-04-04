package main

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
