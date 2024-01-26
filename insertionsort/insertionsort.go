package insertionsort

// InsertionSort sorts the given list using the insertion sort algorithm.
func InsertionSort(list []int) {
	for pos := 0; pos < len(list); pos++ {
		MoveLeft(list, pos)
	}
}
