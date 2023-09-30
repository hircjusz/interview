package main

import "fmt"

type node struct {
	isTerminal bool
	chars      []*node
}

type trie struct {
	R    int
	root *node
}

func newTrie() *trie {
	return &trie{
		root: newNode(),
	}
}
func newNode() *node {
	return &node{
		chars: make([]*node, 128),
	}
}

func (t *trie) Put(key string) {

	t.put(t.root, key, 0)
}

func (t *trie) put(x *node, key string, d int) *node {

	if x == nil {
		x = newNode()
	}

	if len(key) == d {
		x.isTerminal = true
		return x
	}
	ch := key[d]
	x.chars[ch] = t.put(x.chars[ch], key, d+1)
	return x
}

func (t *trie) Contains(key string) bool {

	return t.contains(t.root, key, 0)
}

func (t *trie) contains(x *node, key string, d int) bool {

	if x == nil {
		return false
	}
	if x.isTerminal == true && len(key) == d {
		return true
	}
	if d >= len(key) {
		return false
	}
	ch := key[d]
	return t.contains(x.chars[ch], key, d+1)
}

func (t *trie) keysWithPrefix(prefix string) []string {

	return []string{}
}

func main() {

	tr := newTrie()

	tr.Put("seller")
	tr.Put("sell")
	tr.Put("selfish")
	tr.Put("selector")
	result := tr.Contains("seller")

	fmt.Println(result)
}
