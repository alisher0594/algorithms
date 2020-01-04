package findSmallest

import (
	"testing"
)

type testListForSmallest struct {
	list []int
	result int
}

var fieldsForSmallest = []testListForSmallest {
	{[]int{4, 2, 5, 7, 9}, 1},
	{[]int{2, 1, 5, 8, 0}, 4},
	{[]int{1, 2, 5, 3, 8}, 0},
}

func TestFindSmallest(t *testing.T) {
	for _, field := range fieldsForSmallest {
		item := FindSmallest(field.list)
		if item != field.result {
			t.Errorf("TestFindSmallest for OK Failed - error")
		}
	}
}
