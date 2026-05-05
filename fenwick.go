package goin

import "golang.org/x/exp/constraints"

// Number 约束所有整数和浮点类型。
type Number interface {
	constraints.Integer | constraints.Float
}

// FenwickTree 实现二叉索引树（树状数组），支持单点更新和前缀/区间查询。
// 外部 API 使用 0-based 下标，内部数组 bit 使用 1-based（bit[0] 恒为零值）。
type FenwickTree[T Number] struct {
	n   int
	bit []T
}

// NewFenwickTree 创建一个长度为 n 的树状数组，初始全为零值，外部下标范围 0..n-1。
// n < 0 时 panic；n == 0 时返回空树。
func NewFenwickTree[T Number](n int) *FenwickTree[T] {
	if n < 0 {
		panic("FenwickTree: n must be non-negative")
	}
	return &FenwickTree[T]{
		n:   n,
		bit: make([]T, n+1),
	}
}

// NewFenwickTreeFromSlice 从 0-based 切片 a 构建树状数组，时间复杂度 O(n)。
func NewFenwickTreeFromSlice[T Number](a []T) *FenwickTree[T] {
	n := len(a)
	bit := make([]T, n+1)
	copy(bit[1:], a)
	for i := 1; i <= n; i++ {
		j := i + (i & -i)
		if j <= n {
			bit[j] += bit[i]
		}
	}
	return &FenwickTree[T]{
		n:   n,
		bit: bit,
	}
}

// Update 将下标 i（0-based）处的值增加 delta。
// i 越界时 panic。
func (f *FenwickTree[T]) Update(i int, delta T) {
	f.validate(i)
	i++
	for i <= f.n {
		f.bit[i] += delta
		i += i & -i
	}
}

// Query 返回前缀和 a[0] + a[1] + ... + a[i]，i 为 0-based。
// i 越界时 panic。
func (f *FenwickTree[T]) Query(i int) T {
	f.validate(i)
	var res T
	i++
	for i > 0 {
		res += f.bit[i]
		i -= i & -i
	}
	return res
}

// RangeQuery 返回区间和 a[l] + ... + a[r]，要求 0 <= l <= r < n。
// 下标越界或 l > r 时 panic。
func (f *FenwickTree[T]) RangeQuery(l, r int) T {
	f.validate(l)
	f.validate(r)
	if l > r {
		panic("FenwickTree: index out of bounds")
	}
	if l == 0 {
		return f.Query(r)
	}
	return f.Query(r) - f.Query(l-1)
}

// validate 检查 i 是否在合法范围，否则 panic。
func (f *FenwickTree[T]) validate(i int) {
	if i < 0 || i >= f.n {
		panic("FenwickTree: index out of bounds")
	}
}
