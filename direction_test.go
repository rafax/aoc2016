package aoc2016

import "testing"

var (
	right = [][2]Direction{[2]Direction{U, R}, [2]Direction{R, D}, [2]Direction{D, L}, [2]Direction{L, U}}
	left  = [][2]Direction{[2]Direction{U, L}, [2]Direction{L, D}, [2]Direction{D, R}, [2]Direction{R, U}}
)

func TestTurnRight(t *testing.T) {
	for _, turn := range right {
		after := turn[0].TurnRight()
		if after != turn[1] {
			t.Errorf("%v != %v", after, turn[1])
		}
	}
}

func TestTurnLeft(t *testing.T) {
	for _, turn := range left {
		after := turn[0].TurnLeft()
		if after != turn[1] {
			t.Errorf("%v != %v", after, turn[1])
		}
	}
}
