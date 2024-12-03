package day02

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInputFile(filePath string) ([][]int, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	reports := strings.Split(strings.TrimSpace(string(content)), "\n")
	result := make([][]int, len(reports))

	for i, report := range reports {
		levels := strings.Fields(report)
		result[i] = make([]int, 0, len(levels))

		for _, level := range levels {
			convLevel, err := strconv.Atoi(level)
			if err != nil {
				return nil, fmt.Errorf("error converting string to integer: %v", err)
			}
			result[i] = append(result[i], convLevel)
		}
	}

	return result, nil
}

func removeBadLevel(report []int, index int) []int {
	newReport := make([]int, 0, len(report)-1)
	newReport = append(newReport, report[:index]...)
	newReport = append(newReport, report[index+1:]...)

	return newReport
}

func checkReport(report []int) (bool, int) {
	rg3000 := NewReindeerGuard()
	for i := 1; i < len(report); i++ {
		rg3000.Next(report[i-1], report[i])
		if !rg3000.isValid {
			if i == len(report)-1 {
				return false, len(report) - 1
			} else {
				return false, i - 1
			}
		}
	}

	return rg3000.isValid, 0
}

func Part1(input string) (int, error) {
	reports, err := readInputFile(input)
	if err != nil {
		return 0, err
	}

	safeCount := 0
	for _, report := range reports {
		valid, _ := checkReport(report)
		if valid {
			safeCount++
		}
	}

	return safeCount, nil
}

func Part2(input string) (int, error) {
	reports, err := readInputFile(input)
	if err != nil {
		return 0, err
	}

	safeCount := 0
	for _, report := range reports {
		valid, _ := checkReport(report)
		if valid {
			safeCount++
		} else {
			reportSafe := false
			for i := 0; i < len(report); i++ {
				dampenedReport := removeBadLevel(report, i)
				isValid, _ := checkReport(dampenedReport)
				if isValid {
					reportSafe = true
					break
				}
			}
			if reportSafe {
				safeCount++
			}
		}
	}

	return safeCount, nil
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
