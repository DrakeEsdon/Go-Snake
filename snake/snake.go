package snake

import (
	"fmt"
	"github.com/DrakeEsdon/Go-Snake/datatypes"
	"github.com/DrakeEsdon/Go-Snake/dijkstra"
	"math/rand"
)

var turnsSinceEating = 0

func ChooseMove(request datatypes.GameRequest) (string, string) {
	var move *datatypes.Direction

	const findFoodHealthThreshold = 50

	if request.Turn > 5 {
		if request.You.Health > findFoodHealthThreshold {
			if isLargestSnake(request.You, request) {
				fmt.Printf("Health > %v and largest snake, following tail\n", findFoodHealthThreshold)
				move = FollowTail(&request)
			} else {
				fmt.Printf("Not largest snake, going for food\n")
				move = GoToFood(&request)
			}
		} else {
			fmt.Printf("Health < %v, going for food\n", findFoodHealthThreshold)
			move = GoToFood(&request)
		}
	}

	if move == nil {
		fmt.Println("Move was nil, doing any other move")
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

func canGoForTail() bool {
	return turnsSinceEating > 1
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

	availableMoves = stopHittingYourself(you, availableMoves)
	if len(availableMoves) == 0 {
		// Ruh roh
		fmt.Println("AnyOtherMove: no available moves, picking 'Up'")
		return datatypes.DirectionUp
	} else {
		return availableMoves[rand.Intn(len(availableMoves))]
	}
}

func GoToFood(request *datatypes.GameRequest) *datatypes.Direction {
	graph := dijkstra.GetDijkstraGraph(request, canGoForTail())

	head := request.You.Head
	var food datatypes.Coord
	var move *datatypes.Direction
	for _, food = range request.Board.Food {
		move = dijkstra.GetDijkstraPathDirection(head, food, graph)
		if move != nil {
			fmt.Printf("GoToFood: Path found to food %s\n", datatypes.CoordToString(food))
			break
		}
		fmt.Printf("GoToFood: Path not found to food %s\n", datatypes.CoordToString(food))
	}
	return move
}

func FollowTail(request *datatypes.GameRequest) *datatypes.Direction {
	graph := dijkstra.GetDijkstraGraph(request, canGoForTail())
	head := request.You.Head
	tail := request.You.Body[len(request.You.Body) - 1]
	return dijkstra.GetDijkstraPathDirection(head, tail, graph)
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
	for i, v := range s {
		if v != r {
			resultDirs = append(resultDirs, s[i])
		}
	}
	return resultDirs
}

func stopHittingYourself(you datatypes.Battlesnake, avaliableDirections []datatypes.Direction) []datatypes.Direction{
	var excludedDirections []datatypes.Direction
	for _, dir := range avaliableDirections {
		nextCoord := datatypes.AddDirectionToCoord(you.Head, dir)
		for _, bodyCoord := range you.Body {
			if nextCoord == bodyCoord {
				excludedDirections = append(excludedDirections, dir)
				break
			}
		}
	}
	for _, dir := range excludedDirections {
		avaliableDirections = removeDirection(avaliableDirections, dir)
	}
	return avaliableDirections
}
