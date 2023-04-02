# Data structures

A data structure is a way of organizing and storing data in a computer so that it can be accessed and manipulated
efficiently. A data structure can be thought of as a container that holds data in a particular arrangement, with
specific operations that can be performed on the data it contains. There are many types of data structures, each with
its own advantages and disadvantages. Some of the most common data structures
include `arrays`, `linked lists`, `stacks`, `queues`, `trees`, `graphs`, and `hash tables`.

Implemented data structures:

- [Singly Linked List](singly_linked_list.go)
- [Binary Tree](binary_tree.go)
- [Stack](stack.go)
- [Tree](tree.go)

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

**Stack data structure use cases:**

1. `Function Call Stack`: In programming languages, a stack is used to manage function calls and their execution
   context. When a function is called, its context is pushed onto the call stack, and when it returns, its context is
   popped from the stack.
2. `Expression Evaluation and Syntax Parsing`: Stacks are used in compilers and interpreters to evaluate arithmetic
   expressions and parse the syntax of programming languages. They help manage the order of operations and handle nested
   expressions.
3. `Undo/Redo Operations`: Many applications, like text editors and graphic design software, use stacks to manage undo
   and redo operations. Each action performed is pushed onto the undo stack, and when an undo operation is executed, the
   last action is popped from the stack and reversed. For redo, a separate stack can be maintained.
4. `Backtracking`: Stacks are useful in backtracking problems, like navigating a maze, solving a Sudoku puzzle, or
   implementing search algorithms. They can store the current state or path so that it's easy to backtrack to the
   previous state when needed.
5. `Balancing Symbols`: Stacks can be used to check if the opening and closing symbols (e.g., parentheses, brackets, and
   braces) in an expression are balanced. The algorithm pushes opening symbols onto the stack and pops them when a
   matching closing symbol is encountered.

# Tree

A tree data structure is a non-linear hierarchical data structure that consists of nodes connected by edges. Each node
in the tree has a parent node and zero or more child nodes. The top node of the tree is called the root, and the nodes
with no children are called leaves. Trees are used to represent hierarchical relationships and to enable efficient
searching, insertion, and deletion operations.

**Some important properties of trees are:**

- A tree with N nodes has N-1 edges.
- A tree is a connected graph without cycles.
- Each node in the tree can have at most one parent.

**Real-world use-cases of tree data structures:**

1. `File systems`: Tree structures are used to represent the organization of files and directories in a computer's file
   system. Each directory is a node, and the files and subdirectories within it are its children.
2. `XML and HTML`: Tree structures are used to represent the hierarchical organization of elements in XML and HTML
   documents.
3. `Organizational charts`: Trees can be used to represent the hierarchical structure of an organization, with nodes
   representing employees and their relationships to managers and subordinates.
4. `Decision trees`: In machine learning, decision trees are used to represent a series of decisions that lead to a
   final
   classification or prediction.
5. `Game trees`: In artificial intelligence, tree structures are used to represent game states and possible moves in
   game
   playing algorithms, such as chess and tic-tac-toe.
6. `Abstract syntax trees (AST)`: In compilers, tree structures are used to represent the syntactic structure of the
   source code, enabling semantic analysis and code generation.
7. `Routing algorithms`: In computer networks, tree structures are used in routing algorithms to find the shortest path
   between two nodes.
8. `Binary search trees (BST)`: BSTs are used to store and efficiently search, insert, and delete elements in a sorted
   manner.
9. `Trie (prefix tree)`: Tries are used for efficient string matching and retrieval, commonly used in autocomplete and
   dictionary implementations.
10. `Taxonomies`: Trees are used to represent the hierarchical classification of objects, such as biological species or
    product categories in e-commerce websites.

# Running tests

Tests can be run using Golang built-in `go test` cli:

- `go test ./...` - Run all tests.
- `go test -run TestNewSinglyLinkedList ./...` - Run only tests with name `TestNewSinglyLinkedList`.
