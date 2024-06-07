package heap

// Heap defines the interface for a generic heap data structure.
type Heap[T any] interface {
	// Insert adds an element to the heap
	Insert(T)

	// Extract removes and returns the root element from the heap
	//
	// returns ok = false if heap is empty
	Extract() (t T, ok bool)

	// Peek returns the root element without removing it
	//
	// returns ok = false if heap is empty
	Peek() (t T, ok bool)

	// Size returns the number of elements in the heap
	Size() int
}
