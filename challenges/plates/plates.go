package main

import "fmt"

type coordinates struct {
	start  int
	length int
}
type PlateIndexer struct {
	list map[string]*coordinates
}

func NewPlateIndexer(plates []string) *PlateIndexer {

	s := &PlateIndexer{
		list: make(map[string]*coordinates),
	}
	prev := plates[0]
	s.list[prev] = &coordinates{start: 0, length: 1}

	el := s.list[prev]

	plates = plates[1:]

	for _, p := range plates {
		if p == prev {
			el.length += 1
		} else {
			s.list[p] = &coordinates{start: el.length, length: 1}
			el = s.list[p]
		}
		prev = p
	}
	return s
}

func main() {

	plates := []string{
		"flower-decorated plate",
		"light green plate",
		"light green plate",
		"big blue plate",
		"big blue plate",
		"transparent plate",
	}
	s := NewPlateIndexer(plates)
	fmt.Println(s)
}
