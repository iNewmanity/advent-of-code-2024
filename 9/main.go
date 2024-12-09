package main

import (
	"advent-of-code-2024/9/day09"
	"advent-of-code-2024/util"
	"fmt"
)

func main() {
	data, _ := util.ReadFile("/home/janneumann/Dokumente/Daten/Projekte/Privat/advent-of-code-2024/9/input.txt", "", true)
	storage := day09.ConvertInputToStorageRepresentation(data)
	fs := day09.ConvertStorageRepresentationToFileRepresentation(storage)
	day09.PrintStorage(storage)
	sortedStorage := day09.SortStorage(storage)
	fmt.Println("-----------------------")
	day09.PrintStorage(sortedStorage)
	fmt.Println("\n++++++++++++++++++++++")
	fmt.Println("Solution: ", day09.CalculateStorageChecksum(sortedStorage))

	sortedFileStorage := day09.SortStorageByFile(fs)
	day09.PrintStorage(sortedFileStorage)
	newFS := day09.ConvertFileRepresentationToStorageRepresentation(sortedFileStorage)
	fmt.Println("Solution: ", day09.CalculateStorageChecksum(newFS))
}
