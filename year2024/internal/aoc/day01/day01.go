package day01

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readInputFile(filePath string) ([]int, []int, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("error reading file: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")

	l1 := make([]int, 0, len(lines))
	l2 := make([]int, 0, len(lines))

	for _, line := range lines {
		numbers := strings.Fields(line)
		num1, _ := strconv.Atoi(numbers[0])
		num2, _ := strconv.Atoi(numbers[1])
		l1 = append(l1, num1)
		l2 = append(l2, num2)
	}

	return l1, l2, nil
}

func sumList(list []int) int {
	sum := 0
	for _, v := range list {
		sum += v
	}
	return sum
}

func absDiff(a, b int) int {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	return diff
}

func Part1(input string) (int, error) {
	l1, l2, err := readInputFile(input)
	if err != nil {
		return 0, err
	}

	sort.Ints(l1)
	sort.Ints(l2)

	diff := make([]int, len(l1))
	for i := range l1 {
		diff[i] = absDiff(l1[i], l2[i])
	}
	sum := sumList(diff)

	return sum, nil
}

func Part2(input string) (int, error) {
	l1, l2, err := readInputFile(input)
	if err != nil {
		return 0, err
	}

	sum := 0
	for _, n1 := range l1 {
		for _, n2 := range l2 {
			if n1 == n2 {
				sum += n1
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
	fmt.Printf("Day 01 - Part 1: %d\n", part1)

	part2, err := Part2(input)
	if err != nil {
		return err
	}
	fmt.Printf("Day 01 - Part 2: %d\n", part2)

	return nil
}
