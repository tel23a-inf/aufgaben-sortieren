package bubblesort

import (
	"testing"

	"github.com/tel23a-inf/aufgaben-sortieren/testhelpers"
)

func TestBubbleSort(t *testing.T) {
	list := []int{25, 2, 4, 3, 1, 13, 12, 5}
	BubbleSort(list)
	testhelpers.AssertListsEqual(t, []int{1, 2, 3, 4, 5, 12, 13, 25}, list)
}
