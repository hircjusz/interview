package treebuilder

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node

	leftVisited  bool
	rightVisited bool
}

type Tree struct {
	Root *Node
}

type queue struct {
	a []*Node
}

func (q *queue) front() *Node {
	return q.a[0]
}

func (q *queue) push(n *Node) {
	q.a = append(q.a, n)
}

func (q *queue) pop() *Node {
	var e *Node
	e, q.a = q.a[0], q.a[1:]
	return e
}
func (q *queue) empty() bool {
	return len(q.a) == 0
}

func insertNode(root *Node, value int, q *queue) *Node {

	node := &Node{
		Value: value,
	}
	if root == nil {
		root = node
	} else if value == -1 && !q.front().leftVisited {
		q.front().leftVisited = true
	} else if value == -1 && !q.front().rightVisited {
		q.front().rightVisited = true
	} else if value > 0 && q.front().Left == nil && !q.front().leftVisited {
		q.front().Left = node
	} else if !q.front().rightVisited {
		q.front().Right = node
		q.pop()
	}

	if !q.empty() && q.front().leftVisited && q.front().rightVisited {
		q.pop()
	}
	if value != -1 {
		q.push(node)
	}
	return root
}

func CreateTree(arr []int) *Tree {

	var root *Node
	q := &queue{a: make([]*Node, 0)}

	for i := 0; i < len(arr); i++ {
		root = insertNode(root, arr[i], q)
	}
	return &Tree{
		root,
	}
}

//
//func buildTree(level int, arr []int) *Node {
//
//	var node *Node
//	if level > 0 && level <= len(arr) {
//		value := arr[level-1]
//		//leaf node
//		if value == -1 {
//			return nil
//		}
//
//		left := buildTree(2*level, arr)
//		right := buildTree(2*level+1, arr)
//		node = &Node{
//			Value: value,
//			Left:  left,
//			Right: right,
//		}
//	}
//
//	return node
//}
//
//func NewTree(arr []int) *Tree {
//	root := buildTree(1, arr)
//	tr := Tree{
//		Root: root,
//	}
//	return &tr
//}

func main() {

	arr := []int{1, 2, 3, -1, 10, 5, 6, -1, -1, -1, 8, -1, -1, -1, -1}

	tr := CreateTree(arr)
	fmt.Println(tr)

}
