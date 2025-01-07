package stack

import (
	"testing"
)

func TestPushRight(t *testing.T) {
	var check int
	var exists bool
	s := New[int]()

	if s.Len() != 0 {
		t.Errorf("Length of an empty stack should be 0")
	}

	s.Push(1)

	if s.Len() != 1 {
		t.Errorf("Length should be 0")
	}

	check, exists = s.Peek()
	if !exists {
		t.Errorf("Top item must exists")
	}
	if check != 1 {
		t.Errorf("Top item on the stack should be 1")
	}

	check, exists = s.Pop()
	if !exists {
		t.Errorf("Top item must exists")
	}
	if check != 1 {
		t.Errorf("Top item on the stack should be 1")
	}

	if s.Len() != 0 {
		t.Errorf("Stack should be empty")
	}

	s.Push(1)
	s.Push(2)

	if s.Len() != 2 {
		t.Errorf("Length should be 2")
	}

	check, exists = s.Peek()
	if !exists {
		t.Errorf("Top item must exists")
	}
	if check != 2 {
		t.Errorf("Top item on the stack should be 2")
	}

	check, exists = s.Pop()
	if !exists {
		t.Errorf("Top item must exists")
	}
	if check != 2 {
		t.Errorf("Top item on the stack should be 2")
	}
	check, exists = s.Pop()
	if !exists {
		t.Errorf("Top item must exists")
	}
	if check != 1 {
		t.Errorf("Top item on the stack should be 1")
	}
	if s.Len() != 0 {
		t.Errorf("Stack should be empty")
	}
}

func TestPushLeft(t *testing.T) {
	var check int
	var exists bool
	s := New[int]()

	if s.Len() != 0 {
		t.Errorf("Length of an empty stack should be 0")
	}

	s.PushLeft(1)

	if s.Len() != 1 {
		t.Errorf("Length should be 0")
	}

	check, exists = s.Peek()
	if !exists {
		t.Errorf("Top item must exists")
	}
	if check != 1 {
		t.Errorf("Top item on the stack should be 1")
	}

	check, exists = s.PopLeft()
	if !exists {
		t.Errorf("Top item must exists")
	}
	if check != 1 {
		t.Errorf("Top item on the stack should be 1")
	}

	if s.Len() != 0 {
		t.Errorf("Stack should be empty")
	}

	s.PushLeft(1)
	s.PushLeft(2)

	if s.Len() != 2 {
		t.Errorf("Length should be 2")
	}

	check, exists = s.PeekLeft()
	if !exists {
		t.Errorf("Top item must exists")
	}
	if check != 2 {
		t.Errorf("Top item on the stack should be 2")
	}

	check, exists = s.PopLeft()
	if !exists {
		t.Errorf("Top item must exists")
	}
	if check != 2 {
		t.Errorf("Top item on the stack should be 2")
	}
	check, exists = s.PopLeft()
	if !exists {
		t.Errorf("Top item must exists")
	}
	if check != 1 {
		t.Errorf("Top item on the stack should be 1")
	}
	if s.Len() != 0 {
		t.Errorf("Stack should be empty")
	}
}

func TestPushBoth(t *testing.T) {
	var check int
	var exists bool
	s := New[int]()

	s.PushLeft(1)
	check, exists = s.Peek()
	if !exists {
		t.Errorf("Top item must exists")
	}
	if check != 1 {
		t.Errorf("Top item on the stack should be 1")
	}

	s.PushLeft(2)
	check, exists = s.Peek()
	if !exists {
		t.Errorf("Top item must exists")
	}
	if check != 1 {
		t.Errorf("Top item on the stack should be 1")
	}

	check, exists = s.Pop()
	if !exists {
		t.Errorf("Top item must exists")
	}
	if check != 1 {
		t.Errorf("Top item on the stack should be 1")
	}
}
