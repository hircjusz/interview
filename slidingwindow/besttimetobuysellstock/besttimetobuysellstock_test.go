package besttimetobuysellstock

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxProfit(t *testing.T) {

	arr := []int{7, 1, 5, 3, 6, 4}
	result := maxProfit(arr)

	assert.Equal(t, 5, result)
}
