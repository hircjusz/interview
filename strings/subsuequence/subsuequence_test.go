package subsuequence

import (
	"fmt"
	"testing"
)

func TestLcs(t *testing.T) {
	x := "AGGTAB"
	y := "GXTXAYB"

	result := lcs(x, y)
	fmt.Println(result)
}
