package set

import (
	"github.com/stretchr/testify/suite"
	"slices"
	"testing"
)

func addTestValues(m map[int]empty, values ...int) {
	for _, value := range values {
		m[value] = empty{}
	}
}
func newTestMap(elements ...int) map[int]empty {
	m := map[int]empty{}
	for _, i := range elements {
		m[i] = empty{}
	}
	return m
}

type setTestSuite struct {
	suite.Suite
}

func (suite *setTestSuite) TestNewEmpty() {
	s := New[int]()

	suite.Require().NotNil(s)
	suite.Require().Equal(0, len(s.hash))
}

func (suite *setTestSuite) TestNewSimgle() {
	s := New(1)
	expected := newTestMap(1)
	suite.Require().Equal(expected, s.hash)
}

func (suite *setTestSuite) TestNewSeveral() {
	s := New(1, 2, 3)
	expected := newTestMap(1, 2, 3)
	suite.Require().Equal(expected, s.hash)
}

func (suite *setTestSuite) TestInsert() {
	s := New[int]()
	expected := newTestMap()

	s.Insert(1)
	addTestValues(expected, 1)
	suite.Require().Equal(expected, s.hash)

	s.Insert(1)
	suite.Require().Equal(expected, s.hash)

	s.Insert(1, 2, 3, 4)
	addTestValues(expected, 2, 3, 4)
	suite.Require().Equal(expected, s.hash)

	s.Insert(1, 2, 3, 4)
	suite.Require().Equal(expected, s.hash)
}

func (suite *setTestSuite) TestRemove() {
	s := New(1, 2, 3)

	s.Remove(2)
	suite.Require().Equal(newTestMap(1, 3), s.hash)

	s.Remove(2)
	suite.Require().Equal(newTestMap(1, 3), s.hash)

	s.Remove(1, 3)
	suite.Require().Equal(newTestMap(), s.hash)

	s.Insert(3, 4)
	s.Remove(1, 3, 5, 4)
	suite.Require().Equal(newTestMap(), s.hash)
}

func (suite *setTestSuite) TestEmptyIntersection() {
	left := New[int]()
	right := New[int]()

	suite.Require().Equal(newTestMap(), left.Intersection(right).hash)

	left = New[int]()
	right = New(5)
	suite.Require().Equal(newTestMap(), left.Intersection(right).hash)

	left = New(1)
	right = New[int]()
	suite.Require().Equal(newTestMap(), left.Intersection(right).hash)

	left = New(1)
	right = New(2)
	suite.Require().Equal(newTestMap(), left.Intersection(right).hash)
}

func (suite *setTestSuite) TestIntersection() {
	left := New(1, 2, 3)
	right := New(2, 3, 4, 5)

	suite.Require().Equal(newTestMap(2, 3), left.Intersection(right).hash)
}

func (suite *setTestSuite) TestEmptyDifference() {
	left := New[int]()
	right := New[int]()

	suite.Require().Equal(newTestMap(), left.Difference(right).hash)

	left = New[int]()
	right = New(5)
	suite.Require().Equal(newTestMap(), left.Difference(right).hash)

	left = New(1)
	right = New(1)
	suite.Require().Equal(newTestMap(), left.Difference(right).hash)
}

func (suite *setTestSuite) TestDifference() {
	left := New(1, 2, 3)
	right := New(2, 3, 4, 5)

	suite.Require().Equal(newTestMap(1), left.Difference(right).hash)
	suite.Require().Equal(newTestMap(4, 5), right.Difference(left).hash)

	left = New(1)
	right = New[int]()
	suite.Require().Equal(newTestMap(1), left.Difference(right).hash)
}

func (suite *setTestSuite) TestUnion() {
	left := New[int]()
	right := New[int]()
	suite.Require().Equal(newTestMap(), left.Union(right).hash)

	left = New(1)
	right = New[int]()
	suite.Require().Equal(newTestMap(1), left.Union(right).hash)

	left = New[int]()
	right = New(1)
	suite.Require().Equal(newTestMap(1), left.Union(right).hash)

	left = New(1, 2, 4)
	right = New(2, 3, 4, 5)
	suite.Require().Equal(newTestMap(1, 2, 3, 4, 5), left.Union(right).hash)
}

func (suite *setTestSuite) TestClone() {
	s := New(1, 2, 3)
	clone := s.Clone()
	suite.Equal(s.hash, clone.hash)

	clone.Insert(4)
	s.Insert(5)
	suite.NotEqual(s.hash, clone.hash)
}

func (suite *setTestSuite) TestSubsetOf() {
	left := New[int]()
	right := New(1, 2, 3)
	suite.Require().True(left.SubsetOf(right))
	suite.Require().False(right.SubsetOf(left))
	suite.Require().True(left.SubsetOf(left))
	suite.Require().True(right.SubsetOf(right))

	left = New(2, 3)
	right = New(1, 2, 3)
	suite.Require().True(left.SubsetOf(right))
	suite.Require().False(right.SubsetOf(left))
	suite.Require().True(left.SubsetOf(left))

	left = New[int](1, 2, 3)
	right = New(1, 2, 3)
	suite.Require().True(left.SubsetOf(right))
	suite.Require().True(right.SubsetOf(left))
	suite.Require().True(left.SubsetOf(left))
	suite.Require().True(right.SubsetOf(right))
}

func (suite *setTestSuite) TestProperSubsetOf() {
	left := New[int]()
	right := New(1, 2, 3)
	suite.Require().True(left.ProperSubsetOf(right))
	suite.Require().False(right.ProperSubsetOf(left))
	suite.Require().False(left.ProperSubsetOf(left))
	suite.Require().False(right.ProperSubsetOf(right))

	left = New(2, 3)
	right = New(1, 2, 3)
	suite.Require().True(left.ProperSubsetOf(right))
	suite.Require().False(right.ProperSubsetOf(left))
	suite.Require().False(left.ProperSubsetOf(left))
	suite.Require().False(right.ProperSubsetOf(right))

	left = New[int](1, 2, 3)
	right = New(1, 2, 3)
	suite.Require().False(left.ProperSubsetOf(right))
	suite.Require().False(right.ProperSubsetOf(left))
	suite.Require().False(left.ProperSubsetOf(left))
	suite.Require().False(right.ProperSubsetOf(right))
}

func (suite *setTestSuite) TestEmptySymmetricDifference() {
	left := New[int]()
	right := New[int]()

	ld, rd := left.SymmetricDifference(right)
	suite.Equal(newTestMap(), ld.hash)
	suite.Equal(newTestMap(), rd.hash)

	left = New(1, 2, 3)
	right = New[int](1, 2, 3)
	ld, rd = left.SymmetricDifference(right)
	suite.Equal(newTestMap(), ld.hash)
	suite.Equal(newTestMap(), rd.hash)
	ld, rd = right.SymmetricDifference(left)
	suite.Equal(newTestMap(), ld.hash)
	suite.Equal(newTestMap(), rd.hash)
}
func (suite *setTestSuite) TestSymmetricDifference() {
	left := New(1, 2, 3)
	right := New[int]()
	ld, rd := left.SymmetricDifference(right)
	suite.Equal(newTestMap(1, 2, 3), ld.hash)
	suite.Equal(newTestMap(), rd.hash)
	ld, rd = right.SymmetricDifference(left)
	suite.Equal(newTestMap(), ld.hash)
	suite.Equal(newTestMap(1, 2, 3), rd.hash)

	left = New(1, 2, 3)
	right = New[int](2, 3, 4, 5)
	ld, rd = left.SymmetricDifference(right)
	suite.Equal(newTestMap(1), ld.hash)
	suite.Equal(newTestMap(4, 5), rd.hash)
	ld, rd = right.SymmetricDifference(left)
	suite.Equal(newTestMap(4, 5), ld.hash)
	suite.Equal(newTestMap(1), rd.hash)
}

func (suite *setTestSuite) TestRange() {
	expected := []int{1, 2, 3, 4, 5, 6}
	set := New[int](expected...)

	repeatCheck := make([]int, 0, len(expected))
	for v := range set.Range() {
		suite.Require().NotContains(repeatCheck, v)
		suite.Require().Contains(expected, v)
		repeatCheck = append(repeatCheck, v)
	}
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func (suite *setTestSuite) TestToSlice() {
	suite.Require().Equal([]int{}, New[int]().ToSlice())

	expected := []int{1, 2, 3, 4, 5, 6}
	repeatCheck := make([]int, 0, len(expected))
	dat := New(expected...).ToSlice()
	for _, v := range dat {
		suite.Require().NotContains(repeatCheck, v)
		suite.Require().Contains(expected, v)
		expected = remove(expected, slices.Index(expected, v))
	}
	suite.Require().Empty(expected)
}

func TestSetTestSuite(t *testing.T) {
	suite.Run(t, new(setTestSuite))
}

func (suite *setTestSuite) TestEqual() {
	suite.Require().False(New[int]() == New(1))
	suite.Require().False(New(1) == New[int]())

	set := New[int]()
	suite.Require().True(set.Equal(set))
	suite.Require().True(New[int]().Equal(New[int]()))
	suite.Require().True(New[int](1, 2).Equal(New[int](1, 2)))
	suite.Require().True(New[int](1, 3, 2).Equal(New[int](1, 2, 3)))

	suite.Require().False(New[int](1, 2).Equal(New[int](1, 2, 3)))
	suite.Require().False(New[int](1, 2, 3).Equal(New[int](1, 2)))

}
