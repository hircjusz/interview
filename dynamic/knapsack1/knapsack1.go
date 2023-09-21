package main

import "fmt"

var (
	sizeArr  = []int{4, 1, 2, 3, 2, 2}
	valueArr = []int{5, 8, 4, 0, 5, 3}
	memcache = map[int]map[int]int{}
)

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func knapsack(index, size int) int {

	if index == 0 {
		return 0
	}
	if v, ok := memcache[index][size]; ok {
		return v
	}

	//include
	var ans int
	if sizeArr[index] <= size {
		ans = valueArr[index] + knapsack(index-1, size-sizeArr[index])
	}
	ans = max(ans, knapsack(index-1, size))
	memcache[index] = make(map[int]int)
	memcache[index][size] = ans
	return ans
}

func main() {

	result := knapsack(len(sizeArr)-1, 4)

	fmt.Println(result)

}
