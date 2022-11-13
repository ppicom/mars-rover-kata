package app

import (
	"fmt"
)

type Rover struct {
	d direction
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
		} else {
			r.d = r.left()
		}
	}

	return fmt.Sprintf("0:0:%s", r.d)
}

func (r *Rover) right() direction {
	return r.d.toTheRight()
}

func (r *Rover) left() direction {
	return r.d.toTheLeft()
}
