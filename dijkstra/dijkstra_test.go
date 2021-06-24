package dijkstra

import (
	"github.com/RyanCarrier/dijkstra"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDijkstraShortest(t *testing.T) {
	graph := dijkstra.NewGraph()
	/*
	[2]
	 | \
	 |  \
	[1] [3]
	 */
	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)
	err := graph.AddArc(1, 2, 1)
	assert.Nil(t, err)

	err = graph.AddArc(2, 3, 1)
	assert.Nil(t, err)

	bestPath, err := graph.Shortest(1, 3)
	assert.Nil(t, err)

	assert.Equal(t, len(bestPath.Path), 3)
	assert.Equal(t, bestPath.Path[0], 1)
	assert.Equal(t, bestPath.Path[1], 2)
	assert.Equal(t, bestPath.Path[2], 3)
	assert.Equal(t, bestPath.Distance, int64(2))
}
