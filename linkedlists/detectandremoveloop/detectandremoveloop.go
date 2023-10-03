package detectandremoveloop

type Node struct {
	Value int
	Next  *Node
}

func buildLinkedList(arr []int, i int) *Node {

	if i >= len(arr) || arr[i] == -1 {
		return nil
	}

	return &Node{
		Value: arr[i],
		Next:  buildLinkedList(arr, i+1),
	}
}

func listToArray(node *Node, arr *[]int) {

	if node == nil {
		*arr = append(*arr, -1)
		return
	}

	*arr = append(*arr, node.Value)
	listToArray(node.Next, arr)
}

func detectandremoveloop(root *Node) *Node {

	nodeMap := map[*Node]bool{}

	var nodeToRemove *Node
	var prev *Node
	node := root

	for node != nil {
		if nodeMap[node] == true {
			nodeToRemove = node

			break
		}
		nodeMap[node] = true
		prev = node
		node = node.Next
	}
	if nodeToRemove != nil {
		prev.Next = nil
	}
	return prev
}
