package main

import (
	"advent-of-code-2024/9/day09"
	"advent-of-code-2024/util"
	"fmt"
)

func main() {
	data, _ := util.ReadFile("/home/janneumann/Dokumente/Daten/Projekte/Privat/advent-of-code-2024/9/input.txt", "", true)
	storage := day09.ConvertInputToStorageRepresentation(data)
	fs := day09.ConverStorageRepresentationToFileRepresentation(storage)
	day09.PrintStorage(storage)
	sortedStorage := day09.SortStorage(storage)
	fmt.Println("-----------------------")
	day09.PrintStorage(sortedStorage)
	fmt.Println("\n++++++++++++++++++++++")
	fmt.Println("Solution: ", day09.CalculateStorageChecksum(sortedStorage))
	fmt.Println("\n++++++++++++++++++++++")
	day09.PrintStorage(fs)
	fmt.Println("\n", len(fs[0]))
	fmt.Println("\n++++++++++++++++++++++")
	fmt.Println("Solution: ", day09.CalculateStorageByFileChecksum(fs))
}
