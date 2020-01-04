package binarySearch

import (
	"testing"
)

type testList struct {
	list []int
	item int
	result interface{}
}

var fields = []testList {
	{[]int{1, 3, 5, 7, 9}, 9, 4 },
	{[]int{2, 3, 5, 7, 9}, 0, nil},
	{[]int{0, 2, 3, 5, 8}, 3, 2},
}

func TestBinarySearch(t *testing.T)  {
	for _, field := range fields {
		item := binarySearch(field.list, field.item)
		if item != field.result {
			t.Errorf("TestBinarySearch for OK Failed - error")
		}
	}
}
