package main

import "fmt"

type Stack[T any] struct {
	items []*T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (r *Stack[T]) Push(item T) {
	r.items = append(r.items, &item)
}

func (r *Stack[T]) Pop() (*T, error) {
	if len(r.items) == 0 {
		return nil, fmt.Errorf("stack is empty")
	}

	item := r.items[len(r.items)-1]
	r.items = r.items[:len(r.items)-1]

	return item, nil
}

func (r *Stack[T]) Peek() (*T, error) {
	if len(r.items) == 0 {
		return nil, fmt.Errorf("stack is empty")
	}

	return r.items[len(r.items)-1], nil
}

func (r *Stack[T]) IsEmpty() bool {
	return len(r.items) == 0
}
