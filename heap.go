package goin

import (
	"golang.org/x/exp/constraints"
)

// HeapType 表示堆的类型（最小堆或最大堆）。
type HeapType int

const (
	// MinHeap 表示最小堆，堆顶为最小元素。
	MinHeap HeapType = iota
	// MaxHeap 表示最大堆，堆顶为最大元素。
	MaxHeap
)

// Heap 是一个泛型二叉堆（优先队列），可作最小堆或最大堆使用。
// 元素类型需满足 constraints.Ordered。
type Heap[T constraints.Ordered] struct {
	elements []T
	heapType HeapType
}

// NewMinHeap 创建一个空的最小堆。
func NewMinHeap[T constraints.Ordered]() *Heap[T] {
	return &Heap[T]{
		elements: make([]T, 0),
		heapType: MinHeap,
	}
}

// NewMaxHeap 创建一个空的最大堆。
func NewMaxHeap[T constraints.Ordered]() *Heap[T] {
	return &Heap[T]{
		elements: make([]T, 0),
		heapType: MaxHeap,
	}
}

// Push 将元素压入堆中，时间复杂度 O(log n)。
func (h *Heap[T]) Push(value T) {
	h.elements = append(h.elements, value)
	h.heapifyUp(len(h.elements) - 1)
}

// Pop 弹出并返回堆顶元素（最小堆为最小值，最大堆为最大值），时间复杂度 O(log n)。
// 当堆为空时返回类型零值，调用方需通过 IsEmpty 区分"取出零值"与"堆空"两种情况。
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

// Peek 返回堆顶元素而不弹出，第二个返回值表示堆是否非空。
func (h *Heap[T]) Peek() (T, bool) {
	if len(h.elements) == 0 {
		var zero T
		return zero, false
	}
	return h.elements[0], true
}

// Size 返回堆中元素的数量。
func (h *Heap[T]) Size() int {
	return len(h.elements)
}

// IsEmpty 报告堆是否为空。
func (h *Heap[T]) IsEmpty() bool {
	return len(h.elements) == 0
}

// ToSlice 返回当前堆所有元素的拷贝切片（顺序为内部存储顺序，并非排序结果）。
func (h *Heap[T]) ToSlice() []T {
	result := make([]T, len(h.elements))
	copy(result, h.elements)
	return result
}

// Clear 清空堆中的所有元素，但保留底层数组容量。
func (h *Heap[T]) Clear() {
	h.elements = h.elements[:0]
}

// heapifyUp 通过递归上浮维护堆性质。
func (h *Heap[T]) heapifyUp(index int) {
	parentIndex := (index - 1) / 2
	if index > 0 && h.shouldSwap(index, parentIndex) {
		h.elements[index], h.elements[parentIndex] = h.elements[parentIndex], h.elements[index]
		h.heapifyUp(parentIndex)
	}
}

// heapifyDown 通过递归下沉维护堆性质。
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

// shouldSwap 根据堆类型判断子节点是否应当与父节点交换。
func (h *Heap[T]) shouldSwap(childIndex, parentIndex int) bool {
	if h.heapType == MinHeap {
		return h.elements[childIndex] < h.elements[parentIndex]
	}
	return h.elements[childIndex] > h.elements[parentIndex]
}

// PriorityQueue 是优先队列，对最大堆做的薄封装，提供 Enqueue/Dequeue 命名风格。
// 默认 Dequeue 返回当前最大元素；如需最小优先队列请直接使用 NewMinHeap。
type PriorityQueue[T constraints.Ordered] struct {
	*Heap[T]
}

// NewPriorityQueue 创建一个新的优先队列（基于最大堆）。
func NewPriorityQueue[T constraints.Ordered]() *PriorityQueue[T] {
	return &PriorityQueue[T]{
		Heap: NewMaxHeap[T](),
	}
}

// Enqueue 将元素加入优先队列。
func (pq *PriorityQueue[T]) Enqueue(value T) {
	pq.Push(value)
}

// Dequeue 弹出并返回当前优先级最高（默认即最大）的元素。
func (pq *PriorityQueue[T]) Dequeue() T {
	return pq.Pop()
}
