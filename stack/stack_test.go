package stack

import (
	"github.com/HoskeOwl/ggstruct/list"
	"github.com/stretchr/testify/suite"
	"testing"
)

type StackTestSuite struct {
	suite.Suite
}

func TestRunQueueSuite(t *testing.T) {
	suite.Run(t, new(StackTestSuite))
}

func (s *StackTestSuite) TestInit() {
	var q *Stack[int]
	q = New[int]()
	s.Require().Equal(0, q.Len())
	s.Require().True(q.data.Equal(list.New[int]()))

	q = New[int](2)
	s.Require().Equal(1, q.Len())
	s.Require().True(q.data.Equal(list.New[int](2)))

	q = New[int](2, 3, 4)
	s.Require().Equal(3, q.Len())
	s.Require().True(q.data.Equal(list.New[int](2, 3, 4)))
}

func (s *StackTestSuite) TestLen() {
	var q *Stack[int]
	q = New[int]()
	s.Require().Equal(0, q.data.Len())

	q = New[int](2)
	s.Require().Equal(1, q.data.Len())

	q = New[int](2, 3, 4)
	s.Require().Equal(3, q.data.Len())
	q.Pop()
	s.Require().Equal(2, q.data.Len())
	q.Push(45)
	s.Require().Equal(3, q.data.Len())
	q.Top()
	s.Require().Equal(3, q.data.Len())
}

func (s *StackTestSuite) TestInitLimit() {
	var q *Stack[int]
	q = New[int]().WithLimit(5)
	s.Require().Equal(0, q.Len())
	s.Require().Equal(5, q.limit)
	s.Require().True(q.data.Equal(list.New[int]()))

	q = New[int](2).WithLimit(5)
	s.Require().Equal(1, q.Len())
	s.Require().Equal(5, q.limit)
	s.Require().True(q.data.Equal(list.New[int](2)))

	q = New[int](2, 3, 4).WithLimit(8)
	s.Require().Equal(3, q.Len())
	s.Require().Equal(8, q.limit)
	s.Require().True(q.data.Equal(list.New[int](2, 3, 4)))
}

func (s *StackTestSuite) TestPush() {
	var q *Stack[int]
	var ok bool

	q = New[int]()
	ok = q.Push(1)
	s.Require().True(ok)
	s.Require().Equal(1, q.Len())
	s.Require().True(q.data.Equal(list.New[int](1)))

	ok = q.Push(4)
	s.Require().True(ok)
	s.Require().Equal(2, q.Len())
	s.Require().True(q.data.Equal(list.New[int](4, 1)))

	ok = q.Push(1)
	s.Require().True(ok)
	s.Require().Equal(3, q.Len())
	s.Require().True(q.data.Equal(list.New[int](1, 4, 1)))

	ok = q.Push(4)
	ok = q.Push(5)
	ok = q.Push(6)
	s.Require().True(ok)
	s.Require().Equal(6, q.Len())
	s.Require().True(q.data.Equal(list.New[int](6, 5, 4, 1, 4, 1)))
}

func (s *StackTestSuite) TestPushWithLimit() {
	var q *Stack[int]
	var ok bool

	// 0 and below means no limit
	q = New[int]().WithLimit(0)
	ok = q.Push(1)
	s.Require().True(ok)
	s.Require().Equal(1, q.Len())
	s.Require().True(q.data.Equal(list.New[int](1)))
	ok = q.Push(2)
	ok = q.Push(3)
	ok = q.Push(4)
	s.Require().True(ok)
	s.Require().Equal(4, q.Len())
	s.Require().True(q.data.Equal(list.New[int](4, 3, 2, 1)))

	q = New[int]().WithLimit(1)
	ok = q.Push(1)
	s.Require().True(ok)
	s.Require().Equal(1, q.Len())
	s.Require().True(q.data.Equal(list.New[int](1)))
	ok = q.Push(2)
	s.Require().False(ok)
	s.Require().Equal(1, q.Len())
	s.Require().True(q.data.Equal(list.New[int](1)))

	q = New[int]().WithLimit(2)
	ok = q.Push(4)
	s.Require().True(ok)
	s.Require().Equal(1, q.Len())
	s.Require().True(q.data.Equal(list.New[int](4)))
	ok = q.Push(5)
	s.Require().True(ok)
	s.Require().Equal(2, q.Len())
	s.Require().True(q.data.Equal(list.New[int](5, 4)))
	ok = q.Push(6)
	s.Require().False(ok)
	s.Require().Equal(2, q.Len())
	s.Require().True(q.data.Equal(list.New[int](5, 4)))

	q.Pop()
	ok = q.Push(1)
	s.Require().True(ok)
	s.Require().Equal(2, q.Len())
	s.Require().True(q.data.Equal(list.New[int](1, 4)))

	// yes, can be. Can't push but can pop
	q = New[int](1, 2, 3).WithLimit(1)
	ok = q.Push(8)
	s.Require().False(ok)
	s.Require().Equal(3, q.Len())
	s.Require().True(q.data.Equal(list.New[int](1, 2, 3)))
	q.Pop()
	q.Pop()
	q.Pop()
	ok = q.Push(8)
	s.Require().True(ok)
	s.Require().Equal(1, q.Len())
	s.Require().True(q.data.Equal(list.New[int](8)))
}

func (s *StackTestSuite) TestPop() {
	var q *Stack[int]
	var value int
	var exists bool

	q = New[int]()
	_, exists = q.Pop()
	s.Require().False(exists)
	s.Require().True(q.data.Equal(list.New[int]()))

	q = New[int](1)
	value, exists = q.Pop()
	s.Require().Equal(value, 1)
	s.Require().True(exists)
	s.Require().True(q.data.Equal(list.New[int]()))
	_, exists = q.Pop()
	s.Require().False(exists)
	s.Require().True(q.data.Equal(list.New[int]()))

	expected := []int{1, 2, 3, 4, 5, 6}
	q = New[int](expected...)
	for _, v := range expected {
		value, exists = q.Pop()
		s.Require().Equal(value, v)
		s.Require().True(exists)
	}
	_, exists = q.Pop()
	s.Require().False(exists)
	s.Require().True(q.data.Equal(list.New[int]()))
}

func (s *StackTestSuite) TestTop() {
	var q *Stack[int]
	var value int
	var exists bool

	q = New[int]()
	_, exists = q.Top()
	s.Require().False(exists)

	q.Push(1)
	value, exists = q.Top()
	s.Require().True(exists)
	s.Require().Equal(value, 1)

	q.Push(2)
	q.Push(3)
	q.Push(4)
	value, exists = q.Top()
	s.Require().True(exists)
	s.Require().Equal(value, 4)
	q.Pop()
	value, exists = q.Top()
	s.Require().True(exists)
	s.Require().Equal(value, 3)

	q = New[int](7)
	value, exists = q.Top()
	s.Require().True(exists)
	s.Require().Equal(value, 7)

	q = New[int](7, 8, 9)
	value, exists = q.Top()
	s.Require().True(exists)
	s.Require().Equal(value, 7)
}

func (s *StackTestSuite) TestIsFull() {
	var q *Stack[int]

	q = New[int]()
	s.Require().False(q.IsFull())
	q.limit = -1
	s.Require().False(q.IsFull())
	q.limit = 1
	s.Require().False(q.IsFull())

	q = New(1, 2, 3, 4, 5)
	s.Require().False(q.IsFull())
	q.limit = -1
	s.Require().False(q.IsFull())
	q.limit = 10
	s.Require().False(q.IsFull())
	q.limit = 1
	s.Require().True(q.IsFull())

	q = New[int]()
	q.limit = 1
	s.Require().False(q.IsFull())
	q.Push(1)
	s.Require().True(q.IsFull())
	q.Pop()
	s.Require().False(q.IsFull())

	q = New[int]()
	q.limit = 1
	s.Require().False(q.IsFull())
	q.Push(1)
	s.Require().True(q.IsFull())
	q.Pop()
	s.Require().False(q.IsFull())
}

func (s *StackTestSuite) TestIsEmpty() {
	var q *Stack[int]

	q = New[int](1)
	s.Require().False(q.IsEmpty())

	q = New[int]()
	s.Require().True(q.IsEmpty())
	q.Push(23)
	s.Require().False(q.IsEmpty())
	q.Pop()
	s.Require().True(q.IsEmpty())
}
