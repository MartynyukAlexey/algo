package graph_test

import (
	"testing"

	"go-algo/graph"
)

func TestAdjacencyMatrix(t *testing.T) {
	testGraph(t, func(verticesCnt int) graph.Graph {
		return graph.NewAdjacencyMatrix(verticesCnt)
	})
}
