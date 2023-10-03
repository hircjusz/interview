package main

import (
	"fmt"
	"math"
	"strconv"
)

func digits(n int) int {
	return int(math.Floor(math.Log10(float64(n))) + 1)
}

func toInt(str string) int {
	l, _ := strconv.Atoi(str)
	return l
}

func findMissingNumber(str string, k int) (int, bool) {

	var i int
	var n int

	n = toInt(str[0:k])
	arr := make([]int, 0)
	arr = append(arr, n)

	missing := -1
	found := false
	var invalid bool

	for i = k; i+digits(n) <= len(str); i += digits(n) {
		l := toInt(str[i : i+digits(n+1)])
		if l != n+1 && l == n+2 && !found {
			missing = n + 1
			found = true
			n++
		} else if l != n+1 && found {
			invalid = true
		}
		n++
	}
	return missing, invalid
}

func main() {

	result := findMissingNumber("9899101", 2)
	fmt.Println(result)

	fmt.Println(digits(1))
	fmt.Println(digits(99))
	fmt.Println(digits(100))
	fmt.Println(digits(900))
	fmt.Println(digits(1000))
}
