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
