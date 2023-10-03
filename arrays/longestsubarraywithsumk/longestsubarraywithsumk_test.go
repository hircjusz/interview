package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1(t *testing.T) {

	arr := []int{1, 2, 1, 0, 1}
	result := longestSubarraySumK(arr, 4)
	assert.Equal(t, 4, result)
}
