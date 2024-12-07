package main

import (
	"advent-of-code-2024/6/day6"
	"advent-of-code-2024/util"
	"fmt"
)

type position struct {
	x int
	y int
}

func main() {
	//assignment1()
	assignment2()
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

func assignment2() {
	path := "/home/janneumann/Dokumente/Daten/Projekte/Privat/advent-of-code-2024/6/input.txt"
	values := util.ReadFile(path, "")

	var hashTags int = countHashTags(values)
	var oneFourthHashTags int = deductOneFourth(hashTags)
	var hashTagsMinusTwo int = deductTwo(hashTags)
	var anotherSolution int = hashtagsDividedBy3By2(hashTags)
	var anotherAnotherSolution int = hashtagsDividedBy4By2(hashTags)
	var anotherAnotherAnotherSolution int = customAlgorithm(hashTags)
	fmt.Println(oneFourthHashTags)
	fmt.Println(hashTagsMinusTwo)
	fmt.Println(anotherSolution)
	fmt.Println(anotherAnotherSolution)
	fmt.Println(anotherAnotherAnotherSolution)
}

func hashtagsDividedBy3By2(val int) int {
	return (int(val/3) + 1) * 2
}

func hashtagsDividedBy4By2(val int) int {
	return (int(val/4) + 1) * 2
}

func deductTwo(val int) int {
	return val - 2
}

func countHashTags(values [][]string) int {
	count := 0
	for i := range values {
		for j := range values[i] {
			if values[i][j] == "#" {
				count++
			}
		}
	}
	return count
}

func customAlgorithm(val int) int {
	return val - (val / 4)
}

func deductOneFourth(val int) int {
	var oneForth float64 = 0.75
	var result float64 = float64(val) * oneForth
	return int(result)
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
