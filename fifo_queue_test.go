package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFIFOQueue_Enqueue(t *testing.T) {
	t.Run("should enqueue items in the correct order", func(t *testing.T) {
		queue := NewFIFOQueue()

		queue.Enqueue(FIFOQueueItem{ID: "item1", Value: 1})
		queue.Enqueue(FIFOQueueItem{ID: "item2", Value: 2})
		queue.Enqueue(FIFOQueueItem{ID: "item3", Value: 3})

		assert.Equal(t, "item1", queue.items[0].ID)
		assert.Equal(t, int64(1), queue.items[0].Value)
		assert.Equal(t, "item2", queue.items[1].ID)
		assert.Equal(t, int64(2), queue.items[1].Value)
		assert.Equal(t, "item3", queue.items[2].ID)
		assert.Equal(t, int64(3), queue.items[2].Value)
	})
}
