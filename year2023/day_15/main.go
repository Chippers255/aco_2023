package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readStringsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result []string

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		result = append(result, parts...)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func getStringValue(currentValue int, input string) int {
	newValue := currentValue
	for _, char := range input {
		newValue += int(char)
		newValue = newValue * 17
		newValue = newValue % 256
	}
	return newValue
}

func PartOne(filename string) int {
	input, err := readStringsFromFile(filename)
	if err != nil {
		panic(err)
	}

	ans := 0
	for _, s := range input {
		this := getStringValue(0, s)
		ans += this
		// fmt.Println(s, this)
	}

	return ans
}

type lens struct {
	label  string
	hash   int
	length int
}

func handleEqual(input string, muhMap map[int][]lens) map[int][]lens {
	newMap := make(map[int][]lens)
	for i := 0; i < 256; i++ {
		newMap[i] = []lens{}
	}
	for k, v := range muhMap {
		newMap[k] = v
	}

	parts := strings.Split(input, "=")
	label := parts[0]
	length, _ := strconv.Atoi(parts[1])
	hash := getStringValue(0, label)
	thisLens := lens{label, hash, length}

	notFound := true
	for i, l := range newMap[hash] {
		if l.label == label {
			newMap[hash][i] = thisLens
			notFound = false
			break
		}
	}
	if notFound {
		newMap[hash] = append(newMap[hash], thisLens)
	}

	return newMap
}

func handleDash(input string, muhMap map[int][]lens) map[int][]lens {
	newMap := make(map[int][]lens)
	for i := 0; i < 256; i++ {
		newMap[i] = []lens{}
	}
	for k, v := range muhMap {
		newMap[k] = v
	}

	parts := strings.Split(input, "-")
	label := parts[0]
	hash := getStringValue(0, label)

	for i, l := range newMap[hash] {
		if l.label == label {
			newMap[hash] = append(newMap[hash][:i], newMap[hash][i+1:]...)
			break
		}
	}

	return newMap
}

func PartTwo(filename string) int {
	input, err := readStringsFromFile(filename)
	if err != nil {
		panic(err)
	}
	muhMap := make(map[int][]lens)
	for i := 0; i < 256; i++ {
		muhMap[i] = []lens{}
	}

	for _, s := range input {
		if strings.Contains(s, "-") {
			muhMap = handleDash(s, muhMap)
		} else if strings.Contains(s, "=") {
			muhMap = handleEqual(s, muhMap)
		} else {
			panic("bad input")
		}
	}

	ans := 0
	for k, v := range muhMap {
		for i, l := range v {
			this := (1 + k) * (i + 1) * l.length
			ans += this
		}
	}

	return ans
}

func main() {
	fmt.Println(PartOne("input.txt"))
	fmt.Println(PartTwo("input.txt"))
}
