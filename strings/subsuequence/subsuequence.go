package subsuequence

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func lcs(x, y string) int {

	m := len(x)
	n := len(y)

	dp := [][]int{}

	for i := 0; i <= m; i++ {
		dp = append(dp, make([]int, n+1))

	}
	x = " " + x
	y = " " + y

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if x[i] == y[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	return dp[m][n]

}
