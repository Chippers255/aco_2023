package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"year2024/internal/aoc/day01"
	"year2024/internal/aoc/day02"
	"year2024/internal/aoc/day03"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "aoc",
		Short: "A CLI tool to run Advent of Code solutions for 2024.",
		Run: func(cmd *cobra.Command, args []string) {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter day number (1-25): ")
			dayStr, _ := reader.ReadString('\n')
			dayStr = strings.TrimSpace(dayStr)
			day, err := strconv.Atoi(dayStr)
			if err != nil {
				fmt.Println("Invalid day number.")
				return
			}

			fmt.Print("Enter input file name: ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			if err := runDay(day, input); err != nil {
				fmt.Println("Error:", err)
			}
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runDay(day int, input string) error {
	switch day {
	case 1:
		return day01.Run(input)
	case 2:
		return day02.Run(input)
	case 3:
		return day03.Run(input)
	default:
		return fmt.Errorf("day %d not implemented", day)
	}
}
