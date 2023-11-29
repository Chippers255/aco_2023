# Advent of Code 2023 Solutions

##About Advent of Code
Advent of Code is an annual December event where programmers sharpen their coding skills through a series of daily programming challenges. This repository contains my solutions for the Advent of Code 2023. Each solution is implemented in Go, providing an excellent opportunity to delve into Go's capabilities and features.

## Repository Structure
The repository is organized into separate directories for each day of the challenge, following the naming convention day_xx where xx represents the day number (e.g., day_01, day_02, etc.). Inside each directory, you will find:

- **Solution Files:** The Go source files containing the solutions for the day's challenge.
- **Test Files:** Unit test files for validating the solutions.
- **main.go:** A Go file to run the solution for the day's challenge.

## Getting Started
To get started with these solutions, you need to have Go installed on your machine. Once installed, follow these steps:

### 1. Clone the Repository
First, clone this repository to your local machine using:

```bash
git clone https://github.com/Chippers255/advent-of-code
```

### 2. Navigate to a Challenge Directory
Change into the directory of the day you're interested in:

```bash
cd day_xx
```

### 3. Initialize Go Modules
Initialize a new module (only needed once per day):

```bash
go mod init github.com/chippers255/advent-of-code/day_xx
```

This step creates a new go.mod file in the directory.

### 4. Running Solutions
To run the solution for a particular day, use:
```bash
go run main.go
```

### 5. Building Executables
If you want to compile the program into an executable file, use:
```bash
go build
```

This will create an executable named after the directory in your current folder.

### 6. Running Tests
To run tests and see the coverage, use:

```bash
go test -cover
```

This command will run all tests in the directory and show the percentage of the code covered by these tests.