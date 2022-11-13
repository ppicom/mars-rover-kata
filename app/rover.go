package app

import (
	"fmt"
)

type Rover struct {
	direction  direction
	coordinate coordinate
	grid       grid
}

func NewRover(grid grid) *Rover {
	return &Rover{
		direction: north,
		grid:      grid,
	}
}

func (r *Rover) Execute(commands string) (position string) {
	for i := 0; i < len(commands); i++ {
		command := fmt.Sprintf("%c", commands[i])
		if command == "R" {
			r.direction = r.right()
		} else if command == "L" {
			r.direction = r.left()
		} else {
			r.coordinate = r.grid.getNextCoordinateFor(r.coordinate, r.direction)
		}
	}

	return fmt.Sprintf("%d:%d:%s", r.x(), r.y(), r.direction)
}

func (r *Rover) right() direction {
	return r.direction.toTheRight()
}

func (r *Rover) left() direction {
	return r.direction.toTheLeft()
}

func (r Rover) x() int {
	return r.coordinate.x
}

func (r Rover) y() int {
	return r.coordinate.y
}
