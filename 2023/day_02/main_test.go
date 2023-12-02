package main

import "testing"

func TestDay02PartOne(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"sample_input_1.txt", 8},
		{"input.txt", 2439},
	}

	for _, test := range tests {
		result := PartOne(test.input)
		if result != test.expected {
			t.Errorf("For input '%s', expected %d, got %d", test.input, test.expected, result)
		}
	}
}

func TestDay02PartTwo(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"sample_input_2.txt", 2286},
		{"input.txt", 63711},
	}

	for _, test := range tests {
		result := PartTwo(test.input)
		if result != test.expected {
			t.Errorf("For input '%s', expected %d, got %d", test.input, test.expected, result)
		}
	}
}
