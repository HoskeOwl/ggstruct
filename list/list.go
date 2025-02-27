// Package list provides common two-linked list with iterators
//
// Can store any comparable value. No pointers and type-checks.
// Able to be used with 'range' and some functions of 'slices' package
package list

import (
	"iter"
)

type element[NT comparable] struct {
	data NT
	prev *element[NT]
	next *element[NT]
}

func (n *element[NT]) clear() {
	n.next = nil
	n.prev = nil
}

// List  Two-side linked list
type List[LT comparable] struct {
	root *element[LT]
	len  int
}

// New Create a new instance of List.
// Can be filled through initialization with direct order.
// l := New(1, 2, 3, 4, 5) returns List: 1 <-> 2 <-> 3 <-> 4 <-> 5
func New[LT comparable](data ...LT) *List[LT] {
	l := &List[LT]{}
	for _, d := range data {
		l.PushBack(d)
	}
	return l
}

// Len Returns count of elements in the list.
func (l *List[LT]) Len() int {
	return l.len
}

// Clear Carefully remove all elements from the List. Prevent memory leaks and clear all links between elements. O(N)
func (l *List[LT]) Clear() {
	if l.root == nil {
		return
	}

	e := l.root
	var next *element[LT]
	for i := 0; i < l.len; i++ {
		next = e.next
		e.clear()
		e = next
	}
}

// Seq Return function for value-only sequence. Can be used in slices library and range.
func (l *List[LT]) Seq() iter.Seq[LT] {
	return func(yield func(LT) bool) {
		if l.root == nil {
			return
		}
		cur := l.root
		for i := 0; i < l.len; i++ {
			if !yield(cur.data) {
				return
			}
			cur = cur.next
		}
		return
	}
}

// ReversedSeq Return function for value-only sequence but reversed order. Can be used in slices library and range.
func (l *List[LT]) ReversedSeq() iter.Seq[LT] {
	return func(yield func(LT) bool) {
		if l.root == nil {
			return
		}
		cur := l.root.prev
		for i := 0; i < l.len; i++ {
			if !yield(cur.data) {
				return
			}
			cur = cur.prev
		}
		return
	}
}

// Seq2 Return function for int-value sequence (element number and value). Can be used in slices library and range.
func (l *List[LT]) Seq2() iter.Seq2[int, LT] {
	return func(yield func(int, LT) bool) {
		if l.root == nil {
			return
		}
		cur := l.root
		for i := 0; i < l.len; i++ {
			if !yield(i, cur.data) {
				return
			}
			cur = cur.next
		}
		return
	}
}

// ReversedSeq2 Return function for int-value sequence (element number and value), but reversed order.
// Can be used in slices library and range.
func (l *List[LT]) ReversedSeq2() iter.Seq2[int, LT] {
	return func(yield func(int, LT) bool) {
		if l.root == nil {
			return
		}
		cur := l.root.prev
		for i := l.Len() - 1; i >= 0; i-- {
			if !yield(i, cur.data) {
				return
			}
			cur = cur.prev
		}
		return
	}
}

// PushFront Add value to the start of the list.
func (l *List[T]) PushFront(values ...T) {
	for i := len(values) - 1; i >= 0; i-- {
		l.len += 1
		var newNode *element[T]
		if l.root != nil {
			newNode = &element[T]{data: values[i], next: l.root, prev: l.root.prev}
			newNode.prev.next = newNode
			l.root.prev = newNode
		} else {
			newNode = &element[T]{data: values[i]}
			newNode.next = newNode
			newNode.prev = newNode
		}
		l.root = newNode
	}
	return
}

// PushBack Add value to the end of the list.
func (l *List[T]) PushBack(values ...T) {
	for _, v := range values {
		l.len += 1
		var newNode *element[T]
		if l.root != nil {
			last := l.root.prev
			newNode = &element[T]{data: v, next: l.root, prev: last}
			last.next = newNode
			l.root.prev = newNode
		} else {
			newNode = &element[T]{data: v}
			newNode.next = newNode
			newNode.prev = newNode
			l.root = newNode
		}
	}
	return
}

// AddAfterIndex Add value to the specific position.
// If list is 1 <-> 2 <-> 3 <-> 4 then you call this function with index 2 (for example with value 9).
// You will get 1 <-> 2 <-> 3 <-> 9 <-> 4
func (l *List[T]) AddAfterIndex(index int, values ...T) bool {
	if l.root == nil || index < 0 || l.len <= index {
		return false
	}

	var newNode *element[T]
	// go to index node
	current := l.root
	for i := 0; i < index; i++ {
		current = current.next
	}

	for _, v := range values {
		l.len += 1
		newNode = &element[T]{data: v}
		newNode.next = current.next
		newNode.prev = current
		current.next = newNode
		newNode.next.prev = newNode
		current = newNode
	}
	return true
}

// AddBeforeIndex Add value to the specific position.
// If list is 1 <-> 2 <-> 3 <-> 4 then you call this function with index 2 (for example with value 9).
// You will get 1 <-> 2 <-> 9 <-> 3 <-> 4
func (l *List[T]) AddBeforeIndex(index int, values ...T) bool {
	if l.root == nil || index < 0 || l.len <= index {
		return false
	}

	var newNode *element[T]
	// go to index node
	current := l.root
	for i := 0; i < index; i++ {
		current = current.next
	}

	for i := len(values) - 1; i >= 0; i-- {
		l.len += 1
		newNode = &element[T]{data: values[i]}
		newNode.next = current
		newNode.prev = current.prev
		current.prev = newNode
		newNode.prev.next = newNode
		if current == l.root {
			l.root = newNode
		}
		current = newNode
	}
	return true
}

// Front returns the first element of the list
func (l *List[T]) Front() (val T, exists bool) {
	if l.root == nil {
		exists = false
	} else {
		val = l.root.data
		exists = true
	}

	return
}

// Back returns the last element of the list
func (l *List[T]) Back() (val T, exists bool) {
	if l.root == nil {
		exists = false
	} else {
		val = l.root.prev.data
		exists = true
	}

	return
}

// PeakAt returns an element at the specific of the list and 'true'. If there is no element on this position
// the default value will be returned and 'false' as the second argument.
func (l *List[LT]) PeakAt(index int) (val LT, exists bool) {
	if l.root == nil || index < 0 || l.len <= index {
		exists = false
		return
	}

	e := l.root
	for i := 0; i < index; i++ {
		e = e.next
	}
	val = e.data
	exists = true
	return
}

// PopAt removes an element from the list and return them, the second argument will be 'true'.
// If there is no element on this position the default value will be returned and 'false' as the second argument.
func (l *List[T]) PopAt(index int) (val T, exists bool) {
	if l.root == nil || index < 0 || l.len <= index {
		return val, false
	}

	l.len -= 1
	e := l.root
	if l.len == 0 {
		l.root = nil
	} else {
		for i := 0; i < index; i++ {
			e = e.next
		}
		e.prev.next = e.next
		e.next.prev = e.prev
		if e == l.root {
			l.root = e.next
		}
	}
	e.clear()

	return e.data, true
}

// Delete removes the first found element
func (l *List[T]) Delete(value T) bool {
	if l.root == nil {
		return false
	}

	idx := l.Index(value)
	if idx < 0 {
		return false
	}
	l.PopAt(idx)
	return true
}

// PopFront removes the first element from the list and return them, the second argument will be 'true'.
// If there is no element default value will be returned and 'false' as the second argument.
func (l *List[T]) PopFront() (val T, exists bool) {
	if l.root == nil || l.len == 0 {
		exists = false
		return
	}

	l.len -= 1
	r := l.root
	next := r.next
	next.prev = r.prev
	if l.len == 0 {
		l.root = nil
	} else {
		l.root = next
	}
	r.clear()

	val = r.data
	exists = true
	return
}

// PopBack removes the last element from the list and return them, the second argument will be 'true'.
// If there is no element default value will be returned and 'false' as the second argument.
func (l *List[T]) PopBack() (val T, exists bool) {
	if l.root == nil || l.len == 0 {
		exists = false
		return
	}

	l.len -= 1
	p := l.root.prev
	prev := p.prev
	p.prev = l.root
	if l.len == 0 {
		l.root = nil
	} else {
		l.root.prev = prev
	}
	p.clear()

	val = p.data
	exists = true
	return
}

// Index returns the first index of equal element (from the top).
func (l *List[T]) Index(value T) int {
	if l.root == nil || l.len == 0 {
		return -1
	}

	e := l.root
	for i := 0; i < l.len; i++ {
		if e.data == value {
			return i
		}
		e = e.next
	}
	return -1
}

// RIndex returns the first index of equal element (from the end).
func (l *List[T]) RIndex(value T) int {
	if l.root == nil || l.len == 0 {
		return -1
	}

	e := l.root.prev
	for i := l.len - 1; i >= 0; i-- {
		if e.data == value {
			return i
		}
		e = e.prev
	}
	return -1
}

// Find returns all indexes of equal elements.
func (l *List[T]) Find(value T) []int {
	if l.root == nil || l.len == 0 {
		return nil
	}

	var res []int
	e := l.root
	for i := 0; i < l.len; i++ {
		if e.data == value {
			res = append(res, i)
		}
		e = e.next
	}
	return res
}

// Contains returns true if list has at least one element. In other case returns false.
func (l *List[T]) Contains(value T) bool {
	return l.Index(value) > -1
}

// MoveAfter moves one element from the index 'from' after index 'to'.
func (l *List[T]) MoveAfter(from, to int) bool {
	if l.root == nil || l.len == 0 || l.len == 1 || from < 0 || to < 0 || from >= l.len || to >= l.len {
		return false
	}
	if from == to || from-to == 1 {
		return true
	}

	v, exists := l.PeakAt(from)
	if !exists {
		return false
	}
	l.AddAfterIndex(to, v)
	if to < from {
		from++
	}
	l.PopAt(from)

	return true
}

// MoveBefore moves one element from the index 'from' before index 'to'.
func (l *List[T]) MoveBefore(from, to int) bool {
	if l.root == nil || l.len == 0 || l.len == 1 || from < 0 || to < 0 || from >= l.len || to >= l.len {
		return false
	}
	if from == to || to-from == 1 {
		return true
	}

	v, exists := l.PeakAt(from)
	if !exists {
		return false
	}
	l.AddBeforeIndex(to, v)
	if to < from {
		from++
	}
	l.PopAt(from)

	return true
}

// Equal checks is 'other' list has equal values with the same order.
func (l *List[T]) Equal(other *List[T]) bool {
	if l.Len() != other.Len() {
		return false
	}
	if l.Len() == 0 {
		return true
	}

	// Can't use .Index because of O^2
	next, stop := iter.Pull(other.Seq())
	for v := range l.Seq() {
		o, exists := next()
		if !exists {
			stop()
			return false
		}
		if o != v {
			return false
		}
	}
	return true
}

// Clone creates a new list with equal values with the same order.
func (l *List[T]) Clone() *List[T] {
	result := New[T]()
	for v := range l.Seq() {
		result.PushBack(v)
	}
	return result
}
