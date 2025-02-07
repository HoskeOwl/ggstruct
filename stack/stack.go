package stack

import (
	"github.com/HoskeOwl/ggstruct/list"
)

type Stack[QT comparable] struct {
	data  *list.List[QT]
	Limit int
}

func (q *Stack[QT]) WithLimit(limit int) *Stack[QT] {
	q.Limit = limit
	return q
}

func New[T comparable](initial ...T) *Stack[T] {
	q := &Stack[T]{
		data:  list.New[T](initial...),
		Limit: 0,
	}
	return q
}

func (q *Stack[QT]) Pop() (value QT, exists bool) {
	return q.data.PopFront()
}

func (q *Stack[QT]) Delete(value QT) bool {
	return q.data.Delete(value)
}

func (q *Stack[QT]) Push(value QT) bool {
	if q.Limit > 0 && q.data.Len()+1 > q.Limit {
		return false
	}
	q.data.PushFront(value)
	return true
}

func (q *Stack[QT]) IsFull() bool {
	if q.Limit <= 0 {
		return false
	}
	// >= because we can change limit in any time
	return q.data.Len() >= q.Limit
}

func (q *Stack[QT]) IsEmpty() bool {
	return q.data.Len() == 0
}

// Return the number of items in the queue
func (q *Stack[QT]) Len() int {
	return q.data.Len()
}

// Return the first item in the queue without removing it
func (q *Stack[QT]) Top() (res QT, exists bool) {
	return q.data.Front()
}

func (q *Stack[QT]) Contains(value QT) bool {
	return q.data.Contains(value)
}

func (q *Stack[QT]) Remove(value QT) {
	idx := q.data.Index(value)
	if idx == -1 {
		return
	}
	q.data.PopAt(idx)
}

func (q *Stack[QT]) Clear() {
	q.data = list.New[QT]()
}
