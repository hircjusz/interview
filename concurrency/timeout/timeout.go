package main

import "time"

func main() {

	result := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		result <- 42
	}()

	select {
	case res := <-result:
		println("result:", res)
	case <-time.After(2 * time.Second):
		println("time out")
	}
}
