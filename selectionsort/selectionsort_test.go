package selectionsort

import (
	"testing"

	"github.com/tel23a-inf/aufgaben-sortieren/testhelpers"
)

func TestSelectionSort(t *testing.T) {
	list := []int{3, 2, 1}
	SelectionSort(list)
	testhelpers.AssertListsEqual(t, []int{1, 2, 3}, list)
}
