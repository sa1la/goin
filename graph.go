package goin

import (
	"container/heap"
	"math"
	"sort"
)

// NegInf 是 Bellman-Ford 中用于标记受可达负环影响的顶点的哨兵值。
// 取 math.MinInt64/4 而非 MinInt64 本身，是为了避免在加减边权时发生 int64 溢出。
const NegInf = math.MinInt64 / 4

// Graph 是无权图，支持无向边和有向边。
type Graph struct {
	n   int
	adj [][]int
}

// NewGraph 创建包含 n 个顶点（0..n-1）的空无权图。
func NewGraph(n int) *Graph {
	return &Graph{
		n:   n,
		adj: make([][]int, n),
	}
}

// NewGraphFromAdj 从邻接表 adj 创建无权图。
func NewGraphFromAdj(adj [][]int) *Graph {
	return &Graph{
		n:   len(adj),
		adj: adj,
	}
}

// N 返回图中的顶点数。
func (g *Graph) N() int {
	return g.n
}

// AddEdge 添加一条无向边 (u, v)。
func (g *Graph) AddEdge(u, v int) {
	g.adj[u] = append(g.adj[u], v)
	g.adj[v] = append(g.adj[v], u)
}

// AddDirectedEdge 添加一条有向边 u → v。
func (g *Graph) AddDirectedEdge(u, v int) {
	g.adj[u] = append(g.adj[u], v)
}

// Adj 返回顶点 v 的邻接表；v 越界时返回 nil。
func (g *Graph) Adj(v int) []int {
	if v < 0 || v >= g.n {
		return nil
	}
	return g.adj[v]
}

// WEdge 是带权图中的一条边，To 为目标顶点，W 为边权。
type WEdge struct {
	To int
	W  int
}

// WGraph 是带权图，支持无向边和有向边。
type WGraph struct {
	n   int
	adj [][]WEdge
}

// NewWGraph 创建包含 n 个顶点（0..n-1）的空带权图。
func NewWGraph(n int) *WGraph {
	return &WGraph{
		n:   n,
		adj: make([][]WEdge, n),
	}
}

// NewWGraphFromAdj 从邻接表 adj 创建带权图。
func NewWGraphFromAdj(adj [][]WEdge) *WGraph {
	return &WGraph{
		n:   len(adj),
		adj: adj,
	}
}

// N 返回带权图中的顶点数。
func (g *WGraph) N() int {
	return g.n
}

// AddEdge 添加一条无向边 (u, v, w)。
func (g *WGraph) AddEdge(u, v, w int) {
	g.adj[u] = append(g.adj[u], WEdge{To: v, W: w})
	g.adj[v] = append(g.adj[v], WEdge{To: u, W: w})
}

// AddDirectedEdge 添加一条有向边 u → v，权重为 w。
func (g *WGraph) AddDirectedEdge(u, v, w int) {
	g.adj[u] = append(g.adj[u], WEdge{To: v, W: w})
}

// Adj 返回顶点 v 的带权邻接表；v 越界时返回 nil。
func (g *WGraph) Adj(v int) []WEdge {
	if v < 0 || v >= g.n {
		return nil
	}
	return g.adj[v]
}

// BFS 从 start 出发进行无权图广度优先搜索，返回各顶点的最短距离。
// 不可达顶点距离为 -1；start 越界时所有距离均为 -1。
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

// DFS 从 start 出发进行深度优先搜索，返回访问顺序。
// start 越界时返回空切片。
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

// IsBipartite 判断图是否为二分图。
// 返回 (true, color) 表示是二分图，color 为每个顶点的 0/1 染色；否则返回 (false, color)。
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

// TopologicalSort 对有向图进行拓扑排序。
// 若图存在环则返回 (nil, false)；否则返回拓扑序和 true。
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

// SCCResult 存储强连通分量分解的结果。
type SCCResult struct {
	Components [][]int // 每个内层切片是一个 SCC
	Comp       []int   // Comp[v] 为顶点 v 所属的连通分量编号
}

// SCC 使用 Kosaraju 算法求强连通分量。
// 第一次 DFS 记录后序，第二次在反向图上按逆后序执行 DFS；第二次 DFS 的每棵树即为一个 SCC。
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

	// 复用 cur 缓冲区，避免大量小组件时总拷贝量达到 O(k²)。
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

// TreeDiameter 返回无权树的直径（最长简单路径上的边数）。
// 通过两次 BFS 实现。
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

// BFS01 在边权仅为 0 或 1 的带权图上运行 0-1 BFS，返回从 start 出发的最短距离。
// 不可达顶点距离为 -1；start 越界时所有距离均为 -1。
func (g *WGraph) BFS01(start int) []int {
	dist := make([]int, g.n)
	for i := range dist {
		dist[i] = -1
	}
	if start < 0 || start >= g.n {
		return dist
	}

	dist[start] = 0
	// 预分配切片并用前后双指针模拟双端队列，避免引入 container/deque。
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

// dijkstraHeap 是 Dijkstra 使用的局部最小堆，按距离排序。
// 直接使用 container/heap 而非现有的 Heap[T]，因为 Dijkstra 需要存储 (距离, 顶点) 二元组。
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

// Dijkstra 返回从 start 出发到各顶点的最短距离（非负权图），时间复杂度 O((V+E) log V)。
// 不可达顶点距离为 -1；start 越界时所有距离均为 -1。
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

// BellmanFord 返回从 start 出发的最短距离，以及一个 bool 表示 start 可达范围内是否不存在负环。
//
// 内部使用 inf（非 -1）进行松弛，避免将合法最短距离 -1 与"不可达"混淆。
// 处理完成后，不可达顶点在结果中仍表示为 -1；受负环影响的顶点标记为 NegInf。
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

// FloydWarshall 在 g 上运行 Floyd-Warshall 全源最短路，返回距离矩阵。
// 不可达点对值为 -1。
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

// FloydWarshall 在传入的距离矩阵 dist 上原地运行 Floyd-Warshall 全源最短路。
// dist[i][i] 应初始化为 0，不可达点对应初始化为一个足够大的哨兵值（如 -1）。返回 dist 本身。
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

// MSTKruskal 使用 Kruskal 算法求最小生成森林，返回总权重和选中的边集。
// e.To > u 的判定假设图通过 AddEdge（无向）构建，用于防止每条无向边被加入两次。
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

// TreeDiameter 返回带权树的直径（最长简单路径的权重和）。
// 通过两次 Dijkstra 实现。
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
