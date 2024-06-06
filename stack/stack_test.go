package stack

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewStack(t *testing.T) {
	s := NewStack[any]()

	require.NotNil(t, s)
	require.NotEmpty(t, s)
	require.Equal(t, 0, s.Size())

	// no value to peep
	value, ok := s.Peek()
	require.Zero(t, value)
	require.False(t, ok)

	// no value to pop
	value, ok = s.Pop()
	require.Zero(t, value)
	require.False(t, ok)
}

func TestStack_Push_Pop(t *testing.T) {
	s := NewStack[any]()

	const numberOfElements = 5
	for i := 0; i < numberOfElements; i++ {
		s.Push(i)
		require.Equal(t, i+1, s.Size())
	}

	require.Equal(t, numberOfElements, s.Size())

	for i := 0; i < numberOfElements; i++ {
		value, ok := s.Pop()
		require.True(t, ok)
		require.Equal(t, numberOfElements-i-1, value)
		require.Equal(t, numberOfElements-i-1, s.Size())
	}

	value, ok := s.Pop()
	require.False(t, ok)
	require.Zero(t, value)
	require.Equal(t, 0, s.Size())
}

func TestStack_Peek(t *testing.T) {
	s := NewStack[any]()

	value, ok := s.Peek()
	require.False(t, ok)
	require.Zero(t, value)

	const numberOfElements = 5
	for i := 0; i < numberOfElements; i++ {
		s.Push(i)

		value, ok = s.Peek()
		require.True(t, ok)
		require.Equal(t, i, value)
	}
}
