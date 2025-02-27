// Package stack provides simple structure
package stack

import (
	"github.com/HoskeOwl/ggstruct/list"
)

type Stack[QT comparable] struct {
	data  *list.List[QT]
	limit int
}

// Limit returns current maximum number of elements.
func (s *Stack[QT]) Limit() int { return s.limit }

// WithLimit sets the maximum number of elements in the queue and returns pointer to itself.
// limit=0 mens no limits.
func (q *Stack[QT]) WithLimit(limit int) *Stack[QT] {
	q.limit = limit
	return q
}

// New creates a new stack.
func New[T comparable](initial ...T) *Stack[T] {
	q := &Stack[T]{
		data:  list.New[T](initial...),
		limit: 0,
	}
	return q
}

// Pop removes and return the next element.
func (q *Stack[QT]) Pop() (value QT, exists bool) {
	return q.data.PopFront()
}

// Push adds a new element to stack.
func (q *Stack[QT]) Push(value QT) bool {
	if q.limit > 0 && q.data.Len()+1 > q.limit {
		return false
	}
	q.data.PushFront(value)
	return true
}

// IsFull returns 'true' if elements count equal or greater than limit. With limit <= 0 always return false.
func (q *Stack[QT]) IsFull() bool {
	if q.limit <= 0 {
		return false
	}
	// >= because we can change limit in any time
	return q.data.Len() >= q.limit
}

// IsEmpty returns 'true' if no elements on the queue.
func (q *Stack[QT]) IsEmpty() bool {
	return q.data.Len() == 0
}

// Len returns the number of items in the queue.
func (q *Stack[QT]) Len() int {
	return q.data.Len()
}

// Top returns the first item in the stack without removing it.
func (q *Stack[QT]) Top() (res QT, exists bool) {
	return q.data.Front()
}

// Contains returns 'true' if the element is in the stack.
func (q *Stack[QT]) Contains(value QT) bool {
	return q.data.Contains(value)
}

// Remove removes the first equal element from the stack.
func (q *Stack[QT]) Remove(value QT) {
	idx := q.data.Index(value)
	if idx == -1 {
		return
	}
	q.data.PopAt(idx)
}

// Clear removes all elements from the stack.
func (q *Stack[QT]) Clear() {
	q.data = list.New[QT]()
}

// Clone returns a new stack with same elements with the same order.
func (q *Stack[QT]) Clone() *Stack[QT] {
	return &Stack[QT]{
		data:  q.data.Clone(),
		limit: q.limit,
	}
}
