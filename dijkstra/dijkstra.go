package dijkstra

import (
	"github.com/DrakeEsdon/Go-Snake/datatypes"
	"github.com/RyanCarrier/dijkstra"
	"math"
)

func addGameStateToGraph(request *datatypes.GameRequest, g *dijkstra.Graph, canGoToTail bool) *dijkstra.Graph {
	board := &request.Board
	for x := 0; x < board.Width; x++ {
		for y := 0; y < board.Height; y++ {
			coordA := datatypes.Coord{X: x, Y: y}
			for _, direction := range datatypes.AllDirections {
				coordB := datatypes.AddDirectionToCoord(coordA, direction)
				if !datatypes.IsOutOfBounds(coordB, *board) {
					if !datatypes.IsSnake(coordB, *board) ||
					    (datatypes.IsMyTail(coordB, request.You) && canGoToTail) {
						var distance int64 = 1
						if datatypes.IsHazard(coordB, *board) {
							distance = 15
						}
						g.AddMappedVertex(datatypes.CoordToString(coordA))
						g.AddMappedVertex(datatypes.CoordToString(coordB))
						err := g.AddMappedArc(
							datatypes.CoordToString(coordA),
							datatypes.CoordToString(coordB),
							distance,
						)
						if err != nil {
							return nil
						}
					}
				}
			}
		}
	}
	return g
}

func GetDijkstraGraph(request *datatypes.GameRequest, canGoForTail bool) *dijkstra.Graph {
	graph := dijkstra.NewGraph()
	graph = addGameStateToGraph(request, graph, canGoForTail)
	return graph
}

func GetDijkstraPathDirection(from, to datatypes.Coord, graph *dijkstra.Graph) (*datatypes.Direction, int) {
	const errorLength = math.MaxInt64
	if graph == nil {
		// There was an issue adding board elements to the graph
		return nil, errorLength
	}
	fromID, err := graph.GetMapping(datatypes.CoordToString(from))
	if err != nil {
		return nil, errorLength
	}
	toID, err := graph.GetMapping(datatypes.CoordToString(to))
	if err != nil {
		return nil, errorLength
	}
	bestPath, err := graph.Shortest(fromID, toID)
	if err != nil {
		return nil, errorLength
	}
	firstID := bestPath.Path[0]
	secondID := bestPath.Path[1]
	firstString, err := graph.GetMapped(firstID)
	if err != nil {
		return nil, errorLength
	}
	secondString, err := graph.GetMapped(secondID)
	if err != nil {
		return nil, errorLength
	}
	firstCoord := datatypes.CoordFromString(firstString)
	secondCoord := datatypes.CoordFromString(secondString)
	diffX := secondCoord.X - firstCoord.X
	diffY := secondCoord.Y - firstCoord.Y
	var distance = int(bestPath.Distance)
	if diffX > 0 && diffY == 0 {
		return &datatypes.DirectionRight, distance
	} else if diffX < 0 && diffY == 0 {
		return &datatypes.DirectionLeft, distance
	} else if diffX == 0 && diffY > 0 {
		return &datatypes.DirectionUp, distance
	} else if diffX == 0 && diffY < 0 {
		return &datatypes.DirectionDown, distance
	} else {
		// Something went wrong...
		return nil, errorLength
	}
}