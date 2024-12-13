package go_data_structures_test

import (
	. "github.com/andersonmarin/go-data-structures"
	"github.com/stretchr/testify/require"
	"testing"
)

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
