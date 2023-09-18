package main

import "fmt"

func myPanic() {
	panic("a problem")
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered. error:\n", r)
		}
	}()

	myPanic()
	fmt.Println("After mayPanic()")
}
