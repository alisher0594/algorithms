package selectionSort

import u "algorithms/findSmalles"

func SelectionSort(arr []int)  []int {
	var newArr []int

	for _ = range arr {
		smallestIndex := u.FindSmallest(arr)
		newArr = append(newArr, arr[smallestIndex])
		arr = append(arr[:smallestIndex], arr[smallestIndex+1:]...)
	}
	return newArr
}
