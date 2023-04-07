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

func TestFIFOQueue_Front(t *testing.T) {
	t.Run("should return the front item without dequeuing it", func(t *testing.T) {
		queue := NewFIFOQueue()

		queue.Enqueue(FIFOQueueItem{ID: "item1", Value: 1})
		queue.Enqueue(FIFOQueueItem{ID: "item2", Value: 2})
		queue.Enqueue(FIFOQueueItem{ID: "item3", Value: 3})

		item, err := queue.Front()
		require.NoError(t, err)
		assert.Equal(t, "item1", item.ID)
		assert.Equal(t, int64(1), item.Value)

		item, err = queue.Front()
		require.NoError(t, err)
		assert.Equal(t, "item1", item.ID)
		assert.Equal(t, int64(1), item.Value)
	})

	t.Run("should return an error when getting the front item from an empty queue", func(t *testing.T) {
		queue := NewFIFOQueue()

		_, err := queue.Front()
		assert.Error(t, err)
		assert.Equal(t, "queue is empty", err.Error())
	})
}

func TestFIFOQueue_Size(t *testing.T) {
	t.Run("should return the correct size of the queue", func(t *testing.T) {
		queue := NewFIFOQueue()

		assert.Equal(t, 0, queue.Size())

		queue.Enqueue(FIFOQueueItem{ID: "item1", Value: 1})
		queue.Enqueue(FIFOQueueItem{ID: "item2", Value: 2})
		queue.Enqueue(FIFOQueueItem{ID: "item3", Value: 3})

		assert.Equal(t, 3, queue.Size())

		_, _ = queue.Dequeue()

		assert.Equal(t, 2, queue.Size())
	})
}

func TestFIFOQueue_IsEmpty(t *testing.T) {
	t.Run("should return true when the queue is empty", func(t *testing.T) {
		queue := NewFIFOQueue()

		assert.True(t, queue.IsEmpty())
	})

	t.Run("should return false when the queue is not empty", func(t *testing.T) {
		queue := NewFIFOQueue()

		queue.Enqueue(FIFOQueueItem{ID: "item1", Value: 1})

		assert.False(t, queue.IsEmpty())
	})
}
