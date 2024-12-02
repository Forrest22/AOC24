package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadColumnsFromFileButSecondColumnIsADict reads a file with two columns of integers and returns the values as one array and one map.
func ReadColumnsFromFileButSecondColumnIsADict(filename string) ([]int, map[int]int, error) {
	// Open the input file
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var column1 []int
	frequencyMap := make(map[int]int, len(column1))

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
		frequencyMap[num2] = frequencyMap[num2] + 1
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading file: %w", err)
	}

	return column1, frequencyMap, nil
}

func ComputeColumnDifferencesOfMap(column1 []int, frequencyMap map[int]int) []int {
	frequencies := make([]int, len(column1))
	// Compute the difference for each
	for i := 0; i < len(column1); i++ {
		frequencies[i] = int(column1[i] * frequencyMap[column1[i]])
	}

	// fmt.Println("freq: ", frequencies)

	return frequencies
}

func main() {
	// Call the function to read the file
	column1, frequencyMap, err := ReadColumnsFromFileButSecondColumnIsADict("input")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	diffArray := ComputeColumnDifferencesOfMap(column1, frequencyMap)
	sum := 0

	for _, num := range diffArray {
		sum += num
	}

	// Output the results
	fmt.Println("Sum: ", sum)
	// fmt.Println("Column 1:", column1)
	// fmt.Println("Column 2:", column2)

}
