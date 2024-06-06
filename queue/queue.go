package stack

import "github.com/TheFeij/go-collections/linkedlist"

// queue is a struct representing a generic queue data structure.
type queue[T any] struct {
	// list is the underlying linkedlist used to implement the queue.
	list linkedlist.LinkedList[T]
}

// Size returns the number of elements in the queue.
func (q *queue[T]) Size() int {
	return q.list.Size()
}

// Enqueue adds an element to the end of the queue.
func (q *queue[T]) Enqueue(t T) {
	q.list.AddLast(t)
}

// Dequeue removes and returns the element at the front of the queue.
//
// It returns the front element and ok = true if the queue is not empty,
// otherwise it returns the zero value of type T and ok = false.
func (q *queue[T]) Dequeue() (t T, ok bool) {
	t, ok = q.list.GetFirst()
	if !ok {
		return
	}

	q.list.DeleteFirst()
	return
}

// Peek returns the top element from the queue without removing it.
//
// It returns the top element and ok = true if the queue is not empty,
// otherwise it returns the zero value of type T and ok = false.
func (q *queue[T]) Peek() (t T, ok bool) {
	t, ok = q.list.GetFirst()
	return
}

// NewQueue creates and returns a new queue.
// It initializes the underlying linked list with a new singly linked list.
func NewQueue[T any]() Queue[T] {
	return &queue[T]{
		list: linkedlist.NewSinglyLinkedList[T](),
	}
}
