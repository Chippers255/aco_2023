package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getFilePatterns(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var patterns [][]string
	var currentPattern []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" { // Blank line indicates a new pattern
			if len(currentPattern) > 0 {
				patterns = append(patterns, currentPattern)
				currentPattern = []string{} // Reset the current pattern
			}
		} else {
			currentPattern = append(currentPattern, line)
		}
	}

	// Add the last pattern if it's not empty
	if len(currentPattern) > 0 {
		patterns = append(patterns, currentPattern)
	}

	return patterns
}

func reverseRow(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func smolBoi(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func findVerticalReflections(pattern []string) int {
	size := len(pattern[0])
	for i := 1; i < size; i++ {
		biteSize := smolBoi(i, size-i)
		match := true
		for _, row := range pattern {
			if row[i-biteSize:i] != reverseRow(row[i:i+biteSize]) {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return 0
}

func transpose(slice []string) []string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([]string, xl)

	for i := 0; i < xl; i++ {
		newRow := make([]byte, yl)
		for j := 0; j < yl; j++ {
			newRow[j] = slice[j][i]
		}
		result[i] = string(newRow)
	}

	return result
}

func findHorizontalReflections(pattern []string) int {
	newPattern := transpose(pattern)
	return findVerticalReflections(newPattern)
}

func PartOne(filename string) int {
	patterns := getFilePatterns(filename)

	ans := 0
	for _, pattern := range patterns {
		ans += findVerticalReflections(pattern)
		ans += findHorizontalReflections(pattern) * 100
	}

	return ans
}

func main() {
	fmt.Println(PartOne("input.txt"))
}
