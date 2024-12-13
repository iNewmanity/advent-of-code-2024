package main

import (
	"advent-of-code-2024/11/day11"
	"advent-of-code-2024/util"
	"fmt"
	"runtime"
	"sync"
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
	pDS = blinkNTimesO(pDS, 10)
	runtime.GC()
	pDS = blinkNTimesO(pDS, 10)
	runtime.GC()
	pDS = blinkNTimesO(pDS, 10)
	runtime.GC()
	pDS = blinkNTimesO(pDS, 10)
	runtime.GC()
	pDS = blinkNTimesO(pDS, 10)
	runtime.GC()
	pDS = blinkNTimesO(pDS, 10)
	runtime.GC()
	pDS = blinkNTimesO(pDS, 10)
	runtime.GC()
	pDS = blinkNTimesO(pDS, 5)
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

func blinkNTimesO(
	pDS day11.Stones,
	num int,
) day11.Stones {
	chunkSize := 10000                        // Example chunk size; adjust based on performance tests.
	maxGoroutines := 10                       // Limit the number of concurrent goroutines.
	sem := make(chan struct{}, maxGoroutines) // Semaphore channel to control concurrency.

	for i := 0; i < num; i++ {
		start := time.Now()

		// Split pDS into chunks
		var chunks []day11.Stones
		for startIdx := 0; startIdx < len(pDS); startIdx += chunkSize {
			endIdx := startIdx + chunkSize
			if endIdx > len(pDS) {
				endIdx = len(pDS)
			}
			chunks = append(chunks, pDS[startIdx:endIdx])
		}

		// Blink each chunk in parallel using goroutines
		var updatedChunks []day11.Stones
		var wg sync.WaitGroup
		var mu sync.Mutex // Mutex to protect updatedChunks during appends.

		for _, chunk := range chunks {
			wg.Add(1)
			sem <- struct{}{} // Acquire a slot in the semaphore.

			go func(chunk day11.Stones) {
				defer wg.Done()
				defer func() { <-sem }() // Release the slot in the semaphore.

				blinkedChunk := chunk.Blink()

				// Safely append the result to updatedChunks.
				mu.Lock()
				updatedChunks = append(updatedChunks, blinkedChunk)
				mu.Unlock()
			}(chunk)
		}

		// Wait for all goroutines to finish.
		wg.Wait()

		// Merge the updated chunks back into pDS.
		pDS = nil // Clear previous pDS to manage memory
		for _, chunk := range updatedChunks {
			pDS = append(pDS, chunk...)
		}

		elapsed := time.Since(start)
		fmt.Println("Iteration: \t", i, "\tTime: \t", elapsed, "\tResulting in:", len(pDS), "Stones")
	}
	return pDS
}
