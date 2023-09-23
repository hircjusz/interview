package main

import "fmt"

func initarray(n, m int) [][]int {

	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	return dp
}
func longestcommon(str1, str2 string) int {

	var dp [][]int
	var max int
	n := len(str1)
	m := len(str2)

	dp = initarray(n, m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {

			if i == 0 || j == 0 {
				dp[i][j] = 0
			} else if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}

			if max < dp[i][j] {
				max = dp[i][j]
			}
		}
	}

	return max
}

/*
  dp[i,j]=longest common strin where index str1 i and str2 j
 dp[i,j]=1 s[i]==s[j]

dp[i,j]=dp[i-1,j-1] + 1   s[i-1]==s[j-1]

*/

func main() {
	str1 := "opengeneus"
	str2 := "genius"
	//

	result := longestcommon(str1, str2)
	fmt.Println(result)

}
