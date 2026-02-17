package main

import "errors"

type Queue struct {
	items []any
}

func NewQueue() *Queue {
	return &Queue{
		items: make([]any, 0),
	}
}

func (q *Queue) Enqueue(v any) {
	q.items = append(q.items, v)
}

func (q *Queue) Dequeue() (any, error) {
	if len(q.items) == 0 {
		return nil, errors.New("очередь пуста")
	}
	val := q.items[0]
	q.items = q.items[1:]
	return val, nil
}

func (q *Queue) Peek() (any, error) {
	if len(q.items) == 0 {
		return nil, errors.New("очередь пуста")
	}
	return q.items[0], nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}
