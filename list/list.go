package list

import "iter"

type listNode[NT any] struct {
	data NT
	prev *listNode[NT]
	next *listNode[NT]
}

func (n *listNode[NT]) Clear() {
	n.next = nil
	n.prev = nil
}

type List[LT any] struct {
	start *listNode[LT]
	end   *listNode[LT]
	cnt   int
}

func New[LT any](data ...LT) *List[LT] {
	l := &List[LT]{}
	for _, d := range data {
		l.AddRight(d)
	}
	return l
}

func (l *List[LT]) Len() int {
	return l.cnt
}

func (l *List[LT]) AsSlice() []LT {
	if l.start == nil {
		return nil
	}

	res := make([]LT, 0, l.cnt)
	n := l.start

	for ; n != nil; n = n.next {
		res = append(res, n.data)
	}

	return res
}

func (l *List[LT]) Clear() {
	if l.start == nil {
		return
	}
	var cur *listNode[LT]
	var next *listNode[LT]
	ok := true
	for ok {
		cur = l.start
		next = cur.next
		cur.Clear()
		cur = next
		ok = cur != nil
	}
	return
}

func (l *List[LT]) Range() iter.Seq[LT] {
	return func(yield func(LT) bool) {
		if l.start == nil {
			return
		}
		var cur *listNode[LT]
		ok := true
		for ok {
			cur = l.start
			if !yield(cur.data) {
				return
			}
			cur = cur.next
			ok = cur != nil
		}
		return
	}
}

func (l *List[T]) AddLeft(value T) {
	newNode := &listNode[T]{data: value, next: l.start, prev: nil}
	if l.start == nil {
		l.start = newNode
		l.end = newNode
		l.cnt += 1
		return
	}

	l.cnt += 1
	l.start.prev = newNode
	l.start = newNode
	return
}

func (l *List[T]) AddRight(value T) {
	newNode := &listNode[T]{data: value, next: nil, prev: l.end}
	if l.start == nil {
		l.start = newNode
		l.end = newNode
		l.cnt += 1
		return
	}

	l.cnt += 1
	l.end.next = newNode
	l.end = newNode
	return
}

func (l *List[T]) AddIndex(value T, index int) {
	/*
		if index is the last element - add before them
		if index is the next of the last - add to the end
	*/

}

func (l *List[T]) PopLeft() (value T, exists bool) {
	if l.start == nil {
		return
	}
	exists = true
	value = l.start.data
	switch l.cnt {
	case 1:
		l.start.Clear()
		l.start = nil
		l.end = nil
	case 2:
		l.start.Clear()
		l.start = l.end
	default:
		old := l.start
		l.start = l.start.next
		l.start.prev = nil
		// prevent memory leaks - clear all links
		old.Clear()
	}
	l.cnt -= 1
	return
}

func (l *List[T]) PopRight() (value T, exists bool) {
	if l.end == nil {
		return
	}
	exists = true
	value = l.end.data
	switch l.cnt {
	case 1:
		l.start.Clear()
		l.start = nil
		l.end = nil
	case 2:
		l.end.Clear()
		l.end = l.start
	default:
		old := l.end
		l.end = l.end.prev
		l.end.next = nil
		old.Clear()
	}
	l.cnt -= 1
	return
}

func (l *List[LT]) PopIndex(number int) (result LT, ok bool) {
	if l.start == nil || number >= l.Len() || (number < 0 && l.Len()+number < 0) {
		return
	}
	// there are no options, only one possible index (border was checked on previous step)
	if l.Len() == 1 {
		n := l.start
		l.start = nil
		l.end = nil
		l.cnt -= 1
		n.Clear()
		return n.data, true
	}

	if number < 0 {
		number = l.Len() - 1 + number
	}
	n := l.start

	var i = 0
	for ; n != nil; n = n.next {
		if number == i {
			prev := n.prev
			if prev != nil {
				prev.next = n.next
			} else {
				l.start = n.next
			}

			next := n.next
			if next != nil {
				next.prev = n.prev
			}

			n.Clear()

			l.cnt -= 1
			return n.data, true
		}
		i += 1
	}
	return
}

func (l *List[LT]) PeakIndex(number int) (result LT, ok bool) {
	if l.start == nil || number >= l.Len() || (number < 0 && l.Len()+number < 0) {
		return
	}
	if number < 0 {
		number = l.Len() - 1 + number
	}
	n := l.start

	var i = 0
	for ; n != nil; n = n.next {
		if number == i {
			return n.data, true
		}
		i += 1
	}
	return
}
