package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Pallinder/go-randomdata"
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

func betterAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func betterMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func betterMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func manhattanDistance(a galaxy, b galaxy, thiccRows []int, thiccColumns []int) int {
	// a = [x, y] b = [x, y]
	ans := betterAbs(a.x-b.x) + betterAbs(a.y-b.y)
	for _, row := range thiccRows {
		if row > betterMin(a.y, b.y) && row < betterMax(a.y, b.y) {
			ans += 1
		}
	}
	for _, col := range thiccColumns {
		if col > betterMin(a.x, b.x) && col < betterMax(a.x, b.x) {
			ans += 1
		}
	}
	return ans
}

func nipDistance(a galaxy, b galaxy, thiccRows []int, thiccColumns []int) int {
	// a = [x, y] b = [x, y]
	ans := betterAbs(a.x-b.x) + betterAbs(a.y-b.y)
	for _, row := range thiccRows {
		if row > betterMin(a.y, b.y) && row < betterMax(a.y, b.y) {
			ans += 999999
		}
	}
	for _, col := range thiccColumns {
		if col > betterMin(a.x, b.x) && col < betterMax(a.x, b.x) {
			ans += 999999
		}
	}
	return ans
}

func getThiccBois(space [][]string) ([]int, []int) {
	rows := []int{}
	cols := []int{}

	for i, row := range space {
		if strings.Join(row, "") == strings.Repeat(".", len(row)) {
			rows = append(rows, i)
		}
	}

	newSpace := transpose(space)
	for i, col := range newSpace {
		if strings.Join(col, "") == strings.Repeat(".", len(col)) {
			cols = append(cols, i)
		}
	}

	return rows, cols
}

func transpose(slice [][]string) [][]string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func getGalaxyCoordinates(space [][]string) []galaxy {
	galaxyCoordinates := []galaxy{}
	for i, row := range space {
		for j, col := range row {
			if col == "#" {
				galaxyCoordinates = append(galaxyCoordinates, galaxy{x: j, y: i, name: randomdata.SillyName()})
			}
		}
	}
	return galaxyCoordinates
}

type galaxy struct {
	x, y int
	name string
}

func PartOne(filename string) int {
	SPACE := readInputFile(filename) // https://www.youtube.com/watch?v=niZpcdp2v34
	galaxyCoordinates := getGalaxyCoordinates(SPACE)
	thiccRows, thiccColumns := getThiccBois(SPACE)
	hyperspaceRoutes := make(map[string]int)

	for _, a := range galaxyCoordinates {
		for _, b := range galaxyCoordinates {
			if a.x == b.x && a.y == b.y {
				continue
			}
			hyperspaceRoutes[fmt.Sprintf("%v,%v", a, b)] = manhattanDistance(a, b, thiccRows, thiccColumns)
		}
	}

	ans := 0
	for _, v := range hyperspaceRoutes {
		ans += v
	}
	return ans / 2
}

func PartTwo(filename string) int {
	SPACE := readInputFile(filename) // https://www.youtube.com/watch?v=niZpcdp2v34
	galaxyCoordinates := getGalaxyCoordinates(SPACE)
	thiccRows, thiccColumns := getThiccBois(SPACE)
	hyperspaceRoutes := make(map[string]int)

	for _, a := range galaxyCoordinates {
		for _, b := range galaxyCoordinates {
			if a.x == b.x && a.y == b.y {
				continue
			}
			hyperspaceRoutes[fmt.Sprintf("%v,%v", a, b)] = nipDistance(a, b, thiccRows, thiccColumns)
		}
	}

	ans := 0
	for _, v := range hyperspaceRoutes {
		ans += v
	}
	return ans / 2
}

func main() {
	fmt.Println(PartOne("input.txt"))
	fmt.Println(PartTwo("input.txt"))
}
