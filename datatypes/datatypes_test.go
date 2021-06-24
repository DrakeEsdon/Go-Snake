package datatypes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func getTestBoard() Board {
	return Board{
		Height:  10,
		Width:   10,
		Food:    []Coord{{0, 0}},
		Snakes:  []Battlesnake{
			{
				Body: []Coord{ // facing left at top right
					{7, 9},
					{8, 9},
					{9, 9},
				},
			},
			{
				Body: []Coord{ // facing left at top right
					{7, 8},
					{8, 8},
					{9, 8},
				},
			},
		},
		Hazards: []Coord{{0, 1}},
	}
}

func TestIsOutOfBounds(t *testing.T) {
	b := getTestBoard()

	for x := 0; x < b.Width; x++ {
		for y := 0; y < b.Height; y++ {
			assert.False(t, IsOutOfBounds(Coord{x, y}, b))
		}
	}
	assert.True(t, IsOutOfBounds(Coord{-1, 0}, b))
	assert.True(t, IsOutOfBounds(Coord{-1, -1}, b))
	assert.True(t, IsOutOfBounds(Coord{0, -1}, b))
	assert.True(t, IsOutOfBounds(Coord{9, -1}, b))
	assert.True(t, IsOutOfBounds(Coord{10, -1}, b))
	assert.True(t, IsOutOfBounds(Coord{10, 0}, b))
	assert.True(t, IsOutOfBounds(Coord{10, 9}, b))
	assert.True(t, IsOutOfBounds(Coord{10, 10}, b))
	assert.True(t, IsOutOfBounds(Coord{9, 10}, b))
	assert.True(t, IsOutOfBounds(Coord{0, 10}, b))
	assert.True(t, IsOutOfBounds(Coord{-1, 10}, b))
	assert.True(t, IsOutOfBounds(Coord{-1, 9}, b))
}

func TestIsSnakeOrHazard(t *testing.T) {
	b := getTestBoard()

	assert.False(t, IsSnakeOrHazard(Coord{1, 1}, b))
	assert.True(t, IsSnakeOrHazard(Coord{0, 1}, b))
	assert.False(t, IsSnakeOrHazard(Coord{0, 0}, b)) // Food is ok
	assert.True(t, IsSnakeOrHazard(Coord{7, 9}, b))
}

func TestCoordToString(t *testing.T) {
	coordA := Coord{0, 0}
	assert.Equal(t, "(0,0)", CoordToString(coordA))
	coordB := Coord{1, 10}
	assert.Equal(t, "(1,10)", CoordToString(coordB))
}

func TestCoordFromString(t *testing.T) {
	input := "(1,10)"
	assert.Equal(t, Coord{1, 10}, CoordFromString(input))
}