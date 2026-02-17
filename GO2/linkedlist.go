package main

import "errors"

type Node struct {
	Value any
	Next  *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head: nil,
		Tail: nil,
		Size: 0,
	}
}

func (l *LinkedList) Add(v any) {
	node := &Node{Value: v, Next: nil}

	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		l.Tail.Next = node
		l.Tail = node
	}
	l.Size++
}

func (l *LinkedList) Get(index int) (any, error) {
	if index < 0 || index >= l.Size {
		return nil, errors.New("индекс вне диапазона")
	}

	current := l.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current.Value, nil
}

func (l *LinkedList) Remove(index int) error {
	if index < 0 || index >= l.Size {
		return errors.New("индекс вне диапазона")
	}

	if index == 0 {
		l.Head = l.Head.Next
		if l.Head == nil {
			l.Tail = nil
		}
	} else {
		current := l.Head
		for i := 0; i < index-1; i++ {
			current = current.Next
		}
		current.Next = current.Next.Next
		if current.Next == nil {
			l.Tail = current
		}
	}
	l.Size--
	return nil
}

func (l *LinkedList) Values() []any {
	result := make([]any, l.Size)
	current := l.Head
	for i := 0; i < l.Size; i++ {
		result[i] = current.Value
		current = current.Next
	}
	return result
}
