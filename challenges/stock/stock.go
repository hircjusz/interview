package main

import "fmt"

var memcache map[int]map[bool]int

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func dp(i int, buy bool, prices []int) int {

	if i == len(prices) {
		return 0
	}
	if v, ok := memcache[i][buy]; ok {
		return v
	}

	var ans int
	if buy {
		ans = max(-prices[i]+dp(i+1, !buy, prices), dp(i+1, buy, prices))

	} else {
		ans = max(prices[i]+dp(i+1, !buy, prices), dp(i+1, buy, prices))
	}

	if _, ok := memcache[i]; !ok {
		memcache[i] = make(map[bool]int)
	}

	memcache[i][buy] = ans
	return ans
}

func maxProfit(prices []int) int {
	return dp(0, true, prices)
}

func main() {
	prices := []int{3, 2, 6, 5, 0, 3}
	memcache = map[int]map[bool]int{}

	profit := maxProfit(prices)
	fmt.Println(profit)
}
