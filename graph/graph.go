package graph

import (
	sll "github.com/emirpasic/gods/lists/singlylinkedlist"
)

// Adjacency list representation of a graph

type Graph struct {
	v             int
	AdjacencyList []*sll.List
}

func NewGraph(v int) *Graph {

	adjacencyList := make([]*sll.List, v)
	for i := 0; i < v; i++ {
		adjacencyList[i] = sll.New()
	}

	return &Graph{
		v:             v,
		AdjacencyList: adjacencyList,
	}
}
