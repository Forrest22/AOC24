package day03

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// Reads each line of a list of reports (each line a report) and returns the number of safe reports.
// Safe if:
// - The levels are either all increasing or all decreasing.
// - Any two adjacent levels differ by at least one and at most three.
func readCorruptedMemoryFromFileLineByLine(filename string) (int, error) {
	// Open the input file
	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var answer int
	// Read and process the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		lineSum, err := getLineSum(line)

		answer += lineSum

		if err != nil {
			return 0, fmt.Errorf("error getting answer: %w", err)
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %w", err)
	}

	return answer, nil
}

func getLineSum(line string) (int, error) {
	/*
	 * Regex Explanation:
	 * mul\(:
	 * 	Matches mul( literally.
	 * (\d{1,3}):
	 * 	Captures the first 1-3 digit number.
	 * ,:
	 * 	Matches the comma separating the two numbers.
	 * (\d{1,3}):
	 * 	Captures the second 1-3 digit number.
	 * \):
	 * 	Matches the closing parenthesis.
	 */
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	matches := re.FindAllStringSubmatch(line, -1)

	var sum int
	for _, match := range matches {
		multipland, _ := strconv.Atoi(match[1])
		multiplier, _ := strconv.Atoi(match[2])
		sum += (multipland * multiplier)
	}
	return sum, nil
}

func day3Part1() (string, error) {
	safeReports, err := readCorruptedMemoryFromFileLineByLine("days/day03/input")
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}

	// Output the results
	return strconv.Itoa(safeReports), nil
}
