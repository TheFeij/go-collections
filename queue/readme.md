# Queue

The `queue` subpackage provides a generic implementation of queue data structure using a singly linked list.

## Overview

A queue is a first-in, first-out (FIFO) data structure that allows for efficient insertion
and removal of elements. 

The Queue interface defines the following methods:

| Method                     | Explanation                                                        |
|----------------------------|--------------------------------------------------------------------|
| `Enqueue(T)`               | adds an element to the end of the queue.                           |
| `Peek() (t T, ok bool)`    | returns the element at the front of the queue without removing it. |
| `Dequeue() (t T, ok bool)` | removes and returns the element at the front of the queue.         |
| `Size() int`               | Returns the number of elements in the queue.                       |


## Usage

Here is an example of how to use Queue

```go
package main

import (
	"fmt"
	"github.com/TheFeij/go-collections/queue"
)

func main() {
	// Create a new queue
	q := queue.NewQueue[int]()

	// Enqueue elements to the queue
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	// Dequeue an element from the queue
	element, ok := q.Dequeue()
	if ok {
		fmt.Println("Dequeued element:", element)
	}

	// Peek at the front element in the queue
	element, ok = q.Peek()
	if ok {
		fmt.Println("front element:", element)
	}

	// Get the size of the queue
	size := q.Size()
	fmt.Println("Size of the queue:", size)

}
```

## Time Complexity of the Queue Implementation

| Method                     | Time Complexity |
|----------------------------|-----------------|
| `Enqueue(T)`               | O(1)            |
| `Peek() (t T, ok bool)`    | O(1)            |
| `Dequeue() (t T, ok bool)` | O(1)            |
| `Size() int`               | O(1)            |


## Implementation Details

The queue is implemented using a singly linked list from the `linkedlist` package.
The queue struct contains a linked list as its underlying data structure. 
Each queue operation delegates to the corresponding linked list operation, 
ensuring efficient and correct behavior.