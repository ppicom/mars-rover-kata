package app

import (
	"fmt"
)

type Rover struct {
	d direction
	y int
	x int
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
			if r.d == north {
				r.y++
			}

			if r.d == east {
				r.x++
			}

			if r.d == south {
				if r.y == 0 {
					r.y = 9
				} else {
					r.y--
				}
			}

			if r.d == west {
				if r.x == 0 {
					r.x = 9
				} else {
					r.x--
				}
			}
		}
	}

	return fmt.Sprintf("%d:%d:%s", r.x, r.y, r.d)
}

func (r *Rover) right() direction {
	return r.d.toTheRight()
}

func (r *Rover) left() direction {
	return r.d.toTheLeft()
}
