package go_data_structures_test

import (
	. "github.com/andersonmarin/go-data-structures"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLinkedList_PushFront(t *testing.T) {
	dl := NewLinkedList[int]()

	dl.PushFront(2)
	require.Equal(t, 2, dl.Front())
	require.Equal(t, 2, dl.Back())

	dl.PushFront(4)
	require.Equal(t, 4, dl.Front())
	require.Equal(t, 2, dl.Back())

	dl.PushFront(8)
	require.Equal(t, 8, dl.Front())
	require.Equal(t, 2, dl.Back())
}

func TestLinkedList_PushBack(t *testing.T) {
	dl := NewLinkedList[int]()

	dl.PushBack(2)
	require.Equal(t, 2, dl.Front())
	require.Equal(t, 2, dl.Back())

	dl.PushBack(4)
	require.Equal(t, 2, dl.Front())
	require.Equal(t, 4, dl.Back())

	dl.PushBack(8)
	require.Equal(t, 2, dl.Front())
	require.Equal(t, 8, dl.Back())
}
