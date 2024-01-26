package insertionsort

// MoveLeft moves the element at the given position to the left until it is in the correct position.
func MoveLeft(list []int, pos int) {
	for i := pos; i > 0 && list[i] < list[i-1]; i-- {
		list[i], list[i-1] = list[i-1], list[i]
	}
}
