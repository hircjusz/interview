package main

import "fmt"

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func coinschange(coins []int, m int) int {

	dp := make([]int, m+1)

	for _, c := range coins {
		dp[c] = 1
	}

	for i := 1; i < m+1; i++ {

		for _, c := range coins {

			if i-c > 0 && dp[i] != 1 {

				if dp[i] == 0 {
					dp[i] = dp[i-c] + 1
				} else {
					dp[i] = min(dp[i], dp[i-c]+1)
				}

			}
		}

	}

	return dp[m]
}

func main() {

	coins := []int{1, 3, 7, 10}
	m := 15
	result := coinschange(coins, m)

	fmt.Println(result)
}
