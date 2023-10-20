package prims

import (
	"github.com/stretchr/testify/assert"
	"interview/graphs/graph"
	"testing"
)

func TestPrims(t *testing.T) {

	t.Run("Prime 1", func(t *testing.T) {

		g := graph.NewGraph()

		g.Push(1, 2, 10)
		g.Push(1, 3, 12)
		g.Push(2, 3, 9)
		g.Push(2, 4, 8)
		g.Push(3, 5, 3)
		g.Push(3, 6, 1)
		g.Push(4, 5, 7)
		g.Push(5, 6, 3)
		g.Push(4, 6, 10)
		g.Push(4, 7, 8)
		g.Push(4, 8, 5)
		g.Push(6, 8, 6)
		g.Push(7, 8, 9)
		g.Push(7, 9, 2)
		g.Push(8, 9, 11)

		assert.Equal(t, 9, g.NodesCount())

		ans, mst := primMst(g, 1)
		assert.Equal(t, 43, ans)
		assert.Equal(t, 9, len(mst))

	})
}
