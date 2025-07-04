package goin

// Queue implements a generic FIFO (First In, First Out) queue
type Queue[T any] struct {
	elements []T
	front    int
	rear     int
	size     int
}

// NewQueue creates a new queue with the specified initial capacity
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

// Enqueue adds an element to the rear of the queue
func (q *Queue[T]) Enqueue(value T) {
	if q.size == len(q.elements) {
		q.resize()
	}
	q.elements[q.rear] = value
	q.rear = (q.rear + 1) % len(q.elements)
	q.size++
}

// Dequeue removes and returns the front element of the queue
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

// Front returns the front element without removing it
func (q *Queue[T]) Front() (T, bool) {
	if q.size == 0 {
		var zero T
		return zero, false
	}
	return q.elements[q.front], true
}

// IsEmpty returns true if the queue is empty
func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}

// Size returns the number of elements in the queue
func (q *Queue[T]) Size() int {
	return q.size
}

// ToSlice returns all elements as a slice in FIFO order
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

// Clear removes all elements from the queue
func (q *Queue[T]) Clear() {
	q.front = 0
	q.rear = 0
	q.size = 0
}

// resize doubles the capacity of the queue
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