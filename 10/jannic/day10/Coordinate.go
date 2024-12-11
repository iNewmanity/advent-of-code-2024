package main

type Coordinate struct {
	X int
	Y int
}

func NewCoordinate(x, y int) *Coordinate {
	return &Coordinate{X: x, Y: y}
}

func (p Coordinate) Clone() *Coordinate {
	return &Coordinate{X: p.X, Y: p.Y}
}
