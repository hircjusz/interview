package main

import "fmt"

func search(arr []int, ix int) (int, bool) {

	if len(arr) == 0 {
		return 0, false
	}

	if len(arr) == 1 && ix == 0 {
		return arr[0], true
	}

	left, right := partition(arr, ix)

	if len(left)-1 == ix {
		return left[ix], true
	}

	if ix < len(left) {
		return search(left, ix)
	}
	return search(right, ix-len(left))
}

func partition(arr []int, ix int) (left []int, right []int) {
	pivot := arr[ix]

	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		}
	}

	return append(left, pivot), right
}

func findKthLargest(nums []int, k int) int {
	seen := map[int]bool{}
	distinct := []int{}
	for _, v := range nums {
		if !seen[v] {
			distinct = append(distinct, v)
			seen[v] = true
		}
	}

	el, _ := search(distinct, len(nums)-k)
	return el
}

func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	result := findKthLargest(nums, 4)
	fmt.Println(result)
}
