package main

import "fmt"

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func uniquesubstring(str string) int {

	h := map[uint8]interface{}{}

	var start, end int
	var maxLength int

	h[str[start]] = 0
	maxLength = 1
	//a b c a e b a
	//s
	for end = 1; end < len(str); end = end + 1 {

		if h[str[end]] == nil {
			h[str[end]] = end
			maxLength = max(end-start+1, maxLength)
		} else {
			h[str[end]] = end
			start = start + 1
		}
	}

	return maxLength
}

func main() {
	str := "abcaeba"
	//      0123456
	result := uniquesubstring(str)
	fmt.Println(result)
}
