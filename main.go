package main

import (
	"AOC24/days/day01"
	"AOC24/days/day02"
	"AOC24/days/day03"
	"AOC24/days/day04"
	"AOC24/days/day05"
	"AOC24/days/day06"
	"AOC24/days/day07"
	"AOC24/days/day08"
	"AOC24/days/day09"
	"AOC24/days/day10"
	"AOC24/days/day11"
	"AOC24/days/day12"
	"AOC24/days/day13"
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
	case 7:
		runDay(day07.Solve, part)
	case 8:
		runDay(day08.Solve, part)
	case 9:
		runDay(day09.Solve, part)
	case 10:
		runDay(day10.Solve, part)
	case 11:
		runDay(day11.Solve, part)
	case 12:
		runDay(day12.Solve, part)
	case 13:
		runDay(day13.Solve, part)
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
