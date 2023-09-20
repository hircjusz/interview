package main

import (
	"fmt"
	"math/rand"
	"time"
)

func do(job string) error {
	fmt.Println("doing job", job)
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
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
