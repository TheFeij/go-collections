package linkedlist

// singlyNode represents a node in a singly linked list
type singlyNode[T any] struct {
	value T
	next  *singlyNode[T]
}

// singlyLinkedList is an implementation of the LinkedList interface
type singlyLinkedList[T any] struct {
	first *singlyNode[T]
	last  *singlyNode[T]
	size  int
}

// InsertToIndex inserts input value to the given index
func (l *singlyLinkedList[T]) InsertToIndex(t T, index int) (ok bool) {
	// check if the index is valid
	if index < 0 || index >= l.size {
		return
	}

	if index == 0 {
		l.AddFirst(t)
	} else {
		nodeBeforeIndex := l.get(index - 1)

		nodeAtIndex := nodeBeforeIndex.next

		newNode := &singlyNode[T]{
			value: t,
			next:  nodeAtIndex,
		}

		nodeBeforeIndex.next = newNode

		// update the size
		l.size += 1
	}

	return true
}

// DeleteIndex deletes value at the given index
func (l *singlyLinkedList[T]) DeleteIndex(index int) (ok bool) {
	// check if the index is valid
	if index < 0 || index >= l.size {
		return
	}

	// if index is 0, then call DeleteFirst
	if index == 0 {
		l.deleteFirst()
		return true
	}

	// after the above conditions, we are sure that the index we are
	// about to delete has a previous node

	nodeBeforeIndex := l.get(index - 1)

	nodeAtIndex := nodeBeforeIndex.next

	nodeBeforeIndex.next = nodeAtIndex.next

	if nodeBeforeIndex.next == nil {
		l.last = nodeBeforeIndex
	}

	// clear references to help garbage collection
	nodeAtIndex.next = nil
	nodeAtIndex = nil

	// update the size
	l.size -= 1

	return true
}

// Get returns the value of the element at the input index
//
// worst case O(n)
func (l *singlyLinkedList[T]) Get(index int) (t T, ok bool) {
	if index >= l.size || index < 0 {
		return
	}

	if index == 0 {
		return l.first.value, true
	} else if index == l.size-1 {
		return l.last.value, true
	} else {
		return l.get(index).value, true
	}
}

// get returns the node at the input index
//
// does not check for the validity of the index, should be checked at the caller
// worst case O(n)
func (l *singlyLinkedList[T]) get(index int) *singlyNode[T] {
	var currNode *singlyNode[T]

	currNode = l.first
	for i := 0; i < index; i++ {
		currNode = currNode.next
	}

	return currNode
}

// Size returns the current size of the linked list
func (l *singlyLinkedList[T]) Size() int {
	return l.size
}

// Clear removes all elements from the linked list
func (l *singlyLinkedList[T]) Clear() {
	// Traverse the list and clear references to help garbage collection
	current := l.first
	for current != nil {
		next := current.next
		current.next = nil
		current = next
	}

	l.first = nil
	l.last = nil
	l.size = 0
}

// GetFirst returns the first element of the linked list
//
// ok = false means the linked list is empty and there is no first element
func (l *singlyLinkedList[T]) GetFirst() (t T, ok bool) {
	if l.first == nil {
		return
	}

	return l.first.value, true
}

// GetLast returns the last element of the linked list
//
// ok = false means the linked list is empty and there is no last element
func (l *singlyLinkedList[T]) GetLast() (t T, ok bool) {
	if l.size == 0 {
		return
	}

	return l.last.value, true
}

// Add adds input value to the end of the linked list
func (l *singlyLinkedList[T]) Add(t T) {
	l.AddLast(t)
}

// AddFirst adds input value to the start of the linked list
func (l *singlyLinkedList[T]) AddFirst(t T) {
	newNode := &singlyNode[T]{
		value: t,
		next:  l.first,
	}

	if l.size == 0 {
		l.last = newNode
	}

	l.first = newNode
	l.size += 1
}

// AddLast adds input value to the end of the linked list
func (l *singlyLinkedList[T]) AddLast(t T) {
	newNode := &singlyNode[T]{
		value: t,
		next:  nil,
	}

	if l.size == 0 {
		l.first = newNode
	} else {
		l.last.next = newNode
	}

	l.last = newNode
	l.size += 1
}

// DeleteFirst deletes first element of the linked list
func (l *singlyLinkedList[T]) DeleteFirst() (ok bool) {
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
func (l *singlyLinkedList[T]) deleteFirst() {
	// if the list has only one node
	if l.size == 1 {
		l.first = nil
		l.last = nil
		l.size = 0
		return
	}

	// after the above conditions, we would have at least 2 nodes

	// storing the reference to the element to be deleted
	first := l.first

	l.first = l.first.next

	// clearing references to help garbage collection
	first.next = nil
	first = nil

	l.size -= 1
}

// DeleteLast deletes last element of the linked list
func (l *singlyLinkedList[T]) DeleteLast() (ok bool) {
	// return if the linked list is empty
	if l.first == nil {
		return
	}

	// delete the last element
	l.deleteLast()

	return true
}

// deleteLast deletes last element of the linked list
//
// should not be called on an empty linked list
func (l *singlyLinkedList[T]) deleteLast() {
	// if the list has only one node
	if l.size == 1 {
		l.first = nil
		l.last = nil
		l.size = 0
		return
	}

	// here we would have at least 2 nodes

	// getting the second last node
	secondLast := l.get(l.size - 2)

	// make the second last node the last node
	secondLast.next = nil
	l.last = secondLast

	l.size -= 1
}

// NewSinglyLinkedList returns a new singly linked list
func NewSinglyLinkedList[T any]() LinkedList[T] {
	return &singlyLinkedList[T]{
		first: nil,
		last:  nil,
		size:  0,
	}
}
