package prims

import (
	"github.com/stretchr/testify/assert"
	"interview/graphs/graph"
	"testing"
)

func TestPrims(t *testing.T) {

	t.Run("Prime 1", func(t *testing.T) {

		g := graph.NewGraph()

		g.PushWeighted(1, 2, 10)
		g.PushWeighted(1, 3, 12)
		g.PushWeighted(2, 3, 9)
		g.PushWeighted(2, 4, 8)
		g.PushWeighted(3, 5, 3)
		g.PushWeighted(3, 6, 1)
		g.PushWeighted(4, 5, 7)
		g.PushWeighted(5, 6, 3)
		g.PushWeighted(4, 6, 10)
		g.PushWeighted(4, 7, 8)
		g.PushWeighted(4, 8, 5)
		g.PushWeighted(6, 8, 6)
		g.PushWeighted(7, 8, 9)
		g.PushWeighted(7, 9, 2)
		g.PushWeighted(8, 9, 11)

		assert.Equal(t, 9, g.NodesCount())

		ans, mst := primMst(g, 1)
		assert.Equal(t, 43, ans)
		assert.Equal(t, 9, len(mst))

	})
}
