package datatypes

type Vector Coord

type Direction Vector

var (
	DirectionUp    = Direction{X:  0, Y:  1}
	DirectionRight = Direction{X:  1, Y:  0}
	DirectionDown  = Direction{X:  0, Y: -1}
	DirectionLeft  = Direction{X: -1, Y:  0}
)

var AllDirections = []Direction{DirectionUp, DirectionRight, DirectionDown, DirectionLeft}

func DirectionToStr(d Direction) string {
	if d == DirectionUp {
		return "up"
	}
	if d == DirectionRight {
		return "right"
	}
	if d == DirectionDown {
		return "down"
	}
	if d == DirectionLeft {
		return "left"
	}
	return "up"
}

func AddDirectionToCoord(coord Coord, dir Direction) Coord {
	return Coord{
		X: coord.X + dir.X,
		Y: coord.Y + dir.Y,
	}
}