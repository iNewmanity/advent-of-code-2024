package main

import (
	"advent-of-code-2024/util"
	"strconv"
)

type ValueHolder struct {
	Fields      [][]int
	XMax        int
	yMax        int
	StartAgents []*StartAgent
}

func main() {
	path := "/home/janneumann/Dokumente/Daten/Projekte/Privat/advent-of-code-2024/10/input.txt"
	valuesString, _ := util.ReadFile(path, "", true)
	values := [][]int{}
	for i := range valuesString {
		oneD := []int{}
		for i2 := range valuesString[i] {
			value, _ := strconv.Atoi(valuesString[i][i2])
			oneD = append(oneD, value)
		}
		values = append(values, oneD)
	}
	yMax := len(values) - 1
	xMax := len(values[0]) - 1

	valueHolder := ValueHolder{values, xMax, yMax, make([]*StartAgent, 0)}

	DoTheThing(&valueHolder)
}

func DoTheThing(valueHolder *ValueHolder) {
	populateStartAgents(valueHolder)
	moveAgents(valueHolder)
	println(calculateSumForPartTwo(valueHolder))
}

func calculateSum(valueHolder *ValueHolder) int {
	sum := 0
	for _, startAgent := range valueHolder.StartAgents {
		sum += startAgent.calcSum()
	}
	return sum
}

func calculateSumForPartTwo(valueHolder *ValueHolder) int {
	sum := 0
	for _, startAgent := range valueHolder.StartAgents {
		sum += startAgent.calcSumSecondPart()
	}
	return sum
}

func moveAgents(valueHolder *ValueHolder) {
	x := 0
	for x < 1 {

		if !valueHolder.CanAgentsMove() {
			x = 42
		}

		for _, startAgent := range valueHolder.StartAgents {
			for i := len(startAgent.Agents) - 1; i >= 0; i-- {
				currentAgent := startAgent.Agents[i]

				if currentAgent.CanMove == false {
					continue
				}

				newCoordinates := getNewCoordinates(valueHolder, currentAgent)

				if len(newCoordinates) == 0 {
					currentAgent.CanMove = false
					continue
				}

				if len(newCoordinates) == 1 {
					currentAgent.Coordinate = newCoordinates[0]
					currentAgent.CurrentValue = getValueForCoordinate(valueHolder, newCoordinates[0])
					continue
				}

				if len(newCoordinates) > 1 {
					startAgent.Agents = append(startAgent.Agents[:i], startAgent.Agents[i+1:]...)

					for j := range newCoordinates {
						startAgent.Agents = append(startAgent.Agents, NewAgent(newCoordinates[j], currentAgent.CurrentValue+1))
					}
				}
			}
		}
	}
}

func (valueHolder *ValueHolder) CanAgentsMove() bool {
	for _, startAgent := range valueHolder.StartAgents {

		for i := range startAgent.Agents {
			if startAgent.Agents[i].CanMove == true {
				return true
			}
		}
	}

	return false
}

func getNewCoordinates(valueHolder *ValueHolder, currentAgent *Agent) []*Coordinate {
	var newCoordinates []*Coordinate

	leftCoordinate := getLeftCoordinate(currentAgent)

	if isNewCoordinateValid(valueHolder, currentAgent, leftCoordinate) {
		newCoordinates = append(newCoordinates, leftCoordinate)
	}

	rightCoordinate := getRightCoordinate(valueHolder, currentAgent)

	if isNewCoordinateValid(valueHolder, currentAgent, rightCoordinate) {
		newCoordinates = append(newCoordinates, rightCoordinate)
	}

	topCoordinate := getTopCoordinate(currentAgent)

	if isNewCoordinateValid(valueHolder, currentAgent, topCoordinate) {
		newCoordinates = append(newCoordinates, topCoordinate)
	}

	bottomCoordinate := getBottomCoordinate(valueHolder, currentAgent)

	if isNewCoordinateValid(valueHolder, currentAgent, bottomCoordinate) {
		newCoordinates = append(newCoordinates, bottomCoordinate)
	}

	return newCoordinates
}

func isNewCoordinateValid(valueHolder *ValueHolder, agent *Agent, newCoordinate *Coordinate) bool {
	return newCoordinate.X != -1 && getValueForCoordinate(valueHolder, newCoordinate) == agent.CurrentValue+1
}

func getLeftCoordinate(agent *Agent) *Coordinate {
	if agent.Coordinate.X-1 < 0 {
		return NewCoordinate(-1, -1)
	}

	newCoordinate := agent.Coordinate.Clone()
	newCoordinate.X = newCoordinate.X - 1
	return newCoordinate
}

func getRightCoordinate(valueHolder *ValueHolder, agent *Agent) *Coordinate {
	if agent.Coordinate.X+1 > valueHolder.XMax {
		return NewCoordinate(-1, -1)
	}

	newCoordinate := agent.Coordinate.Clone()
	newCoordinate.X = newCoordinate.X + 1
	return newCoordinate
}

func getBottomCoordinate(valueHolder *ValueHolder, agent *Agent) *Coordinate {
	if agent.Coordinate.Y+1 > valueHolder.yMax {
		return NewCoordinate(-1, -1)
	}

	newCoordinate := agent.Coordinate.Clone()
	newCoordinate.Y = newCoordinate.Y + 1
	return newCoordinate
}

func getTopCoordinate(agent *Agent) *Coordinate {
	if agent.Coordinate.Y-1 < 0 {
		return NewCoordinate(-1, -1)
	}

	newCoordinate := agent.Coordinate.Clone()
	newCoordinate.Y = newCoordinate.Y - 1
	return newCoordinate
}

func getValueForCoordinate(valueHolder *ValueHolder, coordinate *Coordinate) int {
	return valueHolder.Fields[coordinate.Y][coordinate.X]
}

func populateStartAgents(valueHolder *ValueHolder) {
	for i := range valueHolder.Fields {
		column := valueHolder.Fields[i]
		for j := range column {
			if column[j] == 0 {
				coordinate := NewCoordinate(j, i)
				newAgent := NewAgent(coordinate, 0)
				startAgent := NewStartAgent(make([]*Agent, 0))
				startAgent.Agents = append(startAgent.Agents, newAgent)
				valueHolder.StartAgents = append(valueHolder.StartAgents, startAgent)
			}
		}
	}
}
