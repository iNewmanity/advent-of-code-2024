package day10

import (
	"reflect"
	"testing"
)

type convertStringToCoordinatesTestData struct {
	input [][]string
	want  [][]coordinate
}

func TestConvertStringToCoordinates(t *testing.T) {
	tests := []convertStringToCoordinatesTestData{
		convertStringToCoordinatesTestData{
			input: [][]string{
				{"0", "1", "2"},
				{"3", "4", "5"},
			},
			want: [][]coordinate{
				{
					coordinate{
						value: "0",
						x:     0,
						y:     0,
					},
					coordinate{
						value: "1",
						x:     1,
						y:     0,
					},
					coordinate{
						value: "2",
						x:     2,
						y:     0,
					},
				},
				{
					coordinate{
						value: "3",
						x:     0,
						y:     1,
					},
					coordinate{
						value: "4",
						x:     1,
						y:     1,
					},
					coordinate{
						value: "5",
						x:     2,
						y:     1,
					},
				},
			},
		},
	}
	for _, test := range tests {
		got := ConvertStringToCoordinates(test.input)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("ConvertStringToCoordinates(%v) = %v; want %v", test.input, got, test.want)
		}
	}
}

type findTrailHeadsTestData struct {
	input [][]coordinate
	want  []trailhead
}

func TestFindTrailHeads(t *testing.T) {
	tests := []findTrailHeadsTestData{
		findTrailHeadsTestData{
			input: [][]coordinate{
				{
					coordinate{
						value: "0",
						x:     0,
						y:     0,
					},
					coordinate{
						value: "1",
						x:     1,
						y:     0,
					},
					coordinate{
						value: "2",
						x:     2,
						y:     0,
					},
				},
				{
					coordinate{
						value: "3",
						x:     0,
						y:     1,
					},
					coordinate{
						value: "4",
						x:     1,
						y:     1,
					},
					coordinate{
						value: "5",
						x:     2,
						y:     1,
					},
				},
			},
			want: []trailhead{
				{
					start: coordinate{
						value: "0",
						x:     0,
						y:     0,
					},
				},
			},
		},
	}
	for _, test := range tests {
		got := FindTrailHeads(test.input)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("FindTrailHeads(%v) = %v; want %v", test.input, got, test.want)
		}
	}
}

type calculateScoreTestData struct {
	input []trailhead
	want  []trailhead
}

func TestCalculateScore(t *testing.T) {
	tests := []calculateScoreTestData{
		calculateScoreTestData{
			input: []trailhead{
				trailhead{
					start: coordinate{
						value: "0",
						x:     0,
						y:     0,
					},
					trails: []trail{
						trail{
							waypoints: []coordinate{
								coordinate{
									value: "0",
									x:     0,
									y:     0,
								},
								coordinate{
									value: "1",
									x:     1,
									y:     0,
								},
								coordinate{
									value: "2",
									x:     2,
									y:     0,
								},
							},
						},
						trail{
							waypoints: []coordinate{
								coordinate{
									value: "0",
									x:     0,
									y:     0,
								},
								coordinate{
									value: "1",
									x:     0,
									y:     1,
								},
								coordinate{
									value: "2",
									x:     0,
									y:     2,
								},
							},
						},
					},
				},
			},
			want: []trailhead{
				trailhead{
					start: coordinate{
						value: "0",
						x:     0,
						y:     0,
					},
					trails: []trail{
						trail{
							waypoints: []coordinate{
								coordinate{
									value: "0",
									x:     0,
									y:     0,
								},
								coordinate{
									value: "1",
									x:     1,
									y:     0,
								},
								coordinate{
									value: "2",
									x:     2,
									y:     0,
								},
							},
						},
						trail{
							waypoints: []coordinate{
								coordinate{
									value: "0",
									x:     0,
									y:     0,
								},
								coordinate{
									value: "1",
									x:     0,
									y:     1,
								},
								coordinate{
									value: "2",
									x:     0,
									y:     2,
								},
							},
						},
					},
					score: 2,
				},
			},
		},
	}
	for _, test := range tests {
		got := calculateScore(test.input)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("CalculateScore(%v) = %v; want %v", test.input, got, test.want)
		}
	}
}

type calculateSumTestData struct {
	input []trailhead
	want  int
}

func TestCalculateSum(t *testing.T) {
	tests := []calculateSumTestData{
		calculateSumTestData{
			input: []trailhead{
				{
					score: 3,
				},
				{
					score: 10,
				},
				{
					score: 5,
				},
			},
			want: 18,
		},
	}
	for _, test := range tests {
		got := CalculateSum(test.input)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("CalculateSum(%v) = %v; want %v", test.input, got, test.want)
		}
	}
}
