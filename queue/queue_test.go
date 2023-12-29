package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue(t *testing.T) {

	t.Run("Find", func(t *testing.T) {

		q := Queue{
			a: []*Edge{{V: 2}, {V: 5}, {V: 4}, {V: 1}, {V: 0}},
		}

		res := q.FindByEdge(&Edge{V: 4})
		assert.NotNil(t, res)
	})

	t.Run("Erase", func(t *testing.T) {

		q := Queue{
			a: []*Edge{{V: 2}, {V: 5}, {V: 4}, {V: 1}, {V: 0}},
		}

		q.EraseByEdge(&Edge{V: 4})
		assert.Equal(t, 4, len(q.a))
	})

	t.Run("Erase last", func(t *testing.T) {

		q := Queue{
			a: []*Edge{{V: 2}, {V: 5}, {V: 4}, {V: 1}, {V: 0}},
		}

		q.EraseByEdge(&Edge{V: 0})
		assert.Equal(t, 4, len(q.a))
	})

	t.Run("Erase first", func(t *testing.T) {

		q := Queue{
			a: []*Edge{{V: 2}, {V: 5}, {V: 4}, {V: 1}, {V: 0}},
		}

		q.EraseByEdge(&Edge{V: 0})
		assert.Equal(t, 4, len(q.a))
	})
}
