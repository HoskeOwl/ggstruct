// Package queue provides simple structure
package queue

import (
	"github.com/HoskeOwl/ggstruct/list"
)

type Queue[QT comparable] struct {
	data  *list.List[QT]
	limit int
}

// WithLimit sets the maximum number of elements in the queue and returns pointer to itself.
// limit=0 mens no limits.
func (q *Queue[QT]) WithLimit(limit int) *Queue[QT] {
	q.limit = limit
	return q
}

// Limit Returns current limit value.
func (q *Queue[QT]) Limit() int {
	if q.limit < 0 {
		return 0
	}
	return q.limit
}

// New returns new queue instance
func New[T comparable]() *Queue[T] {
	return &Queue[T]{
		data:  list.New[T](),
		limit: 0,
	}
}

// Dequeue get and remove the next element from the queue
func (q *Queue[QT]) Dequeue() (value QT, exists bool) {
	return q.data.PopFront()
}

// Delete remove an element from the queue
func (q *Queue[QT]) Delete(value QT) bool {
	return q.data.Delete(value)
}

// Enqueue add a new element to the queue
func (q *Queue[QT]) Enqueue(value QT) bool {
	if q.limit > 0 && q.data.Len()+1 > q.limit {
		return false
	}
	q.data.PushBack(value)
	return true
}

// Len returns the number of items in the queue
func (q *Queue[QT]) Len() int {
	return q.data.Len()
}

// Peek returns the first item in the queue without removing it
func (q *Queue[QT]) Peek() (res QT, exists bool) {
	return q.data.Front()
}

// Contains check if element is in the queue.
func (q *Queue[QT]) Contains(value QT) bool {
	return q.data.Contains(value)
}

// Clear removes all elements from the queue
func (q *Queue[QT]) Clear() {
	q.data = list.New[QT]()
}

// Clone returns a new queue with the same elements
func (q *Queue[QT]) Clone() *Queue[QT] {
	return &Queue[QT]{
		data:  q.data.Clone(),
		limit: q.limit,
	}
}

// IsFull returns 'true' if elements count equal or greater than limit. With limit <= 0 always returns false.
func (q *Queue[QT]) IsFull() bool {
	if q.limit <= 0 {
		return false
	}
	// >= because we can change limit in any time
	return q.data.Len() >= q.limit
}

// IsEmpty returns 'true' if no elements on the queue.
func (q *Queue[QT]) IsEmpty() bool {
	return q.data.Len() == 0
}
