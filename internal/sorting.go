package grokking

//chapter two

//sort from smallest to largest

type SortNumbers []int
type SortStrings []string

func selectSort(unsortedNumbers []int) []int {
	// takes an unsorted array of ints and sorts them smallest to larges
	sorted := SortNumbers{}
	unsortedNumbersLength := len(unsortedNumbers)

	for i := 0; i < unsortedNumbersLength; i++ {
		smallestNumber := getSmallest(unsortedNumbers)
		sorted = append(sorted, smallestNumber)
	}

	return sorted

}

func getSmallest(unsortedNumbers []int) int {
	//returns the smallest number in an array
	smallest := unsortedNumbers[0]
	smallestIndex := 0
	unsortedNumbersLength := len(unsortedNumbers)
	for i := 0; i < unsortedNumbersLength; i++ {
		if unsortedNumbers[i] < smallest {
			smallest = unsortedNumbers[i]
			smallestIndex = i
		}
	}
	return smallestIndex
}

func reverseOrder(unsortedNumbers []int) []int {
	sorted := SortNumbers{}
	//takes a sorted array and reverses the sorting order
	return sorted

}

func stringSort(unsortedNumbers []string) []string {
	sorted := SortStrings{}

	return sorted
}
