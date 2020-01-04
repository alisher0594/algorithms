package quickSort

import (
	"testing"
)

type testListQuickSort struct {
	list []int
	result []int
}

var fieldsForQuickSort = []testListQuickSort {
	{[]int{5, 3, 6, 2, 10, 1, 9, 4, 8, 12}, []int{1, 2, 3, 4, 5, 6, 8, 9, 10, 12}},
	{[]int{2, 1, 5, 8, 0, 3}, []int{0, 1, 2, 3, 5, 8}},
	{[]int{10, 5, 2, 3}, []int{2, 3, 5, 10}},
}

func TestSQuickSort(t *testing.T) {
	for _, field := range fieldsForQuickSort {
		itemList := QuickSort(field.list)
		for i, val := range field.result {
			if val != itemList[i] {
				t.Errorf("TestSQuickSort for OK Failed - error")
			}
		}
	}
}
