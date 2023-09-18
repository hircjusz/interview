package basicconcurrency

import (
	"fmt"
	"testing"
)

func TestPrintHello(t *testing.T) {

	printHello()
}

func TestManyGoroutines(t *testing.T) {
	manyGoroutines()
}

func TestCallbackFn(t *testing.T) {

	toUpperSync("somewor", func(str string) {
		fmt.Printf("some magic fn %s", str)
	})
}

func TestMessageChannel(t *testing.T) {

	channelMessage()
}

func TestChannelAwait(t *testing.T) {
	//channelAwaitMessage()
}

func TestBlockingChannel(t *testing.T) {
	blockingChannelMessage()
}
