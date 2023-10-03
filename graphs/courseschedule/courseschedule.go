package main

import "fmt"

type digraph struct {
	adj map[int][]int
}

func (g *digraph) push(v, w int) {

	if g.adj[v] == nil {
		g.adj[v] = make([]int, 0)
	}

	g.adj[v] = append(g.adj[v], w)

}

func contains(elems []int, v int) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func dfs(g *digraph, v int, visited *map[int]bool, inStack *[]int) bool {

	if (*visited)[v] == true {
		return false
	}

	if contains(*inStack, v) {
		return true
	}

	(*visited)[v] = true
	*inStack = append(*inStack, v)

	var isCycle bool
	for _, w := range g.adj[v] {
		isCycle = dfs(g, w, visited, inStack)
		if isCycle {
			return true
		}
	}

	return isCycle

}

func main() {

	g := digraph{
		adj: map[int][]int{},
	}

	g.push(4, 0)
	g.push(4, 2)
	g.push(1, 2)
	g.push(2, 3)

	inStack := []int{}
	visited := map[int]bool{}

	isCycle := dfs(&g, 2, &visited, &inStack)
	fmt.Println(isCycle)
}
