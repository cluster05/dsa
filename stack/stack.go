package stack

import (
	"github.com/cluster05/dsa/common/errors"
)

type Stack[T any] struct {
	items  []T
	length int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{items: []T{}, length: 0}
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
	s.length = s.length + 1
}

func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		panic(errors.AccessingInvalidMemoryAddress)
	}
	top := s.items[s.length-1]

	s.items = s.items[:s.length-1]
	s.length = s.length - 1
	return top
}

func (s *Stack[T]) Top() T {
	if s.IsEmpty() {
		panic(errors.AccessingInvalidMemoryAddress)
	}
	return s.items[s.length-1]
}

func (s *Stack[T]) ElementAt(pos int) T {
	if s.IsEmpty() || s.length >= pos {
		panic(errors.AccessingInvalidMemoryAddress)
	}

	return s.items[s.length-pos]
}

func (s *Stack[T]) Size() int {
	return s.length
}

func (s *Stack[T]) IsEmpty() bool {
	return s.length == 0
}
