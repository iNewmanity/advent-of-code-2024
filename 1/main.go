package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	var filePath string = "lists.csv"
	var lists [][]int
	lists, _ = readCSVAsIntArray(filePath)

	fmt.Println(calculateListSize(lists))

	var list_one [1000]int
	var list_two [1000]int

	for i := 0; i < len(lists); i++ {
		list_one[i] = lists[i][0] // First element of each [2]int
		list_two[i] = lists[i][1] // Second element of each [2]int
	}

	slices.Sort(list_one[:])
	slices.Sort(list_two[:])

	var distance int = calculateDistanceBetweenTwoLists(list_one, list_two)
	fmt.Println(distance)

	var similarity int = calculateSimilarityScore(list_one, list_two)
	fmt.Println(similarity)

}

func calculateSimilarityScore(list_one [1000]int, list_two [1000]int) int {
	var similarity int = 0
	for i := 0; i < len(list_one); i++ {
		similarity += multiplyNumber(list_one[i], howOftenInListTwo(list_one[i], list_two))
	}
	return similarity
}

func multiplyNumber(num int, multiplicant int) int {
	return num * multiplicant
}

func howOftenInListTwo(num int, list_two [1000]int) int {
	var multiplicant int = 0
	for i := 0; i < len(list_two); i++ {
		if list_two[i] == num {
			multiplicant++
		}
	}
	return multiplicant
}

func readCSVAsIntArray(filePath string) ([][]int, error) {
	// Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read all rows
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	// Convert rows to a 2D array of integers
	var result [][]int
	for _, row := range rows {
		intRow := []int{}
		for _, value := range row {
			intValue, err := strconv.Atoi(value)
			if err != nil {
				return nil, fmt.Errorf("error converting value to int: %w", err)
			}
			intRow = append(intRow, intValue)
		}
		result = append(result, intRow)
	}

	return result, nil
}

func calculateListSize(lists [][]int) int {
	var horizontalSize int = len(lists)
	return horizontalSize
}

func extractList(lists [][]int, column int) [1000]int {
	var result [1000]int
	for i := 0; i < len(lists); i++ {
		result[i] = lists[i][column]
	}
	return result
}

func sortList(list []int) []int {
	slices.Sort(list)
	return list
}

func calculateDistanceBetweenTwoNumbers(num1 int, num2 int) int {
	biggerNum, smallerNum := compareSizeBetweenNumbers(num1, num2)
	return biggerNum - smallerNum
}

func compareSizeBetweenNumbers(num1 int, num2 int) (int, int) {
	if num1 > num2 {
		return num1, num2
	} else {
		return num2, num1
	}
}

func calculateDistanceBetweenTwoLists(list_one [1000]int, list_two [1000]int) int {
	var sum int = 0
	for i := 0; i < len(list_one); i++ {
		sum += calculateDistanceBetweenTwoNumbers(list_one[i], list_two[i])
	}
	return sum
}
