package breadthFirstSearch

import (
	"fmt"
	"github.com/gammazero/deque"
)

func BreadthFirstSearch(hashTable map[string][]string) bool {
	var searchQueue deque.Deque
	searched := make(map[string]string)

	for _, name := range hashTable["you"] {
		searchQueue.PushBack(name)
	}

	// Consume deque and print elements.
	for searchQueue.Len() != 0 {
		person := fmt.Sprintf("%v", searchQueue.PopFront())

		if _, ok := searched[person]; !ok {
			searched[person] = "added"
			if personIsSeller(person) {
				fmt.Println(person, " is а mango seller!")
				return true
			} else {
				for _, name := range hashTable[person] {
					searchQueue.PushBack(name)
				}
			}
		}
	}
	fmt.Println("no mango seller in Queue")
	return false
}

func personIsSeller (name string) bool {
	if name == "muller" {
		return true
	}
	return false
}