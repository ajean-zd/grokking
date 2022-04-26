package grokking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectSort(t *testing.T) {

	unsorted := SortNumbers{7, 4, 65, 52, 1, 888}
	result := selectSort(unsorted)
	expected := []int{1, 4, 7, 52, 65, 888}

	assert.Equal(t, expected, result, "ints have been sorted")

}

// func TestReverseOrder(t *testing.T) {

// 	sorted := SortNumbers{1, 4, 7, 52, 65, 888}
// 	expected := []int{888, 65, 52, 7, 4, 1}

// 	assert.Equal(t, sorted, expected, "sorted ints have been reversed")

// }

// func TestStringSort(t *testing.T) {

// 	unsorted := SortStrings{"g", "b", "z", "r"}
// 	expected := []string{"b", "g", "r", "z"}
// 	assert.Equal(t, unsorted, expected, "strings sorted")
// }
