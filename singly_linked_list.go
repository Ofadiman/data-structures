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

func (l *SinglyLinkedList) Append(value int) {
	newNode := &SinglyLinkedListNode{Value: value, Next: nil}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}
