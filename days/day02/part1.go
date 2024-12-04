package day02

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
func readReportsFromFileLineByLine(filename string) (int, error) {
	// Open the input file
	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	safeReports := 0

	// Read and process the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		isIncreasing := false
		isUnsafeReport := false
		line := scanner.Text()
		// Split the line into parts
		parts := strings.Fields(line)
		var previousNum int

		// Parse each part into integers
		for i := 0; i < len(parts) && !isUnsafeReport; i++ {
			num, err := strconv.Atoi(parts[i])

			if err != nil {
				return 0, fmt.Errorf("error converting to integer: %v", num)
			}

			// If not the first int in report
			if i > 0 {
				// Sets increasing/decreasing
				if i == 1 && num > previousNum {
					isIncreasing = true
				}

				// Check safety (must differ by at least 1 and at most 3)
				if isIncreasing {
					if (num-previousNum) < 1 || (num-previousNum) > 3 {
						isUnsafeReport = true
					}
				} else {
					if (previousNum-num) < 1 || (previousNum-num) > 3 {
						isUnsafeReport = true
					}
				}
			}
			previousNum = num
		}

		if !isUnsafeReport {
			safeReports++
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %w", err)
	}

	return safeReports, nil
}

func day2Part1() (string, error) {
	safeReports, err := readReportsFromFileLineByLine("days/day02/input")
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}

	// Output the results
	return strconv.Itoa(safeReports), nil
}
