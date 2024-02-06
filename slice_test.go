package goin

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextPermutation(t *testing.T) {
	data := sort.IntSlice{1, 2, 3}
	expected := []sort.IntSlice{
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}

	for i := 0; i < len(expected); i++ {
		ok := NextPermutation(data)
		assert.True(t, ok)
		assert.Equal(t, expected[i], data)
	}

	ok := NextPermutation(data)
	assert.False(t, ok)
}

func TestLastPermutation(t *testing.T) {
	data := sort.IntSlice{3, 2, 1}
	expected := []sort.IntSlice{
		{3, 1, 2},
		{2, 3, 1},
		{2, 1, 3},
		{1, 3, 2},
		{1, 2, 3},
	}

	for i := 0; i < len(expected); i++ {
		ok := LastPermutation(data)
		assert.True(t, ok)
		assert.Equal(t, expected[i], data)
	}

	ok := LastPermutation(data)
	assert.False(t, ok)
}
