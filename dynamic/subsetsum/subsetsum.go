package main

import (
	"fmt"
)

// dp[n,y] [n.....][y]
// dp[0,12]=dp[1][10]+2
// dp[2,12]=dp
var (
	arr       = []int{2, 7, 4, 5, 9}
	memcached = map[int]map[int]bool{}
)

func subsetSum(index, sum int) bool {

	if index < 0 {
		return sum == 0
	}
	if v, ok := memcached[index][sum]; ok {
		return v
	}
	ans := false

	//include
	if sum >= arr[index] {
		ans = ans || subsetSum(index-1, sum-arr[index])
	}

	ans = ans || subsetSum(index-1, sum)
	memcached[index] = map[int]bool{}
	memcached[index][sum] = ans
	return ans
}

func main() {

	result := subsetSum(len(arr)-1, 12)
	fmt.Println(result)
}
