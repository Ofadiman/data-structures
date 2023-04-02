package main

type Stack[T any] struct {
	items []*T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (r *Stack[T]) Push(item T) {
	r.items = append(r.items, &item)
}
