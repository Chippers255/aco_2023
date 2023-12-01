package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func getFirstLastP1(s string) (string, string) {
	var first, last string
	foundFirst := false

	for _, c := range s {
		if unicode.IsNumber(c) {
			if !foundFirst {
				first = string(c)
				foundFirst = true
			}
			last = string(c)
		}
	}
	return first, last
}

func getFirstLastP2(s string) (string, string) {
	subs := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	dictionary := make(map[string]string)
	dictionary["0"] = "0"
	dictionary["1"] = "1"
	dictionary["2"] = "2"
	dictionary["3"] = "3"
	dictionary["4"] = "4"
	dictionary["5"] = "5"
	dictionary["6"] = "6"
	dictionary["7"] = "7"
	dictionary["8"] = "8"
	dictionary["9"] = "9"
	dictionary["zero"] = "0"
	dictionary["one"] = "1"
	dictionary["two"] = "2"
	dictionary["three"] = "3"
	dictionary["four"] = "4"
	dictionary["five"] = "5"
	dictionary["six"] = "6"
	dictionary["seven"] = "7"
	dictionary["eight"] = "8"
	dictionary["nine"] = "9"

	var first, last string
	firstIndex := len(s)
	lastIndex := -1

	for _, sub := range subs {
		index := strings.Index(s, sub)
		if index <= firstIndex && index != -1 {
			firstIndex = index
			first, _ = dictionary[sub]
		}
		index = strings.LastIndex(s, sub)
		if index >= lastIndex && index != -1 {
			lastIndex = index
			last, _ = dictionary[sub]
		}
	}
	fmt.Println(first, last)
	return first, last

}

func readInputFileP1(filename string) []string {
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

func partOne() {
	lines := readInputFileP1("input.txt")
	var numbers []int
	for _, line := range lines {
		first, last := getFirstLastP1(line)
		num, err := strconv.Atoi(first + last)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, num)
	}
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Println(sum)
}

func readInputFileP2(filename string) []string {
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

func partTwo() {
	lines := readInputFileP2("input.txt")
	var numbers []int
	for _, line := range lines {
		first, last := getFirstLastP2(line)
		num, err := strconv.Atoi(first + last)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, num)
	}
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Println(sum)
}

func main() {
	//partOne()
	partTwo()
}
