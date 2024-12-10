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
type ValidEquation struct {
	nums       []int
	operations []rune
	result     int
}

// Reads each line of a input file and starts the wordsearch
func getSumOfValidEquations(filename string) (int, error) {
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
		validEquationFromLine, err := getValidEquationFromLine(line)
		if err != nil {
			panic("error reading equation from line")
		}
		if isValidEquation(validEquationFromLine) {
			sum += validEquationFromLine.result
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %w", err)
	}

	return sum, nil
}

func getValidEquationFromLine(line string) (ValidEquation, error) {
	// parse values
	parts := strings.Split(line, " ")
	testVal, _ := strconv.Atoi(parts[0][:len(parts[0])-1])
	testNums := make([]int, len(parts)-1)
	for i, number := range parts[1:] {
		num, _ := strconv.Atoi(number)
		testNums[i] = num
	}

	// check iterations
	return getAndCheckOperationIterations(testVal, testNums), nil
}

func getAndCheckOperationIterations(testVal int, testNums []int) ValidEquation {
	operations := []rune{'+', '*'}

	n := len(testNums) - 1                                                   // Number of gaps between numbers
	totalCombinations := int(math.Pow(float64(len(operations)), float64(n))) // Total combinations: len(operations)^n
	combinations := make([][]rune, totalCombinations)

	for i := 0; i < totalCombinations; i++ {
		combination := make([]rune, n)
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
		if isOperationSetIsValid(testVal, testNums, testOperationSet) {
			return ValidEquation{
				nums:       testNums,
				operations: testOperationSet,
				result:     testVal,
			}
		}

	}
	return ValidEquation{}
}

func isOperationSetIsValid(testVal int, testNums []int, testOperationSet []rune) bool {
	result := testNums[0]
	for i := 0; i < len(testOperationSet); i++ {
		if testOperationSet[i] == '+' {
			result = result + testNums[i+1]
		} else if testOperationSet[i] == '*' {
			result = result * testNums[i+1]
		} else {
			panic(fmt.Errorf("unknown operation: %v", testOperationSet[i]))
		}
	}
	return testVal == result
}

func isValidEquation(validEquationFromLine ValidEquation) bool {
	for _, operator := range validEquationFromLine.operations {
		if !(operator == '+' || operator == '*') {
			return false
		}
	}
	return true
}

func day7Part1() (string, error) {
	sumOfWordsFound, err := getSumOfValidEquations("days/day07/input")
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}

	// Output the results
	return strconv.Itoa(sumOfWordsFound), nil
}
