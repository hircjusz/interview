package trappingrainwater

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestA(t *testing.T) {

	arr := []int{3, 0, 0, 2, 0, 4}
	result := traprainwater(arr)
	assert.Equal(t, 4, result)
}
