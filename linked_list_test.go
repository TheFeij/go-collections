package collections

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// tests NewLinkedList function
func TestNewLinkedList(t *testing.T) {
	list := NewLinkedList[any]()

	require.Empty(t, list)
	require.NotNil(t, list)

	require.Equal(t, 0, list.Size())

	first, ok := list.GetFirst()
	require.False(t, ok)
	require.Zero(t, first)

	var last any
	last, ok = list.GetLast()
	require.False(t, ok)
	require.Zero(t, last)
}

// tests Add and AddLast methods (since Add uses AddLast internally)
func TestLinkedList_Add_AddLast(t *testing.T) {
	list := NewLinkedList[any]()

	const listSize = 1000
	for i := 0; i < listSize; i++ {
		list.Add(i)

		first, ok := list.GetFirst()
		require.True(t, ok)
		require.Equal(t, 0, first)

		var last any
		last, ok = list.GetLast()
		require.True(t, ok)
		require.Equal(t, i, last)

		require.Equal(t, i+1, list.Size())
	}
}

// tests AddFirst method of the linked list
func TestLinkedList_AddFirst(t *testing.T) {
	list := NewLinkedList[any]()

	const listSize = 1000
	for i := 0; i < listSize; i++ {
		list.AddFirst(i)

		first, ok := list.GetFirst()
		require.True(t, ok)
		require.Equal(t, i, first)

		var last any
		last, ok = list.GetLast()
		require.True(t, ok)
		require.Equal(t, 0, last)

		require.Equal(t, i+1, list.Size())
	}
}

// tests Clear method of the linked list
func TestLinkedList_Clear(t *testing.T) {
	list := NewLinkedList[any]()

	const listSize = 1000
	for i := 0; i < listSize; i++ {
		list.Add(i)

		list.Clear()
		require.Equal(t, 0, list.Size())

		first, ok := list.GetFirst()
		require.False(t, ok)
		require.Zero(t, first)

		var last any
		last, ok = list.GetLast()
		require.False(t, ok)
		require.Zero(t, last)
	}
}

// tests DeleteFirst method of the linked list
func TestLinkedList_DeleteFirst(t *testing.T) {
	list := NewLinkedList[any]()

	list.DeleteFirst()

	const listSize = 1000
	for i := 0; i < listSize; i++ {
		list.Add(i)
	}

	for i := 0; i < listSize; i++ {
		first, ok := list.GetFirst()
		require.True(t, ok)
		require.Equal(t, i, first)
		require.Equal(t, listSize-i, list.Size())

		list.DeleteFirst()
	}
}

// tests DeleteLast method of the linked list
func TestLinkedList_DeleteLast(t *testing.T) {
	list := NewLinkedList[any]()

	list.DeleteLast()

	const listSize = 1000
	for i := 0; i < listSize; i++ {
		list.Add(i)
	}

	for i := 0; i < listSize; i++ {
		last, ok := list.GetLast()
		require.True(t, ok)
		require.Equal(t, listSize-i-1, last)
		require.Equal(t, listSize-i, list.Size())

		list.DeleteLast()
	}
}
