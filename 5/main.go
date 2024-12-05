package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	var inputPath string = "/home/janneumann/Dokumente/Daten/Projekte/Privat/advent-of-code-2024/5/input.txt"
	var ruleSetPath string = "/home/janneumann/Dokumente/Daten/Projekte/Privat/advent-of-code-2024/5/rules.txt"

	var input [][]int = read_file(inputPath, ",")
	var ruleset [][]int = read_file(ruleSetPath, "|")

	var result int = assignmentOne(input, ruleset)
	fmt.Println("Result: ", result)
}

func assignmentOne(data [][]int, ruleSet [][]int) int {
	var result int
	var correctPageOrders [][]int = filterCorrectPageOrders(ruleSet, data)
	fmt.Println(correctPageOrders)
	result = addCorrectPageOrder(correctPageOrders)

	return result
}

// Result Eval Assignment 1
func getMiddelIndexValue(pageOrder []int) int {
	var length int = len(pageOrder)
	var halfIndex float64 = (float64(length/2) - 0.5) + 1
	return pageOrder[int(halfIndex)]
}

func addCorrectPageOrder(correctPageOrders [][]int) int {
	var result int
	for i := range correctPageOrders {
		result += getMiddelIndexValue(correctPageOrders[i])
	}
	return result
}

// Resultation Assignment 1
func filterCorrectPageOrders(ruleset [][]int, pageorders [][]int) [][]int {
	var resultset [][]int
	for i := range pageorders {
		if checkIfRulesApply(ruleset, pageorders[i]) {
			resultset = append(resultset, pageorders[i])
		} else {
			continue
		}
	}
	return resultset
}

func checkIfRulesApply(ruleset [][]int, pageOrder []int) bool {
	var appliableRules [][]int = ruleApplicationFilter(ruleset, pageOrder)
	for i := range appliableRules {
		if checkRuleApplicationByIndex(appliableRules[i], pageOrder) {
			continue
		} else {
			return false
		}
	}
	return true
}

func checkRuleApplicationByIndex(rule []int, pageOrder []int) bool {
	var firstValue int
	for i := range rule {
		if i == 0 {
			firstValue = getIndexOfValue(rule[i], pageOrder)
		} else {
			var val int = getIndexOfValue(rule[i], pageOrder)
			if firstValue < val {
				return true
			} else {
				return false
			}
		}
	}
	return false
}

func getIndexOfValue(value int, pageOrder []int) int {
	return slices.Index(pageOrder, value)
}

func ruleApplicationFilter(ruleset [][]int, pageOrder []int) [][]int {
	var appliableRules [][]int
	for i := range ruleset {
		if checkIfRuleApplies(ruleset[i], pageOrder) {
			appliableRules = append(appliableRules, ruleset[i])
		} else {
			continue
		}
	}
	return appliableRules
}

func checkIfRuleApplies(rule []int, pageOrder []int) bool {
	for i := range rule {
		if !slices.Contains(pageOrder, rule[i]) {
			return false
		} else {
			continue
		}
	}
	return true
}

// Reading in files
func read_file(filePath string, delimiter string) [][]int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error opening file: ", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var data [][]int
	for scanner.Scan() {
		var intArray []int = splitUpLine(scanner.Text(), delimiter)
		data = append(data, intArray)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading file", err)
	}
	return data
}

func splitUpLine(s string, delimiter string) []int {
	var stringArray []string = strings.Split(s, delimiter)
	var intArray []int
	for i := range stringArray {
		intVal, _ := strconv.Atoi(stringArray[i])
		intArray = append(intArray, intVal)
	}
	return intArray
}
