package grokking

// chapter 1
// this function takes a slice of sorted integers and an int and returns true if it's in array

type BT []int

func binary_tree(item int, bt []int) int {
	//returns the location in the slice
	result := -1 // not found
	start := 0
	end := len(bt) - 1
	for start <= end {
		mid := (start + end) / 2
		if bt[mid] == item {
			result = mid // found
			break
		} else if bt[mid] < item {
			start = mid + 1
		} else if bt[mid] > item {
			end = mid - 1
		}
	}
	return result
}
