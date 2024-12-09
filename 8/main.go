package main

import (
	"advent-of-code-2024/8/day08"
	"advent-of-code-2024/util"
	"fmt"
)

func main() {
	path := "/home/janneumann/Dokumente/Daten/Projekte/Privat/advent-of-code-2024/8/input.txt"
	data, _ := util.ReadFile(path, "", true)
	result1 := assignment1(data)
	fmt.Println("Assignment 1: ", result1)
	result2 := assignment2(data)
	fmt.Println("Assignment 2: ", result2)
}

func assignment1(data [][]string) int {
	return day08.Day081(data)
}

func assignment2(data [][]string) int {
	return day08.Day082(data)
}
