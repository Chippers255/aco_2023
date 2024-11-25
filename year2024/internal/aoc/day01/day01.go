package day01

import (
	"fmt"
	"os"
)

func Part1(input string) (int, error) {
	file, err := os.Open(input)
	if err != nil {
		return 0, fmt.Errorf("failed to read input file: %w", err)
	}
	defer file.Close()

	return 42, nil
}

func Part2(input string) (int, error) {
	file, err := os.Open(input)
	if err != nil {
		return 0, fmt.Errorf("failed to read input file: %w", err)
	}
	defer file.Close()

	return 84, nil
}

func Run(input string) error {
	part1, err := Part1(input)
	if err != nil {
		return err
	}
	fmt.Printf("Day 01 - Part 1: %d\n", part1)

	part2, err := Part2(input)
	if err != nil {
		return err
	}
	fmt.Printf("Day 01 - Part 2: %d\n", part2)

	return nil
}
