package dijkstra

import (
	"github.com/DrakeEsdon/Go-Snake/datatypes"
	"github.com/RyanCarrier/dijkstra"
)

func addGameStateToGraph(board *datatypes.Board, g *dijkstra.Graph) *dijkstra.Graph {
	for x := 0; x < board.Width; x++ {
		for y := 0; y < board.Height; y++ {
			coordA := datatypes.Coord{X: x, Y: y}
			for _, direction := range datatypes.AllDirections {
				coordB := datatypes.AddDirectionToCoord(coordA, direction)
				if !datatypes.IsOutOfBounds(coordB, *board) &&
				    !datatypes.IsSnakeOrHazard(coordB, *board) {
					g.AddMappedVertex(datatypes.CoordToString(coordA))
					g.AddMappedVertex(datatypes.CoordToString(coordB))
					err := g.AddMappedArc(
						datatypes.CoordToString(coordA),
						datatypes.CoordToString(coordB),
						1,
					)
					if err != nil {
						return nil
					}
				}
			}
		}
	}
	return g
}

func GetDijkstraGraph(board datatypes.Board) *dijkstra.Graph {
	graph := dijkstra.NewGraph()
	graph = addGameStateToGraph(&board, graph)
	return graph
}

func GetDijkstraPathDirection(from, to datatypes.Coord, graph *dijkstra.Graph) *datatypes.Direction {
	if graph == nil {
		// There was an issue adding board elements to the graph
		return nil
	}
	fromID, err := graph.GetMapping(datatypes.CoordToString(from))
	if err != nil {
		return nil
	}
	toID, err := graph.GetMapping(datatypes.CoordToString(to))
	if err != nil {
		return nil
	}
	bestPath, err := graph.Shortest(fromID, toID)
	if err != nil {
		return nil
	}
	firstID := bestPath.Path[0]
	secondID := bestPath.Path[1]
	firstString, err := graph.GetMapped(firstID)
	if err != nil {
		return nil
	}
	secondString, err := graph.GetMapped(secondID)
	if err != nil {
		return nil
	}
	firstCoord := datatypes.CoordFromString(firstString)
	secondCoord := datatypes.CoordFromString(secondString)
	diffX := secondCoord.X - firstCoord.X
	diffY := secondCoord.Y - firstCoord.Y
	if diffX > 0 && diffY == 0 {
		return &datatypes.DirectionRight
	} else if diffX < 0 && diffY == 0 {
		return &datatypes.DirectionLeft
	} else if diffX == 0 && diffY > 0 {
		return &datatypes.DirectionUp
	} else if diffX == 0 && diffY < 0 {
		return &datatypes.DirectionDown
	} else {
		// Something went wrong...
		return nil
	}
}