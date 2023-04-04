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
