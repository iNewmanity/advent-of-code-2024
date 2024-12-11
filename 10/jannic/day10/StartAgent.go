package main

type StartAgent struct {
	Agents []*Agent
}

func (startAgent *StartAgent) calcSum() int {
	var correctCoordinates []*Coordinate
	for i := range startAgent.Agents {
		currentAgent := startAgent.Agents[i]
		if currentAgent.CurrentValue == 9 {

			isDuplicated := false
			for j := range correctCoordinates {
				currentCoordinate := correctCoordinates[j]

				if currentAgent.Coordinate.X == currentCoordinate.X && currentAgent.Coordinate.Y == currentCoordinate.Y {
					isDuplicated = true
					break
				}
			}

			if !isDuplicated {
				correctCoordinates = append(correctCoordinates, currentAgent.Coordinate)
			}
		}
	}

	return len(correctCoordinates)
}

func (startAgent *StartAgent) calcSumSecondPart() int {
	sum := 0
	for i := range startAgent.Agents {
		currentAgent := startAgent.Agents[i]
		if currentAgent.CurrentValue == 9 {
			sum++
		}
	}

	return sum
}

func NewStartAgent(agents []*Agent) *StartAgent {
	return &StartAgent{Agents: agents}
}
