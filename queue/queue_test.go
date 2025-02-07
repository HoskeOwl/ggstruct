package queue

import (
	"github.com/HoskeOwl/ggstruct/list"
	"github.com/stretchr/testify/suite"
	"testing"
)

type QueueTestSuite struct {
	suite.Suite
}

func TestRunQueueSuite(t *testing.T) {
	suite.Run(t, new(QueueTestSuite))
}

func (suite *QueueTestSuite) TestInit() {
	var q *Queue[int]
	q = New[int]()
	suite.Require().Equal(0, q.Len())
	suite.Require().True(q.data.Equal(list.New[int]()))

	q = New[int](2)
	suite.Require().Equal(1, q.Len())
	suite.Require().True(q.data.Equal(list.New[int](2)))

	q = New[int](2, 3, 4)
	suite.Require().Equal(3, q.Len())
	suite.Require().True(q.data.Equal(list.New[int](2, 3, 4)))
}

func (suite *QueueTestSuite) TestLen() {
	var q *Queue[int]
	q = New[int]()
	suite.Require().Equal(0, q.data.Len())

	q = New[int](2)
	suite.Require().Equal(1, q.data.Len())

	q = New[int](2, 3, 4)
	suite.Require().Equal(3, q.data.Len())
	q.Dequeue()
	suite.Require().Equal(2, q.data.Len())
	q.Enqueue(45)
	suite.Require().Equal(3, q.data.Len())
	q.Peek()
	suite.Require().Equal(3, q.data.Len())
}

func (suite *QueueTestSuite) TestInitLimit() {
	var q *Queue[int]
	q = New[int]().WithLimit(5)
	suite.Require().Equal(0, q.Len())
	suite.Require().Equal(5, q.Limit)
	suite.Require().True(q.data.Equal(list.New[int]()))

	q = New[int](2).WithLimit(5)
	suite.Require().Equal(1, q.Len())
	suite.Require().Equal(5, q.Limit)
	suite.Require().True(q.data.Equal(list.New[int](2)))

	q = New[int](2, 3, 4).WithLimit(8)
	suite.Require().Equal(3, q.Len())
	suite.Require().Equal(8, q.Limit)
	suite.Require().True(q.data.Equal(list.New[int](2, 3, 4)))
}

func (suite *QueueTestSuite) TestEnqueue() {
	var q *Queue[int]
	var ok bool

	q = New[int]()
	ok = q.Enqueue(1)
	suite.Require().True(ok)
	suite.Require().Equal(1, q.Len())
	suite.Require().True(q.data.Equal(list.New[int](1)))

	ok = q.Enqueue(4)
	suite.Require().True(ok)
	suite.Require().Equal(2, q.Len())
	suite.Require().True(q.data.Equal(list.New[int](1, 4)))

	ok = q.Enqueue(1)
	suite.Require().True(ok)
	suite.Require().Equal(3, q.Len())
	suite.Require().True(q.data.Equal(list.New[int](1, 4, 1)))

	ok = q.Enqueue(4)
	ok = q.Enqueue(5)
	ok = q.Enqueue(6)
	suite.Require().True(ok)
	suite.Require().Equal(6, q.Len())
	suite.Require().True(q.data.Equal(list.New[int](1, 4, 1, 4, 5, 6)))
}

func (suite *QueueTestSuite) TestEnqueueWithLimit() {
	var q *Queue[int]
	var ok bool

	// 0 and below means no limit
	q = New[int]().WithLimit(0)
	ok = q.Enqueue(1)
	suite.Require().True(ok)
	suite.Require().Equal(1, q.Len())
	suite.Require().True(q.data.Equal(list.New[int](1)))
	ok = q.Enqueue(2)
	ok = q.Enqueue(3)
	ok = q.Enqueue(4)
	suite.Require().True(ok)
	suite.Require().Equal(4, q.Len())
	suite.Require().True(q.data.Equal(list.New[int](1, 2, 3, 4)))

	q = New[int]().WithLimit(1)
	ok = q.Enqueue(1)
	suite.Require().True(ok)
	suite.Require().Equal(1, q.Len())
	suite.Require().True(q.data.Equal(list.New[int](1)))
	ok = q.Enqueue(2)
	suite.Require().False(ok)
	suite.Require().Equal(1, q.Len())
	suite.Require().True(q.data.Equal(list.New[int](1)))

	q = New[int]().WithLimit(2)
	ok = q.Enqueue(4)
	suite.Require().True(ok)
	suite.Require().Equal(1, q.Len())
	suite.Require().True(q.data.Equal(list.New[int](4)))
	ok = q.Enqueue(5)
	suite.Require().True(ok)
	suite.Require().Equal(2, q.Len())
	suite.Require().True(q.data.Equal(list.New[int](4, 5)))
	ok = q.Enqueue(6)
	suite.Require().False(ok)
	suite.Require().Equal(2, q.Len())
	suite.Require().True(q.data.Equal(list.New[int](4, 5)))

	q.Dequeue()
	ok = q.Enqueue(1)
	suite.Require().True(ok)
	suite.Require().Equal(2, q.Len())
	suite.Require().True(q.data.Equal(list.New[int](5, 1)))

	// yes, can be. Can't enqueue but can dequeue
	q = New[int](1, 2, 3).WithLimit(1)
	ok = q.Enqueue(8)
	suite.Require().False(ok)
	suite.Require().Equal(3, q.Len())
	suite.Require().True(q.data.Equal(list.New[int](1, 2, 3)))
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	ok = q.Enqueue(8)
	suite.Require().True(ok)
	suite.Require().Equal(1, q.Len())
	suite.Require().True(q.data.Equal(list.New[int](8)))
}

func (suite *QueueTestSuite) TestDequeue() {
	var q *Queue[int]
	var value int
	var exists bool

	q = New[int]()
	_, exists = q.Dequeue()
	suite.Require().False(exists)
	suite.Require().True(q.data.Equal(list.New[int]()))

	q = New[int](1)
	value, exists = q.Dequeue()
	suite.Require().Equal(value, 1)
	suite.Require().True(exists)
	suite.Require().True(q.data.Equal(list.New[int]()))
	_, exists = q.Dequeue()
	suite.Require().False(exists)
	suite.Require().True(q.data.Equal(list.New[int]()))

	expected := []int{1, 2, 3, 4, 5, 6}
	q = New[int](expected...)
	for _, v := range expected {
		value, exists = q.Dequeue()
		suite.Require().Equal(value, v)
		suite.Require().True(exists)
	}
	_, exists = q.Dequeue()
	suite.Require().False(exists)
	suite.Require().True(q.data.Equal(list.New[int]()))

}

func (suite *QueueTestSuite) TestPeek() {
	var q *Queue[int]
	var value int
	var exists bool

	q = New[int]()
	_, exists = q.Peek()
	suite.Require().False(exists)

	q.Enqueue(1)
	value, exists = q.Peek()
	suite.Require().True(exists)
	suite.Require().Equal(value, 1)

	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	value, exists = q.Peek()
	suite.Require().True(exists)
	suite.Require().Equal(value, 1)
	q.Dequeue()
	value, exists = q.Peek()
	suite.Require().True(exists)
	suite.Require().Equal(value, 2)

	q = New[int](7)
	value, exists = q.Peek()
	suite.Require().True(exists)
	suite.Require().Equal(value, 7)

	q = New[int](7, 8, 9)
	value, exists = q.Peek()
	suite.Require().True(exists)
	suite.Require().Equal(value, 7)
}
