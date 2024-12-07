package main

import (
	"advent-of-code-2024/7/day07"
	"advent-of-code-2024/util"
	"fmt"
)

func main() {
	_, rawData := util.ReadFile("/home/janneumann/Dokumente/Daten/Projekte/Privat/advent-of-code-2024/7/input.txt", "", false)
	fmt.Println(rawData)
	totalCalibrationResult := day07.GetTotalCalibrationResult(rawData, []day07.Operator{"+", "*", "||"})
	fmt.Println("Assignment 1:", totalCalibrationResult)
}
