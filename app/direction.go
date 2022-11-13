package app

type direction string

const (
	north direction = "N"
	east  direction = "E"
	south direction = "S"
	west  direction = "W"
)

func (d direction) toTheLeft() direction {
	switch d {
	case north:
		return west
	case east:
		return north
	case south:
		return east
	default:
		return south
	}
}

func (d direction) toTheRight() direction {
	switch d {
	case north:
		return east
	case east:
		return south
	case south:
		return west
	default:
		return north
	}
}
