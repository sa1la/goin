package goin

import (
	"golang.org/x/exp/constraints"
)

// HeapType defines the type of heap (min or max)
type HeapType int

const (
	MinHeap HeapType = iota
	MaxHeap
)

// Heap implements a generic binary heap (priority queue)
type Heap[T constraints.Ordered] struct {
	elements []T
	heapType HeapType
}

// NewMinHeap creates a new min heap
func NewMinHeap[T constraints.Ordered]() *Heap[T] {
	return &Heap[T]{
		elements: make([]T, 0),
		heapType: MinHeap,
	}
}

// NewMaxHeap creates a new max heap
func NewMaxHeap[T constraints.Ordered]() *Heap[T] {
	return &Heap[T]{
		elements: make([]T, 0),
		heapType: MaxHeap,
	}
}

// Push adds an element to the heap
func (h *Heap[T]) Push(value T) {
	h.elements = append(h.elements, value)
	h.heapifyUp(len(h.elements) - 1)
}

// Pop removes and returns the top element (min or max depending on heap type)
func (h *Heap[T]) Pop() T {
	if len(h.elements) == 0 {
		var zero T
		return zero
	}

	if len(h.elements) == 1 {
		result := h.elements[0]
		h.elements = h.elements[:0]
		return result
	}

	result := h.elements[0]
	h.elements[0] = h.elements[len(h.elements)-1]
	h.elements = h.elements[:len(h.elements)-1]
	h.heapifyDown(0)
	return result
}

// Peek returns the top element without removing it
func (h *Heap[T]) Peek() (T, bool) {
	if len(h.elements) == 0 {
		var zero T
		return zero, false
	}
	return h.elements[0], true
}

// Size returns the number of elements in the heap
func (h *Heap[T]) Size() int {
	return len(h.elements)
}

// IsEmpty returns true if the heap is empty
func (h *Heap[T]) IsEmpty() bool {
	return len(h.elements) == 0
}

// ToSlice returns all elements as a slice (not in heap order)
func (h *Heap[T]) ToSlice() []T {
	result := make([]T, len(h.elements))
	copy(result, h.elements)
	return result
}

// Clear removes all elements from the heap
func (h *Heap[T]) Clear() {
	h.elements = h.elements[:0]
}

// heapifyUp maintains heap property by moving element up
func (h *Heap[T]) heapifyUp(index int) {
	parentIndex := (index - 1) / 2
	if index > 0 && h.shouldSwap(index, parentIndex) {
		h.elements[index], h.elements[parentIndex] = h.elements[parentIndex], h.elements[index]
		h.heapifyUp(parentIndex)
	}
}

// heapifyDown maintains heap property by moving element down
func (h *Heap[T]) heapifyDown(index int) {
	leftChild := 2*index + 1
	rightChild := 2*index + 2
	targetIndex := index

	if leftChild < len(h.elements) && h.shouldSwap(leftChild, targetIndex) {
		targetIndex = leftChild
	}

	if rightChild < len(h.elements) && h.shouldSwap(rightChild, targetIndex) {
		targetIndex = rightChild
	}

	if targetIndex != index {
		h.elements[index], h.elements[targetIndex] = h.elements[targetIndex], h.elements[index]
		h.heapifyDown(targetIndex)
	}
}

// shouldSwap determines if two elements should be swapped based on heap type
func (h *Heap[T]) shouldSwap(childIndex, parentIndex int) bool {
	if h.heapType == MinHeap {
		return h.elements[childIndex] < h.elements[parentIndex]
	}
	return h.elements[childIndex] > h.elements[parentIndex]
}

// PriorityQueue is an alias for MaxHeap for common use case
type PriorityQueue[T constraints.Ordered] struct {
	*Heap[T]
}

// NewPriorityQueue creates a new priority queue (max heap)
func NewPriorityQueue[T constraints.Ordered]() *PriorityQueue[T] {
	return &PriorityQueue[T]{
		Heap: NewMaxHeap[T](),
	}
}

// Enqueue adds an element with priority
func (pq *PriorityQueue[T]) Enqueue(value T) {
	pq.Push(value)
}

// Dequeue removes and returns the highest priority element
func (pq *PriorityQueue[T]) Dequeue() T {
	return pq.Pop()
}