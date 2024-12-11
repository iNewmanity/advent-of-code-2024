package main

type Agent struct {
	Coordinate   *Coordinate
	CurrentValue int
	CanMove      bool
}

func NewAgent(coordinate *Coordinate, currentValue int) *Agent {
	return &Agent{Coordinate: coordinate, CurrentValue: currentValue, CanMove: true}
}
