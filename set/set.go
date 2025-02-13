// Package set provides simple set with iterator
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

// Seq returns value-iterator.
func (s *Set[T]) Seq() iter.Seq[T] {
	return func(yield func(T) bool) {
		for k := range s.hash {
			if !yield(k) {
				return
			}
		}
	}
}

// Contains checks is element in the set or not.
func (s *Set[ST]) Contains(element ST) bool {
	_, exists := s.hash[element]
	return exists
}

// Insert adds an element to the set.
func (s *Set[ST]) Insert(elements ...ST) {
	for _, e := range elements {
		s.hash[e] = empty{}
	}
}

// Intersection returns elements which are in both sets.
func (s *Set[ST]) Intersection(set *Set[ST]) *Set[ST] {
	n := make(map[ST]empty)

	for k := range s.hash {
		if _, exists := set.hash[k]; exists {
			n[k] = empty{}
		}
	}

	return &Set[ST]{n}
}

// SymmetricDifference returns unique elements for both sets.
func (s *Set[ST]) SymmetricDifference(other *Set[ST]) (*Set[ST], *Set[ST]) {
	left := New[ST]()
	right := New[ST]()
	for k := range s.Seq() {
		if !other.Contains(k) {
			left.Insert(k)
		}
	}
	for k := range other.Seq() {
		if !s.Contains(k) {
			right.Insert(k)
		}
	}
	return left, right
}

// Difference returns unique elements for that set.
func (s *Set[ST]) Difference(other *Set[ST]) *Set[ST] {
	n := make(map[ST]empty)

	for k := range s.hash {
		if _, exists := other.hash[k]; !exists {
			n[k] = empty{}
		}
	}

	return &Set[ST]{n}
}

// Len returns number of items in the set.
func (s *Set[ST]) Len() int {
	return len(s.hash)
}

// SubsetOf checks is this set a subset of "other".
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

// ProperSubsetOf checks is this set a proper subset of "other".
func (s *Set[ST]) ProperSubsetOf(other *Set[ST]) bool {
	if s.Len() >= other.Len() {
		return false
	}
	return s.SubsetOf(other)
}

// Remove removes all element from the set. Missing elements do nothing.
func (s *Set[ST]) Remove(elements ...ST) {
	for _, element := range elements {
		delete(s.hash, element)
	}
}

// Union returns all elements of both sets.
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

// Clone returns a new set with same elements.
func (s *Set[ST]) Clone() *Set[ST] {
	return &Set[ST]{
		hash: maps.Clone(s.hash),
	}
}

// Equal return true if both sets have same elements.
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
