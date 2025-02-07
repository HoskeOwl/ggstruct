package set

import (
	"iter"
	"maps"
)

type (
	empty              struct{}
	Set[ST comparable] struct {
		hash map[ST]empty
	}
)

func New[T comparable](initial ...T) *Set[T] {
	s := &Set[T]{make(map[T]empty)}

	for _, v := range initial {
		s.Insert(v)
	}

	return s
}

func (s *Set[T]) ToSlice() []T {
	result := make([]T, 0, len(s.hash))
	for k := range s.hash {
		result = append(result, k)
	}
	return result
}

func (s *Set[T]) Range() iter.Seq[T] {
	return func(yield func(T) bool) {
		for k := range s.hash {
			if !yield(k) {
				return
			}
		}
	}
}

// Test to see whether or not the element is in the set
func (s *Set[ST]) Contains(element ST) bool {
	_, exists := s.hash[element]
	return exists
}

// Add an element to the set
func (s *Set[ST]) Insert(elements ...ST) {
	for _, e := range elements {
		s.hash[e] = empty{}
	}
}

// Find the intersection of two sets
func (s *Set[ST]) Intersection(set *Set[ST]) *Set[ST] {
	n := make(map[ST]empty)

	for k := range s.hash {
		if _, exists := set.hash[k]; exists {
			n[k] = empty{}
		}
	}

	return &Set[ST]{n}
}

func (s *Set[ST]) SymmetricDifference(other *Set[ST]) (*Set[ST], *Set[ST]) {
	left := New[ST]()
	right := New[ST]()
	for k := range s.Range() {
		if !other.Contains(k) {
			left.Insert(k)
		}
	}
	for k := range other.Range() {
		if !s.Contains(k) {
			right.Insert(k)
		}
	}
	return left, right
}

// Find the difference between two sets
func (s *Set[ST]) Difference(set *Set[ST]) *Set[ST] {
	n := make(map[ST]empty)

	for k := range s.hash {
		if _, exists := set.hash[k]; !exists {
			n[k] = empty{}
		}
	}

	return &Set[ST]{n}
}

// Return the number of items in the set
func (s *Set[ST]) Len() int {
	return len(s.hash)
}

// Test whether or not this set is a subset of "set"
func (s *Set[ST]) SubsetOf(other *Set[ST]) bool {
	if s.Len() == 0 || (other.Len() == 0 && s.Len() == 0) {
		return true
	}
	if s.Len() > other.Len() {
		return false
	}

	for k := range s.hash {
		if _, exists := other.hash[k]; !exists {
			return false
		}
	}
	return true
}

// Test whether or not this set is a proper subset of "set"
func (s *Set[ST]) ProperSubsetOf(other *Set[ST]) bool {
	if s.Len() >= other.Len() {
		return false
	}
	return s.SubsetOf(other)
}

// Remove an element from the set
func (s *Set[ST]) Remove(elements ...ST) {
	for _, element := range elements {
		delete(s.hash, element)
	}
}

// Find the union of two sets
func (s *Set[ST]) Union(set *Set[ST]) *Set[ST] {
	n := make(map[ST]empty, set.Len()+s.Len())

	for k := range s.hash {
		n[k] = empty{}
	}
	for k := range set.hash {
		n[k] = empty{}
	}

	return &Set[ST]{n}
}

func (s *Set[ST]) Clone() *Set[ST] {
	return &Set[ST]{
		hash: maps.Clone(s.hash),
	}
}

func (s *Set[ST]) Equal(other *Set[ST]) bool {
	if s.Len() != other.Len() {
		return false
	}
	for k := range s.hash {
		if !other.Contains(k) {
			return false
		}
	}
	return true
}
