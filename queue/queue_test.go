package stack

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue[any]()

	require.NotNil(t, q)
	require.NotEmpty(t, q)
	require.Equal(t, 0, q.Size())

	// no value to peep
	value, ok := q.Peek()
	require.Zero(t, value)
	require.False(t, ok)

	// no value to dequeue
	value, ok = q.Dequeue()
	require.Zero(t, value)
	require.False(t, ok)
}

func TestQueue_Enqueue_Dequeue(t *testing.T) {
	q := NewQueue[any]()

	const numberOfElements = 5
	for i := 0; i < numberOfElements; i++ {
		q.Enqueue(i)
		require.Equal(t, i+1, q.Size())
	}

	require.Equal(t, numberOfElements, q.Size())

	for i := 0; i < numberOfElements; i++ {
		value, ok := q.Dequeue()
		require.True(t, ok)
		require.Equal(t, i, value)
		require.Equal(t, numberOfElements-i-1, q.Size())
	}

	value, ok := q.Dequeue()
	require.False(t, ok)
	require.Zero(t, value)
	require.Equal(t, 0, q.Size())
}

func TestQueue_Peek(t *testing.T) {
	q := NewQueue[any]()

	value, ok := q.Peek()
	require.False(t, ok)
	require.Zero(t, value)

	const numberOfElements = 5
	for i := 0; i < numberOfElements; i++ {
		q.Enqueue(i)

		value, ok = q.Peek()
		require.True(t, ok)
		require.Equal(t, 0, value)
	}

	for i := 0; i < numberOfElements; i++ {
		value, ok = q.Peek()
		require.True(t, ok)
		require.Equal(t, i, value)

		_, ok := q.Dequeue()
		require.True(t, ok)
	}

	value, ok = q.Peek()
	require.False(t, ok)
	require.Zero(t, value)
}
