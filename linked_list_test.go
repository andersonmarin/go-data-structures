package go_data_structures_test

import (
	. "github.com/andersonmarin/go-data-structures"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	{
		dl := NewLinkedList(4, 2)
		dl.Debug()

		v, ok := dl.Front()
		require.True(t, ok)
		require.Equal(t, v, 4)

		v, ok = dl.Back()
		require.True(t, ok)
		require.Equal(t, v, 2)
	}
	{
		dl := NewLinkedList(3, .14)

		v, ok := dl.Front()
		require.True(t, ok)
		require.Equal(t, v, 3.0)

		v, ok = dl.Back()
		require.True(t, ok)
		require.Equal(t, v, 0.14)
	}
}

func TestLinkedList_Front(t *testing.T) {
	dl := NewLinkedList[int]()

	v, ok := dl.Front()
	require.False(t, ok)
	require.Empty(t, v)

	dl.PushFront(42)
	v, ok = dl.Front()
	require.True(t, ok)
	require.Equal(t, v, 42)
}

func TestLinkedList_Back(t *testing.T) {
	dl := NewLinkedList[int]()

	v, ok := dl.Back()
	require.False(t, ok)
	require.Empty(t, v)

	dl.PushBack(42)
	v, ok = dl.Back()
	require.True(t, ok)
	require.Equal(t, v, 42)
}

func TestLinkedList_At(t *testing.T) {
	dl := NewLinkedList[int]()
	dl.PushBack(4)
	dl.PushBack(2)

	v, ok := dl.At(0)
	require.True(t, ok)
	require.Equal(t, v, 4)

	v, ok = dl.At(1)
	require.True(t, ok)
	require.Equal(t, v, 2)

	v, ok = dl.At(2)
	require.False(t, ok)
	require.Empty(t, v)
}

func TestLinkedList_PushFront(t *testing.T) {
	dl := NewLinkedList[int]()

	dl.PushFront(2)
	f, _ := dl.Front()
	b, _ := dl.Back()
	require.Equal(t, 2, f)
	require.Equal(t, 2, b)

	dl.PushFront(4)
	f, _ = dl.Front()
	b, _ = dl.Back()
	require.Equal(t, 4, f)
	require.Equal(t, 2, b)

	dl.PushFront(8)
	f, _ = dl.Front()
	b, _ = dl.Back()
	require.Equal(t, 8, f)
	require.Equal(t, 2, b)
}

func TestLinkedList_PushBack(t *testing.T) {
	dl := NewLinkedList[int]()

	dl.PushBack(2)
	f, _ := dl.Front()
	b, _ := dl.Back()
	require.Equal(t, 2, f)
	require.Equal(t, 2, b)

	dl.PushBack(4)
	f, _ = dl.Front()
	b, _ = dl.Back()
	require.Equal(t, 2, f)
	require.Equal(t, 4, b)

	dl.PushBack(8)
	f, _ = dl.Front()
	b, _ = dl.Back()
	require.Equal(t, 2, f)
	require.Equal(t, 8, b)
}

func TestLinkedList_PushAt(t *testing.T) {
	dl := NewLinkedList(1, 2, 4, 5)

	v, ok := dl.At(2)
	require.True(t, ok)
	require.Equal(t, v, 4)

	ok = dl.PushAt(2, 3)
	require.True(t, ok)

	v, ok = dl.At(2)
	require.True(t, ok)
	require.Equal(t, v, 3)

	ok = dl.PushAt(0, 0)
	require.True(t, ok)

	v, ok = dl.At(0)
	require.True(t, ok)
	require.Equal(t, 0, v)

	ok = dl.PushAt(5, 6)
	require.True(t, ok)

	v, ok = dl.At(5)
	require.True(t, ok)
	require.Equal(t, 6, v)

	ok = dl.PushAt(7, 7)
	require.False(t, ok)
}

func TestLinkedList_PopFront(t *testing.T) {
	dl := NewLinkedList[int]()

	dl.PushFront(42)
	f, _ := dl.PopFront()
	require.Equal(t, f, 42)

	f, _ = dl.PopFront()
	require.Empty(t, f)
}

func TestLinkedList_PopBack(t *testing.T) {
	dl := NewLinkedList[int]()

	dl.PushBack(42)
	f, _ := dl.PopBack()
	require.Equal(t, f, 42)

	f, _ = dl.PopBack()
	require.Empty(t, f)
}

func TestLinkedList_PopAt(t *testing.T) {
	dl := NewLinkedList(1, 2, 3, 4, 5)

	v, ok := dl.PopAt(2)
	require.True(t, ok)
	require.Equal(t, v, 3)

	v, ok = dl.PopAt(0)
	require.True(t, ok)
	require.Equal(t, v, 1)

	v, ok = dl.At(0)
	require.True(t, ok)
	require.Equal(t, v, 2)

	v, ok = dl.PopAt(2)
	require.True(t, ok)
	require.Equal(t, v, 5)

	v, ok = dl.PopAt(2)
	require.False(t, ok)
	require.Empty(t, v)
}
