package main

func main() {

	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	for val := range ch {
		println(val)
	}

	println("Done")
}
