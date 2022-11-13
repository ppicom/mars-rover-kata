package app

import (
	"fmt"
)

type Rover struct {
	d direction
	y int
}

func NewRover() *Rover {
	return &Rover{
		d: north,
	}
}

func (r *Rover) Execute(commands string) (position string) {
	for i := 0; i < len(commands); i++ {
		command := fmt.Sprintf("%c", commands[i])
		if command == "R" {
			r.d = r.right()
		} else if command == "L" {
			r.d = r.left()
		} else {
			r.y++
		}
	}

	return fmt.Sprintf("0:%d:%s", r.y, r.d)
}

func (r *Rover) right() direction {
	return r.d.toTheRight()
}

func (r *Rover) left() direction {
	return r.d.toTheLeft()
}
