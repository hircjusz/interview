package simple_interview

import (
	"container/list"
	"fmt"
	"strings"
)

func copyMap(a map[string]bool) map[string]bool {

	b := make(map[string]bool)

	for key, value := range a {
		b[key] = value
	}
	return b
}

func reverse(s []int) {

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func concatanateString(arr ...string) string {

	var builder strings.Builder

	for _, val := range arr {
		builder.WriteString(val)
	}
	return builder.String()
}

func useList() {

	l := list.New()

	l.PushBack("1")
	l.PushBack("2")
	l.PushBack("3")

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}

}
