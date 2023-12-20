package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

func parseWorkflow(line string) (string, []func(part map[string]int) string) {
	reGT := regexp.MustCompile(`(\w)>(\d+):(\w+)`)
	reLT := regexp.MustCompile(`(\w)<(\d+):(\w+)`)
	reDefault := regexp.MustCompile(`^(\w+)$`)
	funkyBois := []func(part map[string]int) string{}

	parts := strings.Split(line, "{")
	workflowName := parts[0]
	conditions := strings.Split(strings.TrimRight(parts[1], "}"), ",")
	for _, condition := range conditions {
		matchesGT := reGT.FindStringSubmatch(condition)
		matchesLT := reLT.FindStringSubmatch(condition)
		matchesDefault := reDefault.FindStringSubmatch(condition)

		if len(matchesGT) > 0 {
			expected, _ := strconv.Atoi(matchesGT[2])
			newFunkyBoi := func(part map[string]int) string {
				if part[matchesGT[1]] > expected {
					return matchesGT[3]
				}
				return ""
			}
			funkyBois = append(funkyBois, newFunkyBoi)
		} else if len(matchesLT) > 0 {
			expected, _ := strconv.Atoi(matchesLT[2])
			newFunkyBoi := func(part map[string]int) string {
				if part[matchesLT[1]] < expected {
					return matchesLT[3]
				}
				return ""
			}
			funkyBois = append(funkyBois, newFunkyBoi)
		} else if len(matchesDefault) > 0 {
			newFunkyBoi := func(part map[string]int) string {
				return matchesDefault[1]
			}
			funkyBois = append(funkyBois, newFunkyBoi)
		}
	}

	return workflowName, funkyBois
}

func processPart(line string) map[string]int {
	resp := make(map[string]int)
	re := regexp.MustCompile(`{x=(\d+),m=(\d+),a=(\d+),s=(\d+)}`)
	matches := re.FindStringSubmatch(line)

	resp["x"], _ = strconv.Atoi(matches[1])
	resp["m"], _ = strconv.Atoi(matches[2])
	resp["a"], _ = strconv.Atoi(matches[3])
	resp["s"], _ = strconv.Atoi(matches[4])

	return resp
}

func processInput(lines []string) (map[string][]func(part map[string]int) string, []map[string]int) {
	workflows := make(map[string][]func(part map[string]int) string)
	parts := []map[string]int{}
	whatAreWeDoing := "workflows"

	for _, line := range lines {
		if line == "" {
			whatAreWeDoing = "parts"
			continue
		}
		if whatAreWeDoing == "workflows" {
			key, value := parseWorkflow(line)
			workflows[key] = value
		} else {
			parts = append(parts, processPart(line))
		}
	}
	return workflows, parts
}

func PartOne(filename string) int {
	lines := getFileLines(filename)
	workflows, parts := processInput(lines)
	goodParts := []map[string]int{}

	for _, part := range parts {
		nextFunc := "in"
		for nextFunc != "A" && nextFunc != "R" {
			for _, condition := range workflows[nextFunc] {
				nextFunc = condition(part)
				if nextFunc != "" {
					break
				}
			}
		}
		if nextFunc == "A" {
			goodParts = append(goodParts, part)
		}
	}

	fmt.Println(goodParts)
	return addGoodPuds(goodParts)
}

func addGoodPuds(goodPuds []map[string]int) int {
	sum := 0
	for _, goodPud := range goodPuds {
		sum += goodPud["x"]
		sum += goodPud["m"]
		sum += goodPud["a"]
		sum += goodPud["s"]
	}
	return sum
}

func main() {
	fmt.Println(PartOne("input.txt"))
}
