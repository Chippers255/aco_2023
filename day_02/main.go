package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getFileLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func PartOne(file string) int {
	const maxRed, maxGreen, maxBlue = 12, 13, 14

	games := getFileLines(file)
	sumOfIDs := 0

	for _, game := range games {
		parts := strings.Split(game, ": ")
		gameID, _ := strconv.Atoi(strings.Split(parts[0], " ")[1])
		rounds := strings.Split(parts[1], "; ")

		maxSeenRed, maxSeenGreen, maxSeenBlue := 0, 0, 0
		for _, round := range rounds {
			cubes := strings.Split(round, ", ")
			for _, cube := range cubes {
				cubeDetails := strings.Split(cube, " ")
				count, _ := strconv.Atoi(cubeDetails[0])
				color := cubeDetails[1]

				switch color {
				case "red":
					if count > maxSeenRed {
						maxSeenRed = count
					}
				case "green":
					if count > maxSeenGreen {
						maxSeenGreen = count
					}
				case "blue":
					if count > maxSeenBlue {
						maxSeenBlue = count
					}
				}
			}
		}

		if maxSeenRed <= maxRed && maxSeenGreen <= maxGreen && maxSeenBlue <= maxBlue {
			sumOfIDs += gameID
		}
	}
	return sumOfIDs
}

func PartTwo(file string) int {
	games := getFileLines(file)
	sumOfIDs := 0

	for _, game := range games {
		parts := strings.Split(game, ": ")
		rounds := strings.Split(parts[1], "; ")

		maxSeenRed, maxSeenGreen, maxSeenBlue := 0, 0, 0
		for _, round := range rounds {
			cubes := strings.Split(round, ", ")
			for _, cube := range cubes {
				cubeDetails := strings.Split(cube, " ")
				count, _ := strconv.Atoi(cubeDetails[0])
				color := cubeDetails[1]

				switch color {
				case "red":
					if count > maxSeenRed {
						maxSeenRed = count
					}
				case "green":
					if count > maxSeenGreen {
						maxSeenGreen = count
					}
				case "blue":
					if count > maxSeenBlue {
						maxSeenBlue = count
					}
				}
			}
		}

		sumOfIDs += (maxSeenRed * maxSeenGreen * maxSeenBlue)
	}
	return sumOfIDs
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
