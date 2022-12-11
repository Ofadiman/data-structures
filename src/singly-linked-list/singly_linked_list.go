package singly_linked_list

type Node struct {
	value int
	next  *Node
}

type SinglyLinkedList struct {
	head   *Node
	tail   *Node
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
				node := Node{
					value: values[0],
					next:  nil,
				}

				list.head = &node
				list.tail = &node
			}

			node := Node{
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
