package main

import (
	"fmt"
)

//https://leetcode.com/problems/house-robber/

//[2,7,9,3,1]

// dp(0,0)

// j=1
// dp(i,j)->dp(i+1,0)+prices[i]
// j=0 dp(i,j)->max (dp(i+1,0),dp(i+1,0)+prices[i])

/*
dp(i,0)=max(dp(i+1,0),dp(i+1,1))
dp(i,1)=dp(i+1,0)+prices[i]

profit[i]=max(dp(i,0),dp(i,1))
*/

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

var (
	memcached = map[int]map[bool]int{}
)

func dp(i int, rob bool, nums *[]int) int {

	if i == len(*nums) {
		return 0
	}
	if v, ok := memcached[i][rob]; ok {
		return v
	}

	var ans int
	if rob {
		ans = dp(i+1, !rob, nums) + (*nums)[i]
	} else {
		ans = max(dp(i+1, false, nums), dp(i+1, true, nums))
	}

	memcached[i] = map[bool]int{}
	memcached[i][rob] = ans
	return ans
}

func rob(nums []int) int {
	memcached = map[int]map[bool]int{}
	profit := max(dp(0, false, &nums), dp(0, true, &nums))
	return profit
}

func main() {
	nums := []int{2, 7, 9, 3, 1}
	result := rob(nums)
	fmt.Println(result)
}
