package dijkstra

import (
	"github.com/DrakeEsdon/Go-Snake/datatypes"
	"github.com/RyanCarrier/dijkstra"
	"math"
)

func addGameStateToGraph(request *datatypes.GameRequest, g *dijkstra.Graph, canGoToTail bool) *dijkstra.Graph {
	board := &request.Board
	you := request.You
	dangerousSnakeMoves := GetPossibleMovesOfEqualOrLargerSnakes(request)
	for x := 0; x < board.Width; x++ {
		for y := 0; y < board.Height; y++ {
			coordA := datatypes.Coord{X: x, Y: y}
			for _, direction := range datatypes.AllDirections {
				coordB := datatypes.AddDirectionToCoord(coordA, direction)
				var distance int64 = 1
				if coordA == you.Head {
					_, found := FindCoordInList(dangerousSnakeMoves, coordB)
					if found {
						distance += 99
					}

				}
				if !datatypes.IsOutOfBounds(coordB, *board) {
					if !datatypes.IsSnake(coordB, *board) ||
					    (datatypes.IsMyTail(coordB, request.You) && canGoToTail) {
						if datatypes.IsHazard(coordB, *board) {
							distance += 14
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

func GetPossibleMovesOfSnake(battlesnake datatypes.Battlesnake) []datatypes.Coord {
	possibleCoords := make([]datatypes.Coord, 0)
	for _, dir := range datatypes.AllDirections {
		destCoord := datatypes.AddDirectionToCoord(battlesnake.Head, dir)
		possibleCoords = append(possibleCoords, destCoord)
	}
	return possibleCoords
}

func GetPossibleMovesOfEqualOrLargerSnakes(request *datatypes.GameRequest) []datatypes.Coord {
	possibleCoords := make([]datatypes.Coord, 0)
	for _, snake := range request.Board.Snakes {
		if snake.ID != request.You.ID && snake.Length >= request.You.Length {
			snakeCoords := GetPossibleMovesOfSnake(snake)
			for _, coord := range snakeCoords {
				possibleCoords = append(possibleCoords, coord)
			}
		}
	}
	return possibleCoords
}


func FindCoordInList(slice []datatypes.Coord, element datatypes.Coord) (int, bool) {
	for i, item := range slice {
		if item == element {
			return i, true
		}
	}
	return -1, false
}
