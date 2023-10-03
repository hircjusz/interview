package reverselinkedlist

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

func reverseArray(node *Node) *Node {

	var current *Node
	var previous *Node
	var head *Node

	current = node.Next
	previous = node
	head = node.Next.Next
	// a ->b->c->d
	// p <- c h->

	previous.Next = nil

	for head != nil {

		current.Next = previous
		previous = current
		current = head
		head = head.Next
	}
	current.Next = previous

	return current
}
