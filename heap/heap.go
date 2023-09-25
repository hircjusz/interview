package main

import "fmt"

type pq struct {
	a []int
	n int
}

func NewPQ() *pq {
	return &pq{
		a: make([]int, 1),
		n: 1,
	}
}

func (r *pq) empty() bool {
	return r.n <= 1
}

func (r *pq) insert(e int) {
	r.a = append(r.a, e)
	r.n++
	fixUp(r.a, r.n-1)
}

func (r *pq) getMax() int {

	r.a[1], r.a[r.n-1] = r.a[r.n-1], r.a[1]
	fixDown(r.a, 1, r.n-1)
	ix := r.n - 1
	r.n -= 1
	return r.a[ix]
}

func fixDown(a []int, k int, n int) {

	for 2*k < n {
		j := 2 * k

		if j+1 < n && a[j] < a[j+1] {
			j++
		}
		if a[k] > a[j] {
			break
		} else {
			a[j], a[k] = a[k], a[j]
		}
		k = j
	}

}

func fixUp(a []int, k int) {

	for k > 1 && a[k/2] < a[k] {
		a[k/2], a[k] = a[k], a[k/2]
		k = k / 2
	}
}

func sortHeap(a []int) {

	for i := 1; i < len(a)/2; i++ {
		fixDown(a, i, len(a)/2)
	}

	n := len(a) - 1
	for n > 1 {
		a[1], a[n] = a[n], a[1]
		n--
		fixDown(a, 1, n)
	}
}

func main() {

	a := []int{-1, 100, 90, 80, 70, 60, 50, 40, 110}

	sortHeap(a)
	fmt.Println(a)
}
