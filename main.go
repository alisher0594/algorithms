package main

import (
	"fmt"
	selS "algorithms/selectionSort"
	binS "algorithms/binarySearch"
	qS "algorithms/quickSort"
	bFs "algorithms/breadthFirstSearch"
	rec "algorithms/recursion"
)

var people = map[string][]string{
	"you": {"alice", "bob", "claire"},
	"bob": {"abuja", "peggy"},
	"alice": {"peggy"},
	"claire": {"thom", "johnny"},
	"abuja": {"fed", "def"},
	"peggy": {"red", "der"},
	"thom": {"mike"},
	"johny": {"putin", "muller"},
	"fed": {"muller", "ivan"},
	"def": {},
	"red": {"putin", "ivan", "max"},
	"der": {},
	"mike": {},
	"putin": {},
	"muller": {},
}


func main() {
	sortedList := selS.SelectionSort([]int{5, 3, 6, 2, 10, 1, 9, 4, 8, 12})
	fmt.Println("guess: ", binS.BinarySearch(sortedList, 10))
	fmt.Println(selS.SelectionSort([]int{5, 3, 6, 2, 10, 1, 9, 4, 8, 12}))
	fmt.Println(qS.QuickSort([]int{5, 3, 6, 2, 10, 1, 9, 4, 8, 12}))

	//electorate := []string{"Alice", "Bob", "Mike", "Alice", "Fed", "Jon"}
	//for _, val := range electorate {
	//	checkVoter(val)
	//}

	bFs.BreadthFirstSearch(people)

	fmt.Println(rec.Factorial(50))
	fmt.Println(rec.Sum([]int{5, 3, 6, 2, 10, 1, 9, 4, 8, 12}))
}
