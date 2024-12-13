package main

import (
	"advent-of-code-2024/util"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

func getNumStones(s []string, times int) int {
	strStones := strings.Split(s[0], " ")
	stones := make([]int, len(strStones))
	for i, strStone := range strStones {
		stones[i], _ = strconv.Atoi(strStone)
	}
	for i := 0; i < times; i++ {
		var newStones []int
		for _, stone := range stones {
			if stone == 0 {
				newStones = append(newStones, 1)
			} else if digits := strconv.Itoa(stone); len(digits)%2 == 0 {
				a, _ := strconv.Atoi(digits[:len(digits)/2])
				b, _ := strconv.Atoi(digits[len(digits)/2:])
				newStones = append(newStones, a)
				newStones = append(newStones, b)
			} else {
				newStones = append(newStones, stone*2024)
			}
		}
		stones = newStones
	}
	return len(stones)
}

func getNumStones2(s []string, times int) int {
	strStones := strings.Split(s[0], " ")
	stoneMap := make(map[int]int)
	for _, strStone := range strStones {
		value, _ := strconv.Atoi(strStone)
		stoneMap[value] = stoneMap[value] + 1
	}
	for i := 0; i < times; i++ {
		newStoneMap := make(map[int]int)
		for stone, count := range stoneMap {
			var newStones []int
			if stone == 0 {
				newStones = append(newStones, 1)
			} else if digits := strconv.Itoa(stone); len(digits)%2 == 0 {
				a, _ := strconv.Atoi(digits[:len(digits)/2])
				b, _ := strconv.Atoi(digits[len(digits)/2:])
				newStones = append(newStones, a)
				newStones = append(newStones, b)
			} else {
				newStones = append(newStones, stone*2024)
			}
			for _, newStone := range newStones {
				newStoneMap[newStone] = newStoneMap[newStone] + count // by the number of times that stone appears
			}
		}
		stoneMap = newStoneMap
	}
	result := 0
	for _, count := range stoneMap {
		result += count
	}
	return result
}

func getNumStonesAfter(stone, times int, checked *sync.Map) int {
	result := 0
	if times > 0 {
		if v, ok := checked.Load(fmt.Sprintf("%d:%d", stone, times)); ok {
			result += v.(int)
		} else {

			if stone == 0 {
				result += getNumStonesAfter(1, times-1, checked)
			} else if digits := strconv.Itoa(stone); len(digits)%2 == 0 {
				a, _ := strconv.Atoi(digits[:len(digits)/2])
				b, _ := strconv.Atoi(digits[len(digits)/2:])
				result += getNumStonesAfter(a, times-1, checked)
				result += getNumStonesAfter(b, times-1, checked)
			} else {
				result += getNumStonesAfter(stone*2024, times-1, checked)
			}

			checked.Store(fmt.Sprintf("%d:%d", stone, times), result)
		}
	} else {
		return 1
	}
	return result
}

func getNumStones3(s []string, times int) int {
	wg := sync.WaitGroup{}
	checked := &sync.Map{}
	strStones := strings.Split(s[0], " ")
	stones := make([]int, len(strStones))
	ch := make(chan int, len(stones))
	for i, strStone := range strStones {
		wg.Add(1)
		stones[i], _ = strconv.Atoi(strStone)
		go func(stone int) {
			defer wg.Done()
			ch <- getNumStonesAfter(stone, times, checked)
		}(stones[i])
	}
	wg.Wait()
	close(ch)
	result := 0
	for value := range ch {
		result += value
	}
	return result
}

func main() {
	absPathName, _ := filepath.Abs("/home/janneumann/Dokumente/Daten/Projekte/Privat/advent-of-code-2024/11/input.txt")
	output, _ := util.ReadFile(absPathName, " ", true)

	newOutput := []string{}
	for i := range output {
		for i2 := range output[i] {
			newOutput = append(newOutput, string(output[i][i2]))
		}
	}

	fmt.Println(getNumStones(newOutput, 25))
	fmt.Println(getNumStones2(newOutput, 75))
	fmt.Println(getNumStones3(newOutput, 75))
}
