package main

import (
	"fmt"
	"interview/graphs/graph"
	"interview/queue"
	"math"
)

func dijkstra(g *graph.Graph, src, dest int) int {

	dist := map[int]int{}

	for key, _ := range g.AllEdges() {

		dist[key] = math.MaxInt
	}

	dist[src] = 0
	q := queue.New()

	q.PushNode(src, 0)

	for !q.Empty() {

		it := q.Front()

		node := it.V
		weight := it.Weight
		q.EraseByEdge(it)

		for _, neighbor := range g.Edges(node) {

			neighborNode := neighbor.V
			neighborWeight := neighbor.Weight

			newDist := weight + neighborWeight

			if newDist < dist[neighborNode] {
				q.EraseByNode(neighborNode)
				dist[neighborNode] = newDist
				q.PushNode(neighborNode, dist[neighborNode])
			}
		}

	}
	for i, v := range dist {
		fmt.Printf("Node  %d, Dist %d", i, v)
		fmt.Println()
	}

	return dist[dest]
}

func main() {

	g := graph.NewGraph()
	g.Push(0, 1, 1)
	g.Push(1, 2, 1)
	g.Push(0, 2, 4)
	g.Push(0, 3, 7)
	g.Push(3, 2, 2)
	g.Push(3, 4, 3)

	dijkstra(g, 0, 4)
	//test
}
