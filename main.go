package main

import (
	"fmt"

	"github.com/Arafetki/graph-theory/graph"
)

func main() {

	// (Unweighted directed Graph) v0->v1 v0->v2 v1->v3 v3->v1 v2->v3
	v := 4
	g := graph.NewGraph(v)
	g.AdjacencyList[0].Add(1, 2)
	g.AdjacencyList[1].Add(3)
	g.AdjacencyList[2].Add(3)
	g.AdjacencyList[3].Add(1)
	for i := 0; i < v; i++ {
		fmt.Println(g.AdjacencyList[i].Values())
	}
}
