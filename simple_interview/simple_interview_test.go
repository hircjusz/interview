package simple_interview

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCopyMap(t *testing.T) {

	a := map[string]bool{"A": true, "B": false}
	b := copyMap(a)

	assert.Equal(t, a, b)

}

func TestReverse(t *testing.T) {

	a := []int{1, 2, 3}
	reverse(a)
	fmt.Println(a)
}

func TestConcatanateString(t *testing.T) {

	result := concatanateString("Ala", "ma", "kota")
	assert.Equal(t, "Alamakota", result)
}

func TestUseList(t *testing.T) {
	useList()
}
