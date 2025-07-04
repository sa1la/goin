package goin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueueEnqueueDequeue(t *testing.T) {
	q := NewQueue[int](5)

	// Test empty queue
	assert.True(t, q.IsEmpty())
	assert.Equal(t, 0, q.Size())

	// Test enqueue
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	assert.False(t, q.IsEmpty())
	assert.Equal(t, 3, q.Size())

	// Test dequeue (FIFO order)
	assert.Equal(t, 1, q.Dequeue())
	assert.Equal(t, 2, q.Dequeue())
	assert.Equal(t, 3, q.Dequeue())

	assert.True(t, q.IsEmpty())
	assert.Equal(t, 0, q.Size())
}

func TestQueueFront(t *testing.T) {
	q := NewQueue[string](3)

	// Test front on empty queue
	val, ok := q.Front()
	assert.False(t, ok)
	assert.Equal(t, "", val)

	// Test front with elements
	q.Enqueue("hello")
	q.Enqueue("world")

	val, ok = q.Front()
	assert.True(t, ok)
	assert.Equal(t, "hello", val)

	// Ensure front doesn't modify queue
	assert.Equal(t, 2, q.Size())
	assert.Equal(t, "hello", q.Dequeue())
}

func TestQueueResize(t *testing.T) {
	q := NewQueue[int](2)

	// Fill beyond initial capacity
	for i := 1; i <= 5; i++ {
		q.Enqueue(i)
	}

	assert.Equal(t, 5, q.Size())

	// Verify FIFO order is maintained after resize
	for i := 1; i <= 5; i++ {
		assert.Equal(t, i, q.Dequeue())
	}
}

func TestQueueToSlice(t *testing.T) {
	q := NewQueue[int](5)

	// Test empty queue
	slice := q.ToSlice()
	assert.Empty(t, slice)

	// Test with elements
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	slice = q.ToSlice()
	expected := []int{1, 2, 3}
	assert.Equal(t, expected, slice)

	// Test after partial dequeue
	q.Dequeue()
	slice = q.ToSlice()
	expected = []int{2, 3}
	assert.Equal(t, expected, slice)
}

func TestQueueClear(t *testing.T) {
	q := NewQueue[int](5)

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	assert.Equal(t, 3, q.Size())
	assert.False(t, q.IsEmpty())

	q.Clear()

	assert.Equal(t, 0, q.Size())
	assert.True(t, q.IsEmpty())
}

func TestQueueDequeueEmptyQueue(t *testing.T) {
	q := NewQueue[int](5)

	// Dequeue from empty queue should return zero value
	val := q.Dequeue()
	assert.Equal(t, 0, val)
}

func TestQueueCircularBehavior(t *testing.T) {
	q := NewQueue[int](3)

	// Fill queue
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	// Dequeue some elements
	assert.Equal(t, 1, q.Dequeue())
	assert.Equal(t, 2, q.Dequeue())

	// Add more elements (should use circular buffer)
	q.Enqueue(4)
	q.Enqueue(5)

	// Verify order is correct
	slice := q.ToSlice()
	expected := []int{3, 4, 5}
	assert.Equal(t, expected, slice)
}