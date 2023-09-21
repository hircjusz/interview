package main

import "fmt"

var (
	sizeArr  = []int{4, 1, 2, 3, 2, 2}
	valueArr = []int{5, 8, 4, 0, 5, 3}
	memcache = map[int]int{}
)

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func knapsack(size int) int {

	if size <= 0 {
		return 0
	}
	if v, ok := memcache[size]; ok {
		return v
	}

	var ans int
	for i, s := range sizeArr {
		if s <= size {
			ans = max(ans, valueArr[i]+knapsack(size-s))
		}
	}
	memcache[size] = ans
	return ans
}

func main() {

	result := knapsack(4)

	fmt.Println(result)

}
