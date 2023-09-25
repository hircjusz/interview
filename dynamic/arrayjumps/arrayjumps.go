package main

import "fmt"

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func arrayjumps(a []int) int {

	m := len(a)
	dp := make([]int, m)

	for i := 0; i < m; i++ {
		for j := 0; j < i; j++ {

			if j+a[j] >= i {
				if dp[i] == 0 {
					dp[i] = 1
				} else {
					for k := 1; k <= a[j] && j+a[k] >= i; k++ {
						dp[i] = min(dp[i], dp[i-a[k]]+1)
					}
				}

			}
		}
	}

	return dp[m-1]
}

func main() {

	a := []int{3, 4, 2, 1, 2, 3, 10, 1, 1, 1, 2, 5}
	result := arrayjumps(a)
	fmt.Println(result)
}
