// Package trie provides unicode trie structure
// could hold this one  (🦃) --> (🌯)
//
//	└ --> (🍖)
package trie

import (
	"github.com/HoskeOwl/ggstruct/stack"
	"maps"
	"slices"
)

type (
	Trie[V any] struct {
		root     *node[V]
		valueLen int
		nodesLen int
	}
	node[V any] struct {
		key   rune
		value *V
		leafs map[rune]*node[V]
		//weights map[rune]int|float
		prev *node[V]
	}
)

func (n *node[V]) Clear() {
	n.value = nil
	n.prev = nil
	n.leafs = nil
}

func newNode[V any](key rune, prev *node[V]) *node[V] {
	return &node[V]{
		key:   key,
		value: nil,
		leafs: make(map[rune]*node[V]),
		prev:  prev,
	}
}

func NewFromMap[V any](values map[string]V) *Trie[V] {
	t := &Trie[V]{
		root:     newNode[V](rune(0), nil),
		valueLen: 0,
		nodesLen: 0,
	}
	for k, v := range values {
		t.Insert(k, v)
	}
	return t
}

func New[V any]() *Trie[V] {
	return &Trie[V]{
		root:     newNode[V](rune(0), nil),
		valueLen: 0,
		nodesLen: 0,
	}
}

func (t *Trie[V]) getNode(key string) *node[V] {
	cur := t.root
	for _, r := range key {
		if n, e := cur.leafs[r]; e {
			cur = n
		} else {
			return nil
		}
	}
	return cur
}

// Insert adds element to the three.
func (t *Trie[V]) Insert(key string, value V) {
	cur := t.root
	for _, r := range key {
		if n, exists := cur.leafs[r]; exists {
			cur = n
		} else {
			n = newNode(r, cur)
			cur.leafs[r] = n
			cur = n
			t.nodesLen++
		}
	}
	cur.value = &value
	t.valueLen++
}

// Search finds a node with value. If there is no node returns default value and 'false'.
func (t *Trie[V]) Search(key string) (res V, exists bool) {
	n := t.getNode(key)
	if n == nil {
		exists = false
		return
	}
	if n.value == nil {
		exists = false
		return
	}
	return *n.value, true
}

// Has returns 'true' if the three has a node with not nil value.
func (t *Trie[V]) Has(key string) bool {
	_, exists := t.Search(key)
	return exists
}

// Len returns number of values.
func (t *Trie[V]) Len() int {
	return t.valueLen
}

// NodeLen returns number of nodes (with value and without).
func (t *Trie[V]) NodeLen() int {
	return t.nodesLen
}

// Remove removes node with value if exists. Also remove all empty upper nodes.
func (t *Trie[V]) Remove(key string) (res V, exists bool) {
	if t.valueLen == 0 {
		exists = false
		return
	}

	cur := t.getNode(key)
	if cur == nil || cur.value == nil {
		exists = false
		return
	}

	res = *cur.value
	cur.value = nil
	t.valueLen--
	exists = true

	if len(cur.leafs) > 0 {
		// keep for leafs
		t.valueLen--
		return
	}

	for {
		rm := cur
		cur = cur.prev
		delete(cur.leafs, rm.key)
		t.nodesLen--
		rm.Clear()
		if cur == t.root || len(cur.leafs) != 0 || cur.value != nil {
			return
		}
	}
}

// Clear removes all elements.
func (t *Trie[V]) Clear() {
	if t.valueLen == 0 {
		return
	}

	s := stack.New[*node[V]](slices.Collect(maps.Values(t.root.leafs))...)
	for !s.IsEmpty() {
		n, exists := s.Pop()
		if !exists {
			return
		}
		for _, v := range n.leafs {
			s.Push(v)
		}
		n.Clear()
	}
	t.valueLen = 0
	t.nodesLen = 0
}
