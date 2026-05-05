package goin

import (
	"golang.org/x/exp/constraints"
)

// FillSlice 将切片 s 中所有元素填充为 v。
func FillSlice[T any](s []T, v T) {
	for i := range s {
		s[i] = v
	}
}

// Reverse 原地反转切片 s。
func Reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// NextPermutation 将 s 原地变换为字典序的下一个排列。
// 若已为最大排列（降序），则反转为最小排列并返回 false；否则返回 true。
// 时间复杂度 O(n)。
func NextPermutation[S ~[]E, E constraints.Ordered](s S) bool {
	i := len(s) - 2
	for i >= 0 && s[i] >= s[i+1] {
		i--
	}

	if i < 0 {
		Reverse(s)
		return false
	}

	j := len(s) - 1
	for j > i && s[i] >= s[j] {
		j--
	}

	s[i], s[j] = s[j], s[i]
	Reverse(s[i+1:])

	return true
}

// LastPermutation 将 s 原地变换为字典序的上一个排列。
// 若已为最小排列（升序），则反转为最大排列并返回 false；否则返回 true。
// 时间复杂度 O(n)。
func LastPermutation[S ~[]E, E constraints.Ordered](s S) bool {
	i := len(s) - 2
	for i >= 0 && s[i] <= s[i+1] {
		i--
	}

	if i < 0 {
		Reverse(s)
		return false
	}

	j := len(s) - 1
	for j > i && s[i] <= s[j] {
		j--
	}

	s[i], s[j] = s[j], s[i]
	Reverse(s[i+1:])

	return true
}
