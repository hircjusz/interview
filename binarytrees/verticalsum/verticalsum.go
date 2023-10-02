package main

import (
	"fmt"
	"interview/binarytrees/treebuilder"
)

var sumVertical = make(map[int]int)

func traverseTree(node *treebuilder.Node, hd int) {

	if node == nil {
		return
	}

	traverseTree(node.Left, hd-1)
	sumVertical[hd] += node.Value
	traverseTree(node.Right, hd+1)
}

func verticalSum(arr []int) []int {
	tr := treebuilder.CreateTree(arr)
	traverseTree(tr.Root, 0)

	var ans []int

	for _, v := range sumVertical {
		ans = append(ans, v)
	}
	return ans
}

func main() {
	arr := []int{1, 2, 3, -1, 10, 5, 6, -1, -1, -1, 8, -1, -1, -1, -1}

	result := verticalSum(arr)
	fmt.Println(result)

	arr = []int{20, 8, 22, 5, 3, -1, -1, -1, -1, -1, -1}
	result = verticalSum(arr)
	fmt.Println(result)
}
