package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type gear struct {
	value       string
	x           int
	y           int
	partNumbers []string
}

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

func isSymbol(s string) bool {
	r := rune(s[0])
	return !unicode.IsNumber(r) && r != '.'
}

func isNumber(s string) bool {
	r := rune(s[0])
	return unicode.IsNumber(r)
}

func buildPartNumber(grid [][]string, x, y int) string {
	partNumber := ""
	for dx := -2; dx <= 2; dx++ {
		nx := x + dx
		if nx >= 0 && nx < len(grid[0]) {
			p := isNumber(grid[y][nx])
			if p {
				partNumber += grid[y][nx]
			}
			if !p && partNumber != "" {
				if nx >= x {
					break
				} else {
					partNumber = ""
				}
			}
		}
	}
	return partNumber
}

func getAdjacentNumbers(grid [][]string, x, y int) []string {
	var numbers []string
	dirs := []struct{ dx, dy int }{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	for _, d := range dirs {
		nx, ny := x+d.dx, y+d.dy
		if nx >= 0 && ny >= 0 && ny < len(grid) && nx < len(grid[ny]) {
			if isNumber(grid[ny][nx]) {
				pn := buildPartNumber(grid, nx, ny)
				if !contains(numbers, pn) {
					numbers = append(numbers, pn)
				}
			}
		}
	}

	return numbers
}

func contains(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}

func PartOne(file string) int {
	grid := readInputFile(file)
	var gears []gear

	for y, row := range grid {
		for x, cell := range row {
			if isSymbol(cell) {
				nums := getAdjacentNumbers(grid, x, y)
				g := gear{grid[y][x], x, y, nums}
				gears = append(gears, g)
			}
		}
	}

	total := 0
	for _, c := range gears {
		for _, n := range c.partNumbers {
			p, _ := strconv.Atoi(n)
			total += p
		}
	}
	return total
}

func PartTwo(file string) int {
	grid := readInputFile(file)
	var gears []gear

	for y, row := range grid {
		for x, cell := range row {
			if isSymbol(cell) {
				nums := getAdjacentNumbers(grid, x, y)
				g := gear{grid[y][x], x, y, nums}
				gears = append(gears, g)
			}
		}
	}

	total := 0
	for _, c := range gears {
		if c.value == "*" && len(c.partNumbers) == 2 {
			p1, _ := strconv.Atoi(c.partNumbers[0])
			p2, _ := strconv.Atoi(c.partNumbers[1])
			total += p1 * p2
		}
	}
	return total
}

func main() {
	p1Sample := PartOne("sample_input_1.txt")
	p1Input := PartOne("input.txt")
	fmt.Println(p1Sample)
	fmt.Println(p1Input)

	p2Sample := PartTwo("sample_input_2.txt")
	p2Input := PartTwo("input.txt")
	fmt.Println(p2Sample)
	fmt.Println(p2Input)
}
