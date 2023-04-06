package main

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
