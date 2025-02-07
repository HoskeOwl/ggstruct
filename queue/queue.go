package queue

import (
	"github.com/HoskeOwl/ggstruct/list"
)

type Queue[QT comparable] struct {
	data  *list.List[QT]
	Limit int
}

func (q *Queue[QT]) WithLimit(limit int) *Queue[QT] {
	q.Limit = limit
	return q
}

func New[T comparable](initial ...T) *Queue[T] {
	q := &Queue[T]{
		data:  list.New[T](initial...),
		Limit: 0,
	}
	return q
}

func (q *Queue[QT]) Dequeue() (value QT, exists bool) {
	return q.data.PopFront()
}

func (q *Queue[QT]) Delete(value QT) bool {
	return q.data.Delete(value)
}

func (q *Queue[QT]) Enqueue(value QT) bool {
	if q.Limit > 0 && q.data.Len()+1 > q.Limit {
		return false
	}
	q.data.PushBack(value)
	return true
}

// Return the number of items in the queue
func (q *Queue[QT]) Len() int {
	return q.data.Len()
}

// Return the first item in the queue without removing it
func (q *Queue[QT]) Peek() (res QT, exists bool) {
	return q.data.Front()
}

func (q *Queue[QT]) Contains(value QT) bool {
	return q.data.Contains(value)
}

func (q *Queue[QT]) Remove(value QT) {
	idx := q.data.Index(value)
	if idx == -1 {
		return
	}
	q.data.PopAt(idx)
}

func (q *Queue[QT]) Clear() {
	q.data = list.New[QT]()
}
