package list

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type ListTestSuite struct {
	suite.Suite
}

func TestRunListSuite(t *testing.T) {
	suite.Run(t, new(ListTestSuite))
}

func (l *ListTestSuite) TestCheckInit() {
	var exp = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	lst := New[int](exp...)

	l.Require().Equal(exp, lst.AsSlice())
	l.Require().Equal(len(exp), lst.Len())
}

func (l *ListTestSuite) TestCheckAddRight() {
	var exp = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	lst := New[int]()
	for _, n := range exp {
		lst.AddRight(n)
	}

	l.Require().Equal(exp, lst.AsSlice())
	l.Require().Equal(len(exp), lst.Len())
}

func (l *ListTestSuite) TestCheckAddLeft() {
	var exp = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	var dat = []int{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	lst := New[int]()
	for _, n := range dat {
		lst.AddLeft(n)
	}

	l.Require().Equal(exp, lst.AsSlice())
	l.Require().Equal(len(exp), lst.Len())
}

func (l *ListTestSuite) TestCheckPopLeft() {
	var exp = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	lst := New[int](exp...)

	res := make([]int, 0, len(exp))
	var val int
	exists := true

	for exists {
		val, exists = lst.PopLeft()
		if exists {
			res = append(res, val)
		}
	}

	l.Require().Equal(exp, res)
	l.Require().Equal(0, lst.Len())
}

func (l *ListTestSuite) TestCheckPopRight() {
	var exp = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	var dat = []int{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	lst := New[int](dat...)

	res := make([]int, 0, len(exp))
	var val int
	exists := true

	for exists {
		val, exists = lst.PopRight()
		if exists {
			res = append(res, val)
		}
	}

	l.Require().Equal(exp, res)
	l.Require().Equal(0, lst.Len())
}

func (l *ListTestSuite) TestCheckGetMissing() {
	var exp = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	lst := New[int](exp...)

	_, ok := lst.PeakIndex(100)

	l.Require().Equal(false, ok)
}

func (l *ListTestSuite) TestCheckGetFirst() {
	var exp = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	lst := New[int](exp...)

	val, ok := lst.PeakIndex(0)

	l.Require().Equal(true, ok)
	l.Require().Equal(1, val)
}

func (l *ListTestSuite) TestCheckGetLast() {
	var dat = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	lst := New[int](dat...)

	val, ok := lst.PeakIndex(14)

	l.Require().Equal(true, ok)
	l.Require().Equal(15, val)
}

func (l *ListTestSuite) TestCheckGetMiddle() {
	var dat = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	lst := New[int](dat...)

	val, ok := lst.PeakIndex(7)

	l.Require().Equal(true, ok)
	l.Require().Equal(8, val)
}

func (l *ListTestSuite) TestIterator() {
	var exp = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	dat := make([]int, 0, len(exp))

	yield := func(val int) bool {
		dat = append(dat, val)
		return false
	}

	lst := New[int](exp...)

	it := lst.Iterator()
	var val int
	var ok = true
	for ok {
		if val, ok = it.Next(); ok {
			dat = append(dat, val)
		}
	}

	l.Require().Equal(exp, dat)
}
