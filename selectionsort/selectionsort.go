package selectionsort

// SelectionSort sorts the given list using the selection sort algorithm.
func SelectionSort(list []int) {
	// Bemerkung: Alles ab Position i ist noch nicht sortiert.
	for i := 0; i < len(list); i++ {
		s := SmallestPos(list[i:]) + i
		list[i], list[s] = list[s], list[i]
	}
}
