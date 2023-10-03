package main

type queue struct {
	a []*edge
}

func (q *queue) front() *edge {
	return q.a[0]
}

func (q *queue) push(n *edge) {
	q.a = append(q.a, n)
}

func (q *queue) pop() *edge {
	var e *edge
	e, q.a = q.a[0], q.a[1:]
	return e
}
func (q *queue) empty() bool {
	return len(q.a) == 0
}

type edge struct {
	node   int
	weight int
}
type graph struct {
	adj map[int][]*edge
}

func (g *graph) push(v, w, weight int) {

	if g.adj[v] == nil {
		g.adj[v] = make([]*edge, 0)
	}

	g.adj[v] = append(g.adj[v], &edge{node: w, weight: weight})

	if g.adj[w] == nil {
		g.adj[w] = make([]*edge, 0)
	}

	g.adj[w] = append(g.adj[w], &edge{node: v, weight: weight})
}

//func dijkstra(g *graph, src, dst int) int {
//
//	visited := map[int]bool{}
//	dist := map[int]int{}
//	dist[src] = 0
//	q := queue{a: make([]*edge, 0)}
//
//	q.push(&edge{node: src, weight: 0})
//
//	for !q.empty() {
//
//		it := q.pop()
//		node := it.node
//		weight := it.weight
//
//
//		for _,nbrEdge:=range g.adj[node]{
//
//			nbrNode:=nbrEdge.node
//			nbrWeight:=nbrEdge.weight
//
//			if
//
//		}
//
//
//	}
//}

func main() {

	g := graph{
		adj: map[int][]*edge{},
	}
	g.push(0, 1, 1)
	g.push(1, 2, 2)
	g.push(0, 2, 4)
	g.push(0, 3, 7)
	g.push(3, 2, 2)
	g.push(3, 4, 3)

}
