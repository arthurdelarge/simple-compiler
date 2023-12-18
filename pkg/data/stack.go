package data

import "fmt"

type Stack[T any] struct {
	elements []T
}

func CreateStack[T any]() *Stack[T] {
	return &Stack[T]{elements: make([]T, 0)}
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) Pop() (T, bool) {
	var element T

	if len(s.elements) == 0 {
		return element, false
	}

	element = s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]

	return element, true
}

func (s *Stack[T]) Print() {
	fmt.Printf("%+v\n", s.elements)
}

func (s *Stack[T]) Top() (T, bool) {
	var element T

	if len(s.elements) == 0 {
		return element, false
	}

	element = s.elements[len(s.elements)-1]

	return element, true
}
