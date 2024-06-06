package stack

import "github.com/TheFeij/go-collections/linkedlist"

// stack is a struct representing a generic stack data structure.
type stack[T any] struct {
	// list is the underlying linkedlist used to implement the stack.
	list linkedlist.LinkedList[T]
}

// Size returns the number of elements in the stack.
func (s *stack[T]) Size() int {
	return s.list.Size()
}

// Push adds an element to the top of the stack.
func (s *stack[T]) Push(t T) {
	s.list.AddFirst(t)
}

// Pop removes and returns the top element from the stack.
//
// It returns the popped element and ok = true if the stack is not empty,
// otherwise it returns the zero value of type T and ok = false.
func (s *stack[T]) Pop() (t T, ok bool) {
	t, ok = s.list.GetFirst()
	if !ok {
		return
	}

	s.list.DeleteFirst()
	return
}

// Peek returns the top element from the stack without removing it.
//
// It returns the top element and ok = true if the stack is not empty,
// otherwise it returns the zero value of type T and ok = false.
func (s *stack[T]) Peek() (t T, ok bool) {
	t, ok = s.list.GetFirst()
	return
}

// NewStack creates and returns a new stack.
// It initializes the underlying linked list with a new singly linked list.
func NewStack[T any]() Stack[T] {
	return &stack[T]{
		list: linkedlist.NewSinglyLinkedList[T](),
	}
}
