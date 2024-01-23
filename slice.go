package goin

import "sort"

// slice utility functions
func FillSlice[T any](s []T, v T) {
	for i := range s {
		s[i] = v
	}
}

func SortInts(slice []int) {
	sort.Ints(slice)
}

// permutation
// 向最大排列前进
func NextPermutation(nums []int) bool {
	n := len(nums)

	i := n - 2

	for i >= 0 && nums[i] >= nums[i+1] { //从右侧找到第一个比右边小的数

		i--
	}
	if i < 0 {
		//已经是最大排列
		return false
	}
	j := n - 1
	for j >= 0 && nums[j] <= nums[i] { //从右侧向左找到第一个比i大的数
		j--
	}
	nums[i], nums[j] = nums[j], nums[i]
	for k, l := i+1, n-1; k < l; k, l = k+1, l-1 { //跨过区域颠倒成最小排列
		nums[k], nums[l] = nums[l], nums[k]
	}
	return true
}

// 向最小排列方向前进
func LastPermutation(nums []int) bool {
	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i] <= nums[i+1] { //从右侧找到第一个比右边大的数
		i--
	}
	if i < 0 {
		//已经是最小排列了
		return false
	}
	j := n - 1
	for j >= 0 && nums[j] >= nums[i] { //从右侧找到第一个比小的数
		j--
	}
	nums[i], nums[j] = nums[j], nums[i]
	for k, l := i+1, n-1; k < l; k, l = k+1, l-1 { //跨过区域颠倒成最大排列
		nums[k], nums[l] = nums[l], nums[k]
	}
	return true
}
