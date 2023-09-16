package simple_interview

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCopyMap(t *testing.T) {

	a := map[string]bool{"A": true, "B": false}
	b := copyMap(a)

	assert.Equal(t, a, b)
}
