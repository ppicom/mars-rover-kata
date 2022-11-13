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
	var obstacleString string
	for i := 0; i < len(commands); i++ {
		command := fmt.Sprintf("%c", commands[i])
		if command == "R" {
			r.direction = r.right()
		} else if command == "L" {
			r.direction = r.left()
		} else {
			obstacleString = r.move()
		}
	}

	return fmt.Sprintf("%s%d:%d:%s", obstacleString, r.x(), r.y(), r.direction)
}

func (r *Rover) move() string {
	nextCoordinate := r.grid.getNextCoordinateFor(r.coordinate, r.direction)
	if nextCoordinate != nil {
		r.coordinate = *nextCoordinate
		return ""
	}

	return "O:"
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
