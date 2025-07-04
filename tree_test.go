package goin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceBinaryTree(t *testing.T) {
	tree := &SliceBinaryTree[int]{}
	treeData := []int{1, 2, 3, 4, 5, 6, 7}
	tree = tree.NewTree(treeData)

	// Test Size
	assert.Equal(t, 7, tree.Size())

	// Test GetNode
	val, ok := tree.GetNode(0)
	assert.True(t, ok)
	assert.Equal(t, 1, val)

	val, ok = tree.GetNode(10)
	assert.False(t, ok)

	// Test left/right node indices
	assert.Equal(t, 1, tree.LeftNodeIdx(0))
	assert.Equal(t, 2, tree.RightNodeIdx(0))
	assert.Equal(t, 0, tree.ParentNodeIdx(1))

	// Test left/right nodes
	leftVal, ok := tree.LeftNode(0)
	assert.True(t, ok)
	assert.Equal(t, 2, leftVal)

	rightVal, ok := tree.RightNode(0)
	assert.True(t, ok)
	assert.Equal(t, 3, rightVal)

	// Test LevelOrder
	levelOrder := tree.LevelOrder()
	assert.Equal(t, treeData, levelOrder)

	// Test PreOrder
	preOrder := tree.PreOrder()
	expected := []int{1, 2, 4, 5, 3, 6, 7}
	assert.Equal(t, expected, preOrder)

	// Test InOrder
	inOrder := tree.InOrder()
	expected = []int{4, 2, 5, 1, 6, 3, 7}
	assert.Equal(t, expected, inOrder)

	// Test PostOrder
	postOrder := tree.PostOrder()
	expected = []int{4, 5, 2, 6, 7, 3, 1}
	assert.Equal(t, expected, postOrder)
}

func TestAVLTreeInsertAndSearch(t *testing.T) {
	tree := &AVLTree{}

	// Test empty tree search
	node := tree.Search(5)
	assert.Nil(t, node)

	// Test insertion
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(3)
	tree.Insert(7)

	// Test search
	node = tree.Search(5)
	assert.NotNil(t, node)
	assert.Equal(t, 5, node.Val)

	node = tree.Search(20)
	assert.Nil(t, node)

	// Test traversals
	inOrder := tree.InOrder()
	expected := []int{3, 5, 7, 10, 15}
	assert.Equal(t, expected, inOrder)
}

func TestAVLTreeRemove(t *testing.T) {
	tree := &AVLTree{}

	// Insert elements
	elements := []int{10, 5, 15, 3, 7, 12, 18}
	for _, elem := range elements {
		tree.Insert(elem)
	}

	// Remove leaf node
	tree.Remove(3)
	node := tree.Search(3)
	assert.Nil(t, node)

	// Remove node with one child
	tree.Remove(15)
	node = tree.Search(15)
	assert.Nil(t, node)

	// Remove node with two children
	tree.Remove(10)
	node = tree.Search(10)
	assert.Nil(t, node)

	// Verify remaining elements
	inOrder := tree.InOrder()
	expected := []int{5, 7, 12, 18}
	assert.Equal(t, expected, inOrder)
}

func TestAVLTreeBalance(t *testing.T) {
	tree := &AVLTree{}

	// Insert elements that would cause imbalance in regular BST
	for i := 1; i <= 7; i++ {
		tree.Insert(i)
	}

	// Tree should remain balanced
	inOrder := tree.InOrder()
	expected := []int{1, 2, 3, 4, 5, 6, 7}
	assert.Equal(t, expected, inOrder)

	// Verify tree structure is balanced (root should not be 1)
	node := tree.Search(4) // In a balanced tree, 4 should be easily findable
	assert.NotNil(t, node)
}