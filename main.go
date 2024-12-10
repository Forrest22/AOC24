package main

import (
	"AOC24/days/day01"
	"AOC24/days/day02"
	"AOC24/days/day03"
	"AOC24/days/day04"
	"AOC24/days/day05"
	"AOC24/days/day06"
	"fmt"
	"os"
	"strconv"
	// Add imports for more days as needed
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <day> <part>")
		return
	}

	// Parse arguments
	day, err := strconv.Atoi(os.Args[1])
	if err != nil || day < 1 {
		fmt.Println("Invalid day argument. Please provide a positive integer.")
		return
	}

	part, err := strconv.Atoi(os.Args[2])
	if err != nil || (part != 1 && part != 2) {
		fmt.Println("Invalid part argument. Please provide 1 or 2.")
		return
	}

	// Route to the correct day and part
	switch day {
	case 1:
		runDay(day01.Solve, part)
	case 2:
		runDay(day02.Solve, part)
	case 3:
		runDay(day03.Solve, part)
	case 4:
		runDay(day04.Solve, part)
	case 5:
		runDay(day05.Solve, part)
	case 6:
		runDay(day06.Solve, part)
	// Add more cases for additional days
	default:
		fmt.Printf("Day %d is not implemented yet.\n", day)
	}
}

// Helper function to run a day's solution with error handling
func runDay(solver func(int) (string, error), part int) {
	result, err := solver(part)
	if err != nil {
		fmt.Printf("Error running solution: %v\n", err)
		return
	}
	fmt.Printf("Solution: %s\n", result)
}
