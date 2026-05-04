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

	// Union two different sets
	assert.True(t, uf.Union(0, 1))
	assert.Equal(t, 4, uf.Count())

	// Union same set again should return false
	assert.False(t, uf.Union(0, 1))
	assert.Equal(t, 4, uf.Count())

	// Union another pair
	assert.True(t, uf.Union(2, 3))
	assert.Equal(t, 3, uf.Count())

	// Chain union: connect {0,1} with {2,3}
	assert.True(t, uf.Union(1, 2))
	assert.Equal(t, 2, uf.Count())

	// Union with isolated element
	assert.True(t, uf.Union(0, 4))
	assert.Equal(t, 1, uf.Count())
}

func TestUnionFindIsConnected(t *testing.T) {
	uf := NewUnionFind(5)

	// Before any unions
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

	// Isolated element
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
	// n = 0
	uf0 := NewUnionFind(0)
	assert.Equal(t, 0, uf0.Count())

	// n = 1, self-union returns false
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

// bfsConnected checks if two nodes are connected using BFS on an adjacency list.
func bfsConnected(adj [][]int, start, target int) bool {
	if start == target {
		return true
	}
	visited := make([]bool, len(adj))
	queue := []int{start}
	visited[start] = true
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		for _, v := range adj[u] {
			if v == target {
				return true
			}
			if !visited[v] {
				visited[v] = true
				queue = append(queue, v)
			}
		}
	}
	return false
}

func TestUnionFindStress(t *testing.T) {
	n := 1000
	uf := NewUnionFind(n)
	adj := make([][]int, n)

	r := rand.New(rand.NewSource(42))

	// Perform random unions
	for i := 0; i < 2000; i++ {
		a := r.Intn(n)
		b := r.Intn(n)
		merged := uf.Union(a, b)
		if merged {
			adj[a] = append(adj[a], b)
			adj[b] = append(adj[b], a)
		}
	}

	// Verify connectivity for every pair
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			assert.Equal(t, bfsConnected(adj, i, j), uf.IsConnected(i, j),
				"connectivity mismatch for pair (%d, %d)", i, j)
		}
	}
}
