package day09

import (
	"fmt"
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
	swapInstructions := swaps{}
	for i := len(d) - 1; i > 0; i-- {
		if d[i].id != "." {
			swapPartner := getForwardFreeSpaceForNeededSpace(d, d[i].size)
			fmt.Println(d[i].id, "sp:", swapPartner)
			swapInstructions = append(swapInstructions, swap{
				file: i,
				free: swapPartner,
			})
		}
	}

	for i := range swapInstructions {
		if swapInstructions[i].free == -1 {
			if swapInstructions[i].free != getForwardFreeSpaceForNeededSpace(d, filesize(swapInstructions[i].file)) {
				if swapInstructions[i].free != -1 {
					d = swapFileWithFreeSpace(d, swapInstructions[i].file, getForwardFreeSpaceForNeededSpace(d, filesize(swapInstructions[i].file)))
				}
			}
		} else {
			d = swapFileWithFreeSpace(d, swapInstructions[i].file, swapInstructions[i].free)
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
			fmt.Println(i, "*", id, "=", result)
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

func swapFileWithFreeSpace(d disk, i, j int) disk {
	thefile := d[i]
	remainingSize := int(d[j].size) - int(d[i].size)
	if remainingSize == 0 {
		d[i] = thefile
		d[i] = d[j]
		d[j] = thefile
	} else {
		d[j].size = filesize(int(d[j].size) - remainingSize)
		d[i] = d[j]
		d[j] = thefile
		d = append(d[:j+1], append([]file{file{
			id:   ".",
			size: filesize(remainingSize),
		}}, d[j+1:]...)...)
	}
	return d
}

func getForwardFreeSpaceForNeededSpace(d disk, size filesize) int {
	for i := range d {
		if d[i].id == "." {
			if d[i].size >= size {
				return i
			}
		}
	}
	return -1
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
