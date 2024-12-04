package day03

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Reads each line of a list of reports (each line a report) and returns the number of safe reports.
// Safe if:
// - The levels are either all increasing or all decreasing.
// - Any two adjacent levels differ by at least one and at most three.
func readCorruptedMemoryWithConditionals(filename string) (int, error) {
	// Open the input file
	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var answer int
	// Read and process the file line by line
	scanner := bufio.NewScanner(file)
	var allText string
	for scanner.Scan() {
		allText += scanner.Text()
	}

	lineSum, err := getLineSumConditional(allText)

	answer += lineSum

	if err != nil {
		return 0, fmt.Errorf("error getting answer: %w", err)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %w", err)
	}

	return answer, nil
}

func getLineSumConditional(line string) (int, error) {
	splitDonts := strings.Split(line, "don't()")

	var enabledLines []string
	enabledLines = append(enabledLines, splitDonts[0])

	for _, splitLine := range splitDonts {
		dontLineSplitNowByDos := strings.Split(splitLine, "do()")
		enabledLines = append(enabledLines, dontLineSplitNowByDos[1:]...)
	}

	enabledSum := 0
	for _, enabledLine := range enabledLines {
		enabledLineSum, err := getLineSum(enabledLine)
		if err != nil {
			return -1, fmt.Errorf("error: %v", err)
		}
		enabledSum += enabledLineSum
	}
	return enabledSum, nil
}

func day3Part2() (string, error) {
	safeReports, err := readCorruptedMemoryWithConditionals("days/day03/input")
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}

	// Output the results
	return strconv.Itoa(safeReports), nil
}
