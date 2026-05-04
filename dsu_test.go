package goin

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUnionFind(t *testing.T) {
	uf := NewUnionFind(5)
	assert.NotNil(t, uf)
	assert.Equal(t, 5, uf.Count())

	uf0 := NewUnionFind(0)
	assert.NotNil(t, uf0)
	assert.Equal(t, 0, uf0.Count())
}

func TestUnionFindFind(t *testing.T) {
	uf := NewUnionFind(5)
	for i := 0; i < 5; i++ {
		assert.Equal(t, i, uf.Find(i))
	}
}

func TestUnionFindUnion(t *testing.T) {
	uf := NewUnionFind(5)

	assert.True(t, uf.Union(0, 1))
	assert.Equal(t, 4, uf.Count())

	assert.False(t, uf.Union(0, 1))
	assert.Equal(t, 4, uf.Count())

	assert.True(t, uf.Union(2, 3))
	assert.Equal(t, 3, uf.Count())

	assert.True(t, uf.Union(1, 2))
	assert.Equal(t, 2, uf.Count())

	assert.True(t, uf.Union(0, 4))
	assert.Equal(t, 1, uf.Count())
}

func TestUnionFindIsConnected(t *testing.T) {
	uf := NewUnionFind(5)

	assert.False(t, uf.IsConnected(0, 1))
	assert.False(t, uf.IsConnected(2, 4))

	uf.Union(0, 1)
	assert.True(t, uf.IsConnected(0, 1))
	assert.True(t, uf.IsConnected(1, 0))
	assert.False(t, uf.IsConnected(0, 2))

	uf.Union(1, 2)
	assert.True(t, uf.IsConnected(0, 2))
	assert.True(t, uf.IsConnected(2, 0))
}

func TestUnionFindSize(t *testing.T) {
	uf := NewUnionFind(5)

	for i := 0; i < 5; i++ {
		assert.Equal(t, 1, uf.Size(i))
	}

	uf.Union(0, 1)
	assert.Equal(t, 2, uf.Size(0))
	assert.Equal(t, 2, uf.Size(1))

	uf.Union(1, 2)
	assert.Equal(t, 3, uf.Size(0))
	assert.Equal(t, 3, uf.Size(1))
	assert.Equal(t, 3, uf.Size(2))

	assert.Equal(t, 1, uf.Size(4))
}

func TestUnionFindReset(t *testing.T) {
	uf := NewUnionFind(5)
	uf.Union(0, 1)
	uf.Union(2, 3)
	assert.Equal(t, 3, uf.Count())
	assert.True(t, uf.IsConnected(0, 1))

	uf.Reset()
	assert.Equal(t, 5, uf.Count())
	for i := 0; i < 5; i++ {
		assert.Equal(t, i, uf.Find(i))
		assert.Equal(t, 1, uf.Size(i))
	}
	assert.False(t, uf.IsConnected(0, 1))
}

func TestUnionFindEdgeCases(t *testing.T) {
	uf0 := NewUnionFind(0)
	assert.Equal(t, 0, uf0.Count())

	uf1 := NewUnionFind(1)
	assert.Equal(t, 1, uf1.Count())
	assert.False(t, uf1.Union(0, 0))
	assert.Equal(t, 1, uf1.Count())
}

func TestUnionFindPanic(t *testing.T) {
	assert.Panics(t, func() { NewUnionFind(-1) })

	uf := NewUnionFind(5)
	assert.Panics(t, func() { uf.Find(-1) })
	assert.Panics(t, func() { uf.Find(5) })
	assert.Panics(t, func() { uf.Find(100) })

	assert.Panics(t, func() { uf.Union(-1, 0) })
	assert.Panics(t, func() { uf.Union(0, 5) })
	assert.Panics(t, func() { uf.Union(5, 5) })

	assert.Panics(t, func() { uf.Size(-1) })
	assert.Panics(t, func() { uf.Size(5) })
}

// bfsComponents labels every node with its component id (the smallest node in the component).
func bfsComponents(adj [][]int) []int {
	n := len(adj)
	comp := make([]int, n)
	for i := range comp {
		comp[i] = -1
	}
	q := NewQueue[int](16)
	for start := 0; start < n; start++ {
		if comp[start] != -1 {
			continue
		}
		comp[start] = start
		q.Enqueue(start)
		for !q.IsEmpty() {
			u := q.Dequeue()
			for _, v := range adj[u] {
				if comp[v] == -1 {
					comp[v] = start
					q.Enqueue(v)
				}
			}
		}
	}
	return comp
}

func TestUnionFindStress(t *testing.T) {
	n := 1000
	uf := NewUnionFind(n)
	adj := make([][]int, n)

	r := rand.New(rand.NewSource(42))

	for i := 0; i < 2000; i++ {
		a := r.Intn(n)
		b := r.Intn(n)
		if uf.Union(a, b) {
			adj[a] = append(adj[a], b)
			adj[b] = append(adj[b], a)
		}
	}

	comp := bfsComponents(adj)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			assert.Equal(t, comp[i] == comp[j], uf.IsConnected(i, j),
				"connectivity mismatch for pair (%d, %d)", i, j)
		}
	}
}
