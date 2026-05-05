package goin

// SliceBinaryTree 是基于切片的二叉树（堆式数组存储），适合完全二叉树或近似完全二叉树场景。
// 通过下标公式实现父子节点跳转，无需额外指针。
type SliceBinaryTree[T any] struct {
	tree []T
}

// NewTree 从给定的数组创建 SliceBinaryTree。
func (bTree *SliceBinaryTree[T]) NewTree(treeArr []T) *SliceBinaryTree[T] {
	return &SliceBinaryTree[T]{tree: treeArr}
}

// Size 返回树中节点数量。
func (bTree *SliceBinaryTree[T]) Size() int {
	return len(bTree.tree)
}

// GetNode 返回下标 idx 处的节点值，第二个返回值表示下标是否有效。
func (bTree *SliceBinaryTree[T]) GetNode(idx int) (T, bool) {
	if idx < 0 || idx >= bTree.Size() {
		var zero T
		return zero, false
	}
	return bTree.tree[idx], true
}

// LeftNodeIdx 返回节点 idx 的左子节点下标（2*idx + 1）。
func (bTree *SliceBinaryTree[T]) LeftNodeIdx(idx int) int {
	return 2*idx + 1
}

// LeftNode 返回节点 idx 的左子节点值及其存在性。
func (bTree *SliceBinaryTree[T]) LeftNode(idx int) (T, bool) {
	return bTree.GetNode(bTree.LeftNodeIdx(idx))
}

// RightNodeIdx 返回节点 idx 的右子节点下标（2*idx + 2）。
func (bTree *SliceBinaryTree[T]) RightNodeIdx(idx int) int {
	return 2*idx + 2
}

// RightNode 返回节点 idx 的右子节点值及其存在性。
func (bTree *SliceBinaryTree[T]) RightNode(idx int) (T, bool) {
	return bTree.GetNode(bTree.RightNodeIdx(idx))
}

// ParentNodeIdx 返回节点 idx 的父节点下标（(idx-1)/2）。
func (bTree *SliceBinaryTree[T]) ParentNodeIdx(idx int) int {
	return (idx - 1) / 2
}

// ParentNode 返回节点 idx 的父节点值及其存在性。
func (bTree *SliceBinaryTree[T]) ParentNode(idx int) (T, bool) {
	return bTree.GetNode(bTree.ParentNodeIdx(idx))
}

// LevelOrder 返回按层序遍历（数组顺序）得到的所有节点。
func (bTree *SliceBinaryTree[T]) LevelOrder() []T {
	var res []T
	for i := 0; i < bTree.Size(); i++ {
		if node, success := bTree.GetNode(i); success {
			res = append(res, node)
		}
	}
	return res
}

// OrderType 表示二叉树深度优先遍历的顺序。
type OrderType string

const (
	// PRE 表示前序遍历。
	PRE OrderType = "pre"
	// IN 表示中序遍历。
	IN OrderType = "in"
	// POST 表示后序遍历。
	POST OrderType = "post"
)

// DFS 从节点 i 开始按指定 order 进行深度优先遍历，结果追加写入 res。
func (bTree *SliceBinaryTree[T]) DFS(i int, order OrderType, res *[]T) {
	node, success := bTree.GetNode(i)
	if !success {
		return
	}

	if order == PRE {
		*res = append(*res, node)
	}
	bTree.DFS(bTree.LeftNodeIdx(i), order, res)

	if order == IN {
		*res = append(*res, node)
	}
	bTree.DFS(bTree.RightNodeIdx(i), order, res)

	if order == POST {
		*res = append(*res, node)
	}
}

// PreOrder 返回整树的前序遍历结果。
func (abt *SliceBinaryTree[T]) PreOrder() []T {
	var res []T
	abt.DFS(0, PRE, &res)
	return res
}

// InOrder 返回整树的中序遍历结果。
func (abt *SliceBinaryTree[T]) InOrder() []T {
	var res []T
	abt.DFS(0, IN, &res)
	return res
}

// PostOrder 返回整树的后序遍历结果。
func (abt *SliceBinaryTree[T]) PostOrder() []T {
	var res []T
	abt.DFS(0, POST, &res)
	return res
}

// TreeNode 是 AVL 树的节点，值为 int，包含高度、左右子节点指针。
type TreeNode struct {
	Val    int
	Height int
	Left   *TreeNode
	Right  *TreeNode
}

// AVLTree 是 int 类型的自平衡二叉搜索树。
type AVLTree struct {
	root *TreeNode
}

// getHeight 返回节点高度；空节点高度为 -1，叶节点高度为 0。
func (t *AVLTree) getHeight(node *TreeNode) int {
	if node != nil {
		return node.Height
	}
	return -1
}

// updateHeight 重新计算 node 的高度为左右子树最高者 + 1。
func (t *AVLTree) updateHeight(node *TreeNode) {
	lh := t.getHeight(node.Left)
	rh := t.getHeight(node.Right)
	if lh > rh {
		node.Height = lh + 1
	} else {
		node.Height = rh + 1
	}
}

// balanceFactor 返回 node 的平衡因子（左子树高度 - 右子树高度）。
// 空节点的平衡因子为 0。
func (t *AVLTree) balanceFactor(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return t.getHeight(node.Left) - t.getHeight(node.Right)
}

// rightRotate 对 node 执行右旋操作，返回旋转后的子树根节点。
func (t *AVLTree) rightRotate(node *TreeNode) *TreeNode {
	child := node.Left
	grandChild := child.Right
	child.Right = node
	node.Left = grandChild
	t.updateHeight(node)
	t.updateHeight(child)
	return child
}

// leftRotate 对 node 执行左旋操作，返回旋转后的子树根节点。
func (t *AVLTree) leftRotate(node *TreeNode) *TreeNode {
	child := node.Right
	grandChild := child.Left
	child.Left = node
	node.Right = grandChild
	t.updateHeight(node)
	t.updateHeight(child)
	return child
}

// rotate 检查 node 的平衡因子并执行必要的旋转，使子树恢复平衡。
func (t *AVLTree) rotate(node *TreeNode) *TreeNode {
	bf := t.balanceFactor(node)
	if bf > 1 {
		if t.balanceFactor(node.Left) >= 0 {
			return t.rightRotate(node)
		}
		node.Left = t.leftRotate(node.Left)
		return t.rightRotate(node)
	}
	if bf < -1 {
		if t.balanceFactor(node.Right) <= 0 {
			return t.leftRotate(node)
		}
		node.Right = t.rightRotate(node.Right)
		return t.leftRotate(node)
	}
	return node
}

// Search 在 AVL 树中查找值为 num 的节点，若不存在则返回 nil。
func (bst *AVLTree) Search(num int) *TreeNode {
	node := bst.root
	for node != nil {
		if node.Val == num {
			break
		}
		if node.Val < num {
			node = node.Right
		} else {
			node = node.Left
		}
	}
	return node
}

// Insert 向 AVL 树中插入值为 val 的节点。若已存在则不重复插入。
func (t *AVLTree) Insert(val int) {
	t.root = t.insertHelper(t.root, val)
}

// insertHelper 递归插入节点并重新平衡子树。
func (t *AVLTree) insertHelper(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return &TreeNode{Val: val}
	}
	if val < node.Val {
		node.Left = t.insertHelper(node.Left, val)
	} else if val > node.Val {
		node.Right = t.insertHelper(node.Right, val)
	} else {
		return node
	}
	t.updateHeight(node)
	node = t.rotate(node)
	return node
}

// Remove 从 AVL 树中删除值为 val 的节点。若不存在则无操作。
func (t *AVLTree) Remove(val int) {
	t.root = t.removeHelper(t.root, val)
}

// removeHelper 递归删除节点并重新平衡子树。
func (t *AVLTree) removeHelper(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return nil
	}
	if val < node.Val {
		node.Left = t.removeHelper(node.Left, val)
	} else if val > node.Val {
		node.Right = t.removeHelper(node.Right, val)
	} else {
		if node.Left == nil || node.Right == nil {
			child := node.Left
			if node.Right != nil {
				child = node.Right
			}
			if child == nil {
				return nil
			}
			node = child
		} else {
			temp := node.Right
			for temp.Left != nil {
				temp = temp.Left
			}
			node.Right = t.removeHelper(node.Right, temp.Val)
			node.Val = temp.Val
		}
	}
	t.updateHeight(node)
	node = t.rotate(node)
	return node
}

// DFS 从 node 开始按指定 order 深度优先遍历 AVL 树，结果追加写入 res。
func (bTree *AVLTree) DFS(node *TreeNode, order OrderType, res *[]int) {
	if node == nil {
		return
	}

	if order == PRE {
		*res = append(*res, node.Val)
	}
	bTree.DFS(node.Left, order, res)

	if order == IN {
		*res = append(*res, node.Val)
	}
	bTree.DFS(node.Right, order, res)

	if order == POST {
		*res = append(*res, node.Val)
	}
}

// PreOrder 返回 AVL 树的前序遍历结果。
func (abt *AVLTree) PreOrder() []int {
	var res []int
	abt.DFS(abt.root, PRE, &res)
	return res
}

// InOrder 返回 AVL 树的中序遍历结果（升序）。
func (abt *AVLTree) InOrder() []int {
	var res []int
	abt.DFS(abt.root, IN, &res)
	return res
}

// PostOrder 返回 AVL 树的后序遍历结果。
func (abt *AVLTree) PostOrder() []int {
	var res []int
	abt.DFS(abt.root, POST, &res)
	return res
}
