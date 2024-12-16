package go_data_structures

type linkedListNode[T any] struct {
	data T
	next *linkedListNode[T]
	prev *linkedListNode[T]
}

type LinkedList[T any] struct {
	head *linkedListNode[T]
	tail *linkedListNode[T]
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
	cursor := l.head
	for cursor != nil {
		if index == 0 {
			return cursor.data, true
		}
		cursor = cursor.next
		index--
	}
	return *new(T), false
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
}

func (l *LinkedList[T]) PushAt(index int, data T) bool {
	cursor := l.head
	for cursor != nil {
		if index == 0 {
			node := &linkedListNode[T]{
				data: data,
				prev: cursor.prev,
				next: cursor,
			}
			cursor.next.prev = node
			cursor.prev.next = node
			return true
		}
		cursor = cursor.next
		index--
	}
	return false
}

func (l *LinkedList[T]) PopFront() (T, bool) {
	node := l.head
	if node == nil {
		return *new(T), false
	}
	l.head = node.next
	return node.data, true
}

func (l *LinkedList[T]) PopBack() (T, bool) {
	node := l.tail
	if node == nil {
		return *new(T), false
	}
	l.tail = node.prev
	return node.data, true
}

func (l *LinkedList[T]) PopAt(index int) (T, bool) {
	cursor := l.head
	for cursor != nil {
		if index == 0 {
			if cursor.next != nil {
				cursor.next.prev = cursor.prev
			}
			if cursor.prev != nil {
				cursor.prev.next = cursor.next
			}
			return cursor.data, true
		}
		cursor = cursor.next
		index--
	}
	return *new(T), false
}
