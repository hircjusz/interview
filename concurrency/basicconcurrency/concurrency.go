package basicconcurrency

import (
	"fmt"
	"strings"
	"sync"
)

func printHello() {
	var wait sync.WaitGroup
	wait.Add(1)

	message := "Hello"

	go func(msg string) {
		fmt.Println(msg)
		wait.Done()
	}(message)

	wait.Wait()
}

func manyGoroutines() {
	var wait sync.WaitGroup
	goRoutines := 5

	wait.Add(goRoutines)

	for i := 0; i < goRoutines; i++ {
		go func(goRoutineId int) {
			fmt.Printf("ID : %d: Hello concurrency !\n", i)
			wait.Done()
		}(i)
	}
	wait.Wait()
}

func toUpperSync(word string, f func(str string)) {
	f(strings.ToUpper(word))
}

func channelMessage() {
	channel := make(chan string)
	go func() {
		channel <- "Hello world!"
	}()

	messsage := <-channel
	println(messsage)
}

func channelAwaitMessage() {
	channel := make(chan string)
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	go func() {
		channel <- "Hello world"
		println("Finishing goroutine")
		waitGroup.Done()
	}()

	message := <-channel
	println(message)
	waitGroup.Wait()

}

func blockingChannelMessage() {
	channel := make(chan string, 1)

	go func(ch chan<- string) {
		channel <- "Hello world"
		channel <- "Hello world 2"
	}(channel)

	receivingCh(channel)

}

func receivingCh(ch <-chan string) {
	msg := <-ch
	println(msg)
}
