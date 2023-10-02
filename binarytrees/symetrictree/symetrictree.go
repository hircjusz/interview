package main

import (
	"fmt"
	"interview/binarytrees/treebuilder"
)

type queue struct {
	a []*treebuilder.Node
}

func (q *queue) front() *treebuilder.Node {
	return q.a[0]
}

func (q *queue) push(n *treebuilder.Node) {
	q.a = append(q.a, n)
}

func (q *queue) pop() *treebuilder.Node {
	var e *treebuilder.Node
	e, q.a = q.a[0], q.a[1:]
	return e
}
func (q *queue) empty() bool {
	return len(q.a) == 0
}

func countLevel(node *treebuilder.Node, level int) {

	if node == nil {
		return
	}

	node.Level = level
	countLevel(node.Left, level+1)
	countLevel(node.Right, level+1)
}
func isSymtericArray(arr []int) bool {
	low := 0
	high := len(arr) - 1

	isSymetric := true
	for low < high {

		if arr[low] != arr[high] {
			isSymetric = false
			break
		}
		low++
		high--
	}
	return isSymetric
}

func isSymetricTree(arr []int) bool {

	tr := treebuilder.CreateTree(arr)
	countLevel(tr.Root, 0)

	nodeValues := map[int][]int{}

	q := queue{
		a: make([]*treebuilder.Node, 0),
	}

	q.push(tr.Root)

	for !q.empty() {
		node := q.pop()

		if nodeValues[node.Level] == nil {
			nodeValues[node.Level] = make([]int, 0)
		}
		nodeValues[node.Level] = append(nodeValues[node.Level], node.Value)
		if node.Left != nil {
			q.push(node.Left)
		}
		if node.Right != nil {
			q.push(node.Right)
		}

	}

	isSymetric := true
	for _, arr = range nodeValues {
		isSymetric = isSymtericArray(arr)
		if !isSymetric {
			break
		}
	}
	return isSymetric

}

func main() {

	arr := []int{1, 2, 2, 3, 4, 4, 3, -1, -1, -1, -1, -1, -1, -1, -1}
	//tr := treebuilder.CreateTree(arr)
	//countLevel(tr.Root, 0)

	result := isSymetricTree(arr)
	fmt.Println(result)
}
