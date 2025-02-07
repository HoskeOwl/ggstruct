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

type List[LT comparable] struct {
	root *element[LT]
	cnt  int
}

func New[LT comparable](data ...LT) *List[LT] {
	l := &List[LT]{}
	for _, d := range data {
		l.PushBack(d)
	}
	return l
}

func (l *List[LT]) Len() int {
	return l.cnt
}

func (l *List[LT]) Clear() {
	if l.root == nil {
		return
	}

	e := l.root
	var next *element[LT]
	for i := 0; i < l.cnt; i++ {
		next = e.next
		e.clear()
		e = next
	}
}

func (l *List[LT]) Seq() iter.Seq[LT] {
	return func(yield func(LT) bool) {
		if l.root == nil {
			return
		}
		cur := l.root
		for i := 0; i < l.cnt; i++ {
			if !yield(cur.data) {
				return
			}
			cur = cur.next
		}
		return
	}
}

func (l *List[LT]) ReversedSeq() iter.Seq[LT] {
	return func(yield func(LT) bool) {
		if l.root == nil {
			return
		}
		cur := l.root.prev
		for i := 0; i < l.cnt; i++ {
			if !yield(cur.data) {
				return
			}
			cur = cur.prev
		}
		return
	}
}

func (l *List[LT]) Seq2() iter.Seq2[int, LT] {
	return func(yield func(int, LT) bool) {
		if l.root == nil {
			return
		}
		cur := l.root
		for i := 0; i < l.cnt; i++ {
			if !yield(i, cur.data) {
				return
			}
			cur = cur.next
		}
		return
	}
}

func (l *List[LT]) ReversedSeq2() iter.Seq2[int, LT] {
	return func(yield func(int, LT) bool) {
		if l.root == nil {
			return
		}
		cur := l.root.prev
		for i := 0; i < l.cnt; i++ {
			if !yield(i, cur.data) {
				return
			}
			cur = cur.prev
		}
		return
	}
}

func (l *List[T]) PushFront(values ...T) {
	for i := len(values) - 1; i >= 0; i-- {
		l.cnt += 1
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

func (l *List[T]) PushBack(values ...T) {
	for _, v := range values {
		l.cnt += 1
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

func (l *List[T]) AddAfterIndex(index int, values ...T) bool {
	if l.root == nil || index < 0 || l.cnt <= index {
		return false
	}

	var newNode *element[T]
	// go to index node
	current := l.root
	for i := 0; i < index; i++ {
		current = current.next
	}

	for _, v := range values {
		l.cnt += 1
		newNode = &element[T]{data: v}
		newNode.next = current.next
		newNode.prev = current
		current.next = newNode
		newNode.next.prev = newNode
		current = newNode
	}
	return true
}

func (l *List[T]) AddBeforeIndex(index int, values ...T) bool {
	if l.root == nil || index < 0 || l.cnt <= index {
		return false
	}

	var newNode *element[T]
	// go to index node
	current := l.root
	for i := 0; i < index; i++ {
		current = current.next
	}

	for i := len(values) - 1; i >= 0; i-- {
		l.cnt += 1
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

func (l *List[T]) Front() (val T, exists bool) {
	if l.root == nil {
		exists = false
	} else {
		val = l.root.data
		exists = true
	}

	return
}

func (l *List[T]) Back() (val T, exists bool) {
	if l.root == nil {
		exists = false
	} else {
		val = l.root.prev.data
		exists = true
	}

	return
}

func (l *List[LT]) PeakAt(index int) (val LT, exists bool) {
	if l.root == nil || index < 0 || l.cnt <= index {
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

func (l *List[T]) PopAt(index int) (val T, exists bool) {
	if l.root == nil || index < 0 || l.cnt <= index {
		return val, false
	}

	l.cnt -= 1
	e := l.root
	if l.cnt == 0 {
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

func (l *List[T]) PopFront() (val T, exists bool) {
	if l.root == nil || l.cnt == 0 {
		exists = false
		return
	}

	l.cnt -= 1
	r := l.root
	next := r.next
	next.prev = r.prev
	if l.cnt == 0 {
		l.root = nil
	} else {
		l.root = next
	}
	r.clear()

	val = r.data
	exists = true
	return
}

func (l *List[T]) PopBack() (val T, exists bool) {
	if l.root == nil || l.cnt == 0 {
		exists = false
		return
	}

	l.cnt -= 1
	p := l.root.prev
	prev := p.prev
	p.prev = l.root
	if l.cnt == 0 {
		l.root = nil
	} else {
		l.root.prev = prev
	}
	p.clear()

	val = p.data
	exists = true
	return
}

func (l *List[T]) Index(value T) int {
	if l.root == nil || l.cnt == 0 {
		return -1
	}

	e := l.root
	for i := 0; i < l.cnt; i++ {
		if e.data == value {
			return i
		}
		e = e.next
	}
	return -1
}

func (l *List[T]) Contains(value T) bool {
	return l.Index(value) > -1
}

func (l *List[T]) MoveAfter(from, to int) bool {
	if l.root == nil || l.cnt == 0 || l.cnt == 1 || from < 0 || to < 0 || from >= l.cnt || to >= l.cnt {
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

func (l *List[T]) MoveBefore(from, to int) bool {
	if l.root == nil || l.cnt == 0 || l.cnt == 1 || from < 0 || to < 0 || from >= l.cnt || to >= l.cnt {
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

func (l *List[T]) Clone() *List[T] {
	result := New[T]()
	for v := range l.Seq() {
		result.PushBack(v)
	}
	return result
}
