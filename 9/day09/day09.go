package day09

import (
	"fmt"
	"slices"
	"strconv"
)

type disk []file

type file struct {
	id   string
	size filesize
}

type filesize int

type storage []disk

type swap struct {
	file int
	free int
}

type swaps []swap

func SortStorageByFile(sto storage) storage {
	var sortedStorage storage
	for i := range sto {
		sortedStorage = append(sortedStorage, sortFileSystem(sto[i]))
	}
	return sortedStorage
}

func sortFileSystem(d disk) disk {
	swapped := []string{}
	for i := 0; i < len(d); i++ {
		if d[i].id != "." && !slices.Contains(swapped, d[i].id) {

		}

		for i2 := range d {
			if d[i2].id == "." {
				if d[i2].size == d[i].size {

				}
			}
		}
	}
	return d
}

func SortStorage(sto storage) storage {
	var sortedStorage storage
	for i := range sto {
		sortedStorage = append(sortedStorage, sortFileSystem(sto[i]))
	}
	return sortedStorage
}

func sortDisks(d disk) disk {
	for i := len(d) - 1; i >= 0; i-- {
		if d[i].id != "." {
			for i2 := range d {
				if d[i2].id == "." {
					temp := d[i].id
					d[i].id = d[i2].id
					d[i2].id = temp
				}
			}
		}
		//printDisk(d)
		if isDiskSorted(d) {
			break
		}
	}
	return d
}

func CalculateStorageChecksum(sto storage) int {
	sum := 0
	for i := range sto {
		sum += calculateDiskChecksum(sto[i])
	}
	return sum
}

func calculateDiskChecksum(d disk) int {
	sum := 0
	printDisk(d)
	for i := range d {
		if d[i].id != "." {
			id, _ := strconv.Atoi(d[i].id)
			result := i * id
			sum += result
			fmt.Println(i, "\t*", id, "\t=", result)
		}
	}
	return sum
}

func isDiskSorted(d disk) bool {
	freeSpace := false
	for i := range d {
		if !freeSpace {
			if d[i].id == "." {
				freeSpace = true
			}
		} else {
			if d[i].id != "." {
				return false
			}
		}
	}
	return true
}

func ConvertFileRepresentationToStorageRepresentation(sto storage) storage {
	newStorage := storage{}
	for i := range sto {
		newDisk := disk{}
		for j := range sto[i] {
			for k := 0; k < int(sto[i][j].size); k++ {
				newDisk = append(newDisk, file{id: sto[i][j].id, size: 1})
			}
		}
		newStorage = append(newStorage, newDisk)
	}
	return newStorage
}

func ConvertStorageRepresentationToFileRepresentation(sto storage) storage {
	fs := storage{}

	for _, d := range sto {
		newDisk := disk{}
		i := 0
		for i < len(d) {
			currentID := d[i].id
			counter := 1

			// Count consecutive occurrences of the same id
			for i+counter < len(d) && d[i+counter].id == currentID {
				counter++
			}

			// Append the file with id and its consecutive count
			newDisk = append(newDisk, file{
				id:   currentID,
				size: filesize(counter),
			})

			// Move the index to the next different id
			i += counter
		}
		fs = append(fs, newDisk)
	}

	return fs
}

func ConvertInputToStorageRepresentation(data [][]string) []disk {
	sto := storage{}
	for i := range data {
		d := disk{}
		for i2 := range data[i] {
			if i2%2 == 1 {
				times, _ := strconv.Atoi(data[i][i2])
				if times > 0 {
					for i := 0; i < times; i++ {
						d = append(d, file{".", filesize(times)})
					}
				}

			} else {
				times, _ := strconv.Atoi(data[i][i2])
				index := nextIndex(d)
				for i := 0; i < times; i++ {
					d = append(d, file{strconv.Itoa(index), filesize(times)})
				}
			}
		}
		sto = append(sto, d)
	}
	return sto
}

func printDisk(d disk) {
	fmt.Println("----------------Disk---------------\n")
	for i := range d {
		fmt.Print(d[i].id)
	}
	fmt.Println("\n-----------------------------------\n")
}

func PrintStorage(sto storage) {
	for i := range sto {
		printDisk(sto[i])
	}
}

func nextIndex(d disk) int {
	highestindex := -1
	for i := range d {
		if d[i].id != "." {
			highestindex, _ = strconv.Atoi(d[i].id)
		}
	}
	return highestindex + 1
}
