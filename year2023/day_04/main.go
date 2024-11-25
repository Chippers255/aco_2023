package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func PartOne(filename string) int {
	lines := getFileLines(filename)

	score := 0
	for _, card := range lines {
		c := 0
		parts := strings.Split(card, ": ")
		parts = strings.Split(parts[1], " | ")
		winNums := processString(parts[0])
		myNums := processString(parts[1])

		for _, num := range myNums {
			for _, winNum := range winNums {
				if num == winNum {
					if c == 0 {
						c++
					} else {
						c *= 2
					}
				}
			}
		}
		score += c
	}

	return score
}

func PartTwo(filename string) int {
	lines := getFileLines(filename)

	score := 0
	for _, card := range lines {
		c := 0
		parts := strings.Split(card, ": ")
		parts = strings.Split(parts[1], " | ")
		winNums := processString(parts[0])
		myNums := processString(parts[1])

		for _, num := range myNums {
			for _, winNum := range winNums {
				if num == winNum {
					if c == 0 {
						c++
					} else {
						c *= 2
					}
				}
			}
		}
		score += c
	}

	return score
}

// Helper function to process the string and remove any empty strings
func processString(str string) []string {
	parts := strings.Split(str, " ")
	var result []string
	for _, part := range parts {
		if strings.TrimSpace(part) != "" {
			result = append(result, part)
		}
	}
	return result
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
