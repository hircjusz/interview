package main

import "fmt"

func main() {

	genMap := make([][]int, 10) // Make the outer slice and give it size 10
	for i := 0; i < 10; i++ {
		genMap[i] = make([]int, 10) // Make one inner slice per iteration and give it size 10
		for j := 0; j < 10; j++ {
			genMap[i][j] = i + j
		}
	}

	fmt.Println(genMap)

}
