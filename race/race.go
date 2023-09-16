package main

import (
	"sync"
	"time"
)

type Counter struct {
	sync.Mutex
	value int
}

// go env -w CGO_ENABLED=0
func main() {

	counter := Counter{}

	for i := 0; i < 10; i++ {

		go func(i int) {
			counter.Lock()
			counter.value++
			defer counter.Unlock()
		}(i)
	}

	time.Sleep(time.Second)
	counter.Lock()
	defer counter.Unlock()
	println(counter.value)
}
