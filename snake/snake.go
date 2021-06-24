package snake

import (
	"fmt"
	"github.com/DrakeEsdon/Go-Snake/datatypes"
	"github.com/DrakeEsdon/Go-Snake/dijkstra"
	"math/rand"
)

func ChooseMove(request datatypes.GameRequest) (string, string) {
	var move *datatypes.Direction

	if request.You.Health > 50 {
		fmt.Println("Health > 50, following tail")
		move = FollowTail(&request)
	} else {
		fmt.Println("Health < 50, going for food")
		move = GoToFood(&request)
	}

	if move == nil {
		fmt.Println("Move was nil, doing any other move")
		moveValue := AnyOtherMove(request)
		move = &moveValue
	}

	return datatypes.DirectionToStr(*move), "🤤"
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
	graph := dijkstra.GetDijkstraGraph(request)

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
	graph := dijkstra.GetDijkstraGraph(request)
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
