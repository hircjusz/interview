package prims

import (
	"interview/queue"
	"testing"
)

func TestPrims(t *testing.T) {

	t.Run("Prime 1", func(t *testing.T) {

		g := graph{
			adj: map[int][]*queue.Edge{},
		}
		g.push(0, 1, 1)
		g.push(1, 2, 1)
		g.push(0, 2, 4)
		g.push(0, 3, 7)
		g.push(3, 2, 2)
		g.push(3, 4, 3)

	})
}
