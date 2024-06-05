package collections

// LinkedList represents a linked list
type LinkedList[T any] interface {
	Add(T)
	AddFirst(T)
	AddLast(T)
	GetFirst() (T, bool)
	GetLast() (T, bool)
	Clear()
	DeleteFirst()
	DeleteLast()
	Size() int
}

// node represents a node in linked list
type node[T any] struct {
	value    T
	next     *node[T]
	previous *node[T]
}

// linkedList is the implementation of the LinkedList interface
type linkedList[T any] struct {
	first *node[T]
	last  *node[T]
	size  int
}

// Size returns the current size of the linked list
func (l *linkedList[T]) Size() int {
	return l.size
}

// Clear removes all elements from the linked list
func (l *linkedList[T]) Clear() {
	l.first = nil
	l.last = nil
	l.size = 0
}

// GetFirst returns the first element of the linked list
//
// ok = false means the linked list is empty and there is no first element
func (l *linkedList[T]) GetFirst() (t T, ok bool) {
	if l.first == nil {
		return
	}

	return l.first.value, true
}

// GetLast returns the last element of the linked list
//
// ok = false means the linked list is empty and there is no last element
func (l *linkedList[T]) GetLast() (t T, ok bool) {
	if l.last == nil {
		return
	}

	return l.last.value, true
}

// Add adds input value to the end of the linked list
func (l *linkedList[T]) Add(t T) {
	l.AddLast(t)
}

// AddFirst adds input value to the start of the linked list
func (l *linkedList[T]) AddFirst(t T) {
	newNode := &node[T]{
		value:    t,
		next:     nil,
		previous: nil,
	}

	if l.size == 0 {
		l.last = newNode
	} else {
		newNode.next = l.first
		l.first.previous = newNode
	}

	l.first = newNode
	l.size += 1
}

// AddLast adds input value to the end of the linked list
func (l *linkedList[T]) AddLast(t T) {
	newNode := &node[T]{
		value:    t,
		next:     nil,
		previous: nil,
	}

	if l.size == 0 {
		l.first = newNode
	} else {
		newNode.previous = l.last
		l.last.next = newNode
	}

	l.last = newNode
	l.size += 1
}

// DeleteFirst deletes first element of the linked list
func (l *linkedList[T]) DeleteFirst() {
	// return if the linked list is empty
	if l.first == nil {
		return
	}

	// clear the linked list if it has only one node
	if l.first == l.last {
		l.Clear()
		return
	}

	// here we would have at least 2 nodes

	l.first = l.first.next
	l.first.previous = nil
	l.size -= 1
}

// DeleteLast deletes last element of the linked list
func (l *linkedList[T]) DeleteLast() {
	// return if the linked list is empty
	if l.last == nil {
		return
	}

	// clear the linked list if it has only one node
	if l.last == l.first {
		l.Clear()
		return
	}

	// here we would have at least 2 nodes

	l.last = l.last.previous
	l.last.next = nil
	l.size -= 1
}

// NewLinkedList returns a new linked list
func NewLinkedList[T any]() LinkedList[T] {
	return &linkedList[T]{
		first: nil,
		last:  nil,
		size:  0,
	}
}
