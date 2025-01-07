package trie

import (
	"fmt"

	"github.com/HoskeOwl/ggstruct/stack"
)

type (
	Trie[V any] struct {
		root *node[V]
		size int
	}
	node[V any] struct {
		key   rune
		value *V
		leaf  map[rune]*node[V]
		prev  *node[V]
	}
)

func New[V any]() *Trie[V] {
	return &Trie[V]{nil, 0}
}
func (this *Trie[V]) newNode(key rune, prev *node[V]) *node[V] {
	return &node[V]{key, nil, make(map[rune]*node[V]), prev}
}
func (this *Trie[V]) Get(key string) (res V, exists bool) {
	if this.size == 0 {
		exists = false
		return
	}

	cur := this.root
	for _, r := range key {
		if node, e := cur.leaf[r]; e {
			cur = node
		} else {
			exists = false
		}
	}
	res = *cur.value
	exists = true
	return
}
func (this *Trie[V]) Has(key string) bool {
	_, exists := this.Get(key)
	return exists
}
func (this *Trie[V]) Insert(key string, value V) {
	if this.size == 0 {
		this.root = this.newNode(rune(0), nil)
	}

	cur := this.root
	for _, r := range key {
		if node, exists := cur.leaf[r]; exists {
			cur = node
		} else {
			node = this.newNode(r, cur)
			cur.leaf[r] = node
			cur = node
		}
	}
	if cur.value == nil {
		this.size++
	}
	cur.value = &value
}
func (this *Trie[V]) Len() int {
	return this.size
}
func (this *Trie[V]) Remove(key string) (res V, exists bool) {
	if this.size == 0 {
		exists = false
		return
	}

	cur := this.root

	for _, r := range key {
		if node, e := cur.leaf[r]; e {
			cur = node
		} else {
			exists = false
			return
		}
	}

	if cur.value != nil {
		this.size--
		cur.value = nil
	} else {
		res = *cur.value
	}

	exists = true

	//cleanup
	var prev *node[V]
	var k rune
	for cur != this.root {
		if cur.value == nil {
			prev = cur.prev
			k = cur.key
			cur.prev = nil
			cur.value = nil
			delete(prev.leaf, k)
			cur = prev
		} else {
			return
		}
	}
	return
}
func (this *Trie[V]) String() string {
	str := "{"
	i := 0
	this.Do(func(k string, v V) bool {
		if i > 0 {
			str += ", "
		}
		str += fmt.Sprint(k, ":", v)
		i++
		return true
	})
	str += "}"
	return str
}
func (this *Trie[V]) Do(handler func(string, V) bool) {
	if this.size == 0 {
		return
	}

	var n *node[V]
	nodes := stack.New[*node[V]]()
	nodes.Push(this.root)
	for nodes.Len() > 0 {
		n, _ = nodes.Pop()
		if n.value != nil {
			if !handler(string(n.key), *n.value) {
				return
			}
		}
		for _, ln := range n.leaf {
			nodes.Push(ln)
		}
	}
}
