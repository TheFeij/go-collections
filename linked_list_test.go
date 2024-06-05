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

// tests Get method of the linked list
func TestLinkedList_Get(t *testing.T) {
	list := NewLinkedList[any]()

	const listSize = 1000
	for i := 0; i < listSize; i++ {
		list.Add(i)
	}

	t.Run("OK", func(t *testing.T) {
		for i := 0; i < listSize; i++ {
			value, ok := list.Get(i)
			require.True(t, ok)
			require.Equal(t, i, value)
		}
	})
	t.Run("Index Out of Bound", func(t *testing.T) {
		t.Run("Index < 0", func(t *testing.T) {
			value, ok := list.Get(-1)
			require.False(t, ok)
			require.Zero(t, value)
		})
		t.Run("Index >= list size", func(t *testing.T) {
			value, ok := list.Get(listSize)
			require.False(t, ok)
			require.Zero(t, value)
		})
	})
}
