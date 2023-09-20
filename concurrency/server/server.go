package main

import (
	"fmt"
	"time"
)

type Server struct {
	quit chan bool
}

func NewServer() *Server {
	s := &Server{make(chan bool)}
	go s.run()
	return s
}

func (s *Server) run() {
	for {
		select {
		case <-s.quit:
			println("finishing task")
			time.Sleep(time.Second)
			println("task done")
			s.quit <- true
			return
		case <-time.After(time.Second):
			println("running task")
		}
	}
}

func (s *Server) Stop() {
	fmt.Println("server stopping")
	s.quit <- true
	<-s.quit
	fmt.Println("server stopped")
}

func main() {
	s := NewServer()
	time.Sleep(5 * time.Second)
	s.Stop()
}
