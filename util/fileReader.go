package util

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// Reading in files
func ReadFile(filePath string, delimiter string) [][]string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error opening file: ", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var data [][]string
	for scanner.Scan() {
		var stringArray []string = splitUpLine(scanner.Text(), delimiter)
		data = append(data, stringArray)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading file", err)
	}
	return data
}

func splitUpLine(s string, delimiter string) []string {
	var stringArray []string = strings.Split(s, delimiter)
	return stringArray
}
