package main

import (
	"fmt"
	"math/rand/v2"
)

type Vertex map[uint]struct{}

func NewVertex(labels ...uint) Vertex {
	result := make(Vertex, len(labels))
	for _, label := range labels {
		result[label] = struct{}{}
	}
	return result
}

func (v Vertex) Merge(w Vertex) {
	for key, _ := range w {
		v[key] = struct{}{}
	}
}

func (v Vertex) Has(label uint) bool {
	_, ok := v[label]
	return ok
}

type Edge struct {
	From uint
	To   uint
}

type Graph struct {
	Edges    []Edge
	Vertices map[uint]Vertex
}

// KargerMinCut returns the minimum cut of the graph using the Karger's algorithm.
// Note: It modifies the graph in place.
func (g *Graph) KargerMinCut() []Edge {
	// Iterate until only 2 vertices are left
	for range len(g.Vertices) - 2 {
		edgeIndex := rand.N[int](len(g.Edges))
		edge := g.Edges[edgeIndex]

		v1 := g.Vertices[edge.From]
		v2 := g.Vertices[edge.To]

		// Merge v2 into v1
		for key := range v2 {
			v1[key] = struct{}{}
			// Updatem mapping from vertices to merged vertices
			g.Vertices[key] = v1
		}
		v2 = nil

		// Remove edges between v1 and v2
		newEdges := make([]Edge, 0, len(g.Edges))
		for _, e := range g.Edges {
			if v1.Has(e.From) && v1.Has(e.To) {
			} else {
				newEdges = append(newEdges, e)
			}
		}
		g.Edges = newEdges
	}

	return g.Edges
}

func MapToGraph(m map[uint][]uint) *Graph {
	g := &Graph{
		Vertices: make(map[uint]Vertex, len(m)),
	}

	for from, tos := range m {
		v := NewVertex(from)
		g.Vertices[from] = v
		for _, to := range tos {
			g.Edges = append(g.Edges, Edge{From: from, To: to})
		}
	}

	return g
}

func main() {
	// Example 1 ~ Joanna Kulig

	graph := map[uint][]uint{
		0: {1, 2, 3, 4},
		1: {0, 2, 3, 4},
		2: {0, 1, 3, 4, 7},
		3: {0, 1, 2, 4, 6},
		4: {0, 1, 2, 3, 5},
		5: {4, 6, 7, 8, 9},
		6: {3, 5, 7, 8, 9},
		7: {2, 5, 6, 8, 9},
		8: {5, 6, 7, 9},
		9: {5, 6, 7, 8},
	}

	g := MapToGraph(graph)

	sol := g.KargerMinCut()

	fmt.Printf("Minimum cut of size %d found: %v\n", len(sol), sol)
}
