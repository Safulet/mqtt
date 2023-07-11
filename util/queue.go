package util

import (
	"sync"
)

type Queue struct {
	capacity uint64
	data     []interface{}

	sync.RWMutex
}

func NewQueue(n uint64) (q *Queue) {
	return &Queue{
		capacity: n,
		data:     make([]interface{}, 0, n),
	}
}

func (q *Queue) Enqueue(v interface{}) bool {
	q.Lock()
	defer q.Unlock()

	if q.capacity <= uint64(len(q.data)) {
		return false
	} else {
		q.data = append(q.data, v)
		return true
	}
}

func (q *Queue) Dequeue() interface{} {
	q.Lock()
	defer q.Unlock()

	if len(q.data) == 0 {
		return nil
	}
	v := q.data[0]
	q.data = q.data[1:]

	return v
}

func (q *Queue) Len() int {
	q.RLock()
	defer q.RUnlock()

	return len(q.data)
}

func (q *Queue) Clone() *Queue {
	q.Lock()
	defer q.Unlock()

	clone := &Queue{capacity: q.capacity, data: make([]interface{}, len(q.data))}

	for i, item := range q.data {
		clone.data[i] = item
	}

	return clone
}
