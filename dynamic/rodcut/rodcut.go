package main

import "fmt"

var (
	prices    = []int{1, 3, 4, 5, 7, 9, 10, 11}
	memcached = map[int]int{}
)

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

//dp[n] -maximum prices if rod of length n is cut
//dp[n]=price[1]+dp[7]
//dp[n]=price[2]+dp[6]
//dp[n]=price[3]+dp[5]

func dp(n int) int {

	if n <= 0 {
		return 0
	}
	if v, ok := memcached[n]; ok {
		return v
	}
	var ans int

	for i := 0; i < n; i++ {
		ans = max(ans, dp(n-i-1)+prices[i])
	}
	memcached[n] = ans
	return ans
}

func main() {

	result := dp(8)
	fmt.Println(result)
}
