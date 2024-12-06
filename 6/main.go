package main

import (
	"advent-of-code-2024/6/day6"
	"advent-of-code-2024/util"
	"fmt"
)

func main() {
	var value [][]string = assignment1()
	assignment2(value)
}

func getTransitions() map[day6.State]day6.State {
	return map[day6.State]day6.State{
		day6.Up:    day6.Right,
		day6.Right: day6.Down,
		day6.Down:  day6.Left,
		day6.Left:  day6.Up,
	}
}

func getStartingStuff(values [][]string) (*day6.FSM, *day6.Coordinate) {
	var fsm *day6.FSM
	var coordinate *day6.Coordinate

	for i := range values {
		line := values[i]
		for j := range line {
			character := line[j]
			if character == "^" {
				fsm = day6.NewFSM(day6.Up, getTransitions())
				coordinate = day6.NewCoordinate(j, i)
			}
		}
	}

	return fsm, coordinate
}

func printStuff(values [][]string) {
	for i := range values {
		fmt.Println(values[i])
	}

	fmt.Println("----------------------------------")
}

func assignment2(values [][]string) {

	stateMachine, coordinate := getStartingStuff(values)
	values[coordinate.Y][coordinate.X] = "0"
	min := 0
	max := len(values[0]) - 1

	var coordinates []*day6.Coordinate
	fmt.Println(coordinates)

	i := 0
	for i < 1 {
		printStuff(values)
		state := stateMachine.CurrentState()

		if state == day6.Up {
			if coordinate.Y-1 < min {
				i = 1
				continue
			}

			if values[coordinate.Y-1][coordinate.X] == "#" {
				stateMachine.NextState()
				coordinates = append(coordinates, day6.NewCoordinate(coordinate.Y-1, coordinate.X))

				if len(coordinates) == 3 {
					x := coordinates[0].X
					y := coordinates[2].Y

					values[y][x] = "O"
				}
				continue
			}

			values[coordinate.Y][coordinate.X] = "0"
			coordinate.Y = coordinate.Y - 1

			continue
		}

		if state == day6.Right {
			if coordinate.X+1 > max {
				i = 1
				continue
			}

			if values[coordinate.Y][coordinate.X+1] == "#" {
				stateMachine.NextState()
				coordinates = append(coordinates, day6.NewCoordinate(coordinate.Y, coordinate.X+1))

				if len(coordinates) == 3 {
					x := coordinates[0].X
					y := coordinates[2].Y

					values[y][x] = "O"
				}
				continue
			}

			values[coordinate.Y][coordinate.X] = "0"
			coordinate.X = coordinate.X + 1

			continue
		}

		if state == day6.Down {
			if coordinate.Y+1 > max {
				i = 1
				continue
			}

			if values[coordinate.Y+1][coordinate.X] == "#" {
				stateMachine.NextState()
				coordinates = append(coordinates, day6.NewCoordinate(coordinate.Y, coordinate.X))

				if len(coordinates) == 3 {
					x := coordinates[0].X
					y := coordinates[2].Y

					values[y][x] = "O"
				}
				continue
			}

			values[coordinate.Y][coordinate.X] = "0"
			coordinate.Y = coordinate.Y + 1

			continue
		}

		if state == day6.Left {
			if coordinate.X-1 < min {
				i = 1
				continue
			}

			if values[coordinate.Y][coordinate.X-1] == "#" {
				stateMachine.NextState()
				coordinates = append(coordinates, day6.NewCoordinate(coordinate.Y, coordinate.X-1))

				if len(coordinates) == 3 {
					x := coordinates[0].X
					y := coordinates[2].Y

					values[y][x] = "O"

				}
				continue
			}

			values[coordinate.Y][coordinate.X] = "0"
			coordinate.X = coordinate.X - 1

			continue
		}
	}

	values[coordinate.Y][coordinate.X] = "0"
}

func assignment1() [][]string {
	path := "/home/janneumann/Dokumente/Daten/Projekte/Privat/advent-of-code-2024/6/input.txt"
	values := util.ReadFile(path, "")

	stateMachine, coordinate := getStartingStuff(values)

	min := 0
	max := len(values[0]) - 1

	i := 0
	for i < 1 {
		printStuff(values)
		state := stateMachine.CurrentState()

		if state == day6.Up {
			if coordinate.Y-1 < min {
				i = 1
				continue
			}

			if values[coordinate.Y-1][coordinate.X] == "#" {
				stateMachine.NextState()
				continue
			}

			values[coordinate.Y][coordinate.X] = "X"
			coordinate.Y = coordinate.Y - 1
			values[coordinate.Y][coordinate.X] = stateMachine.GetStateChar()

			continue
		}

		if state == day6.Right {
			if coordinate.X+1 > max {
				i = 1
				continue
			}

			if values[coordinate.Y][coordinate.X+1] == "#" {
				stateMachine.NextState()
				continue
			}

			values[coordinate.Y][coordinate.X] = "X"
			coordinate.X = coordinate.X + 1
			values[coordinate.Y][coordinate.X] = stateMachine.GetStateChar()

			continue
		}

		if state == day6.Down {
			if coordinate.Y+1 > max {
				i = 1
				continue
			}

			if values[coordinate.Y+1][coordinate.X] == "#" {
				stateMachine.NextState()
				continue
			}

			values[coordinate.Y][coordinate.X] = "X"
			coordinate.Y = coordinate.Y + 1
			values[coordinate.Y][coordinate.X] = stateMachine.GetStateChar()

			continue
		}

		if state == day6.Left {
			if coordinate.X-1 < min {
				i = 1
				continue
			}

			if values[coordinate.Y][coordinate.X-1] == "#" {
				stateMachine.NextState()
				continue
			}

			values[coordinate.Y][coordinate.X] = "X"
			coordinate.X = coordinate.X - 1
			values[coordinate.Y][coordinate.X] = stateMachine.GetStateChar()

			continue
		}
	}

	values[coordinate.Y][coordinate.X] = "X"

	count := 0
	for j := range values {
		line := values[j]
		for i := range line {
			character := line[i]
			if character == "X" {
				count++
			}
		}

	}

	fmt.Println("Assignment 1: ", count)
	return values
}
