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
		{commands: "RRRR ", position: "0:0:N"},
	}

	for _, tt := range tests {
		rover := NewRover()
		newPosition := rover.Execute(tt.commands)
		s.Equal(tt.position, newPosition)
	}
}

func Test(t *testing.T) {
	suite.Run(t, new(RoverShould))
}
