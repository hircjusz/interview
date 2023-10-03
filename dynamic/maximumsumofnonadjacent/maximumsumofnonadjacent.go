package maximumsumofnonadjacent

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

//1, 2, 3, 1, 3, 5, 8, 1, 9
// dp[j]- max sum adjacent till j

//1, 2, 3, 1, 3, 5, 8, 1, 9
//           dp[j]=max(dp[j-2] +arr[j],dp[j]
//dp[0]=1

func maximumsumofnonadjacent(arr []int) int {

	var ans int

	dp := make([]int, len(arr))
	dp[0] = 1
	dp[1] = 2

	for j := 2; j < len(arr); j++ {
		dp[j] = max(dp[j-2]+arr[j], dp[j-1])
		ans = max(ans, dp[j])
	}
	return ans
}
