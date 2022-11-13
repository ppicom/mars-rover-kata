package app

type coordinate struct {
	x int
	y int
}

func newCoordinate(x, y int) *coordinate {
	return &coordinate{
		x: x,
		y: y,
	}
}

func (c coordinate) equals(other coordinate) bool {
	return c.x == other.x && c.y == other.y
}
