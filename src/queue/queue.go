// This queue implementation is based on
// https://github.com/hishboy/gocommons/blob/master/lang/queue.go

package queue

import (
	"sync"
)

type queuenode struct {
	data interface{}
	next *queuenode
}

// Queue is a go-routine safe fifo data structure
type Queue struct {
	head  *queuenode
	tail  *queuenode
	count int
	lock  *sync.Mutex
}

// NewQueue creates a pointer to a new queue
func NewQueue() *Queue {
	q := &Queue{}
	q.lock = &sync.Mutex{}
	return q
}

// Len is a go-routine safe method that returns the
// number of elements in a queue
func (q *Queue) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.count
}

// Push inserts a value at the end of the queue
// This mutate the queue and go-routine safe
func (q *Queue) Push(item interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()

	n := &queuenode{data: item}

	if q.tail == nil {
		q.tail = n
		q.head = n
	} else {
		q.tail.next = n
		q.tail = n
	}
	q.count++
}

// Pop returns the front most value in the queue
// This mutate the queue and go-routine safe
func (q *Queue) Pop() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.head == nil {
		return nil
	}

	n := q.head
	q.head = n.next

	if q.head == nil {
		q.tail = nil
	}
	q.count--

	return n.data
}

// Peek reads the value at the front of the queue
// this does NOT mutate the queue and go-routine safe
func (q *Queue) Peek() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()

	n := q.head
	if n == nil || n.data == nil {
		return nil
	}

	return n.data
}

// Get reads the value at a position specified by the index
// this does NOT mutate the queue and go-routine safe
func (q *Queue) Get(index int) (interface{}, bool) {
	q.lock.Lock()
	defer q.lock.Unlock()

	n := q.head

	for i := 0; i < q.count; i++ {
		if i == index {
			return n.data, true
		}

		n = n.next
	}

	return nil, false
}
