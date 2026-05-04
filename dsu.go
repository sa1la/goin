package goin

// UnionFind implements the disjoint-set union (DSU) data structure
// with path compression and union by size.
type UnionFind struct {
	parent []int
	size   []int
	count  int
}

// NewUnionFind creates a new UnionFind with n elements (0..n-1).
// Panics if n < 0. Returns an empty DSU when n == 0.
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

// Find returns the root of x with iterative path compression.
// Panics if x is out of bounds.
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

// Union merges the sets containing x and y.
// Uses union by size (smaller attaches to larger).
// If sizes are equal, y's root attaches to x's root.
// Returns true if a merge happened, false if already in the same set.
// Panics if x or y is out of bounds.
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

// IsConnected returns true if x and y are in the same set.
func (u *UnionFind) IsConnected(x, y int) bool {
	return u.Find(x) == u.Find(y)
}

// Size returns the size of the component containing x.
// Panics if x is out of bounds.
func (u *UnionFind) Size(x int) int {
	return u.size[u.Find(x)]
}

// Count returns the number of disjoint components.
func (u *UnionFind) Count() int {
	return u.count
}

// Reset restores the UnionFind to its initial state. O(n).
func (u *UnionFind) Reset() {
	for i := range u.parent {
		u.parent[i] = i
		u.size[i] = 1
	}
	u.count = len(u.parent)
}

// validate panics if x is out of bounds.
func (u *UnionFind) validate(x int) {
	if x < 0 || x >= len(u.parent) {
		panic("UnionFind: index out of bounds")
	}
}
