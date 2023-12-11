package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readInputFile(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		lines = append(lines, line)
	}

	return lines
}

func findStart(grid [][]string) (int, int) {
	for y, line := range grid {
		for x, char := range line {
			if char == "S" {
				return x, y
			}
		}
	}
	return -1, -1
}

func getNextPosition(grid [][]string, curX, curY, lastX, lastY int) (int, int) {
	curVal := grid[curY][curX]
	if curVal == "S" {
		return curX, curY - 1
	} else if curVal == "|" {
		if lastY < curY {
			return curX, curY + 1
		} else {
			return curX, curY - 1
		}
	} else if curVal == "-" {
		if lastX < curX {
			return curX + 1, curY
		} else {
			return curX - 1, curY
		}
	} else if curVal == "L" {
		if lastY < curY {
			return curX + 1, curY
		} else {
			return curX, curY - 1
		}
	} else if curVal == "J" {
		if lastY < curY {
			return curX - 1, curY
		} else {
			return curX, curY - 1
		}
	} else if curVal == "7" {
		if lastY > curY {
			return curX - 1, curY
		} else {
			return curX, curY + 1
		}
	} else if curVal == "F" {
		if lastY > curY {
			return curX + 1, curY
		} else {
			return curX, curY + 1
		}
	} else {
		panic("Invalid character")
	}

}

func traverse(grid [][]string, startX, startY int) int {
	stepCount := 1
	curX, curY := startX, startY-1
	lastX, lastY := startX, startY
	current := grid[curY][curX]
	for current != "S" {
		nextX, nextY := getNextPosition(grid, curX, curY, lastX, lastY)
		current = grid[nextY][nextX]
		stepCount++
		lastX = curX
		lastY = curY
		curX = nextX
		curY = nextY
	}

	return stepCount
}

func PartOne(filename string) int {
	grid := readInputFile(filename)
	startX, startY := findStart(grid)
	ans := traverse(grid, startX, startY)
	return ans / 2
}

func main() {
	fmt.Println("Part 1:", PartOne("input.txt"))
}
