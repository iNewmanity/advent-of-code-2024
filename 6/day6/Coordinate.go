package day6

import "fmt"

type Coordinate struct {
	X int
	Y int
}

func NewCoordinate(x, y int) *Coordinate {
	return &Coordinate{X: x, Y: y}
}

func (p Coordinate) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}
