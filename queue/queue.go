package queue

type Edge struct {
	V      int
	Weight int
}

type Queue struct {
	a []*Edge
}

func New() *Queue {
	return &Queue{
		a: make([]*Edge, 0),
	}
}

func (q *Queue) Front() *Edge {
	return q.a[0]
}

func (q *Queue) End() *Edge {
	return q.a[len(q.a)-1]
}

func (q *Queue) PushEdge(n *Edge) {
	q.a = append(q.a, n)
}

func (q *Queue) PushNode(node, weight int) {
	q.a = append(q.a, &Edge{V: node, Weight: weight})
}

func (q *Queue) Pop() *Edge {
	var e *Edge
	e, q.a = q.a[0], q.a[1:]
	return e
}
func (q *Queue) Empty() bool {
	return len(q.a) == 0
}

func (q *Queue) FindByEdge(edge *Edge) *Edge {
	j := -1

	for i, node := range q.a {
		if *node == *edge {
			j = i
			break
		}
	}
	if j == -1 {
		return nil
	}
	return q.a[j]
}

func (q *Queue) FindByNode(v int) *Edge {
	j := -1

	for i, node := range q.a {
		if node.V == v {
			j = i
			break
		}
	}
	if j == -1 {
		return nil
	}
	return q.a[j]
}

func (q *Queue) EraseByEdge(erase *Edge) {
	j := -1

	for i, node := range q.a {
		if *node == *erase {
			j = i
			break
		}
	}
	if j == -1 {
		return
	}

	var q2 []*Edge
	q2 = append(q2, q.a[:j]...)
	q2 = append(q2, q.a[j+1:]...)
	q.a = q2

}

func (q *Queue) EraseByNode(v int) {
	j := -1

	for i, node := range q.a {
		if node.V == v {
			j = i
			break
		}
	}
	if j == -1 {
		return
	}

	var q2 []*Edge
	q2 = append(q2, q.a[:j]...)
	q2 = append(q2, q.a[j+1:]...)
	q.a = q2

}
