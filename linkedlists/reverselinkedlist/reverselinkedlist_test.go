package reverselinkedlist

import (
	"fmt"
	"testing"
)

func TestBuildList(t *testing.T) {

	arr := []int{1, 2, 3, 4, 5, 6, -1}
	root := buildLinkedList(arr, 0)
	fmt.Println(root.Value)

	arr2 := []int{}
	listToArray(root, &arr2)
	fmt.Println(arr2)
}

func TestReverse(t *testing.T) {

	arr := []int{1, 2, 3, 4, 5, 6, -1}
	root := buildLinkedList(arr, 0)

	head := reverseArray(root)
	arr2 := make([]int, 0)
	listToArray(head, &arr2)
	fmt.Println(arr2)
}
