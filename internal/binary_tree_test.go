package grokking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryTreeNumberFound(t *testing.T) {

	btslice := BT{1, 2, 3, 4, 6, 33, 55, 66, 77, 88, 99}

	item := 99

	returned_number := binary_tree(item, btslice)

	assert.Equal(t, returned_number, 10, "found location inslice")
}

func TestBinaryTreeNumberFirstFound(t *testing.T) {

	btslice := BT{1, 2, 3, 4, 6, 33, 55, 66, 77, 88, 99}

	item := 1

	returned_number := binary_tree(item, btslice)

	assert.Equal(t, returned_number, 0, "found location inslice")
}

func TestBinaryTreeNumberNotFound(t *testing.T) {

	btslice := BT{1, 2, 3, 4, 6, 33, 55, 66, 77, 88, 99}

	item := 42

	returned_number := binary_tree(item, btslice)

	assert.Equal(t, returned_number, -1, "item is not in slice")
}
