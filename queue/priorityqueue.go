package queue

import "sort"

type PriorityQueue struct {
	q *Queue
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		q: New(),
	}
}

func (pq *PriorityQueue) PushNode(node, weight int) {
	pq.q.PushNode(node, weight)
	sort.Slice(pq.q.a, func(i, j int) bool { return pq.q.a[i].Weight < pq.q.a[j].Weight })
}

func (pq *PriorityQueue) Top() *Edge {
	return pq.q.Front()
}

func (pq *PriorityQueue) Pop() *Edge {
	return pq.q.Pop()
}
