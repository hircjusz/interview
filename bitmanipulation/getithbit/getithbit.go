package main

import "fmt"

func countbits(n int) int {
	ans := 0

	for n > 0 {
		n = n & (n - 1)
		ans++
	}
	return ans
}

func covertBinary(n int) int {
	var ans int
	var p int

	for n > 0 {
		lastBit := n & 1
		ans += p * lastBit
		p = p * 10
		n = n >> 1
	}
	return ans
}

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(countbits(i))
	}
}
