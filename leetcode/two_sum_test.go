package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	expected := []int{0, 1}
	result := twoSum([]int{2, 7, 11, 15}, 9)

	assert.Equal(t, len(expected), len(result))
	assert.Equal(t, expected, result)

	expected = []int{1, 2}
	result = twoSum([]int{3, 2, 4}, 6)

	assert.Equal(t, len(expected), len(result))
	assert.Equal(t, expected, result)

	expected = []int{0, 1}
	result = twoSum([]int{3, 3}, 6)

	assert.Equal(t, len(expected), len(result))
	assert.Equal(t, expected, result)
}
