package app

type grid struct {
	width  int
	height int
}

func newGrid(width, height int) *grid {
	return &grid{
		width:  width,
		height: height,
	}
}

func (g grid) getNextCoordinateFor(coordinate coordinate, direction direction) coordinate {
	nextX := coordinate.x
	nextY := coordinate.y

	if direction == north {
		nextY = (coordinate.y + 1) % g.height
	}

	if direction == east {
		nextX++
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

	return *newCoordinate(nextX, nextY)
}
