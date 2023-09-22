package main

import "fmt"

func longestMountain(a []int) int {

	up := true
	down := false
	//	maxLocal := arr[i]

	var maxLength int
	length := 1

	for i, j := 0, 1; j < len(a); i, j = i+1, j+1 {
		if up && a[i] < a[j] {
			length++
		} else if up {
			up = false
			down = true
		}

		if down && a[i] > a[j] {
			length++
		} else if down {
			up = true
			down = false
			length = 2
		}

		maxLength = max(length, maxLength)
	}
	return maxLength
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {

	//arr := []int{2, 1, 4, 7, 3, 2, 5}
	arr := []int{2, 1, 4, 7, 3, 2, 5, 1, 2, 3, 4, 5, 6, 7, 8, 7, 6, 5, 4, 3, 2, 1, 2, 5, 6}

	result := longestMountain(arr)

	fmt.Println(result)

}
