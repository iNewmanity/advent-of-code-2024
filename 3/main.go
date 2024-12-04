package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var filePath string = "/home/janneumann/Dokumente/Daten/Projekte/Privat/advent-of-code-2024/3/input.txt"

	var data string = read_file(filePath)
	var multiplyCalls []string = filterMultiplyCalls(data)
	var result int = 0
	for i := range multiplyCalls {
		var multiplicant int
		var multiplicator int
		multiplicant, multiplicator = getMultiplyPairs(multiplyCalls[i])
		result += multiply(multiplicant, multiplicator)
	}

	fmt.Println("Result Assignment 1: ", result)
	var result2 int = assignmentTwo(data)
	fmt.Println("Result Assignment 2: ", result2)

}

func multiply(multiplicant int, multiplicator int) int {
	return multiplicator * multiplicant
}

func getMultiplyPairs(value string) (int, int) {
	re := regexp.MustCompile(`\d{1,3}`)

	matches := re.FindAllString(value, -1)

	first, err := strconv.Atoi(matches[0])

	if err != nil {
		panic(err)
	}

	second, err := strconv.Atoi(matches[1])

	if err != nil {
		panic(err)
	}

	return first, second
}

func assignmentTwo(data string) int {
	var matches []string = filterMatches(data)
	fmt.Println(matches)
	var allowedMultiplications []string = filterAllowed(matches)
	var result int
	for i := range allowedMultiplications {
		var multiplicant int
		var multiplicator int
		multiplicant, multiplicator = getMultiplyPairs(allowedMultiplications[i])
		result += multiply(multiplicant, multiplicator)
	}
	return result
}

func filterAllowed(matches []string) []string {
	var allowed bool = true
	var result []string
	for i := range matches {
		if matches[i] == "do()" {
			allowed = true
		} else if matches[i] == "don't()" {
			allowed = false
		} else {
			if allowed {
				result = append(result, matches[i])
			} else {
				continue
			}
		}
	}
	return result
}

func runRegEx(data string, regex string) []string {
	re := regexp.MustCompile(regex)

	matches := re.FindAllString(data, -1)

	var result []string
	for _, match := range matches {
		result = append(result, match)
	}
	return result
}

func filterMultiplyCalls(data string) []string {
	return runRegEx(data, "mul\\((\\d{1,3}),(\\d{1,3})\\)")
}

func filterMatches(data string) []string {
	return runRegEx(data, "mul\\((\\d{1,3}),(\\d{1,3})\\)|do\\(\\)|don't\\(\\)")
}

func read_file(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error opening file: ", err)
		return err.Error()
	}
	defer file.Close()
	var data string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading file", err)
	}
	return data
}
