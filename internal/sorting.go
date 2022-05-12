package grokking

import (
	"sort"
)

//chapter two

//sort from smallest to largest

type SortNumbers []int
type SortStrings []string

func selectSort(unsortedNumbers []int) []int {
	// takes an unsorted array of ints and sorts them smallest to largest
	if len(unsortedNumbers) == 0 {
		return []int{}
	}

	index := len(unsortedNumbers) / 2
	working_sort := unsortedNumbers[index]
	var under []int
	var over []int

	for i := range unsortedNumbers {
		if i == index {
			continue
		}
		if unsortedNumbers[i] <= working_sort {
			under = append(under, unsortedNumbers[i])
		} else {
			over = append(over, unsortedNumbers[i])
		}
	}

	result := append(selectSort(under), working_sort)
	return append(result, selectSort(over)...)
}

func getSmallest(unsortedNumbers []int) int {
	//returns the smallest number in an array
	smallest := unsortedNumbers[0]
	for _, number := range unsortedNumbers {
		if number < smallest {
			smallest = number
		}
	}
	return smallest
}

func reverseSortOrder(sorted []int) []int {
	// takes a sorted array and reverses the sort order
	if len(sorted) == 0 {
		return []int{}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sorted)))
	return sorted
}

func stringSort(unsortedNumbers []string) []string {
	sorted := SortStrings{}

	return sorted
}
