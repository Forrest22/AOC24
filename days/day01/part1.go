package day01

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// readColumnsFromFile reads a file with two columns of integers and returns the values as two slices.
func readColumnsFromFile(filename string) ([]int, []int, error) {
	// Open the input file
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var column1, column2 []int

	// Read and process the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Split the line into parts
		parts := strings.Fields(line)
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("invalid input format: each line must have exactly two integers")
		}

		// Parse each part into integers
		num1, err1 := strconv.Atoi(parts[0])
		num2, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			return nil, nil, fmt.Errorf("error converting to integers: %v, %v", err1, err2)
		}

		// Add the integers to respective columns
		column1 = append(column1, num1)
		column2 = append(column2, num2)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading file: %w", err)
	}

	if len(column1) != len(column2) {
		return nil, nil, fmt.Errorf("error with input: Column lengths different, input error.")
	}

	return column1, column2, nil
}

func computeColumnDifferences(column1 []int, column2 []int) []int {
	// Sort them, ascending
	sort.Ints(column1)
	sort.Ints(column2)

	differences := make([]int, len(column1))
	// Compute the difference for each
	for i := 0; i < len(column1); i++ {
		differences[i] = int(math.Abs(float64(column1[i] - column2[i])))
	}

	return differences
}

func day1Part1() (string, error) {
	// Call the function to read the file
	column1, column2, err := readColumnsFromFile("days/day01/input")
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}

	diffArray := computeColumnDifferences(column1, column2)
	sum := 0

	for _, num := range diffArray {
		sum += num
	}

	return strconv.Itoa(sum), nil
}
