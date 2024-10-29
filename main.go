package main

import (
	"fmt"
	"go-algo/graph"
)

func main() {
	input := [][]int{
		{0, 1, 5},
		{0, 2, 7},
		{1, 3, 9},
		{3, 4, 13},
	}

	var m graph.Graph = graph.NewAdjacencyMatrix(5)

	for i, v := range input {
		if err := m.AddEdge(v[0], v[1], v[2]); err != nil {
			panic(fmt.Sprintf("error while adding edge: %v, iteration %d", err, i))
		}
	}

}
