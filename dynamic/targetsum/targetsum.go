package targetsum

func targetsum(arr []int, k int) int {

	ans := dp(len(arr)-1, 0, true, 0, k) + dp(len(arr)-1, 0, false, 0, k)
	return ans
}

func dp(maxIndex int, i int, add bool, sum int, target int) int {

	//if sum == target && i == maxIndex {
	//	return 1
	//}

	if i > maxIndex {
		return 0
	}
	if add {
		sum = sum + 1
	} else {
		sum = sum - 1
	}
	if sum == target && i == maxIndex {
		return 1
	}

	var ans int
	ans = dp(maxIndex, i+1, true, sum, target) + dp(maxIndex, i+1, false, sum, target)

	return ans

}
