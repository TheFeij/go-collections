# Linked List

The `linkedlist` subpackage provides different implementations of the `LinkedList` interface.

### Methods

The `LinkedList` interface defines the following methods:

| Method                                    | Explanation                                                                               |
|-------------------------------------------|-------------------------------------------------------------------------------------------|
| `Add(T)`                                  | Adds an element to the end of the list.                                                   |
| `AddFirst(T)`                             | Adds an element to the start of the list.                                                 |
| `AddLast(T)`                              | Adds an element to the end of the list.                                                   |
| `GetFirst() (t T, ok bool)`               | Returns the first element of the list. Returns `false` if the list is empty.              |
| `GetLast() (t T, ok bool)`                | Returns the last element of the list. Returns `false` if the list is empty.               |
| `Clear()`                                 | Removes all elements from the list.                                                       |
| `DeleteFirst() (ok bool)`                 | Deletes the first element of the list. Returns `false` if the list is empty.              |
| `DeleteLast() (ok bool)`                  | Deletes the last element of the list. Returns `false` if the list is empty.               |
| `Size() int`                              | Returns the current size of the list.                                                     |
| `Get(index int) (t T, ok bool)`           | Returns the element at the specified index. Returns `false` if the index is out of range. |
| `InsertToIndex(t T, index int) (ok bool)` | Inserts an element at the specified index. Returns `false` if the index is out of range.  |
| `DeleteIndex(index int) (ok bool)`        | Deletes the element at the specified index. Returns `false` if the index is out of range. |

### Usage

Here's an example of how to use the doubly linked list implementation:

```go
package main

import (
	"fmt"
	"github.com/TheFeij/go-collections/linked-list"
)

func main() {
	// Use the New functions to get a desired implementation of the linked list
	// for example here we used NewDoublyLinkedList
	list := linkedlist.NewDoublyLinkedList[int]()

	// Add elements to the list
	list.Add(10)
	list.AddFirst(5)
	list.AddLast(15)

	// Get the first and last elements
	first, _ := list.GetFirst()
	last, _ := list.GetLast()

	fmt.Println("First element:", first)
	fmt.Println("Last element:", last)

	// Insert and delete elements at specific indices
	list.InsertToIndex(7, 1)
	list.DeleteIndex(2)

	// Get the size of the list
	fmt.Println("Size of the list:", list.Size())

	// Clear the list
	list.Clear()
}
```


## Implementations:

### Doubly Linked List

A doubly linked list is a type of linked list in which each node contains
a reference to both the next node and the previous node.

#### Time Complexities of the Doubly Linked List Implementation

| Method                                        | Time Complexity |
|-----------------------------------------------|-----------------|
| `Add(T)`                                      | O(1)            |
| `AddFirst(T)`                                 | O(1)            |
| `AddLast(T)`                                  | O(1)            |
| `GetFirst() (t T, ok bool)`                   | O(1)            |
| `GetLast() (t T, ok bool)`                    | O(1)            |
| `Clear()`                                     | O(n)            |
| `DeleteFirst() (ok bool)`                     | O(1)            |
| `DeleteLast() (ok bool)`                      | O(1)            |
| `Size() int`                                  | O(1)            |
| `Get(index int) (t T, ok bool)`               | O(n/2)          |
| `InsertToIndex(t T, index int) (ok bool)`     | O(n/2)          |
| `DeleteIndex(index int) (ok bool)`            | O(n/2)          |


