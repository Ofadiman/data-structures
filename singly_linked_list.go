package main

type SinglyLinkedListNode struct {
	value int
	next  *SinglyLinkedListNode
}

type SinglyLinkedList struct {
	head   *SinglyLinkedListNode
	tail   *SinglyLinkedListNode
	length int
}

func NewSinglyLinkedList(values ...int) SinglyLinkedList {
	list := SinglyLinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}

	if len(values) > 0 {
		for i := 0; i < len(values); i++ {
			if i == 0 {
				node := SinglyLinkedListNode{
					value: values[0],
					next:  nil,
				}

				list.head = &node
				list.tail = &node
			}

			node := SinglyLinkedListNode{
				value: values[i],
				next:  nil,
			}

			list.tail.next = &node
			list.tail = &node
		}
	}

	list.length = len(values)

	return list
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
