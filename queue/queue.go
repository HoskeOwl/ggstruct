package queue

import (
	"fmt"
	"strings"
)

type (
	Queue[QT any] struct {
		start, end *node[QT]
		length     int
	}
	node[NT any] struct {
		value NT
		next  *node[NT]
	}
)

func New[T any](initial ...T) *Queue[T] {
	q := &Queue[T]{}
	for _, n := range initial {
		q.Enqueue(n)
	}
	return q
}

func (q *Queue[QT]) String() string {
	if q.length < 1 {
		return ""
	}
	var keys []string = make([]string, q.length)
	n := q.start
	var i int = 0
	for {
		keys[i] = fmt.Sprint(n.value)
		i += 1
		n = n.next
		if n == nil {
			break
		}
	}
	return fmt.Sprintf("(%v)", strings.Join(keys, ", "))
}

// Take the next item off the front of the queue
func (q *Queue[QT]) Dequeue() (res QT, exists bool) {
	if q.length == 0 {
		exists = false
		return
	}
	exists = true
	res = q.start.value
	if q.length == 1 {
		q.start = nil
		q.end = nil
	} else {
		q.start = q.start.next
	}
	q.length--
	return
}

// Put an item on the end of a queue
func (q *Queue[QT]) Enqueue(value QT) {
	n := &node[QT]{value, nil}
	if q.length == 0 {
		q.start = n
		q.end = n
	} else {
		q.end.next = n
		q.end = n
	}
	q.length++
}

// Return the number of items in the queue
func (q *Queue[QT]) Len() int {
	return q.length
}

// Return the first item in the queue without removing it
func (q *Queue[QT]) Peek() (res QT, exists bool) {
	if q.length == 0 {
		exists = false
		return
	}
	return q.start.value, true
}
