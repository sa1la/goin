package goin

// Queue 是一个基于循环数组的泛型 FIFO（先进先出）队列。
// Enqueue/Dequeue 均为均摊 O(1)，容量不足时通过 resize 翻倍扩容。
type Queue[T any] struct {
	elements []T
	front    int
	rear     int
	size     int
}

// NewQueue 创建一个具有指定初始容量的队列。
// 当 capacity < 1 时，使用默认容量 10。
func NewQueue[T any](capacity int) *Queue[T] {
	if capacity < 1 {
		capacity = 10
	}
	return &Queue[T]{
		elements: make([]T, capacity),
		front:    0,
		rear:     0,
		size:     0,
	}
}

// Enqueue 在队尾添加一个元素。
// 当底层数组已满时，先调用 resize 将容量翻倍。
func (q *Queue[T]) Enqueue(value T) {
	if q.size == len(q.elements) {
		q.resize()
	}
	q.elements[q.rear] = value
	q.rear = (q.rear + 1) % len(q.elements)
	q.size++
}

// Dequeue 移除并返回队首元素。
// 当队列为空时返回类型零值，调用方需通过 IsEmpty 区分"取出零值"与"队空"两种情况。
func (q *Queue[T]) Dequeue() T {
	if q.size == 0 {
		var zero T
		return zero
	}
	value := q.elements[q.front]
	q.front = (q.front + 1) % len(q.elements)
	q.size--
	return value
}

// Front 返回队首元素而不移除，第二个返回值表示队列是否非空。
func (q *Queue[T]) Front() (T, bool) {
	if q.size == 0 {
		var zero T
		return zero, false
	}
	return q.elements[q.front], true
}

// IsEmpty 报告队列是否为空。
func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}

// Size 返回队列中元素的数量。
func (q *Queue[T]) Size() int {
	return q.size
}

// ToSlice 以 FIFO 顺序（队首到队尾）返回所有元素的拷贝切片。
func (q *Queue[T]) ToSlice() []T {
	if q.size == 0 {
		return []T{}
	}

	result := make([]T, q.size)
	for i := 0; i < q.size; i++ {
		result[i] = q.elements[(q.front+i)%len(q.elements)]
	}
	return result
}

// Clear 清空队列中的所有元素，但保留底层数组容量。
func (q *Queue[T]) Clear() {
	q.front = 0
	q.rear = 0
	q.size = 0
}

// resize 将队列容量翻倍，并把元素重新摆正为从下标 0 开始线性排列。
func (q *Queue[T]) resize() {
	newCapacity := len(q.elements) * 2
	newElements := make([]T, newCapacity)

	for i := 0; i < q.size; i++ {
		newElements[i] = q.elements[(q.front+i)%len(q.elements)]
	}

	q.elements = newElements
	q.front = 0
	q.rear = q.size
}
