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

type filesystem storage

type freespace file

type filesize int

type space file

type storage []disk

func SortStorage(sto storage) storage {
	var sortedStorage storage
	for i := range sto {
		sortedStorage = append(sortedStorage, sortDisks(sto[i]))
	}
	return sortedStorage
}

func SortStorageByFile(sto storage) storage {
	var sortedStorage storage
	for i := range sto {
		sortedStorage = append(sortedStorage, sortFileSystem(sto[i]))
	}
	return sortedStorage
}

func sortFileSystem(d disk) disk {
	for i := len(d) - 1; i >= 0; i-- {
		for i2 := range d {
			if d[i].size <= d[i2].size {
				if d[i].size < d[i2].size {
					difference := d[i2].size - d[i].size
					temp := d[i]
					temp2 := file{
						id:   ".",
						size: difference,
					}
					d = insertValues(d, i2, temp, temp2)

					d[i2].size = d[i2].size - difference
					d[i] = d[i2]

				}
				if d[i].size == d[i2].size {
					temp := d[i]
					d[i] = d[i2]
					d[i2] = temp
				}
			}
		}
	}
	d = RecalcFreeSpace(d)
	return d
}

func insertValues(d disk, index int, files ...file) disk {
	if index < 0 || index > len(d) {
		panic("index out of range")
	}
	// Append the values at the specified index
	return append(d[:index], append(files, d[index:]...)...)
}

func RecalcFreeSpace(d disk) disk {
	var newD disk
	for i := range d {
		if d[i].id == "." {
			if i != len(d)-1 {
				if d[i+1].id == "." {
					f := file{
						id:   ".",
						size: d[i].size + d[i+1].size,
					}
					newD = append(newD, f)
					i++
				}
			}
		} else {
			newD = append(newD, d[i])
		}
	}
	return newD
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

func CalculateStorageByFileChecksum(sto storage) int {
	sum := 0
	for i := range sto {
		sum += calculateDiskByFileChecksum(sto[i])
	}
	return sum
}

func calculateDiskByFileChecksum(d disk) int {
	sum := 0
	printDisk(d)
	for i := range d {
		if d[i].id != "." {
			id, _ := strconv.Atoi(d[i].id)
			for j := i; j < i+int(d[i].size); j++ {
				result := j * id
				sum += result
				fmt.Println(j, "*", id, "=", result)
			}
			i = i + int(d[i].size)
		}
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

func ConverStorageRepresentationToFileRepresentation(sto storage) storage {
	fs := storage{}
	for i := range sto {
		d := disk{}
		currentIndex := ""
		for j := range sto[i] {
			currentIndex = sto[i][j].id
			counter := 0
			fmt.Println(currentIndex)
			for i2 := range sto {
				if sto[i][i2].id == currentIndex {
					counter++
				} else {
					break
				}
			}
			if j != len(sto[i])-1 {
				j += counter - 1
			}
			d = append(d, file{
				id:   currentIndex,
				size: filesize(counter),
			})
		}
		fs = append(fs, d)
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
	fmt.Println("----------------Disk------------\n")
	for i := range d {
		fmt.Print(d[i].id)
	}
	fmt.Println("\n--------------------------------\n")
}

func PrintStorage(sto storage) {
	for i := range sto {
		for i2 := range sto[i] {
			fmt.Print(sto[i][i2].id)
		}
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
