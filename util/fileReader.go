package util

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// Reading in files
func ReadFile(filePath string, delimiter string, useDelimiter bool) ([][]string, []string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error opening file: ", err)
		return nil, nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var data [][]string
	var altdata []string
	for scanner.Scan() {
		if useDelimiter {
			var stringArray []string = splitUpLine(scanner.Text(), delimiter)
			data = append(data, stringArray)
		}
		altdata = append(altdata, scanner.Text())

	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading file", err)
	}
	return data, altdata
}

func splitUpLine(s string, delimiter string) []string {
	var stringArray []string = strings.Split(s, delimiter)
	return stringArray
}
