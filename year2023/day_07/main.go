package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type Hand struct {
	cards []string
	kind  int
	bid   int
}

func readInputFile(filename string) []string {
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

func countCards(hand []string) map[string]int {
	counts := make(map[string]int)
	for _, card := range hand {
		counts[card]++
	}
	return counts
}

// isNOfAKind checks if there is a value in the given map that occurs exactly n times.
// It iterates over the values in the map and returns true if any value occurs n times, otherwise it returns false.
func isNOfAKind(counts map[string]int, n int) bool {
	for _, count := range counts {
		if count == n {
			return true
		}
	}
	return false
}

func bestHand(hand []string) int {
	counts := countCards(hand)

	if isNOfAKind(counts, 5) {
		return 7
	} else if isNOfAKind(counts, 4) {
		return 6
	} else if isNOfAKind(counts, 3) && isNOfAKind(counts, 2) {
		return 5
	} else if isNOfAKind(counts, 3) {
		return 4
	} else if isNOfAKind(counts, 2) && len(counts) == 3 {
		return 3
	} else if isNOfAKind(counts, 2) {
		return 2
	}

	return 1
}

func bestJHand(hand []string) int {
	counts := countCards(hand)

	// check if counts has a key of "J"
	if _, hasJ := counts["J"]; hasJ {
		maxKey := "J"
		maxValue := 0
		delete(counts, "J")
		for key, value := range counts {
			if value > maxValue {
				maxKey = key
				maxValue = value
			}
		}

		// rebuild hand but replace J with maxKey
		newHand := make([]string, 0)
		for _, card := range hand {
			if card == "J" {
				newHand = append(newHand, maxKey)
			} else {
				newHand = append(newHand, card)
			}
		}
		return bestHand(newHand)
	} else {
		return bestHand(hand)
	}
}

var cardMap map[string]int = map[string]int{"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "T": 10, "J": 1, "Q": 12, "K": 13, "A": 14}

func handCompare(a, b Hand) int {
	if a.kind < b.kind {
		return -1
	} else if a.kind > b.kind {
		return 1
	} else {
		for i := 0; i < len(a.cards); i++ {
			if cardMap[a.cards[i]] < cardMap[b.cards[i]] {
				return -1
			} else if cardMap[a.cards[i]] > cardMap[b.cards[i]] {
				return 1
			}
		}
	}
	return 0
}

func main() {
	var hands []Hand
	lines := readInputFile("input.txt")
	for _, line := range lines {
		parts := strings.Split(line, " ")
		hand := strings.Split(parts[0], "")
		bid, _ := strconv.Atoi(parts[1])
		hands = append(hands, Hand{hand, bestJHand(hand), bid})
	}

	handCompare := func(a, b Hand) int {
		return handCompare(a, b)
	}
	slices.SortFunc(hands, handCompare)

	ans := 0
	for i, hand := range hands {
		ans += hand.bid * (i + 1)
	}

	fmt.Println(ans)
}
