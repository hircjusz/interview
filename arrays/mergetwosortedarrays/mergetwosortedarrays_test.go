package mergetwosortedarrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerged(t *testing.T) {
	arr1 := []int{1, 3, 5, 7, 9, 11}
	arr2 := []int{2, 4, 6, 8, 10}

	result := mergetwosortedarrays(arr1, arr2)

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, result)
}
