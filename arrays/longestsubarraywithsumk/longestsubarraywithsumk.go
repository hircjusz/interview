package main

import "fmt"

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func longestSubarraySumK(arr []int, k int) int {

	var ans int

	var startWindIx int
	var endWindIx int

	var sum int
	for ; endWindIx < len(arr); endWindIx++ {

		sum += arr[endWindIx]

		if k == sum {
			ans = max(ans, endWindIx-startWindIx+1)
		}

		for sum > k {
			sum -= arr[startWindIx]
			startWindIx++
		}

	}
	return ans

}

func main() {

	arr := []int{1, 2, 1, 0, 1}
	result := longestSubarraySumK(arr, 4)
	fmt.Println(result)
}
