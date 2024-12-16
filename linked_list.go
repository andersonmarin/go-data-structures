package go_data_structures

import "fmt"

type linkedListNode[T any] struct {
	data T
	next *linkedListNode[T]
	prev *linkedListNode[T]
}

type LinkedList[T any] struct {
	head *linkedListNode[T]
	tail *linkedListNode[T]
	size int
}

func NewLinkedList[T any](datas ...T) *LinkedList[T] {
	l := new(LinkedList[T])
	for _, data := range datas {
		l.PushBack(data)
	}
	return l
}

func (l *LinkedList[T]) Front() (T, bool) {
	if l.head == nil {
		return *new(T), false
	}
	return l.head.data, true
}

func (l *LinkedList[T]) Back() (T, bool) {
	if l.tail == nil {
		return *new(T), false
	}
	return l.tail.data, true
}

func (l *LinkedList[T]) At(index int) (T, bool) {
	cursor := l.at(index)
	if cursor == nil {
		return *new(T), false
	}
	return cursor.data, true
}

func (l *LinkedList[T]) PushFront(data T) {
	node := &linkedListNode[T]{
		data: data,
		next: l.head,
	}
	if l.head == nil {
		l.tail = node
	} else {
		l.head.prev = node
	}
	l.head = node
	l.size++
}

func (l *LinkedList[T]) PushBack(data T) {
	node := &linkedListNode[T]{
		data: data,
		prev: l.tail,
	}
	if l.tail == nil {
		l.head = node
	} else {
		l.tail.next = node
	}
	l.tail = node
	l.size++
}

func (l *LinkedList[T]) PushAt(index int, data T) bool {
	cursor := l.at(index)
	if cursor == nil {
		return false
	}
	node := &linkedListNode[T]{
		data: data,
		next: cursor,
		prev: cursor.prev,
	}
	if cursor.prev == nil {
		l.head = node
	} else {
		cursor.prev.next = node
	}
	cursor.prev = node
	l.size++
	return true
}

func (l *LinkedList[T]) PopFront() (T, bool) {
	node := l.head
	if node == nil {
		return *new(T), false
	}
	l.head = node.next
	l.size--
	return node.data, true
}

func (l *LinkedList[T]) PopBack() (T, bool) {
	node := l.tail
	if node == nil {
		return *new(T), false
	}
	l.tail = node.prev
	l.size--
	return node.data, true
}

func (l *LinkedList[T]) PopAt(index int) (T, bool) {
	cursor := l.at(index)
	if cursor == nil {
		return *new(T), false
	}
	if cursor.next == nil {
		l.tail = cursor.prev
	} else {
		cursor.next.prev = cursor.prev
	}
	if cursor.prev == nil {
		l.head = cursor.next
	} else {
		cursor.prev.next = cursor.next
	}
	l.size--
	return cursor.data, true
}

func (l *LinkedList[T]) Debug() {
	cursor := l.head
	for cursor != nil {
		fmt.Print(cursor.data)
		cursor = cursor.next
		if cursor != nil {
			fmt.Print(" → ")
		}
	}
	fmt.Println()
	cursor = l.tail
	for cursor != nil {
		fmt.Print(cursor.data)
		cursor = cursor.prev
		if cursor != nil {
			fmt.Print(" ← ")
		}
	}
	fmt.Println()
}

func (l *LinkedList[T]) at(index int) *linkedListNode[T] {
	if index >= 0 && index < l.size {
		if index < l.size/2 {
			cursor := l.head
			for cursor != nil {
				if index == 0 {
					return cursor
				}
				cursor = cursor.next
				index--
			}
		} else {
			index = l.size - 1 - index
			cursor := l.tail
			for cursor != nil {
				if index == 0 {
					return cursor
				}
				cursor = cursor.prev
				index--
			}
		}
	}
	return nil
}
