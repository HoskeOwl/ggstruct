package trie

import (
	"testing"
)

func Test(t *testing.T) {
	x := New[int]()
	x.Insert("35", 100)
	if x.Len() != 1 {
		t.Errorf("expected len 1")
	}
	var check int
	var exists bool
	check, exists = x.Get("35")
	if !exists {
		t.Errorf("Key should exists")
	}
	if check != 100 {
		t.Errorf("expected to get 100 for 1")
	}
	x.Remove("35")
	if x.Len() != 0 {
		t.Errorf("expected len 0")
	}
	if len(x.root.leaf) != 0 {
		t.Errorf("empty nodes did not remove")
	}

	x.Insert("ac", 200)
	x.Insert("afg", 100)

	if x.Len() != 2 {
		t.Errorf("expected len 1")
	}

	check, exists = x.Get("afg")
	if !exists {
		t.Errorf("Key should exists")
	}
	if check != 100 {
		t.Errorf("expected to get 100 for 1")
	}
	check, exists = x.Get("ac")
	if !exists {
		t.Errorf("Key should exists")
	}
	if check != 200 {
		t.Errorf("expected to get 100 for 1")
	}

}
