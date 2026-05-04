package goin

import "golang.org/x/exp/constraints"

// Number is a constraint that matches all integer and float types.
type Number interface {
	constraints.Integer | constraints.Float
}

// FenwickTree implements a binary indexed tree for point updates and prefix/range queries.
type FenwickTree[T Number] struct {
	n   int // external visible length (indices 0..n-1)
	bit []T // internal array length n+1, index 0 unused (zero value)
}

// NewFenwickTree creates a new FenwickTree with n elements (indices 0..n-1).
// Panics if n < 0. Returns an empty tree when n == 0.
func NewFenwickTree[T Number](n int) *FenwickTree[T] {
	if n < 0 {
		panic("FenwickTree: n must be non-negative")
	}
	return &FenwickTree[T]{
		n:   n,
		bit: make([]T, n+1),
	}
}

// NewFenwickTreeFromSlice builds a FenwickTree from a 0-indexed slice in O(n).
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

// Update adds delta to the element at index i (0-indexed external API).
// Panics if i < 0 or i >= n.
func (f *FenwickTree[T]) Update(i int, delta T) {
	if i < 0 || i >= f.n {
		panic("FenwickTree: index out of bounds")
	}
	i++
	for i <= f.n {
		f.bit[i] += delta
		i += i & -i
	}
}

// Query returns the prefix sum a[0] + a[1] + ... + a[i].
// 0-indexed external API. Panics if i < 0 or i >= n.
func (f *FenwickTree[T]) Query(i int) T {
	if i < 0 || i >= f.n {
		panic("FenwickTree: index out of bounds")
	}
	var res T
	i++
	for i > 0 {
		res += f.bit[i]
		i -= i & -i
	}
	return res
}

// RangeQuery returns the interval sum a[l] + ... + a[r].
// Must have 0 <= l <= r < n, else panics.
func (f *FenwickTree[T]) RangeQuery(l, r int) T {
	if l < 0 || r >= f.n || l > r {
		panic("FenwickTree: index out of bounds")
	}
	if l == 0 {
		return f.Query(r)
	}
	return f.Query(r) - f.Query(l-1)
}
