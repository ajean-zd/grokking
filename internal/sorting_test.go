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

func TestSelectSortBlank(t *testing.T) {

	unsorted := SortNumbers{}
	result := selectSort(unsorted)
	expected := []int{}
	assert.Equal(t, expected, result, "nothing returns")

}

func TestGetSmallest(t *testing.T) {

	unsorted := SortNumbers{7, 4, 65, 52, 1, 888}
	result := getSmallest(unsorted)
	expected := 1
	assert.Equal(t, expected, result, "smallest number achieved")

}

func TestReverseSortOrder(t *testing.T) {

	input := SortNumbers{1, 7, 4, 52, 65, 888}
	result := reverseSortOrder(input)
	expected := []int{888, 65, 52, 7, 4, 1}

	assert.Equal(t, result, expected, "sorted ints have been reversed")

}

// func TestStringSort(t *testing.T) {

// 	unsorted := SortStrings{"g", "b", "z", "r"}
// 	expected := []string{"b", "g", "r", "z"}
// 	assert.Equal(t, unsorted, expected, "strings sorted")
// }
