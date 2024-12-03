package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var filePath string = "2/input.txt"
	var data [][]int = read_file(filePath)
	var result int = countSafeRecords(data)
	fmt.Println(result)

}

func countSafeRecords(data [][]int) int {
	var counter int
	for i := range data {
		fmt.Println(data[i])
		if isSafe(data[i]) {
			counter++
		} else if withProblemDampenerSuccess(data[i]) {
			counter++
		} else {
		}
	}
	return counter
}

func withProblemDampenerSuccess(numArray []int) bool {
	for i := range numArray {
		subarray := arrayWithoutIndex(numArray, i)

		var isMonotoneDescending bool = isMonotoneDescending(subarray)
		var isMonotoneAscending bool = isMonotoneAscending(subarray)
		var isAppropriateDistance bool = hasAppropriateDistance(subarray)

		if (isMonotoneAscending || isMonotoneDescending) && isAppropriateDistance {
			return true
		}

	}
	return false
}

func arrayWithoutIndex(numArray []int, index int) []int {
	var arrayWithoutIndex []int
	for i := range numArray {
		if i == index {
			continue
		}
		arrayWithoutIndex = append(arrayWithoutIndex, numArray[i])
	}
	return arrayWithoutIndex
}

func hasTooMuchDistance() {

}

func isSafe(numArray []int) bool {
	if isMonotoneAscending(numArray) || isMonotoneDescending(numArray) {
		if hasAppropriateDistance(numArray) {
			return true
		}
		return false
	} else {
		return false
	}
}

func hasAppropriateDistance(numArray []int) bool {
	var previousValue int = numArray[0]
	for i := range numArray {
		if i == 0 {
			continue
		}
		var distanceAppropriate bool = isDistanceAppropriate(previousValue, numArray[i])
		previousValue = numArray[i]
		if !distanceAppropriate {
			return false
		}

	}
	return true
}

func isDistanceAppropriate(num1 int, num2 int) bool {
	var distance = math.Abs(float64(num1) - float64(num2))
	if distance == 0 {
		return false
	} else if distance <= 3 {
		return true
	}
	return false
}

func isMonotoneDescending(numArray []int) bool {
	var previousValue int = numArray[0]
	for i := range numArray {
		if i == 0 {
			continue
		}
		if numArray[i] >= previousValue {
			return false
		}
		previousValue = numArray[i]
	}
	return true
}

func isMonotoneAscending(numArray []int) bool {
	var previousValue int = numArray[0]
	for i := range numArray {
		if i == 0 {
			continue
		}
		if numArray[i] <= previousValue {
			return false
		}
		previousValue = numArray[i]

	}
	return true
}

func read_file(filePath string) [][]int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error opening file: ", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var data [][]int
	for scanner.Scan() {
		var lineArray []int
		var stringArray []string = splitUpLine(scanner.Text())
		for i := range stringArray {
			lineArray = append(lineArray, convertToInt(stringArray[i]))
		}
		data = append(data, lineArray)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading file", err)
	}
	return data
}

func splitUpLine(s string) []string {
	return strings.Split(s, " ")
}

func convertToInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
