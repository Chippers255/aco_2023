package day03

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readInputFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error reading file: %v", err)
	}

	return strings.TrimSpace(string(content)), nil
}

func findMulExpressions(input string) []string {
	pattern := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	return pattern.FindAllString(input, -1)
}

func findAllExpressions(input string) []string {
	pattern := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	return pattern.FindAllString(input, -1)
}

func findExpressionArgs(expression string) (int, int) {
	pattern := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := pattern.FindStringSubmatch(expression)
	arg1, _ := strconv.Atoi(matches[1])
	arg2, _ := strconv.Atoi(matches[2])

	return arg1, arg2
}

func Part1(input string) (int, error) {
	inputString, err := readInputFile(input)
	if err != nil {
		return 0, err
	}

	sum := 0

	mulExpressions := findMulExpressions(inputString)
	for _, expr := range mulExpressions {
		arg1, arg2 := findExpressionArgs(expr)
		sum += arg1 * arg2
	}

	return sum, nil
}

func Part2(input string) (int, error) {
	inputString, err := readInputFile(input)
	if err != nil {
		return 0, err
	}

	sum := 0
	do := true
	allExpressions := findAllExpressions(inputString)
	for _, expr := range allExpressions {
		if expr == "do()" {
			do = true
		} else if expr == "don't()" {
			do = false
		} else {
			if do {
				arg1, arg2 := findExpressionArgs(expr)
				sum += arg1 * arg2
			}
		}
	}
	return sum, nil
}

func Run(input string) error {
	part1, err := Part1(input)
	if err != nil {
		return err
	}
	fmt.Printf("Day 02 - Part 1: %d\n", part1)

	part2, err := Part2(input)
	if err != nil {
		return err
	}
	fmt.Printf("Day 02 - Part 2: %d\n", part2)

	return nil
}
