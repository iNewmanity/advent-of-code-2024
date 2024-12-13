package day11

import (
	"testing"
)

type blinkTestData struct {
	input Stones
	want  Stones
}

func TestBlink(t *testing.T) {
	tests := []blinkTestData{
		{input: Stones{"0", "1", "1234"}, want: Stones{"1", "2024", "12", "34"}},
		{input: Stones{"1", "2024", "12", "34"}, want: Stones{"2024", "20", "24", "1", "2", "3", "4"}},
		{input: Stones{"2024", "20", "24", "1", "2", "3", "4"}, want: Stones{"20", "24", "2", "0", "2", "4", "2024", "4048", "6072", "8096"}},
	}
	for _, test := range tests {
		got := test.input.Blink()
		for i := range got {
			if got[i] != test.want[i] {
				t.Errorf("blink(%v, %v)=%v, want %v", test.input, got[i], test.want[i], test.want[i])
			}
		}
	}
}

type applyRulesTestData struct {
	input Stone
	want  Stones
}

func TestApplyRules(t *testing.T) {
	tests := []applyRulesTestData{
		{input: "0", want: Stones{"1"}},
		{input: "1234", want: Stones{"12", "34"}},
		{input: "1", want: Stones{"2024"}},
	}
	for _, test := range tests {
		got := test.input.applyRules()
		for i := range got {
			if got[i] != test.want[i] {
				t.Errorf("applyRules() got = %v, want %v", got, test.want)
			}
		}
	}
}

type hasEvenNumberOfDigitsTestData struct {
	input Stone
	want  bool
}

func TestHasEvenNumberOfDigits(t *testing.T) {
	tests := []hasEvenNumberOfDigitsTestData{
		{input: "1234", want: true},
		{input: "987654321", want: false},
		{input: "98765432", want: true},
	}
	for _, test := range tests {
		got := test.input.hasEvenNumbersOfDigits()
		if got != test.want {
			t.Errorf("hasEvenNumbersOfDigits() = %v, want %v", got, test.want)
		}
	}
}

type hasValOfZeroTestData struct {
	input Stone
	want  bool
}

func TestHasValOfZero(t *testing.T) {
	tests := []hasValOfZeroTestData{
		{input: "0", want: true},
		{input: "1", want: false},
		{input: "000", want: true},
	}
	for _, test := range tests {
		got := test.input.hasValOfZero()
		if got != test.want {
			t.Errorf("hasValOfZero(%v) = %v; want %v", test.input, got, test.want)
		}
	}
}

type cutLeadingZerosTestData struct {
	input Stone
	want  Stone
}

func TestCutLeadingZeros(t *testing.T) {
	tests := []cutLeadingZerosTestData{
		{input: "1000", want: "1000"},
		{input: "0001", want: "1"},
		{input: "000", want: "0"},
	}
	for _, test := range tests {
		got := test.input.cutLeadingZeros()
		if got != test.want {
			t.Errorf("CutLeadingZeros(%v) = %v, want %v", test.input, got, test.want)
		}
	}
}

type multiplyBy2024TestData struct {
	input Stone
	want  Stone
}

func TestMultiplyBy2024(t *testing.T) {
	tests := []multiplyBy2024TestData{
		{input: "1", want: "2024"},
		{input: "0", want: "0"},
		{input: "256", want: "518144"},
	}
	for _, test := range tests {
		got := test.input.multiplyBy2024()
		if got != test.want {
			t.Errorf("MultiplyBy2024(%v) = %v, want %v", test.input, got, test.want)
		}
	}
}

type toIntTestData struct {
	input Stone
	want  int
}

func TestToInt(t *testing.T) {
	tests := []toIntTestData{
		{input: "0", want: 0},
		{input: "202020", want: 202020},
	}
	for _, test := range tests {
		got := test.input.toInt()
		if got != test.want {
			t.Errorf("toInt(%v) = %v, want %v", test.input, got, test.want)
		}
	}
}

type toStoneTestData struct {
	input number
	want  Stone
}

func TestToStone(t *testing.T) {
	tests := []toStoneTestData{
		{input: 0001, want: "1"},
		{input: 000, want: "0"},
		{input: 1000, want: "1000"},
	}
	for _, test := range tests {
		got := test.input.toStone()
		if got != test.want {
			t.Errorf("toStone(%v) = %v, want %v", test.input, got, test.want)
		}
	}
}

type splitTestData struct {
	input Stone
	want  Stones
}

func TestSplit(t *testing.T) {
	tests := []splitTestData{
		{input: "12", want: Stones{
			"1", "2",
		}},
		{input: "123456", want: Stones{
			"123", "456",
		}},
	}
	for _, test := range tests {
		got := test.input.split()
		for i := range got {
			if got[i] != test.want[i] {
				t.Errorf("split(%v) = %v, want %v", test.input, got, test.want)
			}
		}
	}
}
