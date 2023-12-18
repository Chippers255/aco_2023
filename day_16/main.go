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

type lightPartical struct {
	x   int
	y   int
	dir string
}

func movePartical(lp lightPartical, grid [][]string) []lightPartical {
	if lp.x < 0 || lp.x >= len(grid[0]) || lp.y < 0 || lp.y >= len(grid) {
		return []lightPartical{}
	}
	if gotOne(lp.x, lp.y, lp.dir) {
		return []lightPartical{}
	}

	particalList := []lightPartical{lp}
	curVal := grid[particalList[0].y][particalList[0].x]
	visit(particalList[0].x, particalList[0].y)
	writeGotOne(particalList[0].x, particalList[0].y, particalList[0].dir)

	if particalList[0].dir == "r" {
		if curVal == "-" || curVal == "." {
			particalList[0].x++
		} else if curVal == "/" {
			particalList[0].dir = "u"
			particalList[0].y--
		} else if curVal == "\\" {
			particalList[0].dir = "d"
			particalList[0].y++
		} else if curVal == "|" {
			particalList = append(particalList, lightPartical{particalList[0].x, particalList[0].y + 1, "d"})
			particalList[0].y--
			particalList[0].dir = "u"
		}
	} else if particalList[0].dir == "l" {
		if curVal == "-" || curVal == "." {
			particalList[0].x--
		} else if curVal == "/" {
			particalList[0].dir = "d"
			particalList[0].y++
		} else if curVal == "\\" {
			particalList[0].dir = "u"
			particalList[0].y--
		} else if curVal == "|" {
			particalList = append(particalList, lightPartical{particalList[0].x, particalList[0].y + 1, "d"})
			particalList[0].y--
			particalList[0].dir = "u"
		}
	} else if particalList[0].dir == "u" {
		if curVal == "|" || curVal == "." {
			particalList[0].y--
		} else if curVal == "/" {
			particalList[0].dir = "r"
			particalList[0].x++
		} else if curVal == "\\" {
			particalList[0].dir = "l"
			particalList[0].x--
		} else if curVal == "-" {
			particalList = append(particalList, lightPartical{particalList[0].x + 1, particalList[0].y, "r"})
			particalList[0].x--
			particalList[0].dir = "l"
		}
	} else if particalList[0].dir == "d" {
		if curVal == "|" || curVal == "." {
			particalList[0].y++
		} else if curVal == "/" {
			particalList[0].dir = "l"
			particalList[0].x--
		} else if curVal == "\\" {
			particalList[0].dir = "r"
			particalList[0].x++
		} else if curVal == "-" {
			particalList = append(particalList, lightPartical{particalList[0].x + 1, particalList[0].y, "r"})
			particalList[0].x--
			particalList[0].dir = "l"
		}
	}

	return particalList
}

var visitedLocations = make(map[string]bool)
var alreadyGotOne = make(map[string]bool)

func visit(x int, y int) {
	key := fmt.Sprintf("%d,%d", x, y)
	visitedLocations[key] = true
}

func writeGotOne(x int, y int, dir string) {
	key := fmt.Sprintf("%d,%d,%s", x, y, dir)
	alreadyGotOne[key] = true
}

func gotOne(x int, y int, dir string) bool {
	key := fmt.Sprintf("%d,%d,%s", x, y, dir)
	_, ok := alreadyGotOne[key]
	return ok
}

func PartOne(filename string) int {
	grid := readInputFile(filename)
	energyBoys := []lightPartical{lightPartical{0, 0, "r"}}
	newBoys := []lightPartical{}

	for i := 0; i < 50000; i++ {
		for _, lp := range energyBoys {
			n := movePartical(lp, grid)
			newBoys = append(newBoys, n...)
		}
		energyBoys = newBoys
		newBoys = []lightPartical{}
		if len(energyBoys) == 0 {
			break
		}
	}
	count := 0
	for _, v := range visitedLocations {
		if v {
			count++
		}
	}
	return count
}

func getStart(x int, y int, grid [][]string) []string {
	startDir := []string{}
	if y == 0 {
		startDir = append(startDir, "d")
	} else if y == len(grid)-1 {
		startDir = append(startDir, "u")
	}
	if x == 0 {
		startDir = append(startDir, "r")
	} else if x == len(grid[0])-1 {
		startDir = append(startDir, "l")
	}
	return startDir
}

func PartTwo(filename string) int {
	grid := readInputFile(filename)
	best := 0

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			startDir := getStart(x, y, grid)
			for _, dir := range startDir {
				visitedLocations = make(map[string]bool)
				alreadyGotOne = make(map[string]bool)
				energyBoys := []lightPartical{lightPartical{x, y, dir}}
				newBoys := []lightPartical{}

				for i := 0; i < 50000; i++ {
					for _, lp := range energyBoys {
						n := movePartical(lp, grid)
						newBoys = append(newBoys, n...)
					}
					energyBoys = newBoys
					newBoys = []lightPartical{}
					if len(energyBoys) == 0 {
						break
					}
				}
				count := 0
				for _, v := range visitedLocations {
					if v {
						count++
					}
				}
				if count > best {
					best = count
				}
			}
		}
	}

	return best
}

func main() {
	fmt.Println(PartOne("input.txt"))
	fmt.Println(PartTwo("input.txt"))
}
