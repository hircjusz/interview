package main

import (
	"fmt"
	"interview/binarytrees/treebuilder"
)

func lca(arr []int, a, b int) int {
	tr := treebuilder.CreateTree(arr)

	pathA := []int{}
	pathSearch(tr.Root, a, &pathA)

	pathB := []int{}
	pathSearch(tr.Root, b, &pathB)

	var lc int

	for {

		if len(pathA) == 0 || len(pathB) == 0 {
			break
		}

		a := pathA[len(pathA)-1]
		b := pathB[len(pathB)-1]
		if a == b {
			lc = a
		}
		pathA = pathA[:len(pathA)-1]
		pathB = pathB[:len(pathB)-1]
	}

	return lc
}

func pathSearch(node *treebuilder.Node, v int, path *[]int) bool {
	if node == nil {
		return false
	}

	if node.Value == v {
		*path = append(*path, node.Value)
		return true
	}

	var found bool
	if node.Value < v {
		found = pathSearch(node.Left, v, path)
	} else {
		found = pathSearch(node.Right, v, path)
	}
	if found {
		*path = append(*path, node.Value)
		return true
	}
	return false

}

func main() {
	arr := []int{1, 2, 3, 4, -1, 5, 6, -1, 7, -1, -1, -1, -1, -1, -1}

	result := lca(arr, 4, 2)
	fmt.Println(result)
}
