package main

import (
	"advent-of-code-2024/util"
	"fmt"
	"strconv"
	"strings"
)

type ValueHolder struct {
	InitialInput   []string
	ExtractedInput []string
	SortedInput    []string
	Checksum       int
	CharSets       []*CharSet
}

func main() {
	path := "/home/janneumann/Dokumente/Daten/Projekte/Privat/advent-of-code-2024/9/input.txt"
	valuesDA, _ := util.ReadFile(path, "", true)
	values := valuesDA[0]

	valueHolder := ValueHolder{values, make([]string, 0), make([]string, 0), 0, make([]*CharSet, 0)}

	DoTheThing(&valueHolder)

	fmt.Println(valueHolder.Checksum)
}

func DoTheThing(valueHolder *ValueHolder) {
	ExtractList(valueHolder)
	PutInEmptySpace(valueHolder)
	CalculateChecksum(valueHolder)
}

func CalculateChecksum(valueHolder *ValueHolder) {
	for i := range valueHolder.SortedInput {
		if valueHolder.SortedInput[i] == "." {
			continue
		}

		valueHolder.Checksum += castStringToInt(valueHolder.SortedInput[i]) * i
	}
}

func ExtractList(valueHolder *ValueHolder) {
	for i := range valueHolder.InitialInput {
		value := castStringToInt(valueHolder.InitialInput[i])

		for _ = range value {
			if i%2 == 0 {
				valueHolder.ExtractedInput = append(valueHolder.ExtractedInput, castIntToString(i/2))
			} else {
				valueHolder.ExtractedInput = append(valueHolder.ExtractedInput, ".")
			}
		}
	}
}

func PutInEmptySpace(valueHolder *ValueHolder) {
	createCharSets(valueHolder)
	sort(valueHolder)
	setSortedInput(valueHolder)
}

func sort(holder *ValueHolder) {
	index := len(holder.CharSets) - 1

	x := 0
	for x < 1 {
		if index < 0 || areAllSlicesMoved(holder) {
			x = 42
			continue
		}

		numberCharSet := holder.CharSets[index]
		if numberCharSet.Char == "." || numberCharSet.WasMoved == true {
			index--
			continue
		}

		numberCharSet.WasMoved = true

		for j := 0; j < index; j++ {
			emptyCharSet := holder.CharSets[j]

			if emptyCharSet.Char != "." {
				continue
			}

			if emptyCharSet.Count == numberCharSet.Count {
				emptyCharSet.Char = numberCharSet.Char
				numberCharSet.Char = "."
				index = len(holder.CharSets) - 1
				break
			}

			if emptyCharSet.Count > numberCharSet.Count {
				numberCharSetLength := numberCharSet.Count
				insertCharSet := numberCharSet.Clone()
				numberCharSet.Char = "."
				InsertAt(holder, j, insertCharSet)
				emptyCharSet.Count -= numberCharSetLength

				index = len(holder.CharSets) - 1
				break
			}
		}

		index--
	}
}

func InsertAt(valueHolder *ValueHolder, index int, charSet *CharSet) {
	valueHolder.CharSets = append(valueHolder.CharSets[:index], append([]*CharSet{charSet}, valueHolder.CharSets[index:]...)...)
}

func setSortedInput(valueHolder *ValueHolder) {
	valueHolder.SortedInput = make([]string, 0)
	for i := range valueHolder.CharSets {
		slice := valueHolder.CharSets[i]

		for i := 0; i < slice.Count; i++ {
			valueHolder.SortedInput = append(valueHolder.SortedInput, slice.Char)
		}
	}

	println(strings.Join(valueHolder.SortedInput, ""))
}

func areAllSlicesMoved(valueHolder *ValueHolder) bool {
	notMovedSliceCount := 0

	for i := range valueHolder.CharSets {
		if valueHolder.CharSets[i].Char == "." {
			continue
		}

		if !valueHolder.CharSets[i].WasMoved {
			notMovedSliceCount++
		}
	}

	return notMovedSliceCount == 0
}

func createCharSets(valueHolder *ValueHolder) {
	currentChar := valueHolder.ExtractedInput[0]
	charCount := 1
	for i := 1; i < len(valueHolder.ExtractedInput); i++ {
		if valueHolder.ExtractedInput[i] == currentChar {
			charCount++
			if i == len(valueHolder.ExtractedInput)-1 {
				if currentChar == "." {
					valueHolder.CharSets = append(valueHolder.CharSets, NewCharSet(charCount, currentChar, true))
				} else {
					valueHolder.CharSets = append(valueHolder.CharSets, NewCharSet(charCount, currentChar, false))
				}
			}
			continue
		}

		valueHolder.CharSets = append(valueHolder.CharSets, NewCharSet(charCount, currentChar, false))

		currentChar = valueHolder.ExtractedInput[i]
		charCount = 1
	}
}

func castStringToInt(input string) int {
	atoi, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}

	return atoi
}

func castIntToString(input int) string {
	return strconv.Itoa(input)
}
