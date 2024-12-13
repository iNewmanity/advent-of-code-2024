package day11

import (
	"strconv"
)

type Stone string

type number int

type Stones []Stone

func (physicsDefyingStones Stones) Blink() Stones {
	result := Stones{}
	for i := range physicsDefyingStones {
		result = append(result, physicsDefyingStones[i].applyRules()...)
	}
	return result
}

func (s Stone) applyRules() Stones {
	result := Stones{}
	if s.hasValOfZero() {
		result = append(result, "1")
		return result
	}
	if s.hasEvenNumbersOfDigits() {
		result = s.split()
		for i := range result {
			result[i] = result[i].cutLeadingZeros()
		}
		return result

	}
	if !s.hasEvenNumbersOfDigits() && !s.hasValOfZero() {
		multResult := s.multiplyBy2024()
		result = append(result, multResult)
	}
	return result
}

func (s Stone) hasEvenNumbersOfDigits() bool {
	if len(s)%2 == 0 {
		return true
	}
	return false
}

func (s Stone) hasValOfZero() bool {
	s = s.cutLeadingZeros()
	if s.toInt() == 0 {
		return true
	}
	return false
}

func (s Stone) cutLeadingZeros() Stone {
	return number(s.toInt()).toStone()
}

func (s Stone) multiplyBy2024() Stone {
	intValue := s.toInt()
	result := number(intValue * 2024).toStone()
	return result
}

func (i number) toStone() Stone {
	result := Stone(strconv.Itoa(int(i)))
	return result
}

func (s Stone) toInt() int {
	result, err := strconv.Atoi(string(s))
	if err != nil {
		panic(err)
	}
	return result
}

func (s Stone) split() Stones {
	length := len(s)
	mid := length / 2

	// Split into two parts
	firstPart := s[:mid]
	secondPart := s[mid:]

	result := Stones{}
	result = append(result, firstPart)
	result = append(result, secondPart)
	return result
}
