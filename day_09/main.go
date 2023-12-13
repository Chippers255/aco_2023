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

func getDiffs(nums []int) []int {
	var diffs []int
	for i := 0; i < len(nums)-1; i++ {
		diffs = append(diffs, nums[i+1]-nums[i])
	}
	return diffs
}

func predictotron(seq [][]int) int {
	ans := 0
	for _, s := range seq {
		ans += s[len(s)-1]
	}
	return ans
}

func PartOne(filename string) int {
	score := 0
	lines := getFileLines(filename)
	for _, line := range lines {
		seq := [][]int{}
		splitLine := strings.Split(line, " ")
		var nums []int
		for _, num := range splitLine {
			n, err := strconv.Atoi(num)
			if err != nil {
				log.Fatal(err)
			}
			nums = append(nums, n)
		}
		seq = append(seq, nums)
		d := getDiffs(nums)
		seq = append(seq, d)

		for {
			lastSeq := seq[len(seq)-1]
			allZeros := true
			for _, value := range lastSeq {
				if value != 0 {
					allZeros = false
					break
				}
			}
			if allZeros {
				break
			}
			d = getDiffs(lastSeq)
			seq = append(seq, d)
		}
		ans := predictotron(seq)
		score += ans
	}
	return score
}

func redactotron(seq [][]int) int {
	ans := 0
	for i, s := range seq {
		if i == 0 {
			ans = s[0]
		} else if i%2 == 0 {
			ans += s[0]
		} else {
			ans -= s[0]
		}
	}
	return ans
}

func PartTwo(filename string) int {
	score := 0
	lines := getFileLines(filename)
	for _, line := range lines {
		seq := [][]int{}
		splitLine := strings.Split(line, " ")
		var nums []int
		for _, num := range splitLine {
			n, err := strconv.Atoi(num)
			if err != nil {
				log.Fatal(err)
			}
			nums = append(nums, n)
		}
		seq = append(seq, nums)
		d := getDiffs(nums)
		seq = append(seq, d)

		for {
			lastSeq := seq[len(seq)-1]
			allZeros := true
			for _, value := range lastSeq {
				if value != 0 {
					allZeros = false
					break
				}
			}
			if allZeros {
				break
			}
			d = getDiffs(lastSeq)
			seq = append(seq, d)
		}
		ans := redactotron(seq)
		score += ans
	}
	return score
}

func main() {
	ans := PartOne("input.txt")
	fmt.Println(ans)

	ans = PartTwo("input.txt")
	fmt.Println(ans)
}
