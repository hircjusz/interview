package main

import (
	"container/heap"
	"time"
)

type Request struct {
	fn func() int
	c  chan int
}

type Worker struct {
	requests chan Request
	pending  int
	index    int
}

type myHeap []int

func requester(work chan<- Request) {

	c := make(chan int)
	for {
		time.Sleep(2 * time.Second)
		work <- Request{fn: func() int {
			return 1
		}, c: c}
	}
}

func main() {
	h := heap.Init(myHeap{})
}
