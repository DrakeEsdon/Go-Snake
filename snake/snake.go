package snake

import (
	"fmt"
	"github.com/DrakeEsdon/Go-Snake/datatypes"
	"github.com/DrakeEsdon/Go-Snake/dijkstra"
	dijkstra2 "github.com/RyanCarrier/dijkstra"
	"math"
	"math/rand"
)

var turnsSinceEating = 0

func ChooseMove(request datatypes.GameRequest) (string, string) {
	var move *datatypes.Direction
	var graph *dijkstra2.Graph = dijkstra.GetDijkstraGraph(&request, canGoForTail(&request))


	const findFoodHealthThreshold = 90
	var foodMove = GoToFood(&request, graph)
	var tailMove = FollowTail(&request, graph)

	if foodMove != nil {
		move = foodMove
	} else if tailMove != nil {
		move = tailMove
	} else {
		fmt.Println("Falling back to AnyOtherMove")
		moveValue := AnyOtherMove(request)
		move = &moveValue
	}

	destCoord := datatypes.AddDirectionToCoord(request.You.Head, *move)
	if datatypes.IsFood(destCoord, request.Board) {
		turnsSinceEating = 0
	}
	turnsSinceEating += 1

	return datatypes.DirectionToStr(*move), "🤤"
}

func isLargestSnake(snake datatypes.Battlesnake, request datatypes.GameRequest) bool {
	for _, otherSnake := range request.Board.Snakes {
		if snake.ID != otherSnake.ID {
			if snake.Length <= otherSnake.Length {
				return false
			}
		}
	}
	return true
}

func canGoForTail(request *datatypes.GameRequest) bool {
	return turnsSinceEating > 1 && request.Turn > 4
}

func AnyOtherMove(request datatypes.GameRequest) datatypes.Direction {
	/*
	A simple and buggy algorithm to use if all else fails. Tries not to hit itself
	or run out of bounds.
	 */
	var board = request.Board
	var you = request.You

	availableMoves := datatypes.AllDirections

	availableMoves = borderCheck(you, board, availableMoves)

	availableMoves = stopHittingOtherSnakes(you, &request, availableMoves)
	if len(availableMoves) == 0 {
		// Ruh roh
		fmt.Println("AnyOtherMove: no available moves, picking 'Up'")
		return datatypes.DirectionUp
	} else {
		return availableMoves[rand.Intn(len(availableMoves))]
	}
}

func GoToFood(request *datatypes.GameRequest, graph *dijkstra2.Graph) *datatypes.Direction {

	head := request.You.Head
	var food datatypes.Coord
	var bestFoodDistance = math.MaxInt64
	var bestMove *datatypes.Direction
	for _, food = range request.Board.Food {
		move, distance := dijkstra.GetDijkstraPathDirection(head, food, graph)
		if move != nil {
			fmt.Printf("GoToFood: Path found to food %s\n", datatypes.CoordToString(food))
			if distance < bestFoodDistance {
				bestFoodDistance = distance
				bestMove = move
			}
		}
		fmt.Printf("GoToFood: Path not found to food %s\n", datatypes.CoordToString(food))
	}
	return bestMove
}

func FollowTail(request *datatypes.GameRequest, graph *dijkstra2.Graph) *datatypes.Direction {
	head := request.You.Head
	tail := request.You.Body[len(request.You.Body) - 1]
	move, _ := dijkstra.GetDijkstraPathDirection(head, tail, graph)
	return move
}

func borderCheck(you datatypes.Battlesnake, board datatypes.Board, availableMoves []datatypes.Direction) []datatypes.Direction {
	if datatypes.IsTopEdge(you.Head, board) {
		availableMoves = removeDirection(availableMoves, datatypes.DirectionUp)
	}
	if datatypes.IsRightEdge(you.Head, board) {
		availableMoves = removeDirection(availableMoves, datatypes.DirectionRight)
	}
	if datatypes.IsBottomEdge(you.Head, board) {
		availableMoves = removeDirection(availableMoves, datatypes.DirectionDown)
	}
	if datatypes.IsLeftEdge(you.Head, board) {
		availableMoves = removeDirection(availableMoves, datatypes.DirectionLeft)
	}
	return availableMoves
}

func removeDirection(s []datatypes.Direction, r datatypes.Direction) []datatypes.Direction {
	var resultDirs []datatypes.Direction
	for _, v := range s {
		if v != r {
			resultDirs = append(resultDirs, v)
		}
	}
	return resultDirs
}

func stopHittingOtherSnakes(you datatypes.Battlesnake, request *datatypes.GameRequest, avaliableDirections []datatypes.Direction) []datatypes.Direction{
	var excludedDirections []datatypes.Direction
	for _, dir := range avaliableDirections {
		nextCoord := datatypes.AddDirectionToCoord(you.Head, dir)
		for _, snake := range request.Board.Snakes {
			for _, bodyCoord := range snake.Body {
				if nextCoord == bodyCoord {
					excludedDirections = append(excludedDirections, dir)
					break
				}
			}
		}
	}
	for _, dir := range excludedDirections {
		avaliableDirections = removeDirection(avaliableDirections, dir)
	}
	return avaliableDirections
}
