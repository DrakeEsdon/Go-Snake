package snake

import (
	"github.com/DrakeEsdon/Go-Snake/datatypes"
	"math/rand"
)

func isTopEdge(coord datatypes.Coord, board datatypes.Board) bool {
	return coord.Y == board.Height - 1
}

func isRightEdge(coord datatypes.Coord, board datatypes.Board) bool {
	return coord.X == board.Width - 1
}


func isBottomEdge(coord datatypes.Coord, board datatypes.Board) bool {
	_ = board
	return coord.Y == 0
}

func isLeftEdge(coord datatypes.Coord, board datatypes.Board) bool {
	_ = board
	return coord.X == 0
}

func ChooseMove(g datatypes.GameRequest) string {
	var gameState = g.Board
	var you = g.You

	availableMoves := datatypes.AllDirections

	availableMoves = borderCheck(you, gameState, availableMoves)

	availableMoves = stopHittingYourself(you, availableMoves)

	move := availableMoves[rand.Intn(len(availableMoves))]

	return datatypes.DirectionToStr(move)
}

func borderCheck(you datatypes.Battlesnake, board datatypes.Board, availableMoves []datatypes.Direction) []datatypes.Direction {
	if isTopEdge(you.Head, board) {
		availableMoves = removeDirection(availableMoves, datatypes.DirectionUp)
	}
	if isRightEdge(you.Head, board) {
		availableMoves = removeDirection(availableMoves, datatypes.DirectionRight)
	}
	if isBottomEdge(you.Head, board) {
		availableMoves = removeDirection(availableMoves, datatypes.DirectionDown)
	}
	if isLeftEdge(you.Head, board) {
		availableMoves = removeDirection(availableMoves, datatypes.DirectionLeft)
	}
	return availableMoves
}

func removeDirection(s []datatypes.Direction, r datatypes.Direction) []datatypes.Direction {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func stopHittingYourself(you datatypes.Battlesnake, avaliableDirections []datatypes.Direction) []datatypes.Direction{
	if you.Body[0].X > you.Body[1].X{
		avaliableDirections = removeDirection(avaliableDirections, datatypes.DirectionLeft)
	} else if you.Body[0].X < you.Body[1].X{
		avaliableDirections = removeDirection(avaliableDirections, datatypes.DirectionRight)
	} else if you.Body[0].Y > you.Body[1].Y{
		avaliableDirections = removeDirection(avaliableDirections, datatypes.DirectionDown)
	} else if you.Body[0].Y < you.Body[1].Y{
		avaliableDirections = removeDirection(avaliableDirections, datatypes.DirectionUp)
	}
	return avaliableDirections
}
