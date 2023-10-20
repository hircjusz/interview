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

func (q *Queue) Push(n *Edge) {
	q.a = append(q.a, n)
}

func (q *Queue) Pop() *Edge {
	var e *Edge
	e, q.a = q.a[0], q.a[1:]
	return e
}
func (q *Queue) Empty() bool {
	return len(q.a) == 0
}

func (q *Queue) Find(edge *Edge) *Edge {
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

func (q *Queue) Erase(erase *Edge) {
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
