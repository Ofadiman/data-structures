package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestFIFOQueue_Dequeue(t *testing.T) {
	t.Run("should dequeue items in the correct order", func(t *testing.T) {
		queue := NewFIFOQueue()

		queue.Enqueue(FIFOQueueItem{ID: "item1", Value: 1})
		queue.Enqueue(FIFOQueueItem{ID: "item2", Value: 2})
		queue.Enqueue(FIFOQueueItem{ID: "item3", Value: 3})

		item, err := queue.Dequeue()
		require.NoError(t, err)
		assert.Equal(t, "item1", item.ID)
		assert.Equal(t, int64(1), item.Value)

		item, err = queue.Dequeue()
		require.NoError(t, err)
		assert.Equal(t, "item2", item.ID)
		assert.Equal(t, int64(2), item.Value)

		item, err = queue.Dequeue()
		require.NoError(t, err)
		assert.Equal(t, "item3", item.ID)
		assert.Equal(t, int64(3), item.Value)
	})

	t.Run("should return an error when dequeueing from an empty queue", func(t *testing.T) {
		queue := NewFIFOQueue()

		_, err := queue.Dequeue()
		assert.Error(t, err)
		assert.Equal(t, "queue is empty", err.Error())
	})
}
