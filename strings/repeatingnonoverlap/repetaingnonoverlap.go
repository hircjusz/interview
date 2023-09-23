package main

import "fmt"

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func initarray(n, m int) [][]int {

	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	return dp
}

func repeatingNonOverlap(s string) int {

	n := len(s)
	dp := initarray(n, n)

	result := 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {

			if i == 0 || j == 0 {
				dp[i][j] = 0
			} else if s[i-1] == s[j-1] && dp[i-1][j-1] < j-i {
				dp[i][j] = dp[i-1][j-1] + 1
				//
			}
			result = max(dp[i][j], result)
		}
	}
	return result
}

func main() {

	s := "banana"
	result := repeatingNonOverlap(s)
	fmt.Println(result)
}
