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

	result := assignmentOne(input, ruleset)
	result2 := assignmentTwo(input, ruleset, result)
	fmt.Println("Result: ", result)
	fmt.Println("Result2: ", result2)
}

func assignmentOne(data [][]int, ruleSet [][]int) int {
	var result int
	var correctPageOrders [][]int = filterCorrectPageOrders(ruleSet, data)
	result = addCorrectPageOrder(correctPageOrders)

	return result
}

func assignmentTwo(data [][]int, ruleSet [][]int, value int) int {
	count := 0
	for i := 0; i < len(data); i++ {
		ints := data[i]
		if !isSuccessfulOrder(ints, ruleApplicationFilter(ruleSet, data[i])) {
			orderedInts := reorder(ints, ruleApplicationFilter(ruleSet, data[i]))
			middleIndex := len(orderedInts) / 2
			count += orderedInts[middleIndex]
		}
	}
	return count

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

// Resultation Assignment 2

func filterCorrectPageOrders2(ruleset [][]int, pageorders [][]int) [][]int {
	var resultset [][]int
	for i := range pageorders {
		if checkIfRulesApply2(ruleset, pageorders[i]) {
			resultset = append(resultset, pageorders[i])
		} else {
			continue
		}
	}
	return resultset
}

func checkIfRulesApply2(ruleset [][]int, pageOrder []int) bool {
	var appliableRules [][]int = ruleApplicationFilter(ruleset, pageOrder)
	for i := range appliableRules {
		if checkRuleApplicationByIndex2(appliableRules[i], pageOrder) {
			continue
		} else {
			return false
		}
	}
	return true
}

func checkRuleApplicationByIndex2(rule []int, pageOrder []int) bool {
	var firstValue int
	for i := range rule {
		if i == 0 {
			firstValue = getIndexOfValue(rule[i], pageOrder)
		} else {
			var val int = getIndexOfValue(rule[i], pageOrder)
			if firstValue < val {
				return true
			} else {
				pageOrder = RepairRuleInPageOrder(rule, pageOrder)
				break
			}
		}
	}
	if checkRuleApplicationByIndex(rule, pageOrder) {
		return true
	}
	return false
}

func RepairRuleInPageOrder(rule []int, pageOrder []int) []int {
	for i := range pageOrder {
		if pageOrder[i] == rule[0] {
			// Ensure we don't go out of bounds
			if i+1 < len(pageOrder) {
				// Swap the current value with the next one
				pageOrder[i], pageOrder[i+1] = pageOrder[i+1], pageOrder[i]
				return pageOrder
			} else {
				// If rule is still not satisfied, revert the swap
				if i+1 < len(pageOrder) {
					pageOrder[i], pageOrder[i+1] = pageOrder[i+1], pageOrder[i]
				}
				continue
			}
		}
	}
	return pageOrder
}

func isSuccessfulOrder(arr []int, rules [][]int) bool {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if !isOrderCorrect(arr[i], arr[j], rules) {
				return false
			}
		}
	}
	return true
}

func isOrderCorrect(currentInt, nextInt int, rules [][]int) bool {
	for _, rule := range rules {
		if rule[0] == currentInt && rule[1] == nextInt {
			return true
		}
		if rule[0] == nextInt && rule[1] == currentInt {
			return false
		}
	}
	return true
}

func reorder(arr []int, rules [][]int) []int {
	count := len(arr)
	result := make([]int, count)
	copy(result, arr)

	for i := 0; i < count-1; i++ {
		for j := 0; j < count-1-i; j++ {
			if !isOrderCorrect(result[j], result[j+1], rules) {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	return result
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
