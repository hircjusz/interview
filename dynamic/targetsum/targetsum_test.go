package targetsum

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {

	arr := []int{1, 1, 1, 1, 1}

	result := targetsum(arr, 3)
	fmt.Println(result)
}
