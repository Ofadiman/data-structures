# Data structures

A data structure is a way of organizing and storing data in a computer so that it can be accessed and manipulated
efficiently. A data structure can be thought of as a container that holds data in a particular arrangement, with
specific operations that can be performed on the data it contains. There are many types of data structures, each with
its own advantages and disadvantages. Some of the most common data structures
include `arrays`, `linked lists`, `stacks`, `queues`, `trees`, `graphs`, and `hash tables`.

Implemented data structures:

- [Binary Search Tree](binary_search_tree.go)
- [Singly Linked List](singly_linked_list.go)
- [Stack](stack.go)
- [Tree](tree.go)

# Binary Search Tree

A Binary Search Tree (BST) is a binary tree data structure in which each node has at most two children, arranged in such
a way that the value of the node to the left is less than or equal to the value of the parent node, and the value of the
node to the right is greater than or equal to the value of the parent node. This property is maintained throughout the
entire tree, making it efficient for searching, inserting, and deleting values. The following methods are supported:

- `Insert`: Insert a new node with a given value into the tree.

Here are some real-world use cases of Binary Search Trees:

- `Databases`: BSTs are commonly used in databases for indexing purposes, making it efficient to search, insert, and
  delete records based on a certain key value.
- `Symbol Tables`: In compilers, BSTs can be used to store and manage the symbol table, which keeps track of variable
  and function names during compilation.
- `Autocomplete functionality`: BSTs can be used to efficiently store and search a dictionary of words for implementing
  autocomplete functionality in text editors or search engines.
- `Balanced Partitions`: BSTs can be used to balance partitions in distributed systems, such as partitioning users based
  on their IDs, ensuring an evenly distributed load across the system.
- `Game Trees`: In game theory and artificial intelligence, BSTs can be used to represent and explore possible game
  states in a minimax algorithm, allowing for efficient pruning of non-optimal branches.
- `File Systems`: File systems can use BSTs to store and manage directory structures and file metadata, making it
  efficient to search, insert, and delete files and directories.

# Singly Linked List

A Singly Linked List is a linear data structure in which each element, called a node, contains a reference to the next
node in the sequence. The first node in the list is called the head, and the last node's reference points to nil (or
null), marking the end of the list. The following methods are supported:

- `Add`: Add a new node with a given value to the end of the list.
- `Prepend`: Add a new node to the beginning of the list.
- `Insert`: Insert a new node at a specific position in the list.
- `Remove`: Remove a node with a specific value from the list.
- `RemoveAt`: Remove a node at a specific position in the list.
- `Find`: Find a node with a specific value in the list.
- `IndexOf`: Find the index of the first occurrence of a node with a specific value.
- `Serialize`: Convert a singly linked list into a JSON-formatted string representing an array of nodes with their
  respective values.
- `Deserialize`: Reconstruct a singly linked list from a JSON-formatted string representing an array of nodes, ensuring
  compatibility with the Serialize method.

Here are some real-world use cases of Singly Linked Lists:

1. `Undo/Redo functionality`: Singly Linked Lists can be used to implement undo and redo functionality in applications
   like text editors, image editors, or web browsers. Each action performed by the user can be stored as a node in the
   list, and traversing the list in reverse order can undo the actions.
2. `Dynamic memory allocation`: Singly Linked Lists can be used to manage memory allocation in operating systems. Free
   blocks of memory can be stored as nodes in a linked list, making it easier to allocate and deallocate memory as
   needed.
3. `Implementing data structures`: Singly Linked Lists can be used as a building block for more complex data structures
   like Stacks, Queues, or Symbol Tables.
4. `Representing sparse data`: In situations where data is sparse, Singly Linked Lists can be used to store non-zero
   elements efficiently.
5. `Implementing navigation systems`: In navigation systems, waypoints can be stored in a Singly Linked List to
   represent a route or path. The list can be traversed from the head to the tail to follow the route.
6. `Music playlist`: A music player can use a Singly Linked List to represent a playlist, where each node represents a
   song. This allows for easy insertion and deletion of songs in the playlist.

# Stack

A stack is a linear data structure that follows the Last In, First Out (LIFO) principle, meaning that the last element
added to the stack is the first one to be removed. It can be thought of as a collection of items, where new items can be
added, and existing items can be removed from the top of the stack. It supports the following operations:

- `Empty`: Remove all elements from the stack.
- `IsEmpty`: Check if the stack is empty, returning true if empty and false otherwise.
- `Peek`: Return the top element of the stack without removing it, or return an error if the stack is empty.
- `Pop`: Remove and return the top element of the stack, or return an error if the stack is empty.
- `Push`: Add a new element to the top of the stack.
- `Size`: Return the current number of elements in the stack.
- `Serialize`: Convert a stack into a JSON-formatted string representing an array of its items.
- `Deserialize`: Reconstruct a stack of integers from a JSON-formatted string representing an array of items, ensuring
  compatibility with the Serialize method.

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
searching, insertion, and deletion operations. Tree supports the following operations:

- `FindNodeByID`: Find a node in the tree by its unique ID, returning an error if the node is not found.
- `Insert`: Insert a new child node under a parent node with the specified parent ID.
- `Delete`: Remove a node with a specified ID from the tree, including detaching it from its parent and updating the
  parent's children list.
- `ForEachNodeDepthFirst`: Perform a depth-first traversal of the tree, calling a provided callback function for each
  visited node.
- `ForEachNodeBreadthFirst`: Perform a breadth-first traversal of the tree, calling a provided callback function for
  each visited node.
- `Serialize`: Convert the tree into a JSON-formatted string, including all nodes and their relationships.
- `Deserialize`: Reconstruct a tree from a JSON-formatted string, ensuring compatibility with the Serialize method and
  preserving node relationships.
- `Height`: Calculate and return the height of the tree, which is the maximum number of edges in the longest path from
  the root to a leaf.

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
