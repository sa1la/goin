package goin

// Stack 是一个基于切片的泛型 LIFO（后进先出）栈。
// 容量随 Push 自动扩展，Pop 不会立即收缩底层数组。
// 零值即可使用，无需通过构造函数初始化。
type Stack[T any] struct {
	elements []T
	size     int
}

// Push 将元素压入栈顶。
// 若底层切片仍有空位则原地复用，否则追加并由切片自动扩容。
func (stack *Stack[T]) Push(value T) {
	if stack.size < len(stack.elements) {
		stack.elements[stack.size] = value
	} else {
		stack.elements = append(stack.elements, value)
	}
	stack.size++
}

// Pop 弹出并返回栈顶元素。
// 当栈为空时返回类型零值，调用方需通过 IsEmpty 区分"弹出零值"与"栈空"两种情况。
func (stack *Stack[T]) Pop() T {
	if stack.size < 1 {
		var zero T
		return zero
	}
	value := stack.elements[stack.size-1]
	stack.size--
	return value
}

// Peek 返回栈顶元素而不弹出，第二个返回值表示栈是否非空。
func (stack *Stack[T]) Peek() (T, bool) {
	if stack.size < 1 {
		var zero T
		return zero, false
	}
	value := stack.elements[stack.size-1]
	return value, true
}

// IsEmpty 报告栈是否为空。
func (stack *Stack[T]) IsEmpty() bool {
	return stack.size == 0
}

// ToSlice 以栈底到栈顶的顺序返回内部元素的切片视图。
// 注意：返回的是底层数组的切片视图，调用方修改会影响栈内容。
func (stack *Stack[T]) ToSlice() []T {
	return stack.elements[:stack.size]
}
