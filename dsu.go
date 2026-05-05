package goin

// UnionFind 实现并查集（Disjoint Set Union），支持路径压缩和按大小合并。
type UnionFind struct {
	parent []int
	size   []int
	count  int
}

// NewUnionFind 创建包含 n 个元素（0..n-1）的并查集。
// n < 0 时 panic；n == 0 时返回空并查集。
func NewUnionFind(n int) *UnionFind {
	if n < 0 {
		panic("UnionFind: n must be non-negative")
	}
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &UnionFind{
		parent: parent,
		size:   size,
		count:  n,
	}
}

// Find 返回 x 所在集合的根节点，同时执行迭代路径压缩。
// x 越界时 panic。
func (u *UnionFind) Find(x int) int {
	u.validate(x)
	root := x
	for u.parent[root] != root {
		root = u.parent[root]
	}
	for u.parent[x] != x {
		parent := u.parent[x]
		u.parent[x] = root
		x = parent
	}
	return root
}

// Union 合并 x 和 y 所在的集合，采用按大小合并（小集合挂到大集合）。
// 若两元素已在同一集合则返回 false；合并成功返回 true。
// x 或 y 越界时 panic。
func (u *UnionFind) Union(x, y int) bool {
	rx := u.Find(x)
	ry := u.Find(y)
	if rx == ry {
		return false
	}
	if u.size[rx] < u.size[ry] {
		u.parent[rx] = ry
		u.size[ry] += u.size[rx]
	} else {
		u.parent[ry] = rx
		u.size[rx] += u.size[ry]
	}
	u.count--
	return true
}

// IsConnected 判断 x 和 y 是否属于同一集合。
func (u *UnionFind) IsConnected(x, y int) bool {
	return u.Find(x) == u.Find(y)
}

// Size 返回 x 所在连通分量的大小。
// x 越界时 panic。
func (u *UnionFind) Size(x int) int {
	return u.size[u.Find(x)]
}

// Count 返回当前连通分量的数量。
func (u *UnionFind) Count() int {
	return u.count
}

// Reset 将并查集恢复为初始状态，时间复杂度 O(n)。
func (u *UnionFind) Reset() {
	for i := range u.parent {
		u.parent[i] = i
		u.size[i] = 1
	}
	u.count = len(u.parent)
}

// validate 检查 x 是否在合法范围内，否则 panic。
func (u *UnionFind) validate(x int) {
	if x < 0 || x >= len(u.parent) {
		panic("UnionFind: index out of bounds")
	}
}
