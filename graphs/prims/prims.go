package prims

import (
	"interview/graphs/graph"
	"interview/queue"
)

func primMst(g *graph.Graph, start int) (int, []*queue.Edge) {

	visited := map[int]bool{}
	var ans int
	var mst []*queue.Edge

	q := queue.NewPriorityQueue()

	q.PushNode(start, 0)

	for !q.Empty() {
		best := q.Pop()
		to := best.V
		weight := best.Weight

		if visited[to] {
			continue
		}
		mst = append(mst, best)
		//fmt.Printf("Node %d, Weight %d\n", best.V, best.Weight)
		ans += weight
		visited[to] = true

		for _, x := range g.Edges(to) {
			if !visited[x.V] {
				q.PushNode(x.V, x.Weight)
			}
		}

	}
	return ans, mst
}
