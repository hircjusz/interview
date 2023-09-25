package main

import (
	"fmt"
	"sort"
)

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

type stack[item any] struct {
	a []item
}

func newStack() *stack[interval] {
	return &stack[interval]{a: make([]interval, 0)}
}

func (s *stack[item]) push(i item) {
	s.a = append(s.a, i)
}

func (s *stack[item]) pop() item {
	var v item
	ix := len(s.a) - 1
	v, s.a = s.a[ix], s.a[:ix]
	return v
}

func (s *stack[item]) top() item {
	return s.a[len(s.a)-1]
}
func (s *stack[item]) empty() bool {
	return len(s.a) == 0
}

type interval struct {
	start int
	end   int
}

func mergeIntervals(a []int) []int {

	var intervals []interval

	for len(a) >= 1 {
		x := a[0]
		y := a[1]
		intervals = append(intervals, interval{start: x, end: y})
		a = a[2:]
	}

	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	st := newStack()

	st.push(intervals[0])

	for i := 1; i < len(intervals); i++ {

		if st.top().end <= intervals[i].start {
			st.push(intervals[i])
		} else {
			item := st.pop()
			newStart := min(item.start, intervals[i].start)
			newEnd := max(item.end, intervals[i].end)
			st.push(interval{start: newStart, end: newEnd})
		}
	}

	reconstruct := []int{}

	for !st.empty() {

		item := st.pop()
		reconstruct = append(reconstruct, item.start, item.end)
	}
	return reconstruct
}

func main() {

	//{{1,3},{2,4},{6,8},{9,10}}

	//  [ { ] }
	//
	//intervals := []interval{
	//	{start: 6, end: 8},
	//	{start: 1, end: 9},
	//	{start: 2, end: 4},
	//	{start: 4, end: 7},
	//}
	//sort.SliceStable(intervals, func(i, j int) bool {
	//	return intervals[i].start < intervals[j].start
	//})
	a := []int{1, 3, 2, 4, 6, 8, 9, 10, 1, 11}

	intervals := mergeIntervals(a)
	fmt.Println(intervals)
}
