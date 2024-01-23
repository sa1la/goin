package goin

//slice-binary tree

type SliceBinaryTree[T any] struct {
	tree []T //suggest: use with base type like: int,float,string,rune,byte
}

func (bTree *SliceBinaryTree[T]) NewTree(treeArr []T) *SliceBinaryTree[T] {
	return &SliceBinaryTree[T]{tree: treeArr}
}

func (bTree *SliceBinaryTree[T]) Size() int {
	return len(bTree.tree)
}
func (bTree *SliceBinaryTree[T]) GetNode(idx int) (T, bool) {
	if idx < 0 || idx >= bTree.Size() {
		var zero T
		return zero, false
	}
	return bTree.tree[idx], true
}
func (bTree *SliceBinaryTree[T]) LeftNodeIdx(idx int) int {
	return 2*idx + 1
}
func (bTree *SliceBinaryTree[T]) LeftNode(idx int) (T, bool) {
	return bTree.GetNode(bTree.LeftNodeIdx(idx))
}
func (bTree *SliceBinaryTree[T]) RightNodeIdx(idx int) int {
	return 2*idx + 2
}
func (bTree *SliceBinaryTree[T]) RightNode(idx int) (T, bool) {
	return bTree.GetNode(bTree.RightNodeIdx(idx))
}
func (bTree *SliceBinaryTree[T]) ParentNodeIdx(idx int) int {
	return (idx - 1) / 2
}
func (bTree *SliceBinaryTree[T]) ParentNode(idx int) (T, bool) {
	return bTree.GetNode(bTree.ParentNodeIdx(idx))
}

func (bTree *SliceBinaryTree[T]) LevelOrder() []T {
	var res []T
	// 直接遍历数组
	for i := 0; i < bTree.Size(); i++ {
		if node, success := bTree.GetNode(i); success {
			res = append(res, node)
		}
	}
	return res
}

type OrderType string

const (
	PRE  OrderType = "pre"
	IN   OrderType = "in"
	POST OrderType = "post"
)

/* 深度优先遍历 */
func (bTree *SliceBinaryTree[T]) DFS(i int, order OrderType, res *[]T) {
	node, success := bTree.GetNode(i)
	// 若为空位，则返回
	if !success {
		return
	}

	// 前序遍历
	if order == PRE {
		*res = append(*res, node)
	}
	//left node
	bTree.DFS(bTree.LeftNodeIdx(i), order, res)

	// 中序遍历
	if order == IN {
		*res = append(*res, node)
	}
	//right node
	bTree.DFS(bTree.RightNodeIdx(i), order, res)

	// 后序遍历
	if order == POST {
		*res = append(*res, node)
	}
}

/* 前序遍历 */
func (abt *SliceBinaryTree[T]) PreOrder() []T {
	var res []T
	abt.DFS(0, PRE, &res)
	return res
}

/* 中序遍历 */
func (abt *SliceBinaryTree[T]) InOrder() []T {
	var res []T
	abt.DFS(0, IN, &res)
	return res
}

/* 后序遍历 */
func (abt *SliceBinaryTree[T]) PostOrder() []T {
	var res []T
	abt.DFS(0, POST, &res)
	return res
}

// binary search tree
type TreeNode struct {
	Val    int       // 节点值
	Height int       // 节点高度
	Left   *TreeNode // 左子节点引用
	Right  *TreeNode // 右子节点引用
}
type AVLTree struct {
	root *TreeNode
}

/* 获取节点高度 */
func (t *AVLTree) getHeight(node *TreeNode) int {
	// 空节点高度为 -1 ，叶节点高度为 0
	if node != nil {
		return node.Height
	}
	return -1
}

/* 更新节点高度 */
func (t *AVLTree) updateHeight(node *TreeNode) {
	lh := t.getHeight(node.Left)
	rh := t.getHeight(node.Right)
	// 节点高度等于最高子树高度 + 1
	if lh > rh {
		node.Height = lh + 1
	} else {
		node.Height = rh + 1
	}
}

/* 获取平衡因子 */
func (t *AVLTree) balanceFactor(node *TreeNode) int {
	// 空节点平衡因子为 0
	if node == nil {
		return 0
	}
	// 节点平衡因子 = 左子树高度 - 右子树高度
	return t.getHeight(node.Left) - t.getHeight(node.Right)
}

/* 右旋操作 */
func (t *AVLTree) rightRotate(node *TreeNode) *TreeNode {
	child := node.Left
	grandChild := child.Right
	// 以 child 为原点，将 node 向右旋转
	child.Right = node
	node.Left = grandChild
	// 更新节点高度
	t.updateHeight(node)
	t.updateHeight(child)
	// 返回旋转后子树的根节点
	return child
}

/* 左旋操作 */
func (t *AVLTree) leftRotate(node *TreeNode) *TreeNode {
	child := node.Right
	grandChild := child.Left
	// 以 child 为原点，将 node 向左旋转
	child.Left = node
	node.Right = grandChild
	// 更新节点高度
	t.updateHeight(node)
	t.updateHeight(child)
	// 返回旋转后子树的根节点
	return child
}

/* 执行旋转操作，使该子树重新恢复平衡 */
func (t *AVLTree) rotate(node *TreeNode) *TreeNode {
	// 获取节点 node 的平衡因子
	// Go 推荐短变量，这里 bf 指代 t.balanceFactor
	bf := t.balanceFactor(node)
	// 左偏树
	if bf > 1 {
		if t.balanceFactor(node.Left) >= 0 {
			// 右旋
			return t.rightRotate(node)
		} else {
			// 先左旋后右旋
			node.Left = t.leftRotate(node.Left)
			return t.rightRotate(node)
		}
	}
	// 右偏树
	if bf < -1 {
		if t.balanceFactor(node.Right) <= 0 { // 左旋
			return t.leftRotate(node)
		} else {
			// 先右旋后左旋
			node.Right = t.rightRotate(node.Right)
			return t.leftRotate(node)
		}
	}
	// 平衡树，无须旋转，直接返回
	return node
}

/* 查找节点 */
func (bst *AVLTree) Search(num int) *TreeNode {
	node := bst.root
	// 循环查找，越过叶节点后跳出
	for node != nil {
		if node.Val == num {
			// 找到目标节点，跳出循环
			break
		}
		if node.Val < num {
			// 目标节点在 cur 的右子树中
			node = node.Right
			continue
		}
		// 目标节点在 cur 的左子树中
		node = node.Left
	}
	// 返回目标节点
	return node
}

/* 插入节点 */
func (t *AVLTree) Insert(val int) {
	t.root = t.insertHelper(t.root, val)
}

/* 递归插入节点(辅助函数) */
func (t *AVLTree) insertHelper(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return &TreeNode{Val: val}
	}
	/* 1. 查找插入位置，并插入节点 */
	if val < node.Val {

		node.Left = t.insertHelper(node.Left, val)
	} else if val > node.Val {
		node.Right = t.insertHelper(node.Right, val)
	} else {
		// 重复节点不插入，直接返回
		return node
	}
	// 更新节点高度
	t.updateHeight(node)
	/* 2. 执行旋转操作，使该子树重新恢复平衡 */
	node = t.rotate(node)
	// 返回子树的根节点
	return node
}

/* 删除节点 */
func (t *AVLTree) Remove(val int) {
	t.root = t.removeHelper(t.root, val)
}

/* 递归删除节点(辅助函数) */
func (t *AVLTree) removeHelper(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return nil
	}
	/* 1. 查找节点，并删除之 */
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
				// 子节点数量 = 0 ，直接删除 node 并返回
				return nil

			} else {
				// 子节点数量 = 1 ，直接删除 node
				node = child
			}
		} else {
			// 子节点数量 = 2 ，则将中序遍历的下个节点删除，并用该节点替换当前节点
			temp := node.Right
			for temp.Left != nil {
				temp = temp.Left
			}
			node.Right = t.removeHelper(node.Right, temp.Val)
			node.Val = temp.Val
		}
	}
	// 更新节点高度
	t.updateHeight(node)
	/* 2. 执行旋转操作，使该子树重新恢复平衡 */
	node = t.rotate(node)
	// 返回子树的根节点
	return node
}

/* 深度优先遍历 */
func (bTree *AVLTree) DFS(node *TreeNode, order OrderType, res *[]int) {
	// 若为空位，则返回
	if node == nil {
		return
	}

	// 前序遍历
	if order == PRE {
		*res = append(*res, node.Val)
	}
	//left node
	bTree.DFS(node.Left, order, res)

	// 中序遍历
	if order == IN {
		*res = append(*res, node.Val)
	}
	//right node
	bTree.DFS(node.Right, order, res)

	// 后序遍历
	if order == POST {
		*res = append(*res, node.Val)
	}
}

/* 前序遍历 */
func (abt *AVLTree) PreOrder() []int {
	var res []int
	abt.DFS(abt.root, PRE, &res)
	return res
}

/* 中序遍历, 从小到大 */
func (abt *AVLTree) InOrder() []int {
	var res []int
	abt.DFS(abt.root, IN, &res)
	return res
}

/* 后序遍历 */
func (abt *AVLTree) PostOrder() []int {
	var res []int
	abt.DFS(abt.root, POST, &res)
	return res
}
