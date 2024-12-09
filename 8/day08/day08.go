package day08

import (
	_ "io"
	"slices"
)

type coordinate struct {
	x      int
	y      int
	letter string
}

type distance struct {
	c1 coordinate
	c2 coordinate
	x  int
	y  int
}

func Day081(data [][]string) int {
	height := len(data)
	width := len(data[0])
	coordinates := getCoordinates(data)
	letters := getLetters(coordinates)
	filteredCoordinates := filterCoordinates(coordinates, letters)
	distances := calculateDistances(filteredCoordinates)
	antinodes := createAntinodes(distances, height, width)
	return countAntinodes(antinodes)
}

func Day082(data [][]string) int {
	height := len(data)
	width := len(data[0])
	coordinates := getCoordinates(data)
	letters := getLetters(coordinates)
	filteredCoordinates := filterCoordinates(coordinates, letters)
	distances := calculateDistances(filteredCoordinates)
	antinodes := createAntinodes(distances, height, width)
	return countAntinodes(antinodes)
}

func calculateDistances(coordinates []coordinate) []distance {
	var distances []distance
	letters := getLetters(coordinates)
	for i := range letters {
		letterCoordinates := extrudeLetter(coordinates, letters[i])
		if len(letterCoordinates) > 0 {
			distances = append(distances, calculateDistancesForLetter(letterCoordinates)...)
		}
	}
	return distances
}

func calculateDistancesForLetter(coordinates []coordinate) []distance {
	var distances []distance
	for i := range coordinates {
		for i2 := range coordinates {
			if i != i2 {
				distance := calculateDistance(coordinates[i], coordinates[i2])
				distances = append(distances, distance)
			}
		}
	}

	return distances
}

func createAntinodes(distances []distance, height int, width int) []coordinate {
	antinodes := []coordinate{}
	for i := range distances {
		antinode := createAntinode(distances[i])
		if isAntinodeCorrect(antinode, antinodes, height, width) {
			antinodes = append(antinodes, antinode)
		}
	}
	return antinodes
}

func isAntinodeCorrect(antinode coordinate, antinodes []coordinate, height int, width int) bool {
	if isInBounds(antinode, height, width) && checkIfFree(antinode, antinodes) {
		return true
	}
	return false
}

func getLetters(coordinates []coordinate) []string {
	result := []string{}
	for i := range coordinates {
		if coordinates[i].letter != "." {
			if !slices.Contains(result, coordinates[i].letter) {
				result = append(result, coordinates[i].letter)
			}
		}
	}
	return result
}

func getCoordinates(data [][]string) []coordinate {
	result := []coordinate{}
	for i := range data {
		for i2 := range data[i] {
			result = append(result, coordinate{
				x:      i2,
				y:      i,
				letter: data[i][i2],
			})
		}
	}
	return result
}

func filterCoordinates(coordinates []coordinate, letters []string) []coordinate {
	result := []coordinate{}
	for i := range coordinates {
		if slices.Contains(letters, coordinates[i].letter) {
			result = append(result, coordinates[i])
		}
	}
	return result
}

func extrudeLetter(coordinates []coordinate, letter string) []coordinate {
	result := []coordinate{}
	for i := range coordinates {
		if coordinates[i].letter == letter {
			result = append(result, coordinates[i])
		}
	}
	return result
}

func calculateDistance(c1 coordinate, c2 coordinate) distance {
	return distance{
		c1: c1,
		c2: c2,
		x:  c2.x - c1.x,
		y:  c2.y - c1.y,
	}
}

func createAntinode(dist distance) coordinate {
	return coordinate{
		x:      dist.c2.x + dist.x,
		y:      dist.c2.y + dist.y,
		letter: "#",
	}
}

func isInBounds(antinode coordinate, height int, width int) bool {
	if antinode.x < width && antinode.x >= 0 && antinode.y < height && antinode.y >= 0 {
		return true
	}
	return false
}

func checkIfFree(antinode coordinate, antinodes []coordinate) bool {
	if slices.Contains(antinodes, antinode) {
		return false
	}
	return true
}

func countAntinodes(antinodes []coordinate) int {
	count := 0
	for i := range antinodes {
		if antinodes[i].letter == "#" {
			count++
		}
	}
	return count
}
