package datatypes

type Direction uint8

const (
	DirectionUp    Direction = iota
	DirectionRight Direction = iota
	DirectionDown  Direction = iota
	DirectionLeft  Direction = iota
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
