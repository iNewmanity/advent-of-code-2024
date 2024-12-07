package day07

import (
	"slices"
	"strconv"
	"strings"
)

type Assignment struct {
	result    int
	numbers   []int
	operators []operator
}

type operator string

func GetTotalCalibrationResult(input []string) int {
	assignments := assignmentCreator(input)
	var totalCalibrationResult int
	for i := range assignments {
		if operatorEvaluator(assignments[i]) {
			totalCalibrationResult += assignments[i].result
		}
	}
	return totalCalibrationResult
}

func operatorEvaluator(assignment Assignment) bool {
	return evaluateCombination(assignment)
}

func evaluateCombination(assignment Assignment, allowedOperators []operator) bool {
	targetNumber := assignment.result
	values := assignment.numbers

	operatorList := getPermutationOperators(len(values)-1, []operator{"*", "+"})

	for i := 0; i < len(operatorList); i++ {
		result := values[0]
		operators := operatorList[i]

		for j := 1; j < len(values); j++ {
			if operators[j-1] == "*" {
				result *= values[j]
			}
			if operators[j-1] == "+" {
				result += values[j]
			}
		}

		if result == targetNumber {
			return true
		}

	}

	return false
}

func getPermutationOperators(count int, ops []operator) [][]operator {
	var result [][]operator

	totalPermutations := 1
	for i := 0; i < count; i++ {
		totalPermutations *= len(ops)
	}

	for i := 0; i < totalPermutations; i++ {
		permutation := make([]operator, count)
		index := i
		for j := 0; j < count; j++ {
			permutation[j] = ops[index%len(ops)]
			index /= len(ops)
		}
		result = append(result, permutation)
	}

	return result
}

func isAllowedOperator(op operator, allowed []operator) bool {
	if slices.Contains(allowed, op) {
		return true
	}
	return false
}

func assignmentCreator(input []string) []Assignment {
	assignments := []Assignment{}
	for i := range input {
		assignment := Assignment{}
		assignment.numbers = numberExtractor(input[i])
		assignment.result = resultExtractor(input[i])
		assignments = append(assignments, assignment)
	}
	return assignments
}

func stringToInt(input string) int {
	result, _ := strconv.Atoi(input)
	return result
}

func resultExtractor(input string) int {
	result, _, _ := strings.Cut(input, ":")
	intResult := stringToInt(result)
	return intResult
}

func numberExtractor(input string) []int {
	_, numbers, _ := strings.Cut(input, ":")
	numbers, _ = strings.CutPrefix(numbers, " ")
	numberstring := strings.Split(numbers, " ")
	numbersInt := []int{}
	for i := range numberstring {
		numbersInt = append(numbersInt, stringToInt(numberstring[i]))
	}
	return numbersInt
}
