package day10

import (
	"fmt"
	"strconv"
)

type trailhead struct {
	start  coordinate
	trails []trail
	score  int
}

type trail struct {
	waypoints []coordinate
}

type coordinate struct {
	value string
	x     int
	y     int
}

func ConvertStringToCoordinates(s [][]string) [][]coordinate {
	topographicMap := [][]coordinate{}
	for i := range s {
		coordinates := []coordinate{}
		for i2 := range s[i] {
			coordinates = append(coordinates, coordinate{
				value: s[i][i2],
				x:     i2,
				y:     i,
			})
		}
		topographicMap = append(topographicMap, coordinates)
	}
	return topographicMap
}

func FindTrailHeads(topographicMap [][]coordinate) []trailhead {
	tHeads := []trailhead{}
	for i := range topographicMap {
		for i2 := range topographicMap[i] {
			if topographicMap[i][i2].value == "0" {
				tHeads = append(tHeads, trailhead{
					start: topographicMap[i][i2],
				})
			}
		}
	}
	return tHeads
}

func FindAllWayPoints(topographicMap [][]coordinate, trailheads []trailhead) []trailhead {
	for i := range trailheads {
		lastCoordinate := []coordinate{trailheads[i].start}
		newTrail := trail{}
		newTrail.waypoints = append(newTrail.waypoints, lastCoordinate...)

		for i2 := range topographicMap {
			for i3 := range topographicMap[i2] {
				fmt.Println(i, i2, i3)
				nextCoordinates := []coordinate{}
				for i4 := range lastCoordinate {
					fmt.Println(lastCoordinate[i4])
					calcNextCoordinates := findNextCoordinates(topographicMap, lastCoordinate[i4])
					fmt.Println(calcNextCoordinates)
					nextCoordinates = append(nextCoordinates, calcNextCoordinates...)
				}
				lastCoordinate = nextCoordinates
				newTrail.waypoints = append(newTrail.waypoints, nextCoordinates...)
			}
		}
		newTrail.waypoints = removeDuplicates(newTrail.waypoints)
		trailheads[i].trails = append(trailheads[i].trails, newTrail)
	}

	return trailheads
}

func PrintTrailHeads(trailheads []trailhead) {
	for i := range trailheads {
		fmt.Println("\n--------Trailhead", i, "--------")
		fmt.Println("Trails:\t")
		for i2 := range trailheads[i].trails {
			trail := trailheads[i].trails[i2]
			fmt.Println("\n++++++Trail", i2, "++++++")
			for i3 := range trail.waypoints {
				fmt.Println("Value:\t", trail.waypoints[i3].value, "\tX:\t", trail.waypoints[i3].x, "\tY:\t", trail.waypoints[i3].y)
			}
		}
	}
}

func removeDuplicates(waypoints []coordinate) []coordinate {
	uniqueMap := make(map[coordinate]bool)
	uniqueSlice := []coordinate{}

	for _, coord := range waypoints {
		if !uniqueMap[coord] {
			uniqueMap[coord] = true
			uniqueSlice = append(uniqueSlice, coord)
		}
	}

	return uniqueSlice
}

func findNextCoordinates(topographicMap [][]coordinate, c coordinate) []coordinate {
	value, _ := strconv.Atoi(c.value)
	nextCoordinates := []coordinate{}
	if isNotAtBoundary(topographicMap, c) {
		left, _ := strconv.Atoi(topographicMap[c.y][c.x-1].value)
		right, _ := strconv.Atoi(topographicMap[c.y][c.x+1].value)
		down, _ := strconv.Atoi(topographicMap[c.y-1][c.x].value)
		up, _ := strconv.Atoi(topographicMap[c.y+1][c.x].value)

		if left-value == 1 {
			nextCoordinates = append(nextCoordinates, topographicMap[c.y][c.x-1])
		}
		if right-value == 1 {
			nextCoordinates = append(nextCoordinates, topographicMap[c.x+1][c.y])
		}
		if down-value == 1 {
			nextCoordinates = append(nextCoordinates, topographicMap[c.x][c.y-1])
		}
		if up-value == 1 {
			nextCoordinates = append(nextCoordinates, topographicMap[c.x][c.y+1])
		}
		return nextCoordinates
	} else {
		if isAtDownBoundary(topographicMap, c) {
			if isAtLeftBoundary(topographicMap, c) {
				right, _ := strconv.Atoi(topographicMap[c.x+1][c.y].value)
				up, _ := strconv.Atoi(topographicMap[c.x][c.y+1].value)
				if up-value == 1 {
					nextCoordinates = append(nextCoordinates, topographicMap[c.x][c.y+1])
				}
				if right-value == 1 {
					nextCoordinates = append(nextCoordinates, topographicMap[c.x+1][c.y])
				}
				return nextCoordinates
			} else if isAtRightBoundary(topographicMap, c) {
				left, _ := strconv.Atoi(topographicMap[c.x-1][c.y].value)
				up, _ := strconv.Atoi(topographicMap[c.x][c.y+1].value)
				if up-value == 1 {
					nextCoordinates = append(nextCoordinates, topographicMap[c.x][c.y+1])
				}
				if left-value == 1 {
					nextCoordinates = append(nextCoordinates, topographicMap[c.x-1][c.y])
				}
				return nextCoordinates
			} else {
				left, _ := strconv.Atoi(topographicMap[c.y][c.x-1].value)
				right, _ := strconv.Atoi(topographicMap[c.y][c.x+1].value)
				up, _ := strconv.Atoi(topographicMap[c.y+1][c.x].value)

				if left-value == 1 {
					nextCoordinates = append(nextCoordinates, topographicMap[c.y][c.x-1])
				}
				if right-value == 1 {
					nextCoordinates = append(nextCoordinates, topographicMap[c.y][c.x+1])
				}
				if up-value == 1 {
					nextCoordinates = append(nextCoordinates, topographicMap[c.y+1][c.x])
				}
				return nextCoordinates
			}

		}
		if isAtUPBoundary(topographicMap, c) {
			if isAtLeftBoundary(topographicMap, c) {
				right, _ := strconv.Atoi(topographicMap[c.x+1][c.y].value)
				down, _ := strconv.Atoi(topographicMap[c.x][c.y-1].value)

				if right-value == 1 {
					nextCoordinates = append(nextCoordinates, topographicMap[c.x+1][c.y])
				}
				if down-value == 1 {
					nextCoordinates = append(nextCoordinates, topographicMap[c.x][c.y-1])
				}
				return nextCoordinates
			} else if isAtRightBoundary(topographicMap, c) {
				left, _ := strconv.Atoi(topographicMap[c.x-1][c.y].value)
				down, _ := strconv.Atoi(topographicMap[c.x][c.y-1].value)

				if left-value == 1 {
					nextCoordinates = append(nextCoordinates, topographicMap[c.x-1][c.y])
				}
				if down-value == 1 {
					nextCoordinates = append(nextCoordinates, topographicMap[c.x][c.y-1])
				}
				return nextCoordinates
			} else {
				left, _ := strconv.Atoi(topographicMap[c.x-1][c.y].value)
				right, _ := strconv.Atoi(topographicMap[c.x+1][c.y].value)
				down, _ := strconv.Atoi(topographicMap[c.x][c.y-1].value)

				if left-value == 1 {
					nextCoordinates = append(nextCoordinates, topographicMap[c.x-1][c.y])
				}
				if right-value == 1 {
					nextCoordinates = append(nextCoordinates, topographicMap[c.x+1][c.y])
				}
				if down-value == 1 {
					nextCoordinates = append(nextCoordinates, topographicMap[c.x][c.y-1])
				}
				return nextCoordinates
			}
		}
		if isAtLeftBoundary(topographicMap, c) {
			if isAtDownBoundary(topographicMap, c) {
				right, _ := strconv.Atoi(topographicMap[c.x+1][c.y].value)
				up, _ := strconv.Atoi(topographicMap[c.x][c.y+1].value)

				if right-value == 1 {
					nextCoordinates = append(nextCoordinates, topographicMap[c.x+1][c.y])
				}
				if up-value == 1 {
					nextCoordinates = append(nextCoordinates, topographicMap[c.x][c.y+1])
				}
				return nextCoordinates
			} else if isAtUPBoundary(topographicMap, c) {
				right, _ := strconv.Atoi(topographicMap[c.x+1][c.y].value)
				down, _ := strconv.Atoi(topographicMap[c.x][c.y-1].value)

				if right-value == 1 {
					nextCoordinates = append(nextCoordinates, topographicMap[c.x+1][c.y])
				}
				if down-value == 1 {
					nextCoordinates = append(nextCoordinates, topographicMap[c.x][c.y-1])
				}
				return nextCoordinates
			} else {
				right, _ := strconv.Atoi(topographicMap[c.x+1][c.y].value)
				down, _ := strconv.Atoi(topographicMap[c.x][c.y-1].value)
				up, _ := strconv.Atoi(topographicMap[c.x][c.y+1].value)

				if right-value == 1 {
					nextCoordinates = append(nextCoordinates, topographicMap[c.x+1][c.y])
				}
				if down-value == 1 {
					nextCoordinates = append(nextCoordinates, topographicMap[c.x][c.y-1])
				}
				if up-value == 1 {
					nextCoordinates = append(nextCoordinates, topographicMap[c.x][c.y+1])
				}
				return nextCoordinates
			}

		}
		if isAtRightBoundary(topographicMap, c) {
			if isAtDownBoundary(topographicMap, c) {
				left, _ := strconv.Atoi(topographicMap[c.x-1][c.y].value)
				up, _ := strconv.Atoi(topographicMap[c.x][c.y+1].value)

				if left-value == 1 {
					nextCoordinates = append(nextCoordinates, topographicMap[c.x-1][c.y])
				}
				if up-value == 1 {
					nextCoordinates = append(nextCoordinates, topographicMap[c.x][c.y+1])
				}
				return nextCoordinates
			} else if isAtUPBoundary(topographicMap, c) {
				left, _ := strconv.Atoi(topographicMap[c.x-1][c.y].value)
				down, _ := strconv.Atoi(topographicMap[c.x][c.y-1].value)

				if left-value == 1 {
					nextCoordinates = append(nextCoordinates, topographicMap[c.x-1][c.y])
				}
				if down-value == 1 {
					nextCoordinates = append(nextCoordinates, topographicMap[c.x][c.y-1])
				}
				return nextCoordinates
			} else {
				left, _ := strconv.Atoi(topographicMap[c.x-1][c.y].value)
				down, _ := strconv.Atoi(topographicMap[c.x][c.y-1].value)
				up, _ := strconv.Atoi(topographicMap[c.x][c.y+1].value)

				if left-value == 1 {
					nextCoordinates = append(nextCoordinates, topographicMap[c.x-1][c.y])
				}
				if down-value == 1 {
					nextCoordinates = append(nextCoordinates, topographicMap[c.x][c.y-1])
				}
				if up-value == 1 {
					nextCoordinates = append(nextCoordinates, topographicMap[c.x][c.y+1])
				}
				return nextCoordinates
			}
		}
	}
	return nextCoordinates
}

func isAtUPBoundary(coordinates [][]coordinate, c coordinate) bool {
	if c.y == len(coordinates)-1 {
		return true
	}
	return false
}

func isAtDownBoundary(coordinates [][]coordinate, c coordinate) bool {
	if c.y == 0 {
		return true
	}
	return false
}

func isAtLeftBoundary(coordinates [][]coordinate, c coordinate) bool {
	if c.x == 0 {
		return true
	}
	return false
}

func isAtRightBoundary(coordinates [][]coordinate, c coordinate) bool {
	if c.x == len(coordinates[0])-1 {
		return true
	}
	return false
}

func isNotAtBoundary(coordinates [][]coordinate, c coordinate) bool {
	if c.x > 0 && c.y > 0 && c.x < len(coordinates[0])-1 && c.y < len(coordinates)-1 {
		return true
	}
	return false
}

func calculateScore(trailheads []trailhead) []trailhead {
	for i := range trailheads {
		trailheads[i].score = len(trailheads[i].trails)
	}
	return trailheads
}

func CalculateSum(trailheads []trailhead) int {
	sum := 0
	for i := range trailheads {
		sum += trailheads[i].score
	}
	return sum
}
