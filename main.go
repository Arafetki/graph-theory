package main

import (
	"github.com/Arafetki/graph-theory/graph"
)

func main() {

	// (Unweighted directed Graph) v0->v1 v0->v2 v1->v3 v3->v1 v2->v3
	v := 5
	g := graph.NewGraph(v)
	g.AdjacencyList[0].Add(1, 2)
	g.AdjacencyList[1].Add(0, 3)
	g.AdjacencyList[2].Add(0, 4)
	g.AdjacencyList[3].Add(1, 4)
	g.AdjacencyList[4].Add(2, 3)
	g.PrintShortestPath(3, 0)
}
