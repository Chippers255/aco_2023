package day05

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
)

type Rule struct {
	Before int
	After  int
}

type Update struct {
	Pages []int
}

func (u *Update) FindPageIndex(target int) int {
	for i, value := range u.Pages {
		if value == target {
			return i
		}
	}
	return -1
}

func (u *Update) GetMiddlePage() int {
	length := len(u.Pages)
	middleIndex := length / 2
	return u.Pages[middleIndex]
}

func checkEm(beforeIndex int, afterIndex int) bool {
	if beforeIndex == -1 || afterIndex == -1 {
		return true
	}
	return beforeIndex < afterIndex
}

func parseInput(filePath string) ([]Rule, []Update, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, nil, err
	}

	var rules []Rule
	var updates []Update
	scanner := bufio.NewScanner(strings.NewReader(string(file)))
	section := 1

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			section++
			continue
		}

		if section == 1 {
			parts := strings.Split(line, "|")
			if len(parts) != 2 {
				continue
			}
			before, err1 := strconv.Atoi(parts[0])
			after, err2 := strconv.Atoi(parts[1])
			if err1 != nil || err2 != nil {
				continue
			}
			rules = append(rules, Rule{Before: before, After: after})
		} else if section == 2 {
			strNums := strings.Split(line, ",")
			var nums Update
			for _, s := range strNums {
				num, err := strconv.Atoi(strings.TrimSpace(s))
				if err != nil {
					return nil, nil, err
				}
				nums.Pages = append(nums.Pages, num)
			}
			updates = append(updates, nums)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return rules, updates, nil
}

func Part1(input string) (int, error) {
	rules, updates, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	var ans int
	for _, update := range updates {
		heGood := true
		for _, rule := range rules {
			beforeIndex := update.FindPageIndex(rule.Before)
			afterIndex := update.FindPageIndex(rule.After)
			if !checkEm(beforeIndex, afterIndex) {
				heGood = false
				break
			}
		}
		if heGood {
			middlePage := update.GetMiddlePage()
			ans += middlePage
		}
	}

	return ans, nil
}

func reorderUpdate(rules []Rule, update Update) Update {
	// I used a library for once, because implementing shit from scratch in go is a slog
	g := simple.NewDirectedGraph()

	// nodes
	for _, page := range update.Pages {
		g.AddNode(simple.Node(page))
	}

	// edges
	for _, rule := range rules {
		// gonum doesn't like it when you add an edge to a node that doesn't exist
		if slices.Contains(update.Pages, rule.Before) && slices.Contains(update.Pages, rule.After) {
			g.SetEdge(simple.Edge{F: simple.Node(rule.Before), T: simple.Node(rule.After)})
		}
	}

	sorted, err := topo.Sort(g)
	if err != nil {
		panic(err) // at the disco, i didn't want to handle this error because it should never happen
	}

	result := make([]int, len(sorted))
	for i, node := range sorted {
		result[i] = int(node.ID())
	}
	return Update{Pages: result}
}

func Part2(input string) (int, error) {
	// d'ya like dags?
	// https://www.youtube.com/watch?v=zH64dlgyydM
	rules, updates, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	var ans int
	for _, update := range updates {
		heGood := true
		for _, rule := range rules {
			beforeIndex := update.FindPageIndex(rule.Before)
			afterIndex := update.FindPageIndex(rule.After)
			if !checkEm(beforeIndex, afterIndex) {
				heGood = false
			}
		}
		if !heGood {
			sorted := reorderUpdate(rules, update)
			middlePage := sorted.GetMiddlePage()
			ans += middlePage
		}
	}

	return ans, nil
}

func Run(input string) error {
	part1, err := Part1(input)
	if err != nil {
		return err
	}
	fmt.Printf("Day 05 - Part 1: %d\n", part1)

	part2, err := Part2(input)
	if err != nil {
		return err
	}
	fmt.Printf("Day 05 - Part 2: %d\n", part2)

	return nil
}
