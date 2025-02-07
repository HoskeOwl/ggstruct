package list

import (
	"github.com/stretchr/testify/suite"
	"slices"
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

	res := slices.Collect(lst.Seq())
	l.Require().Equal(exp, res)
	l.Require().Equal(len(exp), lst.Len())
}

func (l *ListTestSuite) TestCheckPushFrontOneByOne() {
	var exp = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	var dat = []int{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	lst := New[int]()
	for _, n := range dat {
		lst.PushFront(n)
	}
	res := slices.Collect(lst.Seq())
	l.Require().Equal(exp, res)
	l.Require().Equal(len(exp), lst.Len())
}

func (l *ListTestSuite) TestCheckPushFront() {
	var exp = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	lst := New[int]()
	lst.PushFront(exp...)
	res := slices.Collect(lst.Seq())
	l.Require().Equal(exp, res)
	l.Require().Equal(len(exp), lst.Len())
}

func (l *ListTestSuite) TestCheckPushBackOneByOne() {
	var exp = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	lst := New[int]()
	for _, n := range exp {
		lst.PushBack(n)
	}
	res := slices.Collect(lst.Seq())
	l.Require().Equal(exp, res)
	l.Require().Equal(len(exp), lst.Len())
}

func (l *ListTestSuite) TestCheckPushBack() {
	var exp = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	lst := New[int]()
	lst.PushBack(exp...)
	res := slices.Collect(lst.Seq())
	l.Require().Equal(exp, res)
	l.Require().Equal(len(exp), lst.Len())
}

func (l *ListTestSuite) TestCheckPopFront() {
	var exp = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	lst := New[int](exp...)

	res := make([]int, 0, len(exp))
	var val int
	exists := true

	for exists {
		val, exists = lst.PopFront()
		if exists {
			res = append(res, val)
		}
	}

	l.Require().Equal(exp, res)
	l.Require().Equal(0, lst.Len())
}

func (l *ListTestSuite) TestCheckPopBack() {
	var exp = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	var dat = []int{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	lst := New[int](dat...)

	res := make([]int, 0, len(exp))
	var val int
	exists := true

	for exists {
		val, exists = lst.PopBack()
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

	_, ok := lst.PeakAt(100)

	l.Require().Equal(false, ok)
}

func (l *ListTestSuite) TestCheckGetFirst() {
	var exp = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	lst := New[int](exp...)

	val, ok := lst.PeakAt(0)

	l.Require().Equal(true, ok)
	l.Require().Equal(1, val)
}

func (l *ListTestSuite) TestCheckGetLast() {
	var dat = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	lst := New[int](dat...)

	val, ok := lst.PeakAt(14)

	l.Require().Equal(true, ok)
	l.Require().Equal(15, val)
}

func (l *ListTestSuite) TestCheckGetMiddle() {
	var dat = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	lst := New[int](dat...)

	val, ok := lst.PeakAt(7)

	l.Require().Equal(true, ok)
	l.Require().Equal(8, val)
}

func (l *ListTestSuite) TestIterator() {
	var exp = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	lst := New[int](exp...)
	for i, v := range lst.Seq2() {
		l.Require().Equal(exp[i], v)
	}

	res := make([]int, 0, len(exp))
	for v := range lst.Seq() {
		res = append(res, v)
	}
	l.Require().Equal(exp, res)
}

func (l *ListTestSuite) TestIteratorReversed() {
	var dat = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	var exp = []int{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	lst := New[int](dat...)
	for i, v := range lst.ReversedSeq2() {
		l.Require().Equal(exp[i], v)
	}

	res := make([]int, 0, len(exp))
	for v := range lst.ReversedSeq() {
		res = append(res, v)
	}
	l.Require().Equal(exp, res)
}

func (l *ListTestSuite) TestEmptyIterator() {
	lst := New[int]()
	for _ = range lst.Seq() {
		l.Require().True(false, "Should not be reached")
	}
}

func (l *ListTestSuite) TestEmptyIteratorReversed() {
	lst := New[int]()
	for _ = range lst.ReversedSeq2() {
		l.Require().True(false, "Should not be reached")
	}
}

func (l *ListTestSuite) TestIndex() {
	var dat = []int{1, 2, 3, 4, 5, 5, 5, 8, 9, 10, 11, 12, 13, 14, 15}

	lst := New[int](dat...)

	for _, v := range dat {
		l.Require().Equal(v-1, lst.Index(v))
	}

	l.Require().Equal(-1, lst.Index(-12))
	l.Require().Equal(-1, lst.Index(123))
}

func (l *ListTestSuite) TestEmptyIndex() {
	var dat = []int{1, 2, 3, 4, 5, 5, 5, 8, 9, 10, 11, 12, 13, 14, 15}

	lst := New[int]()

	for _, v := range dat {
		l.Require().Equal(-1, lst.Index(v))
	}

	l.Require().Equal(-1, lst.Index(-12))
	l.Require().Equal(-1, lst.Index(123))
}

func (l *ListTestSuite) TestContains() {
	var dat = []int{1, 2, 3, 4, 5, 5, 5, 8, 9, 10, 11, 12, 13, 14, 15}

	lst := New[int](dat...)

	for _, v := range dat {
		l.Require().True(lst.Contains(v))
	}

	l.Require().False(lst.Contains(-12))
	l.Require().False(lst.Contains(123))
}

func (l *ListTestSuite) TestEmptyContains() {
	var dat = []int{1, 2, 3, 4, 5, 5, 5, 8, 9, 10, 11, 12, 13, 14, 15}

	lst := New[int]()

	for _, v := range dat {
		l.Require().False(lst.Contains(v))
	}

	l.Require().False(lst.Contains(-12))
	l.Require().False(lst.Contains(123))
}

func (l *ListTestSuite) TestNoAddAfterIndex() {
	var dat = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	var lst *List[int]

	lst = New[int](dat...)
	l.Require().False(lst.AddAfterIndex(-1, 64))
	res := slices.Collect(lst.Seq())
	l.Require().Equal(dat, res)

	l.Require().False(lst.AddAfterIndex(128, 64))

	lst = New[int]()
	l.Require().False(lst.AddAfterIndex(0, 64))
}

func (l *ListTestSuite) TestAddAfterIndex() {
	var dat = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	var exp []int
	var lst *List[int]

	lst = New[int](dat...)
	exp = []int{0, 71, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	l.Require().True(lst.AddAfterIndex(0, 71))
	res := slices.Collect(lst.Seq())
	l.Require().Equal(exp, res)

	lst = New[int](dat...)
	exp = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 71}
	l.Require().True(lst.AddAfterIndex(15, 71))
	res = slices.Collect(lst.Seq())
	l.Require().Equal(exp, res)

	lst = New[int](dat...)
	exp = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 71, 10, 11, 12, 13, 14, 15}
	l.Require().True(lst.AddAfterIndex(9, 71))
	res = slices.Collect(lst.Seq())
	l.Require().Equal(exp, res)
}

func (l *ListTestSuite) TestNoAddBeforeIndex() {
	var dat = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	var lst *List[int]

	lst = New[int](dat...)
	l.Require().False(lst.AddBeforeIndex(-1, 64))
	res := slices.Collect(lst.Seq())
	l.Require().Equal(dat, res)

	l.Require().False(lst.AddBeforeIndex(128, 64))
	res = slices.Collect(lst.Seq())
	l.Require().Equal(dat, res)

	lst = New[int]()
	l.Require().False(lst.AddBeforeIndex(0, 64))
}

func (l *ListTestSuite) TestAddBeforeIndex() {
	var dat = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	var exp []int
	var lst *List[int]

	lst = New[int](dat...)
	exp = []int{71, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	l.Require().True(lst.AddBeforeIndex(0, 71))
	res := slices.Collect(lst.Seq())
	l.Require().Equal(exp, res)

	lst = New[int](dat...)
	exp = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 71, 15}
	l.Require().True(lst.AddBeforeIndex(15, 71))
	res = slices.Collect(lst.Seq())
	l.Require().Equal(exp, res)

	lst = New[int](dat...)
	exp = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 71, 9, 10, 11, 12, 13, 14, 15}
	l.Require().True(lst.AddBeforeIndex(9, 71))
	res = slices.Collect(lst.Seq())
	l.Require().Equal(exp, res)
}

func (l *ListTestSuite) TestNoMoveAfter() {
	var dat = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	lst := New[int](dat...)

	l.Require().False(lst.MoveAfter(-1, 3))
	l.Require().False(lst.MoveAfter(3, -1))
	l.Require().False(lst.MoveAfter(-2, -5))
	l.Require().False(lst.MoveAfter(3, 25))
	l.Require().False(lst.MoveAfter(34, 3))
	l.Require().False(lst.MoveAfter(34, 3))

	lst = New[int]()
	l.Require().False(lst.MoveAfter(0, 3))
}

func (l *ListTestSuite) TestMoveAfter() {
	var dat = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	var exp []int
	var lst *List[int]

	lst = New[int](dat...)
	exp = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	l.Require().True(lst.MoveAfter(8, 8))
	res := slices.Collect(lst.Seq())
	l.Require().Equal(exp, res)

	lst = New[int](dat...)
	exp = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	l.Require().True(lst.MoveAfter(8, 7))
	res = slices.Collect(lst.Seq())
	l.Require().Equal(exp, res)

	lst = New[int](dat...)
	exp = []int{0, 8, 1, 2, 3, 4, 5, 6, 7, 9, 10, 11, 12, 13, 14, 15}
	l.Require().True(lst.MoveAfter(8, 0))
	res = slices.Collect(lst.Seq())
	l.Require().Equal(exp, res)

	lst = New[int](dat...)
	exp = []int{0, 1, 2, 3, 4, 5, 6, 7, 9, 10, 11, 12, 13, 14, 15, 8}
	l.Require().True(lst.MoveAfter(8, 15))
	res = slices.Collect(lst.Seq())
	l.Require().Equal(exp, res)
}

func (l *ListTestSuite) TestNoMoveBefore() {
	var dat = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	lst := New[int](dat...)

	l.Require().False(lst.MoveBefore(-1, 3))
	l.Require().False(lst.MoveBefore(3, -1))
	l.Require().False(lst.MoveBefore(-2, -5))
	l.Require().False(lst.MoveBefore(3, 25))
	l.Require().False(lst.MoveBefore(34, 3))
	l.Require().False(lst.MoveBefore(34, 3))

	lst = New[int]()
	l.Require().False(lst.MoveBefore(0, 0))
}

func (l *ListTestSuite) TestMoveBefore() {
	var dat = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	var exp []int
	var lst *List[int]

	lst = New[int](dat...)
	exp = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	l.Require().True(lst.MoveBefore(8, 8))
	res := slices.Collect(lst.Seq())
	l.Require().Equal(exp, res)

	lst = New[int](dat...)
	exp = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	l.Require().True(lst.MoveBefore(8, 9))
	res = slices.Collect(lst.Seq())
	l.Require().Equal(exp, res)

	lst = New[int](dat...)
	exp = []int{8, 0, 1, 2, 3, 4, 5, 6, 7, 9, 10, 11, 12, 13, 14, 15}
	l.Require().True(lst.MoveBefore(8, 0))
	res = slices.Collect(lst.Seq())
	l.Require().Equal(exp, res)

	lst = New[int](dat...)
	exp = []int{0, 1, 2, 3, 4, 5, 6, 7, 9, 10, 11, 12, 13, 14, 8, 15}
	l.Require().True(lst.MoveBefore(8, 15))
	res = slices.Collect(lst.Seq())
	l.Require().Equal(exp, res)
}

func (l *ListTestSuite) TestEqual() {
	l.Require().False(New[int]() == New(1))
	l.Require().False(New(1) == New[int]())

	lst := New[int]()
	l.Require().True(lst.Equal(lst))
	l.Require().True(New[int]().Equal(New[int]()))
	l.Require().True(New[int](1, 2).Equal(New[int](1, 2)))

	l.Require().False(New[int](1, 2).Equal(New[int](1, 2, 3)))
	l.Require().False(New[int](1, 2, 3).Equal(New[int](1, 2)))
	l.Require().False(New[int](1, 3, 2).Equal(New[int](1, 2, 3)))
}

func (l *ListTestSuite) TestDeleteEmpty() {
	lst := New[int]()
	exists := lst.Delete(23)
	l.Require().False(exists)
}

func (l *ListTestSuite) TestDeleteMissing() {
	dat := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	lst := New(dat...)
	exists := lst.Delete(23)
	l.Require().False(exists)
}

func (l *ListTestSuite) TestDelete() {
	dat := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	exp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	lst := New(dat...)
	exists := lst.Delete(0)
	l.Require().True(exists)
	res := slices.Collect(lst.Seq())
	l.Require().Equal(exp, res)

	dat = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	exp = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	lst = New(dat...)
	exists = lst.Delete(15)
	l.Require().True(exists)
	res = slices.Collect(lst.Seq())
	l.Require().Equal(exp, res)

	dat = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	exp = []int{0, 1, 2, 3, 4, 5, 6, 7, 9, 10, 11, 12, 13, 14, 15}
	lst = New(dat...)
	exists = lst.Delete(8)
	l.Require().True(exists)
	res = slices.Collect(lst.Seq())
	l.Require().Equal(exp, res)
}

func (l *ListTestSuite) TestPopAtEmpty() {
	lst := New[int]()
	_, exists := lst.PopAt(23)
	l.Require().False(exists)
}

func (l *ListTestSuite) TestPopAtMissing() {
	dat := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	lst := New(dat...)
	_, exists := lst.PopAt(23)
	l.Require().False(exists)
}

func (l *ListTestSuite) TestPopAt() {
	dat := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	exp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	lst := New(dat...)
	v, exists := lst.PopAt(0)
	l.Require().True(exists)
	l.Require().Equal(0, v)
	res := slices.Collect(lst.Seq())
	l.Require().Equal(exp, res)

	dat = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	exp = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	lst = New(dat...)
	v, exists = lst.PopAt(15)
	l.Require().True(exists)
	l.Require().Equal(15, v)
	res = slices.Collect(lst.Seq())
	l.Require().Equal(exp, res)

	dat = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	exp = []int{0, 1, 2, 3, 4, 5, 6, 7, 9, 10, 11, 12, 13, 14, 15}
	lst = New(dat...)
	v, exists = lst.PopAt(8)
	l.Require().True(exists)
	l.Require().Equal(8, v)
	res = slices.Collect(lst.Seq())
	l.Require().Equal(exp, res)
}
