package graph

type Graph interface {
	AddEdge(from, to, weight int) error
	EdgeExists(from, to int) (bool, error)
	EdgeWeight(from, to int) (int, error)
	GetNeighboringVertices(source int) ([]int, error)
}
