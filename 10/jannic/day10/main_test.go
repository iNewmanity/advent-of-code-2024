package main

import (
	"awesomeProject/fileReader"
	"testing"
)

func TestAgentGeneration(t *testing.T) {
	path := "/home/jannic/GolandProjects/awesomeProject/day10/values_example.txt"
	values := fileReader.GetLineSplitListBySeparator(path, "")

	yMax := len(values) - 1
	xMax := len(values[0]) - 1

	valueHolder := ValueHolder{values, xMax, yMax, make([]*StartAgent, 0)}

	DoTheThing(&valueHolder)

	if len(valueHolder.StartAgents) != 9 {
		t.Error("Failed something")
	}
}

/*
func TestUnzip(t *testing.T) {
	strings.Split("123456", "")

	valueHolder := ValueHolder{strings.Split("123456", ""), make([]string, 0), make([]string, 0), 0}

	ExtractList(&valueHolder)
	targetString := strings.Split("0..111....22222......", "")

	if strings.Join(valueHolder.ExtractedInput, "") != strings.Join(targetString, "") {
		t.Error("Failed something")
	}
}

func TestUnzipExample(t *testing.T) {
	valueHolder := ValueHolder{strings.Split("2333133121414131402", ""), make([]string, 0), make([]string, 0), 0}

	ExtractList(&valueHolder)
	targetString := strings.Split("00...111...2...333.44.5555.6666.777.888899", "")

	if strings.Join(valueHolder.ExtractedInput, "") != strings.Join(targetString, "") {
		t.Error("Failed something")
	}
}

func TestCompressExample(t *testing.T) {

	valueHolder := ValueHolder{strings.Split("2333133121414131402", ""),
		strings.Split("00...111...2...333.44.5555.6666.777.888899", ""), make([]string, 0), 0}

	SortList(&valueHolder)

	targetString := "0099811188827773336446555566.............."

	if strings.Join(valueHolder.SortedInput, "") != targetString {
		t.Error("Failed something")
	}
}

func TestAllExample(t *testing.T) {
	valueHolder := ValueHolder{strings.Split("2333133121414131402", ""),
		make([]string, 0), make([]string, 0), 0}

	DoTheThing(&valueHolder)
	if valueHolder.Checksum != 1928 {
		t.Error("Failed something")
	}
}

*/
