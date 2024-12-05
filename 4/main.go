package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var filename string = "/home/janneumann/Dokumente/Daten/Projekte/Privat/advent-of-code-2024/4/input.txt"
	var data [][]string = read_file(filename)

	var result int = firstAssignment(data)
	fmt.Println("First Assignment: ", result)
	var result2 int = secondAssignment(data)
	fmt.Println("Second Assignment: ", result2)

}

func secondAssignment(data [][]string) int {
	var count int = searchForAs(data)
	return count
}

func searchForAs(data [][]string) int {
	count := 0
	for i := range data {
		if i == 0 {
			continue
		} else if i == len(data)-1 {
			continue
		} else {
			for i2 := range data[i] {
				if i2 == 0 {
					continue
				} else if i2 == len(data[i])-1 {
					continue
				} else {
					if isA(data[i][i2]) {
						if checkSurroundingsforPattern(data, []int{i, i2}) {
							count++
						} else {

						}
					}
				}
			}
		}
	}
	return count
}

func checkSurroundingsforPattern(data [][]string, coordinate []int) bool {
	var upL, downR, upR, downL = calculateCoordinatesToCheck(coordinate)
	if data[upL[0]][upL[1]] == "M" && data[upR[0]][upR[1]] == "M" {
		if data[downR[0]][downR[1]] == "S" && data[downL[0]][downL[1]] == "S" {
			return true
		}
	} else if data[upL[0]][upL[1]] == "S" && data[upR[0]][upR[1]] == "S" {
		if data[downR[0]][downR[1]] == "M" && data[downL[0]][downL[1]] == "M" {
			return true
		}
	} else if data[upL[0]][upL[1]] == "M" && data[upR[0]][upR[1]] == "S" {
		if data[downL[0]][downL[1]] == "M" && data[downR[0]][downR[1]] == "S" {
			return true
		}
	} else if data[upL[0]][upL[1]] == "S" && data[upR[0]][upR[1]] == "M" {
		if data[downL[0]][downL[1]] == "S" && data[downR[0]][downR[1]] == "M" {
			return true
		}
	} else {
		return false
	}
	return false
}

func calculateCoordinatesToCheck(coordinates []int) ([]int, []int, []int, []int) {
	var upL = []int{coordinates[0] - 1, coordinates[1] - 1}
	var downL = []int{coordinates[0] + 1, coordinates[1] - 1}
	var upR = []int{coordinates[0] - 1, coordinates[1] + 1}
	var downR = []int{coordinates[0] + 1, coordinates[1] + 1}
	fmt.Println(upL, upR)
	fmt.Println(coordinates)
	fmt.Println(downL, downR)
	return upL, downR, upR, downL
}

func isA(text string) bool {
	if text == "A" {
		return true
	} else {
		return false
	}
}

func firstAssignment(data [][]string) int {
	var linearAndExtendedDataStructure [][]string = createAlternativeDataStructure(data)

	var count int = countXMASAppearing(linearAndExtendedDataStructure, "xmas")
	return count
}

func countXMASAppearing(altStructure [][]string, match string) int {
	count := 0

	for _, row := range altStructure {
		// Convert the row to a single string by joining the elements
		rowString := strings.Join(row, "")
		if len(rowString) < 4 {
			continue
		}

		var stringsInRow int = strings.Count(strings.ToLower(rowString), match)

		// Count occurrences of "xmas" in the string (case insensitive)
		count += stringsInRow

	}

	return count
}

func createAlternativeDataStructure(data [][]string) [][]string {
	// Horizontal iterations
	var linearAndExtendedDataStructure [][]string = createExtendedHorizontalForward(data)
	linearAndExtendedDataStructure = createExtendedHorizontalBackwards(data, linearAndExtendedDataStructure)

	// Vertical Iterations
	linearAndExtendedDataStructure = createExtendedVerticalForward(data, linearAndExtendedDataStructure)
	linearAndExtendedDataStructure = createExtendedVerticalBackwards(data, linearAndExtendedDataStructure)

	// Diagonal Traversals
	linearAndExtendedDataStructure = createDiagonalTraversal(data, linearAndExtendedDataStructure, "leftTopToRightBottom")
	linearAndExtendedDataStructure = createDiagonalTraversal(data, linearAndExtendedDataStructure, "rightTopToLeftBottom")
	linearAndExtendedDataStructure = createDiagonalTraversal(data, linearAndExtendedDataStructure, "leftBottomToRightTop")
	linearAndExtendedDataStructure = createDiagonalTraversal(data, linearAndExtendedDataStructure, "rightBottomToLeftTop")
	return linearAndExtendedDataStructure
}

func createExtendedHorizontalForward(data [][]string) [][]string {
	var linearAndExtendedDataStructure [][]string
	for i := range data {
		var extendedRow []string
		for i2 := range data[i] {
			extendedRow = append(extendedRow, data[i][i2])
			var maxIndex int = len(data[i]) - 1
			if i2 == maxIndex {
			} else {
				continue
			}
		}
		linearAndExtendedDataStructure = append(linearAndExtendedDataStructure, extendedRow)
	}
	return linearAndExtendedDataStructure
}

func createExtendedHorizontalBackwards(data [][]string, linearAndExtendedDataStructure [][]string) [][]string {
	for i := range data {
		var extendedRow []string
		for i2 := len(data[i]) - 1; i2 >= 0; i2-- {

			extendedRow = append(extendedRow, data[i][i2])
		}
		linearAndExtendedDataStructure = append(linearAndExtendedDataStructure, extendedRow)
	}
	return linearAndExtendedDataStructure
}

func createExtendedVerticalForward(data [][]string, linearAndExtendedDataStructure [][]string) [][]string {
	// Get the number of rows and columns in the matrix
	rows := len(data)
	if rows == 0 {
		return linearAndExtendedDataStructure
	}
	cols := len(data[0])

	// Iterate over the matrix column by column
	for col := 0; col < cols; col++ {
		var extendedRow []string

		for row := 0; row < rows; row++ {
			// Extract data for the current column
			extendedRow = append(extendedRow, data[row][col])
		}

		// Add the constructed extendedRow to the linearAndExtendedDataStructure
		linearAndExtendedDataStructure = append(linearAndExtendedDataStructure, extendedRow)
	}

	return linearAndExtendedDataStructure
}

func createExtendedVerticalBackwards(data [][]string, linearAndExtendedDataStructure [][]string) [][]string {
	// Get the number of rows and columns in the matrix
	rows := len(data)
	if rows == 0 {
		return linearAndExtendedDataStructure
	}
	cols := len(data[0])

	// Iterate over the matrix column by column, bottom to top
	for col := 0; col < cols; col++ {
		var extendedRow []string

		for row := rows - 1; row >= 0; row-- {

			extendedRow = append(extendedRow, data[row][col])
		}

		// Add the constructed extendedRow to the linearAndExtendedDataStructure
		linearAndExtendedDataStructure = append(linearAndExtendedDataStructure, extendedRow)
	}

	return linearAndExtendedDataStructure
}

func createDiagonalTraversal(
	data [][]string,
	linearDataStructure [][]string,
	direction string,
) [][]string {
	rows := len(data)
	if rows == 0 {
		return linearDataStructure
	}
	cols := len(data[0])

	switch direction {
	case "leftTopToRightBottom":
		// Traverse diagonals from left-top to right-bottom
		for start := 0; start < rows+cols-1; start++ {
			var diagonal []string
			for i := 0; i <= start; i++ {
				j := start - i
				if i < rows && j < cols {
					diagonal = append(diagonal, data[i][j])
				}
			}
			linearDataStructure = append(linearDataStructure, diagonal)
		}

	case "rightTopToLeftBottom":
		// Traverse diagonals from right-top to left-bottom
		for start := 0; start < rows+cols-1; start++ {
			var diagonal []string
			for i := 0; i <= start; i++ {
				j := cols - 1 - start + i
				if i < rows && j >= 0 && j < cols {
					diagonal = append(diagonal, data[i][j])
				}
			}
			linearDataStructure = append(linearDataStructure, diagonal)
		}

	case "leftBottomToRightTop":
		// Traverse diagonals from left-bottom to right-top
		for start := 0; start < rows+cols-1; start++ {
			var diagonal []string
			for i := 0; i <= start; i++ {
				j := start - i
				row := rows - 1 - i
				if row >= 0 && j < cols {
					diagonal = append(diagonal, data[row][j])
				}
			}
			linearDataStructure = append(linearDataStructure, diagonal)
		}

	case "rightBottomToLeftTop":
		// Traverse diagonals from right-bottom to left-top
		for start := 0; start < rows+cols-1; start++ {
			var diagonal []string
			for i := 0; i <= start; i++ {
				j := cols - 1 - start + i
				row := rows - 1 - i
				if row >= 0 && j >= 0 && j < cols {
					diagonal = append(diagonal, data[row][j])
				}
			}
			linearDataStructure = append(linearDataStructure, diagonal)
		}

	default:
		panic("Invalid direction specified")
	}

	return linearDataStructure
}

func read_file(filePath string) [][]string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error opening file: ", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var data [][]string
	for scanner.Scan() {
		var stringArray []string = splitUpLine(scanner.Text())
		data = append(data, stringArray)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading file", err)
	}
	return data
}

func splitUpLine(s string) []string {
	return strings.Split(s, "")
}
