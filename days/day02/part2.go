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
// includes Problem Dampener: The Problem Dampener is a reactor-mounted module that lets the reactor
// safety systems tolerate a single bad level in what would otherwise be a safe report.
// It's like the bad level never happened!
func readReportsFromFileLineByLineWithProblemDampener(filename string) (int, error) {
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
		isSafeReport, err := isReportSafe(scanner.Text(), false)

		if err != nil {
			return -1, err
		}

		if isSafeReport {
			safeReports++
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %w", err)
	}

	return safeReports, nil
}

func isReportSafe(line string, hasAlreadyRecursed bool) (bool, error) {
	isIncreasing := false
	isUnsafeReport := false
	// Split the line into parts
	parts := strings.Fields(line)
	var previousNum int

	// Parse each part into integers
	for i := 0; i < len(parts) && !isUnsafeReport; i++ {
		num, err := strconv.Atoi(parts[i])

		if err != nil {
			return false, fmt.Errorf("error converting to integer: %v", num)
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

	// Check if removing a level will be okay
	if !hasAlreadyRecursed && isUnsafeReport {

		for i := 0; i < len(parts); i++ {
			// Create a completely new slice excluding the element at index i
			newParts := make([]string, 0, len(parts)-1)

			// Append elements before the index
			newParts = append(newParts, parts[:i]...)

			// Append elements after the index
			newParts = append(newParts, parts[i+1:]...)

			// Join the new parts to form the modified line
			newLine := strings.Join(newParts, " ")

			// Check if the new line is safe
			isSafe, err := isReportSafe(newLine, true)
			// fmt.Printf("Checking safety for: %v, IsSafe: %v\n", newLine, isSafe)

			if err != nil {
				return false, err
			}

			// If a safe report is found, return
			if isSafe {
				return true, nil
			}
		}

	}
	return !isUnsafeReport, nil
}

func day2Part2() (string, error) {
	safeReports, err := readReportsFromFileLineByLineWithProblemDampener("days/day02/input")
	if err != nil {
		return "", fmt.Errorf("Error: %v", err)
	}

	// Output the results
	return strconv.Itoa(safeReports), nil
}
