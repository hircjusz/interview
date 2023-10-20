package dsu

type Dsu struct {
	parent []int
}

func NewDsu() *Dsu {
	return &Dsu{
		parent: make([]int, 0),
	}
}

func (d *Dsu) findSet(i int) int {
	if d.parent[i] == -1 {
		return i
	}
	return d.findSet(d.parent[i])
}

func (d *Dsu) UnionSet(v, w int) {
	s1 := d.findSet(v)
	s2 := d.findSet(w)

	if s1 != s2 {
		d.parent[s1] = s2
	}
}
