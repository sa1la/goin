package goin

import (
	"golang.org/x/exp/constraints"
)

func FillSlice[T any](s []T, v T) {
	for i := range s {
		s[i] = v
	}
}
func Reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

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
