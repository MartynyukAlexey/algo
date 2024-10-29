package graph

import "fmt"

type AdjacencyMatrix struct {
	verticesCnt int
	matrix      [][]int
}

func NewAdjacencyMatrix(verticesCnt int) *AdjacencyMatrix {
	matrix := make([][]int, verticesCnt)

	for i := 0; i < verticesCnt; i++ {
		matrix[i] = make([]int, verticesCnt)
	}

	return &AdjacencyMatrix{
		verticesCnt: verticesCnt,
		matrix:      matrix,
	}
}

func (m *AdjacencyMatrix) isValidVertice(v int) bool {
	return v >= 0 && v < m.verticesCnt
}

func (m *AdjacencyMatrix) AddEdge(from, to, weight int) error {
	if !m.isValidVertice(from) || !m.isValidVertice(to) {
		return fmt.Errorf("edge (%d, %d) out of bounds of matrix (size %d)", from, to, m.verticesCnt)
	}

	if weight == 0 {
		return fmt.Errorf("edge weight cannot be 0")
	}

	m.matrix[from][to] = weight

	return nil
}

func (m *AdjacencyMatrix) EdgeExists(from, to int) (bool, error) {
	if !m.isValidVertice(from) || !m.isValidVertice(to) {
		return false, fmt.Errorf("edge (%d, %d) out of bounds of matrix (size %d)", from, to, m.verticesCnt)
	}

	return m.matrix[from][to] != 0, nil
}

func (m *AdjacencyMatrix) EdgeWeight(from, to int) (int, error) {
	if !m.isValidVertice(from) || !m.isValidVertice(to) {
		return 0, fmt.Errorf("edge (%d, %d) out of bounds of matrix (size %d)", from, to, m.verticesCnt)
	}

	if m.matrix[from][to] == 0 {
		return 0, fmt.Errorf("edge (%d, %d) does not exist", from, to)
	}

	return m.matrix[from][to], nil
}

func (m *AdjacencyMatrix) GetNeighboringVertices(source int) ([]int, error) {
	if !m.isValidVertice(source) {
		return nil, fmt.Errorf("vertice %d out of bounds of matrix (size %d)", source, m.verticesCnt)
	}

	var vertices []int
	for i := 0; i < m.verticesCnt; i++ {
		if m.matrix[source][i] != 0 {
			vertices = append(vertices, i)
		}
	}

	return vertices, nil
}
