package graph

import "interview/queue"

type Graph struct {
	adj map[int][]*queue.Edge
}

func NewGraph() *Graph {
	return &Graph{make(map[int][]*queue.Edge)}
}

func (g *Graph) NodesCount() int {
	return len(g.adj)
}
func (g *Graph) AllEdges() map[int][]*queue.Edge {
	return g.adj
}

func (g *Graph) Edges(v int) []*queue.Edge {
	return g.adj[v]
}

func (g *Graph) Push(v, w, weight int) {

	if g.adj[v] == nil {
		g.adj[v] = make([]*queue.Edge, 0)
	}

	g.adj[v] = append(g.adj[v], &queue.Edge{V: w, Weight: weight})

	if g.adj[w] == nil {
		g.adj[w] = make([]*queue.Edge, 0)
	}

	g.adj[w] = append(g.adj[w], &queue.Edge{V: v, Weight: weight})
}
