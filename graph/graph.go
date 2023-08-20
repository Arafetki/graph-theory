package graph

import (
	"fmt"
	"math"

	sll "github.com/emirpasic/gods/lists/singlylinkedlist"
	queue "github.com/emirpasic/gods/queues/arrayqueue"
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

func (g *Graph) BFS(root int) {

	visited := make([]bool, g.v)
	q := queue.New()
	visited[root] = true
	q.Enqueue(root)

	for !q.Empty() {
		curr, _ := q.Dequeue()
		fmt.Printf("%d ", curr)
		it := g.AdjacencyList[curr.(int)].Iterator()
		for it.Next() {
			if !visited[it.Value().(int)] {
				visited[it.Value().(int)] = true
				q.Enqueue(it.Value().(int))
			}
		}
	}
}

// Shortest path for unweighted graph
func (g *Graph) PrintShortestPath(srcNode, destNode int) {

	q := queue.New()
	visited := make([]bool, g.v)
	dist, pred := make([]int, g.v), make([]int, g.v)
	for i := 0; i < g.v; i++ {
		dist[i] = math.MaxInt
		pred[i] = -1
	}

	q.Enqueue(srcNode)
	visited[srcNode] = true
	dist[srcNode] = 0
loop:
	for !q.Empty() {
		currNode, _ := q.Dequeue()
		it := g.AdjacencyList[currNode.(int)].Iterator()
		for it.Next() {
			if !visited[it.Value().(int)] {
				visited[it.Value().(int)] = true
				dist[it.Value().(int)] = dist[currNode.(int)] + 1
				pred[it.Value().(int)] = currNode.(int)
				q.Enqueue(it.Value().(int))
				if it.Value().(int) == destNode {
					break loop
				}
			}
		}
	}

	path := []int{}
	crawl := destNode
	path = append(path, crawl)
	for pred[crawl] != -1 {
		path = append(path, pred[crawl])
		crawl = pred[crawl]
	}
	fmt.Println(pred, path)

}
