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

# Stack

A stack is a linear data structure that follows the Last In, First Out (LIFO) principle, meaning that the last element
added to the stack is the first one to be removed. It can be thought of as a collection of items, where new items can be
added, and existing items can be removed from the top of the stack. It supports three primary operations:

- `Empty` - This operation clears all items from the stack.
- `IsEmpty` - This operation returns true if the stack does not have any items.
- `Peek` - This operation returns the top element of the stack without removing it.
- `Pop` - This operation removes the top element from the stack.
- `Push` - This operation adds an element to the stack.
- `Size` - This operation return the information about current number of items in the stack.

Stack data structure use cases:

1. `Function Call Stack` - In programming languages, a stack is used to manage function calls and their execution
   context. When a function is called, its context is pushed onto the call stack, and when it returns, its context is
   popped from the stack.
2. `Expression Evaluation and Syntax Parsing` - Stacks are used in compilers and interpreters to evaluate arithmetic
   expressions and parse the syntax of programming languages. They help manage the order of operations and handle nested
   expressions.
3. `Undo/Redo Operations` - Many applications, like text editors and graphic design software, use stacks to manage undo
   and redo operations. Each action performed is pushed onto the undo stack, and when an undo operation is executed, the
   last action is popped from the stack and reversed. For redo, a separate stack can be maintained.
4. `Backtracking` - Stacks are useful in backtracking problems, like navigating a maze, solving a Sudoku puzzle, or
   implementing search algorithms. They can store the current state or path so that it's easy to backtrack to the
   previous state when needed.
5. `Balancing Symbols` - Stacks can be used to check if the opening and closing symbols (e.g., parentheses, brackets,
   and braces) in an expression are balanced. The algorithm pushes opening symbols onto the stack and pops them when a
   matching closing symbol is encountered.