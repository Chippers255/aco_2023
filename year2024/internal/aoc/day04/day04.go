package day04

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const targetWord = "XMAS"

type direction struct {
	dx int
	dy int
}

type pattern struct {
	expected map[[2]int]rune
}

func definePatterns() []pattern {
	// yes I hard coded the patterns, eat my ass. Copilot did it anyways.
	return []pattern{
		{
			expected: map[[2]int]rune{
				{0, 0}: 'M', {0, 2}: 'S',
				{1, 1}: 'A',
				{2, 0}: 'M', {2, 2}: 'S',
			},
		},
		{
			expected: map[[2]int]rune{
				{0, 0}: 'S', {0, 2}: 'M',
				{1, 1}: 'A',
				{2, 0}: 'S', {2, 2}: 'M',
			},
		},
		{
			expected: map[[2]int]rune{
				{0, 0}: 'M', {0, 2}: 'M',
				{1, 1}: 'A',
				{2, 0}: 'S', {2, 2}: 'S',
			},
		},
		{
			expected: map[[2]int]rune{
				{0, 0}: 'S', {0, 2}: 'S',
				{1, 1}: 'A',
				{2, 0}: 'M', {2, 2}: 'M',
			},
		},
	}
}

func readGrid(filePath string) ([][]rune, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(strings.NewReader(string(file)))
	var grid [][]rune
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	return grid, nil
}

func search(grid [][]rune, x, y int, dir direction, index int, rows, cols int) bool {
	if index == len(targetWord) {
		return true
	}

	if x < 0 || y < 0 || x >= rows || y >= cols {
		return false
	}

	if grid[x][y] != rune(targetWord[index]) {
		return false
	}

	nextX := x + dir.dx
	nextY := y + dir.dy

	return search(grid, nextX, nextY, dir, index+1, rows, cols)
}

func Part1(input string) (int, error) {
	// I learned from last year, recurssion is the way to go
	grid, err := readGrid(input)
	if err != nil {
		return 0, err
	}

	count := 0
	rows := len(grid)
	cols := len(grid[0])

	directions := []direction{
		{dx: 0, dy: 1},
		{dx: 1, dy: 1},
		{dx: 1, dy: 0},
		{dx: 1, dy: -1},
		{dx: 0, dy: -1},
		{dx: -1, dy: -1},
		{dx: -1, dy: 0},
		{dx: -1, dy: 1},
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == rune(targetWord[0]) {
				for _, dir := range directions {
					if search(grid, i, j, dir, 0, rows, cols) {
						count++
					}
				}
			}
		}
	}

	return count, nil
}

func extractWindow(grid [][]rune, startRow, startCol int) [][]rune {
	window := make([][]rune, 3)
	for i := 0; i < 3; i++ {
		window[i] = make([]rune, 3)
		for j := 0; j < 3; j++ {
			window[i][j] = grid[startRow+i][startCol+j]
		}
	}
	return window
}

func matchesPattern(window [][]rune, p pattern) bool {
	for pos, expectedChar := range p.expected {
		row, col := pos[0], pos[1]
		if window[row][col] != expectedChar {
			return false
		}
	}
	return true
}

func Part2(input string) (int, error) {
	// I think this is the first time I have used a kernel outside of AI
	grid, err := readGrid(input)
	if err != nil {
		return 0, err
	}

	count := 0
	rows := len(grid)
	cols := len(grid[0])
	patterns := definePatterns()

	for i := 0; i <= rows-3; i++ {
		for j := 0; j <= cols-3; j++ {
			window := extractWindow(grid, i, j)
			for _, p := range patterns {
				if matchesPattern(window, p) {
					count++
					break
				}
			}
		}
	}

	return count, nil
}

func Run(input string) error {
	part1, err := Part1(input)
	if err != nil {
		return err
	}
	fmt.Printf("Day 04 - Part 1: %d\n", part1)

	part2, err := Part2(input)
	if err != nil {
		return err
	}
	fmt.Printf("Day 04 - Part 2: %d\n", part2)

	return nil
}
