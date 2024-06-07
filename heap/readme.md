# Heap

The `heap` subpackage provides a generic implementation of a binary heap data structure.

## Overview

A binary heap is a complete binary tree where each node satisfies the heap property.
In a min-heap, for every node `i` other than the root, `heap[i] <= heap[parent(i)]`.
In a max-heap, `heap[i] >= heap[parent(i)]`.

The Heap interface defines the following methods:

| Method                      | Explanation                                                                                                   |
|-----------------------------|---------------------------------------------------------------------------------------------------------------|
| `Insert(T)`                 | Adds an element to the heap.                                                                                  |
| `Extract() (t T, ok bool)`  | Removes and returns the root element from the heap (min element in a min-heap and max in a max heap).         | 
| `Peek() (t T, ok bool)`     | Returns the root element from the heap without removing it (min element in a min-heap and max in a max heap). |
| `Size() int`                | Returns the number of elements in the heap.                                                                   |


## Usage

Here is an example of how to use Heap:

```go
package main

import (
	"fmt"
	"github.com/TheFeij/go-collections/heap"
)

func main() {
	// Create a new min heap
	h := heap.NewHeap[int](func(t1, t2 int) bool {
	    return t1 < t2

		// in order to get a max heap:  return t1 > t2
	})

	// Insert elements into the min heap
	h.Insert(5)
	h.Insert(10)
	h.Insert(3)

	// Extract the root element from the heap (here in a min heap it will be the minimum value)
	minimum, ok := h.Extract()
	if ok {
		fmt.Println("Minimum element:", minimum)
	}

	// Peek at the root element in the heap
	minimum, ok = h.Peek()
	if ok {
		fmt.Println("Maximum element:", minimum)
	}

	// Get the size of the heap
	size := h.Size()
	fmt.Println("Size of the heap:", size)
}
```

## Time Complexity of the Heap Implementation

| Function                                                            | Time Complexity                                                               |
|---------------------------------------------------------------------|-------------------------------------------------------------------------------|
| `NewHeap[T any](comparator func(t1, t2 T) bool, data ...T) Heap[T]` | O(n)  (if initial data is provided)<br/>O(1) (if no initial data is provided) |

| Method                                                              | Time Complexity                                                               |
|---------------------------------------------------------------------|-------------------------------------------------------------------------------|
| `Insert(T)`                                                         | O(log(n))                                                                     |
| `Extract() (t T, ok bool)`                                          | O(log(n))                                                                     | 
| `Peek() (t T, ok bool)`                                             | O(1)                                                                          |
| `Size() int`                                                        | O(1)                                                                          |



## Implementation Details
The heap is implemented using a binary tree. The heap struct contains an array/slice as
its underlying data structure. Each heap operation maintains the heap property, 
ensuring efficient and correct behavior.