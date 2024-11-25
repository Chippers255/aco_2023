package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var memoTable = make(map[string]int)

func makeAMemoKey(s []string, c []int, last_block []string) string {
	return fmt.Sprintf("%v %v %v", s, c, last_block)
}

func readMemo(s []string, c []int, last_block []string) (int, bool) {
	memoKey := makeAMemoKey(s, c, last_block)
	val, ok := memoTable[memoKey]
	return val, ok
}

func writeMemo(s []string, c []int, last_block []string, val int) {
	memoKey := makeAMemoKey(s, c, last_block)
	memoTable[memoKey] = val
}

func makeItTiny(s []string, c []int, last_block []string) int {
	// fmt.Println("s", s, "c", c, "last_block", last_block)
	if val, ok := readMemo(s, c, last_block); ok {
		return val
	}

	if len(s) == 0 && len(c) == 0 && len(last_block) == 0 {
		// fmt.Println("Winner Gagnant!")
		return 1
	}

	if len(s) == 0 {
		if len(last_block) == c[len(c)-1] {
			val := makeItTiny([]string{}, c[:len(c)-1], []string{})
			writeMemo([]string{}, c[:len(c)-1], []string{}, val)
			return val
		} else {
			// fmt.Println("JUST A BIT O' BANTER")
			return 0
		}
	} else {
		if s[len(s)-1] == "#" {
			if len(c) == 0 {
				// fmt.Println("JUST A BIT O' BANTER")
				return 0
			}
			newBlock := make([]string, len(last_block))
			copy(newBlock, last_block)
			newBlock = append(newBlock, "#")
			if len(newBlock) > c[len(c)-1] {
				// fmt.Println("JUST A BIT O' BANTER")
				return 0
			} else {
				val := makeItTiny(s[:len(s)-1], c, newBlock)
				writeMemo(s[:len(s)-1], c, newBlock, val)
				return val
			}
		}

		if s[len(s)-1] == "." {
			if len(last_block) == 0 {
				val := makeItTiny(s[:len(s)-1], c, last_block)
				writeMemo(s[:len(s)-1], c, last_block, val)
				return val
			}
			if len(last_block) == c[len(c)-1] {
				val := makeItTiny(s[:len(s)-1], c[:len(c)-1], []string{})
				writeMemo(s[:len(s)-1], c[:len(c)-1], []string{}, val)
				return val
			} else {
				// fmt.Println("JUST A BIT O' BANTER")
				return 0
			}
		}

		if s[len(s)-1] == "?" {
			s1 := make([]string, len(s)-1)
			copy(s1, s[:len(s)-1])
			s1 = append(s1, "#")

			s2 := make([]string, len(s)-1)
			copy(s2, s[:len(s)-1])
			s2 = append(s2, ".")

			val1 := makeItTiny(s1, c, last_block)
			writeMemo(s1, c, last_block, val1)

			val2 := makeItTiny(s2, c, last_block)
			writeMemo(s2, c, last_block, val2)
			return val1 + val2
		}
	}

	panic("Shouldn't get here")
}

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

func main() {
	lines := getFileLines("input.txt")
	ans := 0

	for _, line := range lines {
		parts := strings.Split(line, " ")

		s := strings.Repeat(parts[0]+"?", 4) + parts[0]
		newS := strings.Split(s, "")

		c := strings.Repeat(parts[1]+",", 4) + parts[1]
		nC := strings.Split(c, ",")
		newC := make([]int, len(nC))
		for i, v := range nC {
			num, err := strconv.Atoi(v)
			if err != nil {
				panic(err) // handle the error appropriately
			}
			newC[i] = num
		}

		last_block := []string{}
		a := makeItTiny(newS, newC, last_block)
		fmt.Println(s, a)
		ans += a
	}
	fmt.Println(ans)
}
