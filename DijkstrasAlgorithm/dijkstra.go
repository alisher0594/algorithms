package main

import (
	"fmt"
)

func main() {
	processed := make(map[string]string)
	graph := make(map[string]map[string]int)
// все связи и их веса [a]-->[b]
	graph["start"] = map[string]int{}
	graph["start"]["a"] = 1
	graph["start"]["b"] = 1

	graph["a"] = map[string]int{}
	graph["a"]["e"] = 3
	graph["a"]["c"] = 2
	graph["a"]["d"] = 2

	graph["b"] = map[string]int{}
	graph["b"]["h"] = 2
	graph["b"]["f"] = 3

	graph["c"] = map[string]int{}
	graph["c"]["d"] = 1

	graph["d"] = map[string]int{}
	graph["d"]["g"] = 4

	graph["e"] = map[string]int{}
	graph["e"]["g"] = 5
	graph["e"]["g"] = 1

	graph["f"] = map[string]int{}
	graph["f"]["n"] = 2

	graph["g"] = map[string]int{}
	graph["g"]["m"] = 2

	graph["h"] = map[string]int{}
	graph["h"]["p"] = 1

	graph["n"] = map[string]int{}
	graph["n"]["p"] = 5
	graph["n"]["r"] = 1

	graph["m"] = map[string]int{}
	graph["m"]["p"] = 1
	graph["m"]["s"] = 1

	graph["p"] = map[string]int{}
	graph["p"]["s"] = 2
	graph["p"]["t"] = 2
	graph["p"]["o"] = 3

	graph["h"] = map[string]int{}
	graph["h"]["p"] = 1

	graph["r"] = map[string]int{}
	graph["r"]["o"] = 5

	graph["s"] = map[string]int{}
	graph["s"]["x"] = 2
	graph["s"]["y"] = 3

	graph["t"] = map[string]int{}
	graph["t"]["2t"] = 6

	graph["o"] = map[string]int{}
	graph["o"]["v"] = 1
	graph["o"]["w"] = 1

	graph["2t"] = map[string]int{}
	graph["2t"]["z"] = 3

	graph["v"] = map[string]int{}
	graph["v"]["z"] = 2
	graph["v"]["2a"] = 3

	graph["w"] = map[string]int{}
	graph["w"]["2a"] = 5

	graph["x"] = map[string]int{}
	graph["x"]["2f"] = 1

	graph["y"] = map[string]int{}
	graph["y"]["2f"] = 2

	graph["z"] = map[string]int{}
	graph["z"]["2b"] = 2
	graph["z"]["2c"] = 2

	graph["2a"] = map[string]int{}
	graph["2a"]["2c"] = 3
	graph["2a"]["2d"] = 1

	graph["2c"] = map[string]int{}
	graph["2c"]["2d"] = 3

	graph["2b"] = map[string]int{}
	graph["2b"]["2f"] = 1
	graph["2b"]["2e"] = 3

	graph["2d"] = map[string]int{}
	graph["2d"]["fin"] = 1

	graph["2f"] = map[string]int{}
	graph["2f"]["fin"] = 2

	graph["fin"] = map[string]int{}

	// веса узлов
	infinity := 0x7FF0000000000000
	costs := make(map[string]int)
	costs["a"] = 3
	costs["b"] = 2
	costs["c"] = infinity
	costs["2f"] = infinity
	costs["2d"] = infinity
	costs["2b"] = infinity
	costs["2c"] = infinity
	costs["z"] = infinity
	costs["y"] = infinity
	costs["x"] = infinity
	costs["w"] = infinity
	costs["v"] = infinity
	costs["2t"] = infinity
	costs["o"] = infinity
	costs["t"] = infinity
	costs["s"] = infinity
	costs["r"] = infinity
	costs["h"] = infinity
	costs["p"] = infinity
	costs["m"] = infinity
	costs["h"] = infinity
	costs["n"] = infinity
	costs["g"] = infinity
	costs["f"] = infinity
	costs["e"] = infinity
	costs["c"] = infinity
	costs["d"] = infinity

	costs["fin"] = infinity

	//
	parents := make(map[string]string)
	parents["a"] = "start"
	parents["b"] = "start"
	parents["fin"] = ""


	node, isNotEmpty  := findLowestCostNode(costs, processed)

	for isNotEmpty {
		cost := costs[node]
		neighbors := graph[node]
		for key, n := range neighbors {
			newCost := cost + n
			if costs[key] > newCost {
				costs[key] = newCost
				parents[key] = node
			}
		}
		processed[node] = "processed"
		node, isNotEmpty = findLowestCostNode(costs, processed)
	}
	key := "fin"
	fmt.Print("fin")
	for key != "start" {
		if val, ok := parents[key]; ok {
			fmt.Print("<--", val)
			key = val
		}
	}
}

func findLowestCostNode(costs map[string]int, processed map[string]string) (string, bool) {
	var lowestCostNode string
	var isNotEmpty bool
	lowestCost := 0x7FF0000000000000 // infinity

	for node, cost := range costs {
		if _, ok := processed[node]; !ok { // node not in processed
			if cost < lowestCost {
				lowestCost = cost
				lowestCostNode = node
				isNotEmpty = true
			}
		}
	}
	return lowestCostNode, isNotEmpty
}
