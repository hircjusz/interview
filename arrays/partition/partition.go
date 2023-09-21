package main

import "fmt"

var memcache map[int]int

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

func dp(i int, arr []int, k int) int {

	if i == len(arr) {
		return 0
	}
	if v, ok := memcache[i]; ok {
		return v
	}
	ans := 0
	maxEl := 0

	for j := i; j < min(i+k, len(arr)); j++ {
		maxEl = max(maxEl, arr[j])
		ans = max(ans, maxEl*(j-i+1)+dp(j+1, arr, k))
	}
	memcache[i] = ans
	return ans

}

func maxPartition(arr []int, k int) int {
	return dp(0, arr, k)
}

func main() {
	memcache = map[int]int{}
	memcache[1] = 2
	arr := []int{1, 15, 7, 9, 2, 5, 10}
	k := 3
	result := maxPartition(arr, k)

	fmt.Println("result ", result)
}
