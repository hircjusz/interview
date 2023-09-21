package main

import (
	"fmt"
	"math"
)

var memcache map[int]map[int]int

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

func dp(i int, arr []int, m int) int {

	if i == len(arr) {
		if m == 0 {
			return 0
		}
		return math.MaxInt
	}

	if v, ok := memcache[i][m]; ok {
		return v
	}

	if m <= 0 {
		return math.MaxInt
	}

	var sum int
	ans := math.MaxInt

	for j := i; j < len(arr); j++ {
		sum += arr[j]
		ans = min(ans, max(sum, dp(j+1, arr, m-1)))
	}
	memcache[i] = make(map[int]int)
	memcache[i][m] = ans

	return ans
}

func maxSplitSum(arr []int, m int) int {
	return dp(0, arr, m)
}

func main() {
	memcache = make(map[int]map[int]int)
	arr := []int{7, 2, 5, 10, 8}
	result := maxSplitSum(arr, 2)

	fmt.Println(result)
}
