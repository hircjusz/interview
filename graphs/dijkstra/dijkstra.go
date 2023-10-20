package main

import (
	"fmt"
	"interview/queue"
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
	dist[src] = 0
	q := queue.New()

	q.Push(&queue.Edge{V: src, Weight: 0})

	for !q.Empty() {

		it := q.Front()

		node := it.V
		weight := it.Weight
		q.Erase(it)

		for _, nbrEdge := range g.adj[node] {

			nbrNode := nbrEdge.V
			nbrWeight := nbrEdge.Weight

			if weight+nbrWeight < dist[nbrNode] {
				f := q.Find(&queue.Edge{V: nbrNode, Weight: dist[nbrNode]})
				if *f != *q.End() {
					q.Erase(f)
				}
				dist[nbrNode] = weight + nbrWeight
				q.Push(&queue.Edge{V: nbrNode, Weight: dist[nbrNode]})
			}

		}

	}
	for i, v := range dist {
		fmt.Printf("Node i %d, Dist %d", i, v)
	}

	return dist[dest]
}

func main() {

	g := graph{
		adj: map[int][]*queue.Edge{},
	}
	g.push(0, 1, 1)
	g.push(1, 2, 2)
	g.push(0, 2, 4)
	g.push(0, 3, 7)
	g.push(3, 2, 2)
	g.push(3, 4, 3)

	dijkstra(&g, 0, 4)
}
