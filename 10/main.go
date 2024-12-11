package main

import (
	"advent-of-code-2024/10/day10"
	"advent-of-code-2024/util"
)

func main() {
	path := "/home/janneumann/Dokumente/Daten/Projekte/Privat/advent-of-code-2024/10/input.txt"
	data := getData(path)
	topographicMap := day10.ConvertStringToCoordinates(data)
	trailheads := day10.FindTrailHeads(topographicMap)
	trailheads = day10.FindAllWayPoints(topographicMap, trailheads)
	day10.PrintTrailHeads(trailheads)
}

func getData(path string) [][]string {
	result, _ := util.ReadFile(path, "", true)
	return result
}
