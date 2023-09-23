package main

import "fmt"

//https://leetcode.com/problems/max-consecutive-ones-iii/solutions/1638665/go-sliding-window/

// func longestOnes(nums []int, k int) int {
//
// }
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func longestOnes(nums []int) int {

	var windEnd int
	var maxLength int
	var length int

	for windEnd = 0; windEnd < len(nums); windEnd++ {

		if nums[windEnd] == 1 {
			length++
		}

		if nums[windEnd] == 0 {
			//	winStart=windEnd
			length = 0
		}

		maxLength = max(length, maxLength)

	}

	return maxLength
}

func main() {

	var nums = []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	//k := 2

	result := longestOnes(nums)

	fmt.Println(result)

}
