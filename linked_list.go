package go_data_structures

type linkedListNode[T any] struct {
	value    T
	next     *linkedListNode[T]
	previous *linkedListNode[T]
}

type LinkedList[T any] struct {
	head *linkedListNode[T]
	tail *linkedListNode[T]
}

func NewLinkedList[T any](items ...T) *LinkedList[T] {
	l := new(LinkedList[T])
	for _, item := range items {
		l.PushBack(item)
	}
	return l
}

func (l *LinkedList[T]) Front() T {
	return l.head.value
}

func (l *LinkedList[T]) Back() T {
	return l.tail.value
}

func (l *LinkedList[T]) PushFront(item T) {
	node := &linkedListNode[T]{
		value: item,
		next:  l.head,
	}
	if l.head == nil {
		l.tail = node
	} else {
		l.head.previous = node
	}
	l.head = node
}

func (l *LinkedList[T]) PushBack(item T) {
	node := &linkedListNode[T]{
		value:    item,
		previous: l.tail,
	}
	if l.tail == nil {
		l.head = node
	} else {
		l.tail.next = node
	}
	l.tail = node
}

func (l *LinkedList[T]) PopFront() T {
	node := l.head
	l.head = node.next
	return node.value
}

func (l *LinkedList[T]) PopBack() T {
	node := l.tail
	l.tail = node.previous
	return node.value
}
