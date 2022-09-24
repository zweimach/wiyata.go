package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	result := addTwoNumbers(newListNode(2, 4, 3), newListNode(5, 6, 4))
	expected := newListNode(7, 0, 8)

	assert.True(t, compareListNode(result, expected))

	result = addTwoNumbers(newListNode(0), newListNode(0))
	expected = newListNode(0)

	assert.True(t, compareListNode(result, expected))

	result = addTwoNumbers(newListNode(9, 9, 9, 9, 9, 9, 9), newListNode(9, 9, 9, 9))
	expected = newListNode(8, 9, 9, 9, 0, 0, 0, 1)

	assert.True(t, compareListNode(result, expected))
}
