package detectandremoveloop

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {

	arr := []int{1, 2, 3, 4, 5, 6}

	root := buildLinkedList(arr, 0)
	//var last *Node
	head := root

	for head.Next != nil {
		head = head.Next
	}

	head.Next = root.Next

	fmt.Println(arr)

	detectandremoveloop(root)
	arr2 := make([]int, 0)
	listToArray(root, &arr2)

	fmt.Println(arr2)
}
