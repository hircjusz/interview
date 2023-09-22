package main

import (
	"fmt"
	"math"
)

func findClosestElements(arr []int, k int, x int) []int {

	var closest []int

	for _, v := range arr {
		closest = insertClosestInOrder(v, arr, x)
	}
	return closest[:k]
}

func insertClosestInOrder(v int, arr []int, x int) []int {

	inserted := false
	var temp []int
	for i := 0; i < len(arr); i++ {
		if math.Abs(float64(v-x)) <= math.Abs(float64(arr[i]-x)) {
			temp = append(temp, arr[:i]...)
			temp = append(temp, v)
			temp = append(temp, arr[i:]...)
			inserted = true
			break
		}
	}
	if !inserted {
		temp = append(arr, v)
	}
	return temp

}

func main() {
	//arr = [1,2,3,4,5], k = 4, x = 3

	//[2,1,0,1,2]
	//
	//Output: [1,2,3,4]
	arr := []int{-9, 10}
	k := 2
	x := 1

	result := findClosestElements(arr, k, x)
	//result := sortInOder(arr)

	fmt.Println(result)
}
