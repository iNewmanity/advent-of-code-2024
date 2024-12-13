package main

import (
	"advent-of-code-2024/11/day11"
	"advent-of-code-2024/util"
	"fmt"
	"time"
)

func main() {
	data, _ := util.ReadFile("/home/janneumann/Dokumente/Daten/Projekte/Privat/advent-of-code-2024/11/input.txt", " ", true)
	pDS := day11.Stones{}
	for i := range data {
		for i2 := range data[i] {
			pDS = append(pDS, day11.Stone(data[i][i2]))
		}
	}
	numOfBlinks := 75
	pDS = blinkNTimes(pDS, numOfBlinks)
	fmt.Println("Part 1:", countStones(pDS))
}

func countStones(pDS day11.Stones) int {
	return len(pDS)
}

func blinkNTimes(pDS day11.Stones, num int) day11.Stones {
	for i := 0; i < num; i++ {
		start := time.Now()
		pDS = pDS.Blink()
		elapsed := time.Since(start)
		fmt.Println("Iteration: \t", i, "\tTime: \t", elapsed, "\tResulting in:", len(pDS), "Stones")
	}
	return pDS
}
