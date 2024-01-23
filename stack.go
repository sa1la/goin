package goin

// Stack
type Stack[T any] struct {
	elements []T
	size     int
}

func (stack *Stack[T]) Push(value T) {
	if stack.size < len(stack.elements) {
		stack.elements[stack.size] = value
	} else {
		stack.elements = append(stack.elements, value)
	}
	stack.size++
}

func (stack *Stack[T]) Pop() T {
	if stack.size < 1 {
		var zero T
		return zero
	}
	value := stack.elements[stack.size-1]
	stack.size--
	return value
}

func (stack *Stack[T]) Peek() (T, bool) {
	if stack.size < 1 {
		var zero T
		return zero, false
	}
	value := stack.elements[stack.size-1]
	return value, true
}

func (stack *Stack[T]) IsEmpty() bool {
	return stack.size == 0
}

func (stack *Stack[T]) ToSlice() []T {
	return stack.elements[:stack.size]
}
