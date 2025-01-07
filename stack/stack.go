package stack

import (
	"fmt"
	"strings"
)

type (
	Stack[ST any] struct {
		left   *node[ST]
		right  *node[ST]
		length int
	}
	node[NT any] struct {
		value NT
		prev  *node[NT]
		next  *node[NT]
	}
)

func (this *Stack[QT]) String() string {
	if this.length < 1 {
		return ""
	}
	var keys []string = make([]string, this.length)
	n := this.left
	var i int = 0
	for n != nil {
		keys[i] = fmt.Sprint(n.value)
		i += 1
		n = n.next
		if n == nil {
			break
		}
	}
	return fmt.Sprintf("[%v]", strings.Join(keys, ", "))
}

// Create a new stack
func New[T any]() *Stack[T] {
	return &Stack[T]{nil, nil, 0}
}

// Return the number of items in the stack
func (this *Stack[ST]) Len() int {
	return this.length
}

// View the right(top) item on the stack
func (this *Stack[ST]) Peek() (res ST, exists bool) {
	if this.length == 0 {
		exists = false
		return
	}
	return this.right.value, true
}

// View the right(top) item on the stack
func (this *Stack[ST]) PeekRight() (res ST, exists bool) {
	return this.Peek()
}

// View the bottom(left) item on the stack
func (this *Stack[ST]) PeekLeft() (res ST, exists bool) {
	if this.length == 0 {
		exists = false
		return
	}
	return this.left.value, true
}

// Pop the right(top) item of the stack and return it
func (this *Stack[ST]) Pop() (res ST, exists bool) {
	if this.length == 0 {
		exists = false
		return
	}

	n := this.right
	this.right = n.prev
	this.right.next = nil
	this.length--
	if this.length == 0 {
		this.left = nil
	}
	return n.value, true
}

// Pop the right(top) item of the stack and return it
func (this *Stack[ST]) PopRight() (res ST, exists bool) {
	return this.Pop()
}

// Pop the left (bottom) item of the stack and return it
func (this *Stack[ST]) PopLeft() (res ST, exists bool) {
	if this.length == 0 {
		exists = false
		return
	}

	n := this.left
	this.left = n.next
	this.right.prev = nil
	this.length--
	if this.length == 0 {
		this.right = nil
	}
	return n.value, true
}

// Push a value onto the right(top) of the stack
func (this *Stack[ST]) Push(value ST) {
	n := &node[ST]{value, this.right, nil}
	if this.length == 0 {
		this.left = n
	} else {
		this.right.next = n
	}
	this.right = n
	this.length++
}

// Push a value onto the right(top) of the stack
func (this *Stack[ST]) PushRight(value ST) {
	this.Push(value)
}

// Push a value onto the left(bottom) of the stack
func (this *Stack[ST]) PushLeft(value ST) {
	n := &node[ST]{value, nil, this.left}
	if this.length == 0 {
		this.right = n
	} else {
		this.left.prev = n
	}
	this.left = n
	this.length++
}
