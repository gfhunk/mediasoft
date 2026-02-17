package main

import "errors"

type Stack struct {
	items []any
}

func NewStack() *Stack {
	return &Stack{
		items: make([]any, 0),
	}
}

func (s *Stack) Push(v any) {
	s.items = append(s.items, v)
}

func (s *Stack) Pop() (any, error) {
	if len(s.items) == 0 {
		return nil, errors.New("стек пуст")
	}
	val := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val, nil
}

func (s *Stack) Peek() (any, error) {
	if len(s.items) == 0 {
		return nil, errors.New("стек пуст")
	}
	return s.items[len(s.items)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}
