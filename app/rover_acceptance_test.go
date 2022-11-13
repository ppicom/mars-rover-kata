package app

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type RoverShould struct {
	suite.Suite
}

func (RoverShould) SetupTest() {
}

func (s *RoverShould) TestTurnRight() {

	tests := []struct {
		commands string
		position string
	}{
		{commands: "R", position: "0:0:E"},
		{commands: "RR", position: "0:0:S"},
		{commands: "RRR", position: "0:0:W"},
		{commands: "RRRR", position: "0:0:N"},
	}

	for _, tt := range tests {
		rover := NewRover(*newGrid(10, 10))
		newPosition := rover.Execute(tt.commands)
		s.Equal(tt.position, newPosition)
	}
}

func (s *RoverShould) TestTurnLeft() {

	tests := []struct {
		commands string
		position string
	}{
		{commands: "L", position: "0:0:W"},
		{commands: "LL", position: "0:0:S"},
		{commands: "LLL", position: "0:0:E"},
		{commands: "LLLL", position: "0:0:N"},
	}

	for _, tt := range tests {
		rover := NewRover(*newGrid(10, 10))
		newPosition := rover.Execute(tt.commands)
		s.Equal(tt.position, newPosition)
	}
}

func (s *RoverShould) TestMoveNorth() {

	tests := []struct {
		commands string
		position string
	}{
		{commands: "M", position: "0:1:N"},
		{commands: "MM", position: "0:2:N"},
		{commands: "MMMMM", position: "0:5:N"},
	}

	for _, tt := range tests {
		rover := NewRover(*newGrid(10, 10))
		newPosition := rover.Execute(tt.commands)
		s.Equal(tt.position, newPosition)
	}
}

func (s *RoverShould) TestMoveEast() {

	tests := []struct {
		commands string
		position string
	}{
		{commands: "RM", position: "1:0:E"},
		{commands: "RMMMM", position: "4:0:E"},
	}

	for _, tt := range tests {
		rover := NewRover(*newGrid(10, 10))
		newPosition := rover.Execute(tt.commands)
		s.Equal(tt.position, newPosition)
	}
}

func (s *RoverShould) TestMoveSouth() {

	tests := []struct {
		commands string
		position string
	}{
		{commands: "RRM", position: "0:9:S"},
		{commands: "RRMMMM", position: "0:6:S"},
		{commands: "RRMMMMMM", position: "0:4:S"},
	}

	for _, tt := range tests {
		rover := NewRover(*newGrid(10, 10))
		newPosition := rover.Execute(tt.commands)
		s.Equal(tt.position, newPosition)
	}
}

func (s *RoverShould) TestMoveWest() {

	tests := []struct {
		commands string
		position string
	}{
		{commands: "LM", position: "9:0:W"},
		{commands: "LMMMMM", position: "5:0:W"},
	}

	for _, tt := range tests {
		rover := NewRover(*newGrid(10, 10))
		newPosition := rover.Execute(tt.commands)
		s.Equal(tt.position, newPosition)
	}
}

func (s *RoverShould) TestWrapFromTopToBottomWhenMovinNorth() {

	tests := []struct {
		commands string
		position string
	}{
		{commands: "MMMMMMMMMM", position: "0:0:N"},
		{commands: "MMMMMMMMMMMMMMM", position: "0:5:N"},
	}

	for _, tt := range tests {
		rover := NewRover(*newGrid(10, 10))
		newPosition := rover.Execute(tt.commands)
		s.Equal(tt.position, newPosition)
	}
}

func Test(t *testing.T) {
	suite.Run(t, new(RoverShould))
}
