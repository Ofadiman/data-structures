# Data structures

A data structure is a way of organizing and storing data in a computer so that it can be accessed and manipulated
efficiently. A data structure can be thought of as a container that holds data in a particular arrangement, with
specific operations that can be performed on the data it contains. There are many types of data structures, each with
its own advantages and disadvantages. Some of the most common data structures
include `arrays`, `linked lists`, `stacks`, `queues`, `trees`, `graphs`, and `hash tables`.

Implemented data structures:

- [Singly Linked List](singly_linked_list.go)
- [Binary Tree](binary_tree.go)

### Running tests

Tests can be run using Golang built-in `go test` cli:

- `go test ./...` - Run all tests.
- `go test -run TestNewSinglyLinkedList ./...` - Run only tests with name `TestNewSinglyLinkedList`.
