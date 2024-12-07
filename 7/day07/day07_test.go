package day07

import (
	"testing"
)

type resultExtractorTestData struct {
	input string
	want  int
}

type numberExtractorTestData struct {
	input string
	want  []int
}

type stringToIntTestData struct {
	input string
	want  int
}

type assignmentCreatorTestData struct {
	input []string
	want  []Assignment
}

type allowedOperatorTestData struct {
	input   operator
	allowed []operator
	want    bool
}

type operatorEvaluatorTestData struct {
	input Assignment
	want  bool
}

type totalCalibrationResultTestData struct {
	input []string
	want  int
}

func TestGetTotalCalibrationResult(t *testing.T) {
	tests := []totalCalibrationResultTestData{
		totalCalibrationResultTestData{
			input: []string{
				"190: 10 19",
				"3267: 81 40 27",
				"83: 17 5",
				"156: 15 6",
				"7290: 6 8 6 15",
				"161011: 16 10 13",
				"192: 17 8 14",
				"21037: 9 7 18 13",
				"292: 11 6 16 20",
			},
			want: 3749,
		},
	}
	for _, test := range tests {
		got := GetTotalCalibrationResult(test.input)
		if got != test.want {
			t.Errorf("getTotalCalibrationResult(%q) = %d, want %d", test.input, got, test.want)
		}
	}
}

func TestOperatorEvaluator(t *testing.T) {
	tests := []operatorEvaluatorTestData{
		operatorEvaluatorTestData{input: Assignment{result: 190, numbers: []int{10, 19}}, want: true},
		operatorEvaluatorTestData{input: Assignment{result: 3267, numbers: []int{81, 40, 27}}, want: true},
		operatorEvaluatorTestData{input: Assignment{result: 292, numbers: []int{11, 6, 16, 20}}, want: true},
		operatorEvaluatorTestData{input: Assignment{result: 83, numbers: []int{17, 5}}, want: false},
	}
	for _, test := range tests {
		got := operatorEvaluator(test.input)
		if got != test.want {
			t.Errorf("operatorEvaluator(%v) = %v, want %v", test.input, got, test.want)
		}
	}
}

func TestAllowedOperators(t *testing.T) {
	tests := []allowedOperatorTestData{
		allowedOperatorTestData{input: "+", allowed: []operator{"+", "*"}, want: true},
		allowedOperatorTestData{input: "-", allowed: []operator{"+", "*"}, want: false},
		allowedOperatorTestData{input: "*", allowed: []operator{"+", "*"}, want: true},
		allowedOperatorTestData{input: "/", allowed: []operator{"+", "*"}, want: false},
	}
	for _, test := range tests {
		got := isAllowedOperator(test.input, test.allowed)
		if got != test.want {
			t.Errorf("input: %q, allowed: %v, got: %v, want: %v", test.input, test.allowed, got, test.want)
		}
	}
}

func TestAssignmentCreator(t *testing.T) {
	tests := []assignmentCreatorTestData{
		assignmentCreatorTestData{input: []string{"190: 10 19", "3267: 81 40 27"}, want: []Assignment{{result: 190, numbers: []int{10, 19}}, {result: 3267, numbers: []int{81, 40, 27}}}},
	}
	for _, test := range tests {
		got := assignmentCreator(test.input)
		for i := range got {
			if got[i].numbers[i] != test.want[i].numbers[i] && got[i].result != test.want[i].result {
				t.Errorf("input: %s, got: %v, want: %v", test.input, got, test.want)
			}
		}
	}
}

func TestResultExtractor(t *testing.T) {
	tests := []resultExtractorTestData{
		resultExtractorTestData{input: "190: 10 19", want: 190},
		resultExtractorTestData{input: "3267: 81 40 27", want: 3267},
		resultExtractorTestData{input: "83: 17 5", want: 83},
		resultExtractorTestData{input: "156: 15 6", want: 156},
		resultExtractorTestData{input: "7290: 6 8 6 15", want: 7290},
		resultExtractorTestData{input: "161011: 16 10 13", want: 161011},
		resultExtractorTestData{input: "192: 17 8 14", want: 192},
		resultExtractorTestData{input: "21037: 9 7 18 13", want: 21037},
		resultExtractorTestData{input: "292: 11 6 16 20", want: 292},
	}
	for _, test := range tests {
		got := resultExtractor(test.input)
		if got != test.want {
			t.Errorf("input: %s\nwant: %d\ngot: %d", test.input, test.want, got)
		}
	}
}

func TestOperatorExtractor(t *testing.T) {
	tests := []numberExtractorTestData{
		numberExtractorTestData{input: "190: 10 19", want: []int{10, 19}},
		numberExtractorTestData{input: "3267: 81 40 27", want: []int{81, 40, 27}},
		numberExtractorTestData{input: "83: 17 5", want: []int{17, 5}},
		numberExtractorTestData{input: "156: 15 6", want: []int{15, 6}},
		numberExtractorTestData{input: "7290: 6 8 6 15", want: []int{6, 8, 6, 15}},
		numberExtractorTestData{input: "161011: 16 10 13", want: []int{16, 10, 13}},
		numberExtractorTestData{input: "192: 17 8 14", want: []int{17, 8, 14}},
		numberExtractorTestData{input: "21037: 9 7 18 13", want: []int{9, 7, 18, 13}},
		numberExtractorTestData{input: "292: 11 6 16 20", want: []int{11, 6, 16, 20}},
	}
	for _, test := range tests {
		got := numberExtractor(test.input)
		for i := range got {
			if got[i] != test.want[i] {
				t.Errorf("input: %s\nwant: %d\ngot: %d", test.input, test.want[i], got[i])
			}
		}
	}
}

func TestStringToInt(t *testing.T) {
	tests := []stringToIntTestData{
		stringToIntTestData{input: "3124", want: 3124},
		stringToIntTestData{input: "-1234", want: -1234},
		stringToIntTestData{input: "0", want: 0},
	}
	for _, test := range tests {
		got := stringToInt(test.input)
		if got != test.want {
			t.Errorf("stringToInt(%q) = %d, want %d", test.input, got, test.want)
		}
	}
}