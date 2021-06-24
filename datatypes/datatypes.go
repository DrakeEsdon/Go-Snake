package datatypes

import (
	"fmt"
)

type Game struct {
	ID      string `json:"id"`
	Timeout int32  `json:"timeout"`
}

type Coord struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Battlesnake struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Health int32   `json:"health"`
	Body   []Coord `json:"body"`
	Head   Coord   `json:"head"`
	Length int32   `json:"length"`
	Shout  string  `json:"shout"`
}

type Board struct {
	Height  int           `json:"height"`
	Width   int           `json:"width"`
	Food    []Coord       `json:"food"`
	Snakes  []Battlesnake `json:"snakes"`
	Hazards []Coord       `json:"hazards"`
}

type BattlesnakeInfoResponse struct {
	APIVersion string `json:"apiversion"`
	Author     string `json:"author"`
	Color      string `json:"color"`
	Head       string `json:"head"`
	Tail       string `json:"tail"`
}

type GameRequest struct {
	Game  Game        `json:"game"`
	Turn  int         `json:"turn"`
	Board Board       `json:"board"`
	You   Battlesnake `json:"you"`
}

type MoveResponse struct {
	Move  string `json:"move"`
	Shout string `json:"shout,omitempty"`
}

func IsSnakeOrHazard(coord Coord, board Board) bool {
	for _, snake := range board.Snakes {
		for _, snakeCoord := range snake.Body {
			if coord == snakeCoord {
				return true
			}
		}
	}
	if board.Hazards != nil {
		for _, hazardCoord := range board.Hazards {
			if coord == hazardCoord {
				return true
			}
		}
	}
	return false
}

func IsFood(coord Coord, board Board) bool {
	for _, food := range board.Food {
		if coord == food {
			return true
		}
	}
	return false
}

func IsOutOfBounds(coord Coord, board Board) bool {
	x := coord.X
	y := coord.Y
	return x < 0 || x >= board.Width || y < 0 || y >= board.Height
}

func IsMyTail(coord Coord, myself Battlesnake) bool {
	return coord == myself.Body[len(myself.Body) - 1]
}

func IsTopEdge(coord Coord, board Board) bool {
	return coord.Y == board.Height-1
}

func IsRightEdge(coord Coord, board Board) bool {
	return coord.X == board.Width-1
}

func IsBottomEdge(coord Coord, board Board) bool {
	_ = board
	return coord.Y == 0
}

func IsLeftEdge(coord Coord, board Board) bool {
	_ = board
	return coord.X == 0
}

func CoordToString(coord Coord) string {
	return fmt.Sprintf("(%d,%d)", coord.X, coord.Y)
}

func CoordFromString(string string) Coord {
	var x, y int
	_, err := fmt.Sscanf(string, "(%d,%d)", &x, &y)
	if err != nil {
		return Coord{}
	}
	return Coord{x, y}
}