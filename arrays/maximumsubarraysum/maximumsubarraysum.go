package maximumsubarraysum

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func maximumsubarraysum(arr []int) int {

	var ans int
	var sum int

	//{1, 2, 7, -4, 3, 2, -10, 9, 1}
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		ans = max(ans, sum)

		if sum < 0 {
			sum = 0
		}
	}

	return ans
}
