package main

import (
	"fmt"
	"time"
)

func do(job string) error {
	fmt.Println("doing job", job)
	time.Sleep(1 * time.Second)
	return fmt.Errorf("something went wrong %s", job)
}

func main() {

	jobs := []string{"one", "two", "three"}

	errc := make(chan error)

	for _, job := range jobs {
		go func(j string) {
			errc <- do(j)
		}(job)
	}
	for _ = range jobs {
		if err := <-errc; err != nil {
			fmt.Println(err)
		}
	}
}
