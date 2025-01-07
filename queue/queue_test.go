package queue

import (
	"testing"
)

func Test(t *testing.T) {
	q := New[int]()
	var check int
	var exists bool

	if q.Len() != 0 {
		t.Errorf("Length should be 0")
	}

	q.Enqueue(1)
	if q.Len() != 1 {
		t.Errorf("Length should be 1")
	}

	check, exists = q.Peek()
	if !exists {
		t.Errorf("Top element shold exists")
	}
	if check != 1 {
		t.Errorf("Enqueued value should be 1")
	}

	check, exists = q.Dequeue()
	if !exists {
		t.Errorf("Top element shold exists")
	}
	if check != 1 {
		t.Errorf("Dequeued value should be 1")
	}

	_, exists = q.Peek()
	if exists {
		t.Errorf("Queue have to be empty")
	}
	_, exists = q.Dequeue()
	if exists {
		t.Errorf("Queue have to be empty")
	}

	q.Enqueue(1)
	q.Enqueue(2)

	if q.Len() != 2 {
		t.Errorf("Length should be 2")
	}

	if c, _ := q.Peek(); c != 1 {
		t.Errorf("First value should be 1")
	}

	q.Dequeue()

	if c, _ := q.Peek(); c != 2 {
		t.Errorf("First value should be 2")
	}
}
