package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("/home/tomos/personal/github/advent-of-code/year2024/inputs/2024/day01_part01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var left, right []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}
		a, err := strconv.Atoi(fields[0])
		if err != nil {
			log.Fatal(err)
		}
		b, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatal(err)
		}
		left = append(left, a)
		right = append(right, b)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Part 1
	sort.Ints(left)
	sort.Ints(right)

	totalDistance := 0
	for i := 0; i < len(left) && i < len(right); i++ {
		diff := left[i] - right[i]
		if diff < 0 {
			diff = -diff
		}
		totalDistance += diff
	}

	// Part 2
	counts := make(map[int]int)
	for _, num := range right {
		counts[num]++
	}

	similarityScore := 0
	for _, num := range left {
		similarityScore += num * counts[num]
	}

	fmt.Println(totalDistance)
	fmt.Println(similarityScore)
}
