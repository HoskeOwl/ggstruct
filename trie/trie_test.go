package trie

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestTrie(t *testing.T) {
	suite.Run(t, new(TrieStructTestSuite))
}

type TrieStructTestSuite struct {
	suite.Suite
}

func (s *TrieStructTestSuite) TestLen() {
	t := New[int]()
	s.Equal(0, t.Len())
	s.Equal(0, t.NodeLen())
	t.Insert("35", 100)
	s.Equal(1, t.Len())
	s.Equal(2, t.NodeLen())
	t.Insert("368", 120)
	s.Equal(2, t.Len())
	s.Equal(4, t.NodeLen())
	t.Insert("3", 2)
	s.Equal(3, t.Len())
	s.Equal(4, t.NodeLen())

	t.Remove("368")
	s.Equal(2, t.Len())
	s.Equal(2, t.NodeLen())
}

func (s *TrieStructTestSuite) TestNewFromMap() {
	values := map[string]int{
		"ab":  12,
		"adc": 13,
		"ad":  14,
	}
	t := NewFromMap(values)

	n, exists := t.root.leafs['a']
	s.True(exists)
	s.Nil(n.value)

	l, exists := n.leafs['b']
	s.True(exists)
	s.Require().NotNil(l.value)
	s.Equal(12, *l.value)

	l, exists = n.leafs['d']
	s.True(exists)
	s.Require().NotNil(l.value)
	s.Equal(14, *l.value)

	l, exists = l.leafs['c']
	s.True(exists)
	s.Require().NotNil(l.value)
	s.Equal(13, *l.value)
}

func (s *TrieStructTestSuite) TestInsert() {
	t := New[int]()
	t.Insert("🍿🌮🍖", 12)

	a, exists := t.root.leafs['🍿']
	s.Require().True(exists)
	s.Require().Nil(a.value)

	b, exists := a.leafs['🌮']
	s.Require().True(exists)
	s.Require().Nil(b.value)

	c, exists := b.leafs['🍖']
	s.Require().True(exists)
	s.Require().NotNil(c.value)
	s.Equal(12, *c.value)

	t.Insert("🍿🌮", 25)
	s.Require().NotNil(b.value)
	s.Equal(25, *b.value)
}

func (s *TrieStructTestSuite) TestRemove() {
	values := map[string]int{
		"ab":  12,
		"adc": 13,
		"adf": 14,
	}
	t := NewFromMap(values)

	// prepare
	a, exists := t.root.leafs['a']
	s.Require().True(exists)
	ab, exists := a.leafs['b']
	s.Require().True(exists)
	ad, exists := a.leafs['d']
	s.Require().True(exists)
	adc, exists := ad.leafs['c']
	s.Require().True(exists)
	adf, exists := ad.leafs['f']
	s.Require().True(exists)
	s.Equal(5, t.NodeLen())
	s.Equal(3, t.Len())

	// start test
	t.Remove("ab")
	_, exists = a.leafs['b']
	s.False(exists)
	s.Require().Nil(ab.value)
	s.Equal(4, t.NodeLen())
	s.Equal(2, t.Len())

	t.Remove("adc")
	_, exists = ad.leafs['c']
	s.Require().Nil(adc.value)
	s.Equal(3, t.NodeLen())
	s.Equal(1, t.Len())

	t.Remove("adf")
	_, exists = ad.leafs['f']
	s.Require().Nil(adf.value)
	s.Equal(0, t.NodeLen())
	s.Equal(0, t.Len())

	s.Equal(0, len(t.root.leafs))
}

func (s *TrieStructTestSuite) TestSearchHas() {
	values := map[string]int{
		"ab":  12,
		"adc": 13,
		"adf": 14,
	}
	t := NewFromMap(values)
	_, exists := t.Search("qwerty")
	s.False(exists)
	s.False(t.Has("qwerty"))

	_, exists = t.Search("a")
	s.False(exists)
	s.False(t.Has("a"))

	_, exists = t.Search("adc")
	s.True(exists)
	s.True(t.Has("adc"))
}

func (s *TrieStructTestSuite) TestClear() {
	values := map[string]int{
		"ab":  12,
		"adc": 13,
		"adf": 14,
	}
	t := NewFromMap(values)

	// prepare
	a, exists := t.root.leafs['a']
	s.Require().True(exists)
	ab, exists := a.leafs['b']
	s.Require().True(exists)
	ad, exists := a.leafs['d']
	s.Require().True(exists)
	adc, exists := ad.leafs['c']
	s.Require().True(exists)
	adf, exists := ad.leafs['f']
	s.Require().True(exists)
	s.Equal(5, t.NodeLen())
	s.Equal(3, t.Len())

	// start test
	t.Clear()
	s.Equal(0, t.NodeLen())
	s.Equal(0, t.Len())
	s.Empty(a.leafs)
	s.Nil(a.value)
	s.Empty(ab.leafs)
	s.Nil(ab.value)
	s.Empty(adc.leafs)
	s.Nil(adc.value)
	s.Empty(adf.leafs)
	s.Nil(adf.value)
}
