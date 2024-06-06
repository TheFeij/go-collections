package stack

// Stack defines the interface for a generic stack data structure.
type Stack[T any] interface {
	// Push adds an element to the top of the stack.
	Push(T)

	// Pop removes and returns the top element from the stack.
	//
	// It returns the popped element and ok = true if the stack is not empty,
	// otherwise it returns the zero value of type T and ok = false.
	Pop() (t T, ok bool)

	// Peek returns the top element from the stack without removing it.
	//
	// It returns the top element and ok = true if the stack is not empty,
	// otherwise it returns the zero value of type T and ok = false.
	Peek() (t T, ok bool)

	// Size returns the number of elements in the stack.
	Size() int
}
