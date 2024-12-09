package day08

import (
	"reflect"
	"testing"
)

type calculateDistancesTestData struct {
	input     []coordinate
	unlimited bool
	runs      int
	want      []distance
}

func TestCalculateDistances(t *testing.T) {
	tests := []calculateDistancesTestData{
		calculateDistancesTestData{
			input: []coordinate{
				coordinate{x: 2, y: 1, letter: "A"},
				coordinate{x: 1, y: 2, letter: "A"},
				coordinate{x: 3, y: 1, letter: "b"},
				coordinate{x: 2, y: 2, letter: "b"},
			},
			unlimited: false,
			runs:      5,
			want: []distance{
				distance{
					c1: coordinate{x: 2, y: 1, letter: "A"},
					c2: coordinate{x: 1, y: 2, letter: "A"},
					x:  -1,
					y:  1,
				},
				distance{
					c1: coordinate{x: 1, y: 2, letter: "A"},
					c2: coordinate{x: 2, y: 1, letter: "A"},
					x:  1,
					y:  -1,
				},
				distance{
					c1: coordinate{x: 3, y: 1, letter: "b"},
					c2: coordinate{x: 2, y: 2, letter: "b"},
					x:  -1,
					y:  1,
				},
				distance{
					c1: coordinate{x: 2, y: 2, letter: "b"},
					c2: coordinate{x: 3, y: 1, letter: "b"},
					x:  1,
					y:  -1,
				},
			},
		},
	}
	for _, test := range tests {
		got := calculateDistances(test.input, test.unlimited, test.runs)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("calculateDistances(%v) = %v; want %v", test.input, got, test.want)
		}
	}
}

type calculateDistancesForLetterTestData struct {
	input     []coordinate
	unlimited bool
	runs      int
	want      []distance
}

func TestCalculateDistancesForLetter(t *testing.T) {
	tests := []calculateDistancesForLetterTestData{
		calculateDistancesForLetterTestData{
			input: []coordinate{
				coordinate{x: 2, y: 1, letter: "A"},
				coordinate{x: 1, y: 2, letter: "A"},
			},
			unlimited: false,
			runs:      5,
			want: []distance{
				distance{
					c1: coordinate{x: 2, y: 1, letter: "A"},
					c2: coordinate{x: 1, y: 2, letter: "A"},
					x:  -1,
					y:  1,
				},
				distance{
					c1: coordinate{x: 1, y: 2, letter: "A"},
					c2: coordinate{x: 2, y: 1, letter: "A"},
					x:  1,
					y:  -1,
				},
			},
		},
	}
	for _, test := range tests {
		got := calculateDistancesForLetter(test.input, test.unlimited, test.runs)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("calculateDistancesForLetter(%v) = %v; want %v", test.input, got, test.want)
		}
	}
}

type createAntinodesTestData struct {
	input  []distance
	height int
	width  int
	want   []coordinate
}

func TestCreateAntinodes(t *testing.T) {
	tests := []createAntinodesTestData{
		createAntinodesTestData{
			input: []distance{
				distance{
					c1: coordinate{x: 2, y: 1, letter: "A"},
					c2: coordinate{x: 1, y: 2, letter: "A"},
					x:  -1,
					y:  1,
				},
				distance{
					c1: coordinate{x: 1, y: 2, letter: "A"},
					c2: coordinate{x: 2, y: 1, letter: "A"},
					x:  1,
					y:  -1,
				},
				distance{
					c1: coordinate{x: 3, y: 1, letter: "b"},
					c2: coordinate{x: 2, y: 2, letter: "b"},
					x:  -1,
					y:  1,
				},
				distance{
					c1: coordinate{x: 2, y: 2, letter: "b"},
					c2: coordinate{x: 3, y: 1, letter: "b"},
					x:  1,
					y:  -1,
				},
			},
			height: 5,
			width:  5,
			want: []coordinate{
				coordinate{x: 0, y: 3, letter: "#"},
				coordinate{x: 3, y: 0, letter: "#"},
				coordinate{x: 1, y: 3, letter: "#"},
				coordinate{x: 4, y: 0, letter: "#"},
			},
		},
		createAntinodesTestData{
			input: []distance{
				distance{
					c1: coordinate{x: 2, y: 1, letter: "A"},
					c2: coordinate{x: 1, y: 2, letter: "A"},
					x:  -1,
					y:  1,
				},
				distance{
					c1: coordinate{x: 1, y: 2, letter: "A"},
					c2: coordinate{x: 2, y: 1, letter: "A"},
					x:  1,
					y:  -1,
				},
				distance{
					c1: coordinate{x: 3, y: 1, letter: "b"},
					c2: coordinate{x: 2, y: 2, letter: "b"},
					x:  -1,
					y:  1,
				},
				distance{
					c1: coordinate{x: 2, y: 2, letter: "b"},
					c2: coordinate{x: 3, y: 1, letter: "b"},
					x:  1,
					y:  -1,
				},
				distance{
					c1: coordinate{x: 1, y: 0, letter: "C"},
					c2: coordinate{x: 4, y: 0, letter: "C"},
					x:  3,
					y:  0,
				},
				distance{
					c1: coordinate{x: 4, y: 0, letter: "C"},
					c2: coordinate{x: 1, y: 0, letter: "C"},
					x:  -3,
					y:  0,
				},
			},
			height: 5,
			width:  5,
			want: []coordinate{
				coordinate{x: 0, y: 3, letter: "#"},
				coordinate{x: 3, y: 0, letter: "#"},
				coordinate{x: 1, y: 3, letter: "#"},
				coordinate{x: 4, y: 0, letter: "#"},
			},
		},
	}
	for _, test := range tests {
		got := createAntinodes(test.input, test.height, test.width)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("createAntinodes got %v want %v", got, test.want)
		}
	}
}

type isAntinodeCorrectTestData struct {
	input     coordinate
	antinodes []coordinate
	height    int
	width     int
	want      bool
}

func TestIsAntiNodeCorrect(t *testing.T) {
	tests := []isAntinodeCorrectTestData{
		isAntinodeCorrectTestData{
			input: coordinate{x: 0, y: 0, letter: "#"},
			antinodes: []coordinate{
				{x: 1, y: 0, letter: "."},
				{x: 2, y: 0, letter: "a"},
				{x: 0, y: 1, letter: "A"},
				{x: 1, y: 1, letter: "."},
				{x: 2, y: 1, letter: "."},
				{x: 0, y: 2, letter: "b"},
				{x: 1, y: 2, letter: "."},
				{x: 2, y: 2, letter: "b"},
			},
			height: 3,
			width:  3,
			want:   true,
		},
		isAntinodeCorrectTestData{
			input: coordinate{x: 0, y: 0, letter: "#"},
			antinodes: []coordinate{
				{x: 0, y: 0, letter: "#"},
				{x: 1, y: 0, letter: "."},
				{x: 2, y: 0, letter: "a"},
				{x: 0, y: 1, letter: "A"},
				{x: 1, y: 1, letter: "."},
				{x: 2, y: 1, letter: "."},
				{x: 0, y: 2, letter: "b"},
				{x: 1, y: 2, letter: "."},
				{x: 2, y: 2, letter: "b"},
			},
			height: 3,
			width:  3,
			want:   false,
		},
		isAntinodeCorrectTestData{
			input: coordinate{x: 2, y: 0, letter: "#"},
			antinodes: []coordinate{
				{x: 0, y: 0, letter: "#"},
				{x: 1, y: 0, letter: "."},
				{x: 0, y: 1, letter: "A"},
				{x: 1, y: 1, letter: "."},
				{x: 2, y: 1, letter: "."},
				{x: 0, y: 2, letter: "b"},
				{x: 1, y: 2, letter: "."},
				{x: 2, y: 2, letter: "b"},
			},
			height: 3,
			width:  3,
			want:   true,
		},
		isAntinodeCorrectTestData{
			input: coordinate{x: 3, y: 3, letter: "#"},
			antinodes: []coordinate{
				{x: 0, y: 0, letter: "#"},
				{x: 1, y: 0, letter: "."},
				{x: 2, y: 0, letter: "a"},
				{x: 0, y: 1, letter: "A"},
				{x: 1, y: 1, letter: "."},
				{x: 2, y: 1, letter: "."},
				{x: 0, y: 2, letter: "b"},
				{x: 1, y: 2, letter: "."},
				{x: 2, y: 2, letter: "b"},
			},
			height: 3,
			width:  3,
			want:   false,
		},
	}
	for _, test := range tests {
		got := isAntinodeCorrect(test.input, test.antinodes, test.height, test.width)
		if got != test.want {
			t.Errorf("isAntinodeCorrect(%v, %v, %d, %d) = %v; want %v", test.input, test.antinodes, test.height, test.width, got, test.want)
		}
	}
}

type getLettersTestData struct {
	input []coordinate
	want  []string
}

func TestGetLetters(t *testing.T) {
	tests := []getLettersTestData{
		getLettersTestData{
			input: []coordinate{
				{x: 0, y: 0, letter: "."},
				{x: 1, y: 0, letter: "."},
				{x: 2, y: 0, letter: "a"},
				{x: 0, y: 1, letter: "A"},
				{x: 1, y: 1, letter: "."},
				{x: 2, y: 1, letter: "."},
				{x: 0, y: 2, letter: "b"},
				{x: 1, y: 2, letter: "."},
				{x: 2, y: 2, letter: "b"},
			},
			want: []string{"a", "A", "b"},
		},
	}
	for _, test := range tests {
		got := getLetters(test.input)
		if !reflect.DeepEqual(test.want, got) {
			t.Errorf("getLetters(%v) = %v; want %v", test.input, got, test.want)
		}
	}
}

type getCoordinatesTestData struct {
	input [][]string
	want  []coordinate
}

func TestGetCoordinates(t *testing.T) {
	tests := []getCoordinatesTestData{
		getCoordinatesTestData{
			input: [][]string{
				{".", ".", "a"},
				{"A", ".", "."},
				{"b", ".", "b"},
			},
			want: []coordinate{
				{x: 0, y: 0, letter: "."},
				{x: 1, y: 0, letter: "."},
				{x: 2, y: 0, letter: "a"},
				{x: 0, y: 1, letter: "A"},
				{x: 1, y: 1, letter: "."},
				{x: 2, y: 1, letter: "."},
				{x: 0, y: 2, letter: "b"},
				{x: 1, y: 2, letter: "."},
				{x: 2, y: 2, letter: "b"},
			},
		},
	}
	for _, test := range tests {
		got := getCoordinates(test.input)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("getCoordinates(%v) = %v; want %v", test.input, got, test.want)
		}
	}
}

type filterCoordinatesTestData struct {
	input   []coordinate
	letters []string
	want    []coordinate
}

func TestFilterCoordinates(t *testing.T) {
	tests := []filterCoordinatesTestData{
		filterCoordinatesTestData{
			input:   []coordinate{coordinate{letter: "a"}, coordinate{letter: "b"}, coordinate{letter: "."}, coordinate{letter: "."}, coordinate{letter: "a"}, coordinate{letter: "b"}, coordinate{letter: "."}},
			letters: []string{"a", "b"},
			want:    []coordinate{coordinate{letter: "a"}, coordinate{letter: "b"}, coordinate{letter: "a"}, coordinate{letter: "b"}},
		},
	}
	for i := range tests {
		got := filterCoordinates(tests[i].input, tests[i].letters)
		if !reflect.DeepEqual(got, tests[i].want) {
			t.Errorf("tests[%d] want %v, got %v", i, tests[i].want, got)
		}
	}
}

type extrudeLetterTestData struct {
	input  []coordinate
	letter string
	want   []coordinate
}

func TestExtrudeLetter(t *testing.T) {
	tests := []extrudeLetterTestData{
		extrudeLetterTestData{
			input:  []coordinate{coordinate{letter: "a"}, coordinate{letter: "b"}, coordinate{letter: "c"}, coordinate{letter: "a"}, coordinate{letter: "b"}, coordinate{letter: "c"}},
			letter: "a",
			want:   []coordinate{coordinate{letter: "a"}, coordinate{letter: "a"}},
		},
	}
	for _, test := range tests {
		got := extrudeLetter(test.input, test.letter)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("extrudeLetter(%v,%v) = %v; want %v", test.input, test.letter, got, test.want)
		}
	}
}

type calculateDistanceTestData struct {
	input1 coordinate
	input2 coordinate
	want   distance
}

func TestCalculateDistance(t *testing.T) {
	tests := []calculateDistanceTestData{
		calculateDistanceTestData{
			input1: coordinate{x: 2, y: 2, letter: "A"},
			input2: coordinate{x: 1, y: 1, letter: "A"},
			want: distance{
				c1: coordinate{x: 2, y: 2, letter: "A"},
				c2: coordinate{x: 1, y: 1, letter: "A"},
				x:  -1,
				y:  -1,
			},
		},
		calculateDistanceTestData{
			input1: coordinate{x: 0, y: 0, letter: "a"},
			input2: coordinate{x: 5, y: 7, letter: "a"},
			want: distance{
				c1: coordinate{x: 0, y: 0, letter: "a"},
				c2: coordinate{x: 5, y: 7, letter: "a"},
				x:  5,
				y:  7,
			},
		},
		calculateDistanceTestData{
			input1: coordinate{x: 3, y: 3, letter: "a"},
			input2: coordinate{x: 2, y: 1, letter: "a"},
			want: distance{
				c1: coordinate{x: 3, y: 3, letter: "a"},
				c2: coordinate{x: 2, y: 1, letter: "a"},
				x:  -1,
				y:  -2,
			},
		},
		calculateDistanceTestData{
			input1: coordinate{x: 3, y: 8, letter: "a"},
			input2: coordinate{x: 4, y: 7, letter: "a"},
			want: distance{
				c1: coordinate{x: 3, y: 8, letter: "a"},
				c2: coordinate{x: 4, y: 7, letter: "a"},
				x:  1,
				y:  -1,
			},
		},
	}
	for _, test := range tests {
		got := calculateDistance(test.input1, test.input2)
		if got.x != test.want.x && got.y != test.want.y {
			t.Errorf("calculateDistance(%v, %v) = %v, want %v", test.input1, test.input2, test.want, got)
		}
	}

}

type createAntiNodeTestData struct {
	input distance
	want  coordinate
}

func TestCreateAntinode(t *testing.T) {
	tests := []createAntiNodeTestData{
		createAntiNodeTestData{input: distance{
			c1: coordinate{x: 2, y: 2, letter: "A"},
			c2: coordinate{x: 1, y: 1, letter: "A"},
			x:  -1,
			y:  -1,
		},
			want: coordinate{x: 0, y: 0, letter: "#"}},
		createAntiNodeTestData{input: distance{
			c1: coordinate{x: 0, y: 0, letter: "a"},
			c2: coordinate{x: 5, y: 7, letter: "a"},
			x:  5,
			y:  7,
		},
			want: coordinate{x: 10, y: 14, letter: "#"}},
	}
	for _, test := range tests {
		got := createAntinode(test.input)
		if got.x != test.want.x && got.y != test.want.y && got.letter != test.want.letter {
			t.Errorf("createAntinode(%v) = %v; want %v", test.input, got, test.want)
		}
	}
}

type isInBoundsTestData struct {
	input  coordinate
	height int
	width  int
	want   bool
}

func TestIsInBounds(t *testing.T) {
	tests := []isInBoundsTestData{
		isInBoundsTestData{
			input:  coordinate{x: 0, y: 0, letter: "#"},
			height: 7,
			width:  5,
			want:   true,
		},
		isInBoundsTestData{
			input:  coordinate{x: -2, y: 6, letter: "#"},
			height: 7,
			width:  5,
			want:   false,
		},
		isInBoundsTestData{
			input:  coordinate{x: -2, y: 8, letter: "#"},
			height: 7,
			width:  5,
			want:   false,
		},
		isInBoundsTestData{
			input:  coordinate{x: 77, y: 4, letter: "#"},
			height: 7,
			width:  5,
			want:   false,
		},
		isInBoundsTestData{
			input:  coordinate{x: 3, y: 4, letter: "#"},
			height: 7,
			width:  5,
			want:   true,
		},
	}
	for _, test := range tests {
		got := isInBounds(test.input, test.height, test.width)
		if got != test.want {
			t.Errorf("isInBounds(%v, %v, %v): got %v, want %v", test.input, test.height, test.width, got, test.want)
		}
	}

}

type checkIfFreeTestData struct {
	input     coordinate
	antinodes []coordinate
	want      bool
}

func TestCheckIfFree(t *testing.T) {
	tests := []checkIfFreeTestData{
		checkIfFreeTestData{
			input: coordinate{x: 0, y: 0, letter: "#"},
			antinodes: []coordinate{
				coordinate{x: 1, y: 0, letter: "."},
				coordinate{x: 0, y: 1, letter: "A"},
				coordinate{x: 1, y: 1, letter: "."},
			},
			want: true,
		},
		checkIfFreeTestData{
			input: coordinate{x: 0, y: 0, letter: "#"},
			antinodes: []coordinate{
				coordinate{x: 0, y: 0, letter: "#"},
				coordinate{x: 1, y: 0, letter: "."},
				coordinate{x: 0, y: 1, letter: "A"},
				coordinate{x: 1, y: 1, letter: "."},
			},
			want: false,
		},
	}
	for _, test := range tests {
		got := checkIfFree(test.input, test.antinodes)
		if got != test.want {
			t.Errorf("checkIfFree(%v, %v) = %v, want %v", test.input, test.antinodes, got, test.want)
		}
	}

}

type countAntiNodesTestData struct {
	input []coordinate
	want  int
}

func TestCountAntinodes(t *testing.T) {
	tests := []countAntiNodesTestData{
		countAntiNodesTestData{input: []coordinate{{x: 1, y: 2, letter: "#"}, {x: 4, y: 5, letter: "#"}, {x: 6, y: 7, letter: "#"}}, want: 3},
		countAntiNodesTestData{input: []coordinate{{x: 1, y: 2, letter: "#"}, {x: 4, y: 5, letter: "#"}, {x: 6, y: 7, letter: "#"}, {x: 8, y: 9, letter: "#"}, {x: 10, y: 9, letter: "#"}}, want: 5},
		countAntiNodesTestData{input: []coordinate{{x: 1, y: 2, letter: "#"}, {x: 4, y: 5, letter: "A"}, {x: 6, y: 7, letter: "#"}, {x: 8, y: 9, letter: "#"}, {x: 10, y: 9, letter: "#"}}, want: 4},
	}
	for _, test := range tests {
		got := countAntinodes(test.input)
		if got != test.want {
			t.Errorf("countAntinodes(%v) = %v; want %v", test.input, got, test.want)
		}
	}
}
