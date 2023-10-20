package main

import (
	"fmt"
	"interview/queue"
	"math"
)

type graph struct {
	adj map[int][]*queue.Edge
}

func (g *graph) push(v, w, weight int) {

	if g.adj[v] == nil {
		g.adj[v] = make([]*queue.Edge, 0)
	}

	g.adj[v] = append(g.adj[v], &queue.Edge{V: w, Weight: weight})

	if g.adj[w] == nil {
		g.adj[w] = make([]*queue.Edge, 0)
	}

	g.adj[w] = append(g.adj[w], &queue.Edge{V: v, Weight: weight})
}

func dijkstra(g *graph, src, dest int) int {

	dist := map[int]int{}

	for key, _ := range g.adj {

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

		for _, neighbor := range g.adj[node] {

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

	g := graph{
		adj: map[int][]*queue.Edge{},
	}
	g.push(0, 1, 1)
	g.push(1, 2, 1)
	g.push(0, 2, 4)
	g.push(0, 3, 7)
	g.push(3, 2, 2)
	g.push(3, 4, 3)

	dijkstra(&g, 0, 4)
	//test
}
