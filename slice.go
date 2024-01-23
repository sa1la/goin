package goin

func FillSlice[T any](s []T, v T) {
	for i := range s {
		s[i] = v
	}
}

func NextPermutation(nums []int) bool {
	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i < 0 {
		return false
	}
	j := n - 1
	for j >= 0 && nums[j] <= nums[i] {
		j--
	}
	nums[i], nums[j] = nums[j], nums[i]
	for k, l := i+1, n-1; k < l; k, l = k+1, l-1 {
		nums[k], nums[l] = nums[l], nums[k]
	}
	return true
}

func LastPermutation(nums []int) bool {
	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i] <= nums[i+1] {
		i--
	}
	if i < 0 {
		return false
	}
	j := n - 1
	for j >= 0 && nums[j] >= nums[i] {
		j--
	}
	nums[i], nums[j] = nums[j], nums[i]
	for k, l := i+1, n-1; k < l; k, l = k+1, l-1 {
		nums[k], nums[l] = nums[l], nums[k]
	}
	return true
}
