package goin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFenwickTree(t *testing.T) {
	ft := NewFenwickTree[int](5)
	assert.NotNil(t, ft)
	assert.Equal(t, 0, ft.Query(0))
	assert.Equal(t, 0, ft.Query(2))
	assert.Equal(t, 0, ft.Query(4))

	ft0 := NewFenwickTree[int](0)
	assert.NotNil(t, ft0)
}

func TestFenwickTreeUpdateAndQuery(t *testing.T) {
	ft := NewFenwickTree[int](5)

	for i := 0; i < 5; i++ {
		assert.Equal(t, 0, ft.Query(i))
	}

	ft.Update(0, 3)
	assert.Equal(t, 3, ft.Query(0))
	assert.Equal(t, 3, ft.Query(1))
	assert.Equal(t, 3, ft.Query(4))

	ft.Update(2, 5)
	assert.Equal(t, 3, ft.Query(0))
	assert.Equal(t, 3, ft.Query(1))
	assert.Equal(t, 8, ft.Query(2))
	assert.Equal(t, 8, ft.Query(3))
	assert.Equal(t, 8, ft.Query(4))

	ft.Update(4, 2)
	assert.Equal(t, 3, ft.Query(0))
	assert.Equal(t, 3, ft.Query(1))
	assert.Equal(t, 8, ft.Query(2))
	assert.Equal(t, 8, ft.Query(3))
	assert.Equal(t, 10, ft.Query(4))

	ft.Update(0, 2)
	assert.Equal(t, 5, ft.Query(0))
	assert.Equal(t, 12, ft.Query(4))
}

func TestFenwickTreeRangeQuery(t *testing.T) {
	ft := NewFenwickTree[int](5)
	ft.Update(0, 3)
	ft.Update(1, 2)
	ft.Update(2, 5)
	ft.Update(3, 1)
	ft.Update(4, 4)

	assert.Equal(t, 3, ft.RangeQuery(0, 0))
	assert.Equal(t, 5, ft.RangeQuery(0, 1))
	assert.Equal(t, 15, ft.RangeQuery(0, 4))

	assert.Equal(t, 2, ft.RangeQuery(1, 1))
	assert.Equal(t, 7, ft.RangeQuery(1, 2))
	assert.Equal(t, 8, ft.RangeQuery(1, 3))
	assert.Equal(t, 10, ft.RangeQuery(2, 4))
	assert.Equal(t, 4, ft.RangeQuery(4, 4))
}

func TestFenwickTreeFromSlice(t *testing.T) {
	a := []int{3, 2, 5, 1, 4}
	ft := NewFenwickTreeFromSlice(a)

	assert.Equal(t, 3, ft.Query(0))
	assert.Equal(t, 5, ft.Query(1))
	assert.Equal(t, 10, ft.Query(2))
	assert.Equal(t, 11, ft.Query(3))
	assert.Equal(t, 15, ft.Query(4))

	assert.Equal(t, 7, ft.RangeQuery(1, 2))
	assert.Equal(t, 10, ft.RangeQuery(2, 4))

	ft.Update(0, 2)
	assert.Equal(t, 5, ft.Query(0))
	assert.Equal(t, 17, ft.Query(4))
}

func TestFenwickTreeNegative(t *testing.T) {
	ft := NewFenwickTree[int](5)
	ft.Update(0, 5)
	ft.Update(1, -3)
	ft.Update(2, -2)
	ft.Update(3, 4)
	ft.Update(4, -1)

	assert.Equal(t, 5, ft.Query(0))
	assert.Equal(t, 2, ft.Query(1))
	assert.Equal(t, 0, ft.Query(2))
	assert.Equal(t, 4, ft.Query(3))
	assert.Equal(t, 3, ft.Query(4))

	assert.Equal(t, -5, ft.RangeQuery(1, 2))
	assert.Equal(t, 3, ft.RangeQuery(3, 4))
}

func TestFenwickTreeInt64(t *testing.T) {
	ft := NewFenwickTree[int64](5)
	ft.Update(0, 1000000000)
	ft.Update(1, 2000000000)
	ft.Update(2, 3000000000)

	assert.Equal(t, int64(1000000000), ft.Query(0))
	assert.Equal(t, int64(3000000000), ft.Query(1))
	assert.Equal(t, int64(6000000000), ft.Query(2))
	assert.Equal(t, int64(6000000000), ft.RangeQuery(0, 2))
	assert.Equal(t, int64(2000000000), ft.RangeQuery(1, 1))
}

func TestFenwickTreeFloat64(t *testing.T) {
	ft := NewFenwickTree[float64](5)
	ft.Update(0, 1.5)
	ft.Update(1, 2.5)
	ft.Update(2, 3.0)
	ft.Update(3, -1.0)
	ft.Update(4, 0.5)

	assert.InDelta(t, 1.5, ft.Query(0), 1e-9)
	assert.InDelta(t, 4.0, ft.Query(1), 1e-9)
	assert.InDelta(t, 7.0, ft.Query(2), 1e-9)
	assert.InDelta(t, 6.0, ft.Query(3), 1e-9)
	assert.InDelta(t, 6.5, ft.Query(4), 1e-9)

	assert.InDelta(t, 7.0, ft.RangeQuery(0, 2), 1e-9)
	assert.InDelta(t, -0.5, ft.RangeQuery(3, 4), 1e-9)
}

func TestFenwickTreeEdgeCases(t *testing.T) {
	ft0 := NewFenwickTree[int](0)
	assert.NotNil(t, ft0)

	ft1 := NewFenwickTree[int](1)
	assert.Equal(t, 0, ft1.Query(0))
	ft1.Update(0, 42)
	assert.Equal(t, 42, ft1.Query(0))
	assert.Equal(t, 42, ft1.RangeQuery(0, 0))

	ftEmpty := NewFenwickTreeFromSlice([]int{})
	assert.NotNil(t, ftEmpty)

	ftSingle := NewFenwickTreeFromSlice([]int{7})
	assert.Equal(t, 7, ftSingle.Query(0))
}

func TestFenwickTreePanic(t *testing.T) {
	assert.Panics(t, func() { NewFenwickTree[int](-1) })

	ft := NewFenwickTree[int](5)

	assert.Panics(t, func() { ft.Update(-1, 1) })
	assert.Panics(t, func() { ft.Update(5, 1) })
	assert.Panics(t, func() { ft.Update(100, 1) })

	assert.Panics(t, func() { ft.Query(-1) })
	assert.Panics(t, func() { ft.Query(5) })
	assert.Panics(t, func() { ft.Query(100) })

	assert.Panics(t, func() { ft.RangeQuery(3, 2) })

	assert.Panics(t, func() { ft.RangeQuery(-1, 2) })
	assert.Panics(t, func() { ft.RangeQuery(0, 5) })
	assert.Panics(t, func() { ft.RangeQuery(0, 100) })
	assert.Panics(t, func() { ft.RangeQuery(5, 5) })
}
