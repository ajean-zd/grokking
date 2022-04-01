package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBinaryTreeNumberFound(t *testing.T) {

	btslice := BT{1, 2, 3, 4, 6, 33, 55, 66, 77, 88, 99}

	item := 99

	result, actualErr := binary_tree(item, btslice)
	expected := true

	require.NoError(t, actualErr)
	assert.Equal(t, expected, result, "binary search true")
}
