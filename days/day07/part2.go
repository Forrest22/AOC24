package day07

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Define a struct to represent equations
type ValidEquationPlus struct {
	nums       []int
	operations []string
	result     int
}

// Reads each line of a input file and starts the wordsearch
func getSumOfValidEquationsWithConcat(filename string) (int, error) {
	// Open the input file
	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Read and process the file line by line
	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		validEquationFromLine, err := getValidEquationFromLineWithConcat(line)
		if err != nil {
			panic("error reading equation from line")
		}
		if isValidEquationWithConcat(validEquationFromLine) {
			sum += validEquationFromLine.result
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %w", err)
	}

	return sum, nil
}

func getValidEquationFromLineWithConcat(line string) (ValidEquationPlus, error) {
	// parse values
	parts := strings.Split(line, " ")
	testVal, _ := strconv.Atoi(parts[0][:len(parts[0])-1])
	testNums := make([]int, len(parts)-1)
	for i, number := range parts[1:] {
		num, _ := strconv.Atoi(number)
		testNums[i] = num
	}

	// check iterations
	return getAndCheckOperationIterationsWithConcat(testVal, testNums), nil
}

func getAndCheckOperationIterationsWithConcat(testVal int, testNums []int) ValidEquationPlus {
	operations := []string{
		"+",
		"*",
		"||",
	}

	n := len(testNums) - 1                                                   // Number of gaps between numbers
	totalCombinations := int(math.Pow(float64(len(operations)), float64(n))) // Total combinations: len(operations)^n
	combinations := make([][]string, totalCombinations)

	for i := 0; i < totalCombinations; i++ {
		combination := make([]string, n)
		temp := i
		for j := 0; j < n; j++ {
			// Select the operation
			combination[j] = operations[temp%len(operations)]
			// Move to the next "digit"
			temp /= len(operations)
		}
		combinations[i] = combination
	}

	for _, testOperationSet := range combinations {
		if isOperationSetIsValidWithConcat(testVal, testNums, testOperationSet) {
			return ValidEquationPlus{
				nums:       testNums,
				operations: testOperationSet,
				result:     testVal,
			}
		}

	}
	return ValidEquationPlus{}
}

func isOperationSetIsValidWithConcat(testVal int, testNums []int, testOperationSet []string) bool {
	result := testNums[0]
	for i := 0; i < len(testOperationSet); i++ {
		if testOperationSet[i] == "+" {
			result = result + testNums[i+1]
		} else if testOperationSet[i] == "*" {
			result = result * testNums[i+1]
		} else if testOperationSet[i] == "||" {
			left := strconv.Itoa(result)
			right := strconv.Itoa(testNums[i+1])
			tempResult := left + right
			result, _ = strconv.Atoi(tempResult)
		} else {
			panic(fmt.Errorf("unknown operation: %v", testOperationSet[i]))
		}
	}
	return testVal == result
}

func isValidEquationWithConcat(validEquationFromLine ValidEquationPlus) bool {
	for _, operator := range validEquationFromLine.operations {
		if !(operator == "+" || operator == "*" || operator == "||") {
			return false
		}
	}
	return true
}

func day7Part2() (string, error) {
	sumOfWordsFound, err := getSumOfValidEquationsWithConcat("days/day07/input")
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}

	// Output the results
	return strconv.Itoa(sumOfWordsFound), nil
}
