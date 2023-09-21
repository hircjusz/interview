package main

import (
	"fmt"
	"hash/fnv"
)

//{ [  }         [] { ] }  [         ]

type line struct {
	x int
	y int
}

func (r *line) overlap(start, end int) bool {
	//start        //end
	//{     [r.x     }
	var overlap bool
	if r.x > start && r.x < end {
		r.x = start
		overlap = true
	}
	// start   r.y  end
	//{       ]    }
	if r.y > start && r.y < end {
		r.y = end
		overlap = true
	}

	if r.x >= start && r.y <= end {
		overlap = true
	}

	return overlap

}

func unpaint(paints []int) []int {

	points := []*line{}

	x := paints[0]
	y := paints[1]
	paints = paints[2:]

	points = append(points, &line{x, y})
	var overlap bool

	for {
		if len(paints) < 2 {
			break
		}
		overlap = false

		x := paints[0]
		y := paints[1]
		paints = paints[2:]

		for _, item := range points {
			overlap = item.overlap(x, y)
			if overlap {
				break
			}
		}
		if overlap == false {
			points = append(points, &line{x, y})
		}

	}

	var arr []int
	for _, x := range points {
		arr = append(arr, x.x, x.y)
	}
	return arr
}

func main() {

	h := fnv.New32a()
	h.Write([]byte("Hello"))
	fmt.Println(h.Sum32())

	paints := []int{3, 10, 14, 20, 1, 5}

	result := unpaint(paints)

	fmt.Println(result)
}
