package goin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraphBase(t *testing.T) {
	g := NewGraph(4)
	assert.Equal(t, 4, g.N())

	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddDirectedEdge(2, 3)

	assert.ElementsMatch(t, []int{1}, g.Adj(0))
	assert.ElementsMatch(t, []int{0, 2}, g.Adj(1))
	assert.ElementsMatch(t, []int{1, 3}, g.Adj(2))
	assert.Nil(t, g.Adj(3))
}

func TestGraphFromAdj(t *testing.T) {
	adj := [][]int{{1}, {0, 2}, {1}, {}}
	g := NewGraphFromAdj(adj)
	assert.Equal(t, 4, g.N())
	assert.Equal(t, []int{1}, g.Adj(0))
	assert.Equal(t, []int{0, 2}, g.Adj(1))
}

func TestWGraphBase(t *testing.T) {
	g := NewWGraph(4)
	assert.Equal(t, 4, g.N())

	g.AddEdge(0, 1, 5)
	g.AddDirectedEdge(2, 3, 7)

	assert.Equal(t, []WEdge{{To: 1, W: 5}}, g.Adj(0))
	assert.Equal(t, []WEdge{{To: 0, W: 5}}, g.Adj(1))
	assert.Equal(t, []WEdge{{To: 3, W: 7}}, g.Adj(2))
	assert.Nil(t, g.Adj(3))
}

func TestWGraphFromAdj(t *testing.T) {
	adj := [][]WEdge{{{To: 1, W: 5}}, {{To: 0, W: 5}}, {{}}, {{}}}
	g := NewWGraphFromAdj(adj)
	assert.Equal(t, 4, g.N())
	assert.Equal(t, []WEdge{{To: 1, W: 5}}, g.Adj(0))
}

func TestEmptyGraph(t *testing.T) {
	g := NewGraph(0)
	assert.Equal(t, 0, g.N())
	assert.Nil(t, g.Adj(0))

	wg := NewWGraph(0)
	assert.Equal(t, 0, wg.N())
	assert.Nil(t, wg.Adj(0))
}

func TestGraphBFS(t *testing.T) {
	// Linear chain: 0-1-2-3
	g := NewGraph(4)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)

	dist := g.BFS(0)
	assert.Equal(t, []int{0, 1, 2, 3}, dist)

	// Disconnected: 0-1, 2-3
	g2 := NewGraph(4)
	g2.AddEdge(0, 1)
	g2.AddEdge(2, 3)

	dist2 := g2.BFS(0)
	assert.Equal(t, []int{0, 1, -1, -1}, dist2)

	// Single node
	g3 := NewGraph(1)
	assert.Equal(t, []int{0}, g3.BFS(0))

	// Empty graph
	g4 := NewGraph(0)
	assert.Equal(t, []int{}, g4.BFS(0))
}

func TestGraphDFS(t *testing.T) {
	// Linear chain: 0-1-2-3
	g := NewGraph(4)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)

	order := g.DFS(0)
	assert.Equal(t, []int{0, 1, 2, 3}, order)

	// Cyclic graph: 0-1, 1-2, 2-0 (triangle)
	g2 := NewGraph(3)
	g2.AddEdge(0, 1)
	g2.AddEdge(1, 2)
	g2.AddEdge(2, 0)

	order2 := g2.DFS(0)
	assert.Equal(t, 3, len(order2))
	assert.Contains(t, order2, 0)
	assert.Contains(t, order2, 1)
	assert.Contains(t, order2, 2)

	// Single node
	g3 := NewGraph(1)
	assert.Equal(t, []int{0}, g3.DFS(0))

	// Empty graph
	g4 := NewGraph(0)
	assert.Equal(t, []int{}, g4.DFS(0))
}

func TestGraphIsBipartite(t *testing.T) {
	// Bipartite: 0-1, 1-2, 2-3 (chain)
	g := NewGraph(4)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)

	ok, color := g.IsBipartite()
	assert.True(t, ok)
	assert.Equal(t, 4, len(color))
	assert.NotEqual(t, color[0], color[1])
	assert.NotEqual(t, color[1], color[2])
	assert.NotEqual(t, color[2], color[3])

	// Non-bipartite: triangle
	g2 := NewGraph(3)
	g2.AddEdge(0, 1)
	g2.AddEdge(1, 2)
	g2.AddEdge(2, 0)

	ok2, _ := g2.IsBipartite()
	assert.False(t, ok2)

	// Single node (bipartite)
	g3 := NewGraph(1)
	ok3, color3 := g3.IsBipartite()
	assert.True(t, ok3)
	assert.Equal(t, []int{0}, color3)

	// Empty graph (bipartite)
	g4 := NewGraph(0)
	ok4, color4 := g4.IsBipartite()
	assert.True(t, ok4)
	assert.Equal(t, []int{}, color4)
}

func TestGraphTopologicalSort(t *testing.T) {
	// DAG: 0 -> 1 -> 2, 0 -> 2
	g := NewGraph(3)
	g.AddDirectedEdge(0, 1)
	g.AddDirectedEdge(1, 2)
	g.AddDirectedEdge(0, 2)

	order, ok := g.TopologicalSort()
	assert.True(t, ok)
	assert.Equal(t, []int{0, 1, 2}, order)

	// Cyclic graph
	g2 := NewGraph(3)
	g2.AddDirectedEdge(0, 1)
	g2.AddDirectedEdge(1, 2)
	g2.AddDirectedEdge(2, 0)

	_, ok2 := g2.TopologicalSort()
	assert.False(t, ok2)

	// Empty graph
	g3 := NewGraph(0)
	order3, ok3 := g3.TopologicalSort()
	assert.True(t, ok3)
	assert.Equal(t, []int{}, order3)

	// Single node
	g4 := NewGraph(1)
	order4, ok4 := g4.TopologicalSort()
	assert.True(t, ok4)
	assert.Equal(t, []int{0}, order4)
}

func TestGraphSCC(t *testing.T) {
	// Graph: 0 -> 1 -> 2 -> 0 (SCC {0,1,2}), 3 <-> 4 (SCC {3,4})
	g := NewGraph(5)
	g.AddDirectedEdge(0, 1)
	g.AddDirectedEdge(1, 2)
	g.AddDirectedEdge(2, 0)
	g.AddDirectedEdge(3, 4)
	g.AddDirectedEdge(4, 3)

	res := g.SCC()
	assert.Equal(t, 2, len(res.Components))

	// Verify Comp mapping
	for _, compID := range res.Comp {
		assert.True(t, compID >= 0 && compID < 2)
	}
	// 0,1,2 should be in same component
	assert.Equal(t, res.Comp[0], res.Comp[1])
	assert.Equal(t, res.Comp[1], res.Comp[2])
	// 3,4 should be in same component
	assert.Equal(t, res.Comp[3], res.Comp[4])
	// {0,1,2} and {3,4} should be different components
	assert.NotEqual(t, res.Comp[0], res.Comp[3])

	// Single node
	g2 := NewGraph(1)
	res2 := g2.SCC()
	assert.Equal(t, 1, len(res2.Components))
	assert.Equal(t, []int{0}, res2.Components[0])

	// Empty graph
	g3 := NewGraph(0)
	res3 := g3.SCC()
	assert.Equal(t, 0, len(res3.Components))
}

func TestGraphTreeDiameter(t *testing.T) {
	// Chain: 0-1-2-3-4 (diameter = 4 edges)
	g := NewGraph(5)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	assert.Equal(t, 4, g.TreeDiameter())

	// Star: center 0, leaves 1,2,3,4 (diameter = 2 edges)
	g2 := NewGraph(5)
	g2.AddEdge(0, 1)
	g2.AddEdge(0, 2)
	g2.AddEdge(0, 3)
	g2.AddEdge(0, 4)
	assert.Equal(t, 2, g2.TreeDiameter())

	// Single node
	g3 := NewGraph(1)
	assert.Equal(t, 0, g3.TreeDiameter())
}

func TestWGraphBFS01(t *testing.T) {
	g := NewWGraph(5)
	g.AddDirectedEdge(0, 1, 0)
	g.AddDirectedEdge(1, 2, 0)
	g.AddDirectedEdge(0, 3, 1)
	g.AddDirectedEdge(1, 4, 1)
	g.AddDirectedEdge(3, 4, 0)

	dist := g.BFS01(0)
	assert.Equal(t, []int{0, 0, 0, 1, 1}, dist)

	g2 := NewWGraph(4)
	g2.AddDirectedEdge(0, 1, 0)
	g2.AddDirectedEdge(2, 3, 1)

	dist2 := g2.BFS01(0)
	assert.Equal(t, []int{0, 0, -1, -1}, dist2)

	g3 := NewWGraph(1)
	assert.Equal(t, []int{0}, g3.BFS01(0))
}

func TestWGraphDijkstra(t *testing.T) {
	g := NewWGraph(4)
	g.AddDirectedEdge(0, 1, 4)
	g.AddDirectedEdge(1, 2, 1)
	g.AddDirectedEdge(0, 3, 2)
	g.AddDirectedEdge(3, 2, 3)

	dist := g.Dijkstra(0)
	assert.Equal(t, []int{0, 4, 5, 2}, dist)

	g2 := NewWGraph(4)
	g2.AddDirectedEdge(0, 1, 1)
	dist2 := g2.Dijkstra(0)
	assert.Equal(t, []int{0, 1, -1, -1}, dist2)

	g3 := NewWGraph(1)
	assert.Equal(t, []int{0}, g3.Dijkstra(0))
}

func TestWGraphBellmanFord(t *testing.T) {
	g := NewWGraph(4)
	g.AddDirectedEdge(0, 1, 5)
	g.AddDirectedEdge(1, 2, -3)
	g.AddDirectedEdge(2, 3, 3)

	dist, ok := g.BellmanFord(0)
	assert.True(t, ok)
	assert.Equal(t, []int{0, 5, 2, 5}, dist)

	g2 := NewWGraph(3)
	g2.AddDirectedEdge(0, 1, 1)
	g2.AddDirectedEdge(1, 2, -1)
	g2.AddDirectedEdge(2, 1, -1)

	dist2, ok2 := g2.BellmanFord(0)
	assert.False(t, ok2)
	assert.Equal(t, 0, dist2[0])
	assert.Equal(t, NegInf, dist2[1])
	assert.Equal(t, NegInf, dist2[2])

	g3 := NewWGraph(4)
	g3.AddDirectedEdge(0, 1, 1)
	dist3, ok3 := g3.BellmanFord(0)
	assert.True(t, ok3)
	assert.Equal(t, []int{0, 1, -1, -1}, dist3)
}

func TestWGraphFloydWarshall(t *testing.T) {
	g := NewWGraph(3)
	g.AddDirectedEdge(0, 1, 2)
	g.AddDirectedEdge(1, 2, 3)
	g.AddDirectedEdge(0, 2, 6)

	dist := g.FloydWarshall()
	assert.Equal(t, []int{0, 2, 5}, dist[0])
	assert.Equal(t, []int{-1, 0, 3}, dist[1])
	assert.Equal(t, []int{-1, -1, 0}, dist[2])

	// Negative cycle: 0->1 (1), 1->2 (-3), 2->0 (1) => total = -1 < 0
	g2 := NewWGraph(3)
	g2.AddDirectedEdge(0, 1, 1)
	g2.AddDirectedEdge(1, 2, -3)
	g2.AddDirectedEdge(2, 0, 1)

	dist2 := g2.FloydWarshall()
	assert.True(t, dist2[0][0] < 0)
}

func TestFloydWarshallStandalone(t *testing.T) {
	dist := [][]int{
		{0, 2, 6},
		{-1, 0, 3},
		{-1, -1, 0},
	}

	result := FloydWarshall(dist)
	assert.Equal(t, []int{0, 2, 5}, result[0])
	assert.Equal(t, []int{-1, 0, 3}, result[1])
	assert.Equal(t, []int{-1, -1, 0}, result[2])
}

func TestWGraphMSTKruskal(t *testing.T) {
	g := NewWGraph(4)
	g.AddEdge(0, 1, 1)
	g.AddEdge(0, 2, 3)
	g.AddEdge(1, 2, 2)
	g.AddEdge(1, 3, 4)

	total, edges := g.MSTKruskal()
	assert.Equal(t, 7, total)
	assert.Equal(t, 3, len(edges))

	g2 := NewWGraph(4)
	g2.AddEdge(0, 1, 1)
	g2.AddEdge(2, 3, 2)

	total2, edges2 := g2.MSTKruskal()
	assert.Equal(t, 3, total2)
	assert.Equal(t, 2, len(edges2))

	g3 := NewWGraph(0)
	total3, edges3 := g3.MSTKruskal()
	assert.Equal(t, 0, total3)
	assert.Equal(t, 0, len(edges3))

	g4 := NewWGraph(1)
	total4, edges4 := g4.MSTKruskal()
	assert.Equal(t, 0, total4)
	assert.Equal(t, 0, len(edges4))
}

func TestWGraphTreeDiameter(t *testing.T) {
	g := NewWGraph(4)
	g.AddEdge(0, 1, 2)
	g.AddEdge(1, 2, 3)
	g.AddEdge(2, 3, 4)
	assert.Equal(t, 9, g.TreeDiameter())

	g2 := NewWGraph(5)
	g2.AddEdge(0, 1, 5)
	g2.AddEdge(0, 2, 3)
	g2.AddEdge(0, 3, 2)
	g2.AddEdge(0, 4, 1)
	assert.Equal(t, 8, g2.TreeDiameter())

	g3 := NewWGraph(1)
	assert.Equal(t, 0, g3.TreeDiameter())
}
