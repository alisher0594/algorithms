package main

import (
	//"fmt"
	//qS "algorithms/quickSort"
	bFs "algorithms/breadthFirstSearch"
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
	//sortedList := selectionSort([]int{5, 3, 6, 2, 10, 1, 9, 4, 8, 12})
	//fmt.Println("guess: ", binarySearch(sortedList, 10))
	//fmt.Println(factorial(50))
	//fmt.Println(sum([]int{5, 3, 6, 2, 10, 1, 9, 4, 8, 12}))
	//fmt.Println(selectionSort([]int{5, 3, 6, 2, 10, 1, 9, 4, 8, 12}))
	//fmt.Println(qS.QuickSort([]int{5, 3, 6, 2, 10, 1, 9, 4, 8, 12}))

	//electorate := []string{"Alice", "Bob", "Mike", "Alice", "Fed", "Jon"}
	//for _, val := range electorate {
	//	checkVoter(val)
	//}

	bFs.BreadthFirstSearch(people)
}
