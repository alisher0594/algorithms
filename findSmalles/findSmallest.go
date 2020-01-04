package findSmallest

func FindSmallest(list []int) int {
	smallest := list[0]
	smallestIndex := 0

	for i := range list {
		if list[i] < smallest {
			smallest = list[i]
			smallestIndex = i
		}
	}
	return smallestIndex
}