package main

import "fmt"

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func slidingwindow(a []int, val int) int {

	var i, j int

	sum := 0
	var cnt int
	move := true
	for i <= j && j < len(a) {

		if move {
			sum += a[j]
		}

		if sum == val {
			cnt++
			move = true
			j++
		}

		if sum < val {
			j++
			move = true
		}
		if sum > val {
			sum -= a[i]
			i++
			j = max(i, j)
			move = false
		}

	}
	return cnt
}

func main() {
	//a := []int{1, 7, 4, 3, 1, 2, 5, 1}
	a := []int{1, 1, 1, 1, 1, 1, 1, 7, 1, 6}

	result := slidingwindow(a, 7)

	fmt.Println(result)
}
