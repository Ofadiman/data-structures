package main

import "errors"

type FIFOQueueItem struct {
	ID    string
	Value int64
}

type FIFOQueue struct {
	items []FIFOQueueItem
}

func NewFIFOQueue() *FIFOQueue {
	return &FIFOQueue{items: make([]FIFOQueueItem, 0)}
}

func (r *FIFOQueue) Enqueue(item FIFOQueueItem) {
	r.items = append(r.items, item)
}

func (r *FIFOQueue) Dequeue() (FIFOQueueItem, error) {
	if len(r.items) == 0 {
		return FIFOQueueItem{}, errors.New("queue is empty")
	}

	item := r.items[0]
	r.items = r.items[1:]

	return item, nil
}

func (r *FIFOQueue) Front() (FIFOQueueItem, error) {
	if len(r.items) == 0 {
		return FIFOQueueItem{}, errors.New("queue is empty")
	}

	return r.items[0], nil
}
