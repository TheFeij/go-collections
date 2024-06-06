package linked_list

// node represents a node in linked list
type node[T any] struct {
	value    T
	next     *node[T]
	previous *node[T]
}

// doublyLinkedList is an implementation of the LinkedList interface
type doublyLinkedList[T any] struct {
	first *node[T]
	last  *node[T]
	size  int
}

// InsertToIndex inserts input value to the given index
func (l *doublyLinkedList[T]) InsertToIndex(t T, index int) (ok bool) {
	// check if the index is valid
	if index < 0 || index >= l.size {
		return
	}

	// if index is 0, then call AddFirst
	if index == 0 {
		l.AddFirst(t)
		return true
	}

	// after the above conditions are passed, we are sure that the index we are
	// about to add has both next and previous elements

	// get the element at the input index
	currNodeAtIndex := l.get(index)

	newNodeAtIndex := &node[T]{
		value:    t,
		previous: currNodeAtIndex.previous,
		next:     currNodeAtIndex,
	}

	currNodeAtIndex.previous.next = newNodeAtIndex
	currNodeAtIndex.previous = newNodeAtIndex

	// update the size
	l.size += 1

	return true
}

// DeleteIndex deletes value at the given index
func (l *doublyLinkedList[T]) DeleteIndex(index int) (ok bool) {
	// check if the index is valid
	if index < 0 || index >= l.size {
		return
	}

	// if index is 0, then call DeleteFirst
	if index == 0 {
		l.deleteFirst()
		return true
	}

	// if index is the last index, then call DeleteLast
	if index == l.size-1 {
		l.deleteLast()
		return true
	}

	// after the above conditions, we are sure that the index we are
	// about to delete has both next and previous elements

	currNodeAtIndex := l.get(index)

	currNodeAtIndex.previous.next = currNodeAtIndex.next
	currNodeAtIndex.next.previous = currNodeAtIndex.previous

	// update the size
	l.size -= 1

	// dereference the node to help with garbage collection
	currNodeAtIndex.next = nil
	currNodeAtIndex.previous = nil
	currNodeAtIndex = nil

	return true
}

// Get returns the value of the element at the input index
//
// worst case O(n/2)
func (l *doublyLinkedList[T]) Get(index int) (t T, ok bool) {
	if index >= l.size || index < 0 {
		return
	}

	// get the node at the input index
	node := l.get(index)

	return node.value, true
}

// get returns the node at the input index
//
// does not check for the validity of the index, should be checked at the caller
// worst case O(n/2)
func (l *doublyLinkedList[T]) get(index int) *node[T] {
	distanceToLast := l.size - index
	distanceToFirst := index

	var currNode *node[T]
	if distanceToFirst < distanceToLast {
		currNode = l.first
		for i := 0; i < distanceToFirst; i++ {
			currNode = currNode.next
		}
	} else {
		currNode = l.last
		for i := 0; i < distanceToLast-1; i++ {
			currNode = currNode.previous
		}
	}

	return currNode
}

// Size returns the current size of the linked list
func (l *doublyLinkedList[T]) Size() int {
	return l.size
}

// Clear removes all elements from the linked list
func (l *doublyLinkedList[T]) Clear() {
	// Traverse the list and clear references to help garbage collection
	current := l.first
	for current != nil {
		next := current.next
		current.next = nil
		current.previous = nil
		current = next
	}

	l.first = nil
	l.last = nil
	l.size = 0
}

// GetFirst returns the first element of the linked list
//
// ok = false means the linked list is empty and there is no first element
func (l *doublyLinkedList[T]) GetFirst() (t T, ok bool) {
	if l.first == nil {
		return
	}

	return l.first.value, true
}

// GetLast returns the last element of the linked list
//
// ok = false means the linked list is empty and there is no last element
func (l *doublyLinkedList[T]) GetLast() (t T, ok bool) {
	if l.last == nil {
		return
	}

	return l.last.value, true
}

// Add adds input value to the end of the linked list
func (l *doublyLinkedList[T]) Add(t T) {
	l.AddLast(t)
}

// AddFirst adds input value to the start of the linked list
func (l *doublyLinkedList[T]) AddFirst(t T) {
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
func (l *doublyLinkedList[T]) AddLast(t T) {
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
func (l *doublyLinkedList[T]) DeleteFirst() (ok bool) {
	// return if the linked list is empty
	if l.first == nil {
		return
	}

	// delete the first element
	l.deleteFirst()

	return true
}

// deleteFirst deletes first element of the linked list
//
// should not be called on an empy linked list
func (l *doublyLinkedList[T]) deleteFirst() {
	// clear the linked list if it has only one node
	if l.first == l.last {
		l.Clear()
		return
	}

	// after the above conditions, we would have at least 2 nodes

	// storing the reference to the element to be deleted
	first := l.first

	l.first = l.first.next
	l.first.previous = nil

	// clearing references to help garbage collection
	first.next = nil
	first.previous = nil
	first = nil

	l.size -= 1
}

// DeleteLast deletes last element of the linked list
func (l *doublyLinkedList[T]) DeleteLast() (ok bool) {
	// return if the linked list is empty
	if l.last == nil {
		return
	}

	// delete the last element
	l.deleteLast()

	return true
}

// deleteLast deletes last element of the linked list
//
// should not be called on an empty linked list
func (l *doublyLinkedList[T]) deleteLast() {
	// clear the linked list if it has only one node
	if l.last == l.first {
		l.Clear()
		return
	}

	// here we would have at least 2 nodes

	// storing the reference to the element to be deleted
	last := l.last

	l.last = l.last.previous
	l.last.next = nil

	// clearing references to help garbage collection
	last.next = nil
	last.previous = nil
	last = nil

	l.size -= 1
}

// NewDoublyLinkedList returns a new linked list
func NewDoublyLinkedList[T any]() LinkedList[T] {
	return &doublyLinkedList[T]{
		first: nil,
		last:  nil,
		size:  0,
	}
}
