package graph_test

import (
	"go-algo/graph"
	"testing"
)

type Graph interface {
	AddEdge(from, to, weight int) error
	EdgeExists(from, to int) (bool, error)
	EdgeWeight(from, to int) (int, error)
	GetNeighboringVertices(source int) ([]int, error)
}

// from, to, weight
type EdgeInput [3]int

type GraphInput struct {
	VerticesCnt int
	Edges       []EdgeInput
}

var tests = []GraphInput{
	{
		5, []EdgeInput{
			{0, 1, 5},
			{0, 2, 7},
			{1, 3, 9},
			{3, 4, 13},
		},
	},
}

func testGraph(t *testing.T, constructor func(int) graph.Graph) {
	for _, test := range tests {
		g := constructor(test.VerticesCnt)

		for _, edge := range test.Edges {
			if err := g.AddEdge(edge[0], edge[1], edge[2]); err != nil {
				t.Errorf("error while adding edge: %v", err)
			}
		}

		for _, edge := range test.Edges {
			if w, err := g.EdgeWeight(edge[0], edge[1]); err != nil {
				t.Errorf("error while getting edge weight: %v", err)
			} else if w != edge[2] {
				t.Errorf("expected weight %d, got %d", edge[2], w)
			}
		}
	}
}
