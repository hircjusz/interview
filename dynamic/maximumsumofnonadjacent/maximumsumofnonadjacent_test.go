package maximumsumofnonadjacent

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1(t *testing.T) {

	arr := []int{1, 2, 3, 1, 3, 5, 8, 1, 9}
	result := maximumsumofnonadjacent(arr)
	assert.Equal(t, 24, result)
}

func Test2(t *testing.T) {

	arr := []int{1, 2, 3, 5, 4}
	result := maximumsumofnonadjacent(arr)
	assert.Equal(t, 8, result)
}
