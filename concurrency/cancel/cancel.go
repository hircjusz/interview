package main

import (
	"fmt"
	"time"
)

func main() {

	cancel := make(chan struct{})

	go func() {
		for {
			select {
			case <-time.After(time.Second):
				println("working...")
			case <-cancel:
				println("Cancelled!")
				return
			}
		}
	}()

	println("Press enter to cancel...")
	fmt.Scanln()
	close(cancel)

}
