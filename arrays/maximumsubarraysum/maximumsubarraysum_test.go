package maximumsubarraysum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1(t *testing.T) {

	arr := []int{1, 2, 7, -4, 3, 2, -10, 9, 1}
	result := maximumsubarraysum(arr)
	assert.Equal(t, 11, result)
}
