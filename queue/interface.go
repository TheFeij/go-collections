package stack

// Queue defines the interface for a generic queue data structure.
type Queue[T any] interface {
	// Enqueue adds an element to the end of the queue.
	Enqueue(T)

	// Dequeue removes and returns the element at the front of the queue.
	//
	// It returns the front element and ok = true if the queue is not empty,
	// otherwise it returns the zero value of type T and ok = false.
	Dequeue() (t T, ok bool)

	// Peek returns the element at the front of the queue without removing it.
	//
	// It returns the front element and ok = true if the queue is not empty,
	// otherwise it returns the zero value of type T and ok = false.
	Peek() (t T, ok bool)

	// Size returns the number of elements in the queue.
	Size() int
}
