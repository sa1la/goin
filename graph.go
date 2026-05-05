package goin

import (
	"container/heap"
	"math"
	"sort"
)

// NegInf is a sentinel value used by Bellman-Ford to mark vertices
// affected by a reachable negative cycle. It is math.MinInt64/4 rather
// than MinInt64 itself so that adding/subtracting edge weights cannot
// overflow int64 in a single operation.
const NegInf = math.MinInt64 / 4

type Graph struct {
	n   int
	adj [][]int
}

func NewGraph(n int) *Graph {
	return &Graph{
		n:   n,
		adj: make([][]int, n),
	}
}

func NewGraphFromAdj(adj [][]int) *Graph {
	return &Graph{
		n:   len(adj),
		adj: adj,
	}
}

func (g *Graph) N() int {
	return g.n
}

func (g *Graph) AddEdge(u, v int) {
	g.adj[u] = append(g.adj[u], v)
	g.adj[v] = append(g.adj[v], u)
}

func (g *Graph) AddDirectedEdge(u, v int) {
	g.adj[u] = append(g.adj[u], v)
}

func (g *Graph) Adj(v int) []int {
	if v < 0 || v >= g.n {
		return nil
	}
	return g.adj[v]
}

type WEdge struct {
	To int
	W  int
}

type WGraph struct {
	n   int
	adj [][]WEdge
}

func NewWGraph(n int) *WGraph {
	return &WGraph{
		n:   n,
		adj: make([][]WEdge, n),
	}
}

func NewWGraphFromAdj(adj [][]WEdge) *WGraph {
	return &WGraph{
		n:   len(adj),
		adj: adj,
	}
}

func (g *WGraph) N() int {
	return g.n
}

func (g *WGraph) AddEdge(u, v, w int) {
	g.adj[u] = append(g.adj[u], WEdge{To: v, W: w})
	g.adj[v] = append(g.adj[v], WEdge{To: u, W: w})
}

func (g *WGraph) AddDirectedEdge(u, v, w int) {
	g.adj[u] = append(g.adj[u], WEdge{To: v, W: w})
}

func (g *WGraph) Adj(v int) []WEdge {
	if v < 0 || v >= g.n {
		return nil
	}
	return g.adj[v]
}

func (g *Graph) BFS(start int) []int {
	dist := make([]int, g.n)
	for i := range dist {
		dist[i] = -1
	}
	if start < 0 || start >= g.n {
		return dist
	}

	q := NewQueue[int](16)
	dist[start] = 0
	q.Enqueue(start)

	for !q.IsEmpty() {
		u := q.Dequeue()
		for _, v := range g.adj[u] {
			if dist[v] == -1 {
				dist[v] = dist[u] + 1
				q.Enqueue(v)
			}
		}
	}
	return dist
}

func (g *Graph) DFS(start int) []int {
	visited := make([]bool, g.n)
	order := make([]int, 0, g.n)

	if start < 0 || start >= g.n {
		return order
	}

	var dfs func(int)
	dfs = func(u int) {
		visited[u] = true
		order = append(order, u)
		for _, v := range g.adj[u] {
			if !visited[v] {
				dfs(v)
			}
		}
	}

	dfs(start)
	return order
}

func (g *Graph) IsBipartite() (bool, []int) {
	color := make([]int, g.n)
	for i := range color {
		color[i] = -1
	}

	for start := 0; start < g.n; start++ {
		if color[start] != -1 {
			continue
		}

		q := NewQueue[int](16)
		color[start] = 0
		q.Enqueue(start)

		for !q.IsEmpty() {
			u := q.Dequeue()
			for _, v := range g.adj[u] {
				if color[v] == -1 {
					color[v] = 1 - color[u]
					q.Enqueue(v)
				} else if color[v] == color[u] {
					return false, color
				}
			}
		}
	}
	return true, color
}

func (g *Graph) TopologicalSort() ([]int, bool) {
	indeg := make([]int, g.n)
	for u := 0; u < g.n; u++ {
		for _, v := range g.adj[u] {
			indeg[v]++
		}
	}

	q := NewQueue[int](16)
	for i := 0; i < g.n; i++ {
		if indeg[i] == 0 {
			q.Enqueue(i)
		}
	}

	order := make([]int, 0, g.n)
	for !q.IsEmpty() {
		u := q.Dequeue()
		order = append(order, u)
		for _, v := range g.adj[u] {
			indeg[v]--
			if indeg[v] == 0 {
				q.Enqueue(v)
			}
		}
	}

	if len(order) != g.n {
		return nil, false
	}
	return order, true
}

type SCCResult struct {
	Components [][]int // each inner slice is one SCC
	Comp       []int   // Comp[v] = component ID of vertex v
}

// SCC finds strongly connected components using Kosaraju's algorithm.
// First DFS records vertices in post-order; second DFS runs on the
// reversed graph in reverse post-order. Each tree in the second DFS
// is one SCC.
func (g *Graph) SCC() SCCResult {
	visited := make([]bool, g.n)
	order := make([]int, 0, g.n)

	var dfs1 func(int)
	dfs1 = func(u int) {
		visited[u] = true
		for _, v := range g.adj[u] {
			if !visited[v] {
				dfs1(v)
			}
		}
		order = append(order, u)
	}

	for i := 0; i < g.n; i++ {
		if !visited[i] {
			dfs1(i)
		}
	}

	rev := make([][]int, g.n)
	for u := 0; u < g.n; u++ {
		for _, v := range g.adj[u] {
			rev[v] = append(rev[v], u)
		}
	}

	comp := make([]int, g.n)
	for i := range comp {
		comp[i] = -1
	}
	components := make([][]int, 0)

	// Reuse cur buffer across DFS calls to avoid O(k²) total copying
	// when there are many small components.
	cur := make([]int, 0, g.n)
	var dfs2 func(int)
	dfs2 = func(u int) {
		comp[u] = len(components)
		cur = append(cur, u)
		for _, v := range rev[u] {
			if comp[v] == -1 {
				dfs2(v)
			}
		}
	}

	for i := g.n - 1; i >= 0; i-- {
		v := order[i]
		if comp[v] == -1 {
			cur = cur[:0]
			dfs2(v)
			components = append(components, append([]int(nil), cur...))
		}
	}

	return SCCResult{
		Components: components,
		Comp:       comp,
	}
}

func (g *Graph) TreeDiameter() int {
	if g.n == 0 {
		return 0
	}

	dist := g.BFS(0)
	far := 0
	for i := 1; i < g.n; i++ {
		if dist[i] > dist[far] {
			far = i
		}
	}

	dist2 := g.BFS(far)
	maxDist := 0
	for i := 0; i < g.n; i++ {
		if dist2[i] > maxDist {
			maxDist = dist2[i]
		}
	}
	return maxDist
}

func (g *WGraph) BFS01(start int) []int {
	dist := make([]int, g.n)
	for i := range dist {
		dist[i] = -1
	}
	if start < 0 || start >= g.n {
		return dist
	}

	dist[start] = 0
	// Preallocate a slice and use front/back indices as a deque.
	// This avoids importing container/deque, which matters in CP
	// where every import line counts against the source size limit.
	deque := make([]int, g.n*2)
	front, back := g.n, g.n
	deque[back] = start
	back++

	for front < back {
		u := deque[front]
		front++
		for _, e := range g.adj[u] {
			v := e.To
			w := e.W
			if dist[v] == -1 || dist[u]+w < dist[v] {
				dist[v] = dist[u] + w
				if w == 0 {
					front--
					deque[front] = v
				} else {
					deque[back] = v
					back++
				}
			}
		}
	}
	return dist
}

// dijkstraHeap is a local min-heap ordered by distance.
// We use container/heap directly instead of the existing Heap[T]
// because that type stores single values; Dijkstra needs
// (distance, vertex) pairs ordered by distance.
type dijkstraItem struct {
	dist int
	v    int
}

type dijkstraHeap []dijkstraItem

func (h dijkstraHeap) Len() int           { return len(h) }
func (h dijkstraHeap) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h dijkstraHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *dijkstraHeap) Push(x any) {
	*h = append(*h, x.(dijkstraItem))
}

func (h *dijkstraHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func (g *WGraph) Dijkstra(start int) []int {
	dist := make([]int, g.n)
	for i := range dist {
		dist[i] = -1
	}
	if start < 0 || start >= g.n {
		return dist
	}

	dist[start] = 0
	h := &dijkstraHeap{{dist: 0, v: start}}
	heap.Init(h)

	for h.Len() > 0 {
		item := heap.Pop(h).(dijkstraItem)
		u := item.v
		if item.dist > dist[u] {
			continue
		}
		for _, e := range g.adj[u] {
			v := e.To
			w := e.W
			if dist[v] == -1 || dist[u]+w < dist[v] {
				dist[v] = dist[u] + w
				heap.Push(h, dijkstraItem{dist: dist[v], v: v})
			}
		}
	}
	return dist
}

// BellmanFord returns shortest distances from start and a bool indicating
// whether any negative cycle is reachable from start.
//
// An internal inf constant (not -1) is used during relaxation so that
// a legitimate shortest-path distance of -1 is not confused with
// "unreachable". After processing, unreachable vertices are translated
// back to -1 in the returned slice.
func (g *WGraph) BellmanFord(start int) ([]int, bool) {
	const inf = math.MaxInt64 / 4

	dist := make([]int, g.n)
	for i := range dist {
		dist[i] = inf
	}
	if start < 0 || start >= g.n {
		result := make([]int, g.n)
		for i := range result {
			result[i] = -1
		}
		return result, true
	}

	dist[start] = 0
	for i := 0; i < g.n-1; i++ {
		updated := false
		for u := 0; u < g.n; u++ {
			if dist[u] == inf {
				continue
			}
			for _, e := range g.adj[u] {
				v := e.To
				w := e.W
				if dist[u]+w < dist[v] {
					dist[v] = dist[u] + w
					updated = true
				}
			}
		}
		if !updated {
			break
		}
	}

	inNegativeCycle := make([]bool, g.n)
	q := make([]int, 0)
	for u := 0; u < g.n; u++ {
		if dist[u] == inf {
			continue
		}
		for _, e := range g.adj[u] {
			if dist[u]+e.W < dist[e.To] {
				if !inNegativeCycle[e.To] {
					inNegativeCycle[e.To] = true
					q = append(q, e.To)
				}
			}
		}
	}

	head := 0
	for head < len(q) {
		u := q[head]
		head++
		for _, e := range g.adj[u] {
			if !inNegativeCycle[e.To] {
				inNegativeCycle[e.To] = true
				q = append(q, e.To)
			}
		}
	}

	hasNegCycle := false
	result := make([]int, g.n)
	for i := 0; i < g.n; i++ {
		if inNegativeCycle[i] {
			result[i] = NegInf
			hasNegCycle = true
		} else if dist[i] == inf {
			result[i] = -1
		} else {
			result[i] = dist[i]
		}
	}

	return result, !hasNegCycle
}

func (g *WGraph) FloydWarshall() [][]int {
	dist := make([][]int, g.n)
	for i := 0; i < g.n; i++ {
		dist[i] = make([]int, g.n)
		for j := 0; j < g.n; j++ {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = -1
			}
		}
	}

	for u := 0; u < g.n; u++ {
		for _, e := range g.adj[u] {
			v := e.To
			w := e.W
			if dist[u][v] == -1 || w < dist[u][v] {
				dist[u][v] = w
			}
		}
	}

	return FloydWarshall(dist)
}

// FloydWarshall computes all-pairs shortest paths in-place on the given distance matrix.
func FloydWarshall(dist [][]int) [][]int {
	n := len(dist)
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			if dist[i][k] == -1 {
				continue
			}
			for j := 0; j < n; j++ {
				if dist[k][j] == -1 {
					continue
				}
				if dist[i][j] == -1 || dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}
	return dist
}

// MSTKruskal returns the total weight and edges of a minimum spanning forest.
// The e.To > u guard assumes the graph was built with AddEdge (undirected);
// it keeps each undirected edge from being added twice.
func (g *WGraph) MSTKruskal() (int, []WEdge) {
	type edge struct {
		u, v int
		w    int
	}

	edges := make([]edge, 0)
	for u := 0; u < g.n; u++ {
		for _, e := range g.adj[u] {
			if e.To > u {
				edges = append(edges, edge{u: u, v: e.To, w: e.W})
			}
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].w < edges[j].w
	})

	uf := NewUnionFind(g.n)
	total := 0
	mstEdges := make([]WEdge, 0)

	for _, e := range edges {
		if uf.Union(e.u, e.v) {
			total += e.w
			mstEdges = append(mstEdges, WEdge{To: e.v, W: e.w})
		}
	}

	return total, mstEdges
}

func (g *WGraph) TreeDiameter() int {
	if g.n == 0 {
		return 0
	}

	dist := g.Dijkstra(0)
	far := 0
	for i := 1; i < g.n; i++ {
		if dist[i] > dist[far] {
			far = i
		}
	}

	dist2 := g.Dijkstra(far)
	maxDist := 0
	for i := 0; i < g.n; i++ {
		if dist2[i] > maxDist {
			maxDist = dist2[i]
		}
	}
	return maxDist
}
