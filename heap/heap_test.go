package heap

import (
	"github.com/stretchr/testify/require"
	"math"
	"sort"
	"testing"
)

var data = []int{2, 4, 6, 7, -53, -1, 34, 68, 0, 0}

func TestNewHeap(t *testing.T) {
	t.Run("Heap Without Comparator", func(t *testing.T) {
		h := NewHeap[int](nil)

		require.Nil(t, h)
	})
	t.Run("Min Heap Without Initial Values", func(t *testing.T) {
		h := NewHeap[int](func(t1, t2 int) bool {
			return t1 < t2
		})

		require.NotNil(t, h)
		require.NotEmpty(t, h)
		require.Equal(t, 0, h.Size())

		value, ok := h.Peek()
		require.False(t, ok)
		require.Zero(t, value)

		value, ok = h.Extract()
		require.False(t, ok)
		require.Zero(t, value)

		require.Equal(t, 0, h.Size())
	})
	t.Run("Min Heap With Initial Values", func(t *testing.T) {
		h := NewHeap[int](func(t1, t2 int) bool {
			return t1 < t2
		}, data...)

		require.NotNil(t, h)
		require.NotEmpty(t, h)
		require.Equal(t, len(data), h.Size())

		sort.Ints(data)

		for _, number := range data {
			value, ok := h.Peek()
			require.True(t, ok)
			require.Equal(t, number, value)

			value, ok = h.Extract()
			require.True(t, ok)
			require.Equal(t, number, value)
		}

		value, ok := h.Peek()
		require.False(t, ok)
		require.Zero(t, value)

		value, ok = h.Extract()
		require.False(t, ok)
		require.Zero(t, value)

		require.Equal(t, 0, h.Size())
	})
	t.Run("Max Heap Without Initial Values", func(t *testing.T) {
		h := NewHeap[int](func(t1, t2 int) bool {
			return t1 > t2
		})

		require.NotNil(t, h)
		require.NotEmpty(t, h)
		require.Equal(t, 0, h.Size())

		value, ok := h.Peek()
		require.False(t, ok)
		require.Zero(t, value)

		value, ok = h.Extract()
		require.False(t, ok)
		require.Zero(t, value)

		require.Equal(t, 0, h.Size())
	})
	t.Run("Min Heap With Initial Values", func(t *testing.T) {
		h := NewHeap[int](func(t1, t2 int) bool {
			return t1 > t2
		}, data...)

		require.NotNil(t, h)
		require.NotEmpty(t, h)
		require.Equal(t, len(data), h.Size())

		sort.Slice(data, func(i, j int) bool {
			return data[i] > data[j]
		})

		for _, number := range data {
			value, ok := h.Peek()
			require.True(t, ok)
			require.Equal(t, number, value)

			value, ok = h.Extract()
			require.True(t, ok)
			require.Equal(t, number, value)
		}

		value, ok := h.Peek()
		require.False(t, ok)
		require.Zero(t, value)

		value, ok = h.Extract()
		require.False(t, ok)
		require.Zero(t, value)

		require.Equal(t, 0, h.Size())
	})
}

func TestHeap_Insert(t *testing.T) {
	t.Run("Insert To Min Heap Without Initial Values", func(t *testing.T) {
		h := NewHeap[int](func(t1, t2 int) bool {
			return t1 < t2
		})

		require.NotNil(t, h)
		require.NotEmpty(t, h)
		require.Equal(t, 0, h.Size())

		value, ok := h.Peek()
		require.False(t, ok)
		require.Zero(t, value)

		sort.Slice(data, func(i, j int) bool {
			return data[i] > data[j]
		})

		for _, number := range data {
			h.Insert(number)

			value, ok := h.Peek()
			require.True(t, ok)
			require.Equal(t, number, value)
		}
	})
	t.Run("Insert To Min Heap With Initial Values", func(t *testing.T) {
		h := NewHeap[int](func(t1, t2 int) bool {
			return t1 < t2
		}, math.MaxInt, math.MaxInt)

		require.NotNil(t, h)
		require.NotEmpty(t, h)
		require.Equal(t, 2, h.Size())

		value, ok := h.Peek()
		require.True(t, ok)
		require.Equal(t, math.MaxInt, value)

		sort.Slice(data, func(i, j int) bool {
			return data[i] > data[j]
		})

		for index, number := range data {
			h.Insert(number)

			value, ok := h.Peek()
			require.True(t, ok)
			require.Equal(t, number, value)
			require.Equal(t, index+3, h.Size())
		}
	})
	t.Run("Insert To Max Heap Without Initial Values", func(t *testing.T) {
		h := NewHeap[int](func(t1, t2 int) bool {
			return t1 > t2
		})

		require.NotNil(t, h)
		require.NotEmpty(t, h)
		require.Equal(t, 0, h.Size())

		value, ok := h.Peek()
		require.False(t, ok)
		require.Zero(t, value)

		sort.Slice(data, func(i, j int) bool {
			return data[i] < data[j]
		})

		for _, number := range data {
			h.Insert(number)

			value, ok := h.Peek()
			require.True(t, ok)
			require.Equal(t, number, value)
		}
	})
	t.Run("Insert To Min Heap With Initial Values", func(t *testing.T) {
		h := NewHeap[int](func(t1, t2 int) bool {
			return t1 > t2
		}, math.MinInt, math.MinInt)

		require.NotNil(t, h)
		require.NotEmpty(t, h)
		require.Equal(t, 2, h.Size())

		value, ok := h.Peek()
		require.True(t, ok)
		require.Equal(t, math.MinInt, value)

		sort.Slice(data, func(i, j int) bool {
			return data[i] < data[j]
		})

		for index, number := range data {
			h.Insert(number)

			value, ok := h.Peek()
			require.True(t, ok)
			require.Equal(t, number, value)
			require.Equal(t, index+3, h.Size())
		}
	})
}

func TestHeap_Extract(t *testing.T) {
	t.Run("Extract From Min Heap Without Initial Values", func(t *testing.T) {
		h := NewHeap[int](func(t1, t2 int) bool {
			return t1 < t2
		})

		require.NotNil(t, h)
		require.NotEmpty(t, h)
		require.Equal(t, 0, h.Size())

		value, ok := h.Extract()
		require.False(t, ok)
		require.Zero(t, value)
	})
	t.Run("Extract From Min Heap With Initial Values", func(t *testing.T) {
		h := NewHeap[int](func(t1, t2 int) bool {
			return t1 < t2
		}, math.MaxInt, math.MaxInt)

		require.NotNil(t, h)
		require.NotEmpty(t, h)
		require.Equal(t, 2, h.Size())

		value, ok := h.Extract()
		require.True(t, ok)
		require.Equal(t, math.MaxInt, value)

		value, ok = h.Extract()
		require.True(t, ok)
		require.Equal(t, math.MaxInt, value)

		value, ok = h.Extract()
		require.False(t, ok)
		require.Zero(t, value)
	})
	t.Run("Extract From Max Heap Without Initial Values", func(t *testing.T) {
		h := NewHeap[int](func(t1, t2 int) bool {
			return t1 > t2
		})

		require.NotNil(t, h)
		require.NotEmpty(t, h)
		require.Equal(t, 0, h.Size())

		value, ok := h.Extract()
		require.False(t, ok)
		require.Zero(t, value)
	})
	t.Run("Extract From Min Heap With Initial Values", func(t *testing.T) {
		h := NewHeap[int](func(t1, t2 int) bool {
			return t1 > t2
		}, math.MinInt, math.MinInt)

		require.NotNil(t, h)
		require.NotEmpty(t, h)
		require.Equal(t, 2, h.Size())

		value, ok := h.Extract()
		require.True(t, ok)
		require.Equal(t, math.MinInt, value)

		value, ok = h.Extract()
		require.True(t, ok)
		require.Equal(t, math.MinInt, value)

		value, ok = h.Extract()
		require.False(t, ok)
		require.Zero(t, value)
	})
}
