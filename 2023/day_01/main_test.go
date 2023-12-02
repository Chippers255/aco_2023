package main

import "testing"

func TestDay01PartOne(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"sample_input_1.txt", 142},
		{"input.txt", 55123},
	}

	for _, test := range tests {
		result := PartOne(test.input)
		if result != test.expected {
			t.Errorf("For input '%s', expected %d, got %d", test.input, test.expected, result)
		}
	}
}

func TestDay01PartTwo(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"sample_input_2.txt", 281},
		{"input.txt", 55260},
	}

	for _, test := range tests {
		result := PartTwo(test.input)
		if result != test.expected {
			t.Errorf("For input '%s', expected %d, got %d", test.input, test.expected, result)
		}
	}
}
