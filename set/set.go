package set

import (
	"fmt"
	"strings"
)

type (
	Set[ST comparable] struct {
		hash map[ST]empty
	}
	empty struct{}
)

func New[T comparable](initial ...T) *Set[T] {
	s := &Set[T]{make(map[T]empty, len(initial))}

	for _, v := range initial {
		s.Insert(v)
	}

	return s
}

func (this *Set[QT]) String() string {
	if len(this.hash) < 1 {
		return ""
	}
	var keys []string = make([]string, len(this.hash))
	var i int = 0
	for val := range this.hash {
		keys[i] = fmt.Sprint(val)
		i += 1
	}
	return fmt.Sprintf("{%v}", strings.Join(keys, ", "))
}

// Find the difference between two sets
func (this *Set[ST]) Difference(set *Set[ST]) *Set[ST] {
	n := make(map[ST]empty)

	for k := range this.hash {
		if _, exists := set.hash[k]; !exists {
			n[k] = empty{}
		}
	}

	return &Set[ST]{n}
}

// Call f for each item in the set
func (this *Set[ST]) Do(f func(ST)) {
	for k := range this.hash {
		f(k)
	}
}

// Test to see whether or not the element is in the set
func (this *Set[ST]) Has(element ST) bool {
	_, exists := this.hash[element]
	return exists
}

// Add an element to the set
func (this *Set[ST]) Insert(element ST) {
	this.hash[element] = empty{}
}

// Find the intersection of two sets
func (this *Set[ST]) Intersection(set *Set[ST]) *Set[ST] {
	n := make(map[ST]empty)

	for k := range this.hash {
		if _, exists := set.hash[k]; exists {
			n[k] = empty{}
		}
	}

	return &Set[ST]{n}
}

// Return the number of items in the set
func (this *Set[ST]) Len() int {
	return len(this.hash)
}

// Test whether or not this set is a proper subset of "set"
func (this *Set[ST]) ProperSubsetOf(set *Set[ST]) bool {
	return this.Len() < set.Len() && this.SubsetOf(set)
}

// Remove an element from the set
func (this *Set[ST]) Remove(element ST) {
	delete(this.hash, element)
}

// Test whether or not this set is a subset of "set"
func (this *Set[ST]) SubsetOf(set *Set[ST]) bool {
	if this.Len() > set.Len() {
		return false
	}
	for k := range this.hash {
		if _, exists := set.hash[k]; !exists {
			return false
		}
	}
	return true
}

// Find the union of two sets
func (this *Set[ST]) Union(set *Set[ST]) *Set[ST] {
	n := make(map[ST]empty, set.Len()+this.Len())

	for k := range this.hash {
		n[k] = empty{}
	}
	for k := range set.hash {
		n[k] = empty{}
	}

	return &Set[ST]{n}
}
