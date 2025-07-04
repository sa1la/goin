package goin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinHeap(t *testing.T) {
	heap := NewMinHeap[int]()

	// Test empty heap
	assert.True(t, heap.IsEmpty())
	assert.Equal(t, 0, heap.Size())

	// Test push and peek
	heap.Push(5)
	heap.Push(2)
	heap.Push(8)
	heap.Push(1)

	assert.False(t, heap.IsEmpty())
	assert.Equal(t, 4, heap.Size())

	val, ok := heap.Peek()
	assert.True(t, ok)
	assert.Equal(t, 1, val) // Min should be at top

	// Test pop (should return elements in ascending order)
	assert.Equal(t, 1, heap.Pop())
	assert.Equal(t, 2, heap.Pop())
	assert.Equal(t, 5, heap.Pop())
	assert.Equal(t, 8, heap.Pop())

	assert.True(t, heap.IsEmpty())
}

func TestMaxHeap(t *testing.T) {
	heap := NewMaxHeap[int]()

	// Test push and peek
	heap.Push(5)
	heap.Push(2)
	heap.Push(8)
	heap.Push(1)

	val, ok := heap.Peek()
	assert.True(t, ok)
	assert.Equal(t, 8, val) // Max should be at top

	// Test pop (should return elements in descending order)
	assert.Equal(t, 8, heap.Pop())
	assert.Equal(t, 5, heap.Pop())
	assert.Equal(t, 2, heap.Pop())
	assert.Equal(t, 1, heap.Pop())
}

func TestHeapWithStrings(t *testing.T) {
	heap := NewMinHeap[string]()

	heap.Push("zebra")
	heap.Push("apple")
	heap.Push("banana")

	// Should return in lexicographical order
	assert.Equal(t, "apple", heap.Pop())
	assert.Equal(t, "banana", heap.Pop())
	assert.Equal(t, "zebra", heap.Pop())
}

func TestHeapToSlice(t *testing.T) {
	heap := NewMinHeap[int]()

	// Test empty heap
	slice := heap.ToSlice()
	assert.Empty(t, slice)

	// Test with elements
	heap.Push(3)
	heap.Push(1)
	heap.Push(4)

	slice = heap.ToSlice()
	assert.Len(t, slice, 3)
	assert.Contains(t, slice, 1)
	assert.Contains(t, slice, 3)
	assert.Contains(t, slice, 4)
}

func TestHeapClear(t *testing.T) {
	heap := NewMaxHeap[int]()

	heap.Push(1)
	heap.Push(2)
	heap.Push(3)

	assert.Equal(t, 3, heap.Size())
	assert.False(t, heap.IsEmpty())

	heap.Clear()

	assert.Equal(t, 0, heap.Size())
	assert.True(t, heap.IsEmpty())
}

func TestHeapPopEmpty(t *testing.T) {
	heap := NewMinHeap[int]()

	// Pop from empty heap should return zero value
	val := heap.Pop()
	assert.Equal(t, 0, val)

	// Peek on empty heap
	val, ok := heap.Peek()
	assert.False(t, ok)
	assert.Equal(t, 0, val)
}

func TestPriorityQueue(t *testing.T) {
	pq := NewPriorityQueue[int]()

	// Test enqueue and dequeue
	pq.Enqueue(5)
	pq.Enqueue(2)
	pq.Enqueue(8)
	pq.Enqueue(1)

	// Should dequeue in descending order (highest priority first)
	assert.Equal(t, 8, pq.Dequeue())
	assert.Equal(t, 5, pq.Dequeue())
	assert.Equal(t, 2, pq.Dequeue())
	assert.Equal(t, 1, pq.Dequeue())
}

func TestHeapProperty(t *testing.T) {
	heap := NewMinHeap[int]()

	// Add many elements
	elements := []int{15, 10, 20, 8, 25, 5, 30, 12}
	for _, elem := range elements {
		heap.Push(elem)
	}

	// Extract all elements - should be in sorted order
	var result []int
	for !heap.IsEmpty() {
		result = append(result, heap.Pop())
	}

	// Verify sorted order
	for i := 1; i < len(result); i++ {
		assert.LessOrEqual(t, result[i-1], result[i])
	}
}

func TestHeapDuplicates(t *testing.T) {
	heap := NewMaxHeap[int]()

	// Test with duplicate values
	heap.Push(5)
	heap.Push(5)
	heap.Push(3)
	heap.Push(5)

	assert.Equal(t, 5, heap.Pop())
	assert.Equal(t, 5, heap.Pop())
	assert.Equal(t, 5, heap.Pop())
	assert.Equal(t, 3, heap.Pop())
}