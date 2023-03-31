package queue

import (
	"github.com/cluster05/dsa/common/errors"
)

type Queue[T any] struct {
	items  []T
	length int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{items: []T{}, length: 0}
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
	q.length = q.length + 1
}

func (q *Queue[T]) Dequeue() T {
	if q.IsEmpty() {
		panic(errors.AccessingInvalidMemoryAddress)
	}

	front := q.items[0]

	q.items = q.items[1:]
	q.length = q.length - 1
	return front
}

func (q *Queue[T]) Front() T {
	if q.IsEmpty() {
		panic("queue is empty")
	}

	return q.items[0]
}

func (q *Queue[T]) ElementAt(pos int) T {
	if q.IsEmpty() || pos >= len(q.items) {
		panic(errors.AccessingInvalidMemoryAddress)
	}
	return q.items[pos]
}

func (q *Queue[T]) Size() int {
	return q.length
}

func (q *Queue[T]) IsEmpty() bool {
	return q.length == 0
}
