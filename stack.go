package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

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

func (r *Stack[T]) Size() int {
	return len(r.items)
}

func (r *Stack[T]) Empty() {
	if len(r.items) == 0 {
		return
	}

	var emptyItems []*T
	r.items = emptyItems
}

func (r *Stack[T]) Serialize() (string, error) {
	if len(r.items) == 0 {
		return "[]", nil
	}

	data, err := json.Marshal(r.items)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (r *Stack[T]) Deserialize(data string) error {
	if r == nil {
		return errors.New("cannot deserialize into a nil stack")
	}

	if data == "[]" {
		r.items = []*T{}
		return nil
	}

	var items []*T
	err := json.Unmarshal([]byte(data), &items)
	if err != nil {
		return err
	}

	r.items = items
	return nil
}
