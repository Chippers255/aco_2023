# Advent of Code 2024

## About Advent of Code
Advent of Code is an annual December event where programmers sharpen their coding skills through a series of daily programming challenges. This repository contains my solutions for the Advent of Code 2024 implemented in golang.

## Repository tructure
This years solutions are organized into a single CLI tool. The tool can be started and the user will be prompted for a day and an input file.
- `cmd/`: CLI application entry point.
- `internal/aoc/`: Contains solutions for each day.
- `inputs/`: Input files for each day.
- `.github/workflows/`: GitHub Actions workflows for CI.

## Usage

```bash
go build -o aoc ./cmd
./aoc
```