package main

import "fmt"

func stockspannaive(price []int) []int {

	spans := make([]int, len(price))

	for i := len(price) - 1; i >= 0; i-- {
		spans[i] = 1
		j := i - 1

		for j >= 0 && price[j] < price[i] {
			spans[i]++
			j--
		}
	}

	return spans

}

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

func stockspan(price []int) []int {
	spans := make([]int, len(price))

	st := newStack()

	spans[0] = 1
	st.push(0)

	for i := 1; i < len(price); i++ {

		for !st.empty() && price[st.top()] <= price[i] {
			st.pop()
		}

		if st.empty() {
			spans[i] = i + 1

		} else {
			spans[i] = i - st.top()
		}

		st.push(i)
	}

	return spans
}

func main() {

	prices := []int{100, 80, 60, 70, 60, 75, 85}
	//prices := []int{10, 4, 5, 90, 120, 80}

	result := stockspannaive(prices)
	fmt.Println(result)

	result = stockspan(prices)
	fmt.Println(result)
}
