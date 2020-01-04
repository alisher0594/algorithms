package binarySearch

func binarySearch(list []int, item int) interface{} {
	low := 0
	high := len(list)-1

	for low <= high {

		mid := (low + high) / 2
		guess := list[mid]

		if guess == item {
			return mid
		} else if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return nil
}

