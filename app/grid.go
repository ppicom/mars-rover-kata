package app

type grid struct {
	width     int
	height    int
	obstacles []coordinate
}

const (
	MAX_HEIGHT = 10
	MAX_WIDTH  = 10
)

func newGrid(obstacles []coordinate) *grid {
	return &grid{
		width:     MAX_WIDTH,
		height:    MAX_HEIGHT,
		obstacles: obstacles,
	}
}

func (g grid) getNextCoordinateFor(coordinate coordinate, direction direction) *coordinate {
	nextX := coordinate.x
	nextY := coordinate.y

	if direction == north {
		nextY = (coordinate.y + 1) % g.height
	}

	if direction == east {
		nextX = (coordinate.x + 1) % g.width
	}

	if direction == south {
		if coordinate.y == 0 {
			nextY = g.height - 1
		} else {
			nextY--
		}
	}

	if direction == west {
		if coordinate.x == 0 {
			nextX = g.height - 1
		} else {
			nextX--
		}
	}

	nextCoodinate := newCoordinate(nextX, nextY)
	for _, obstacle := range g.obstacles {
		if obstacle.equals(*nextCoodinate) {
			return nil
		}
	}

	return nextCoodinate
}
