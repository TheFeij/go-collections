package heap

// heap represents a generic heap data structure
type heap[T any] struct {
	data       []T
	comparator func(t1, t2 T) bool
}

// heapifyDown performs heapify down operation from input root index
//
// O(log(n))
func (h *heap[T]) heapifyDown(root int) {
	largest := root
	leftChild := 2*root + 1
	rightChild := leftChild + 1

	size := len(h.data)
	if leftChild < size && h.comparator(h.data[leftChild], h.data[largest]) {
		largest = leftChild
	}
	if rightChild < size && h.comparator(h.data[rightChild], h.data[largest]) {
		largest = rightChild
	}

	if largest != root {
		h.data[root], h.data[largest] = h.data[largest], h.data[root]

		h.heapifyDown(largest)
	}
}

// buildHeap builds heap on h.data
//
// O(n)
func (h *heap[T]) buildHeap() {
	for i := len(h.data)/2 - 1; i >= 0; i-- {
		h.heapifyDown(i)
	}
}

// heapifyUp performs heapify up operation from input index
//
// O(log(n))
func (h *heap[T]) heapifyUp(index int) {
	parent := (index - 1) / 2

	if parent >= 0 && h.comparator(h.data[index], h.data[parent]) {
		h.data[parent], h.data[index] = h.data[index], h.data[parent]

		h.heapifyUp(parent)
	}
}

// Insert adds an element to the heap
func (h *heap[T]) Insert(t T) {
	h.data = append(h.data, t)

	h.heapifyUp(len(h.data) - 1)
}

// Extract removes and returns the root element from the heap
//
// returns ok = false if heap is empty
func (h *heap[T]) Extract() (t T, ok bool) {
	size := len(h.data)
	if size == 0 {
		return
	}

	if size == 1 {
		t = h.data[0]
		h.data = make([]T, 0)

		return t, true
	}

	lastIndex := size - 1
	t = h.data[0]
	h.data[0] = h.data[lastIndex]

	h.data = h.data[:lastIndex]
	h.heapifyDown(0)

	return t, true
}

// Peek Returns the root element without removing it
//
// returns ok = false if heap is empty
func (h *heap[T]) Peek() (t T, ok bool) {
	if len(h.data) == 0 {
		return
	}

	return h.data[0], true
}

// Size returns the number of elements in the heap
func (h *heap[T]) Size() int {
	return len(h.data)
}

// NewHeap creates a new heap with the given comparator and optional initial elements
//
// # Returns nil if comparator is nil
//
// The comparator function defines the heap property:
// - For a max-heap, comparator should return true if the first argument is greater than the second
// - For a min-heap, comparator should return true if the first argument is less than the second
//
// Example usage:
// - Max-Heap: NewHeap(func(a, b int) bool { return a > b }, 3, 1, 6, 5, 2, 4)
// - Min-Heap: NewHeap(func(a, b int) bool { return a < b }, 3, 1, 6, 5, 2, 4)
func NewHeap[T any](comparator func(t1, t2 T) bool, data ...T) (h Heap[T]) {
	if comparator == nil {
		return
	}

	heap := &heap[T]{
		data:       make([]T, len(data)),
		comparator: comparator,
	}
	// copy input variadic to prevent storing reference to the input slice in the case where input is an unpacked slice
	copy(heap.data, data)

	heap.buildHeap()

	return heap
}
