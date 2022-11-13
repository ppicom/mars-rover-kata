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
			switch r.d {
			case north:
				r.d = east
			case east:
				r.d = south
			case south:
				r.d = west
			default:
				r.d = north
			}
		}
	}

	return fmt.Sprintf("0:0:%s", r.d)
}
