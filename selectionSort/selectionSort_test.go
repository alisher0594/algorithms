package selectionSort

import "testing"

type testListSelectionSort struct {
	list []int
	result []int
}

var fieldsForSelectionSort = []testListSelectionSort {
	{[]int{1, 4, 2, 5, 7, 9}, []int{1, 2, 4, 5, 7, 9}},
	{[]int{2, 1, 5, 8, 0, 3}, []int{0, 1, 2, 3, 5, 8}},
	{[]int{10, 5, 2, 3}, []int{2, 3, 5, 10}},
}

func TestSelectionSort(t *testing.T) {
	for _, field := range fieldsForSelectionSort {
		itemList := selectionSort(field.list)
		for i, val := range field.result {
			if val != itemList[i] {
				t.Errorf("TestSelectionSort for OK Failed - error")
			}
		}
	}
}

