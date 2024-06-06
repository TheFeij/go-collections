package linked_list

// LinkedList represents a linked list
type LinkedList[T any] interface {
	// Add adds input value to the end of the linked list.
	Add(T)

	// AddFirst adds input value to the start of the linked list.
	AddFirst(T)

	// AddLast adds input value to the end of the linked list.
	AddLast(T)

	// GetFirst returns the first element of the linked list.
	//
	// ok = false means the linked list is empty and there is no first element.
	GetFirst() (t T, ok bool)

	// GetLast returns the last element of the linked list.
	//
	// ok = false means the linked list is empty and there is no last element.
	GetLast() (t T, ok bool)

	// Clear removes all elements from the linked list.
	Clear()

	// DeleteFirst deletes the first element of the linked list.
	//
	// ok = false means the linked list is empty and no element was deleted.
	DeleteFirst() (ok bool)

	// DeleteLast deletes the last element of the linked list.
	//
	// ok = false means the linked list is empty and no element was deleted.
	DeleteLast() (ok bool)

	// Size returns the current size of the linked list.
	Size() int

	// Get returns the value of the element at the input index.
	//
	// ok = false means the index is out of range.
	Get(index int) (t T, ok bool)

	// InsertToIndex inserts input value at the given index.
	//
	// ok = false means the index is out of range.
	InsertToIndex(t T, index int) (ok bool)

	// DeleteIndex deletes the value at the given index.
	//
	// ok = false means the index is out of range.
	DeleteIndex(index int) (ok bool)
}
