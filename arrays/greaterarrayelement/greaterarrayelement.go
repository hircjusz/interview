package main

import "fmt"

type stack struct {
	a []int
}

func newStack() *stack {
	return &stack{a: make([]int, 0)}
}

func (s *stack) push(i int) {
	s.a = append(s.a, i)
}

func (s *stack) pop() int {
	var v int
	v, s.a = s.a[len(s.a)-1], s.a[:len(s.a)-1]
	return v
}

func (s *stack) top() int {
	return s.a[len(s.a)-1]
}
func (s *stack) empty() bool {
	return len(s.a) == 0
}

func greaterelement(arr []int) []int {
	greater := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		greater[i] = -1
	}

	st := newStack()
	st.push(0)

	for i := 0; i < len(arr); i++ {

		for !st.empty() && arr[st.top()] < arr[i] {
			k := st.pop()
			greater[k] = arr[i]
		}

		st.push(i)
	}

	return greater
}

func main() {
	//Input:  [2, 7, 3, 5, 4, 6, 8]
	//	Output: [7, 8, 5, 6, 6, 8, -1]
	//
	//
	//	Input:  [5, 4, 3, 2, 1]
	//	Output: [-1, -1, -1, -1, -1]

	//arr := []int{2, 7, 3, 5, 4, 6, 8}
	arr := []int{5, 4, 3, 2, 1}
	gr := greaterelement(arr)
	fmt.Println(gr)
}
