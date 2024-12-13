package day05

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// Reads each line of a input file and starts the wordsearch
func sumOfMiddlePagesOfOnlyFixedLines(filename string) (int, error) {
	// Open the input file
	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Read and process the file line by line
	scanner := bufio.NewScanner(file)
	var isSecondPart = false
	var rulesDict = make(map[int][]int)
	var middlePagesToSum []int

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			isSecondPart = true
		}

		if !isSecondPart {
			// checking rules
			orderingRulePages := strings.Split(line, "|")
			firstPageKey, err := strconv.Atoi(orderingRulePages[0])
			if err != nil {
				return 0, fmt.Errorf("error parsing 1 int: %v", err)
			}

			followingPage, err := strconv.Atoi(orderingRulePages[1])
			if err != nil {
				return 0, fmt.Errorf("error parsing 2 int: %v", err)
			}
			rulesDict[firstPageKey] = append(rulesDict[firstPageKey], followingPage)
		}
		if isSecondPart && line != "" {
			// creating rules (second part)
			pagesSplitAsString := strings.Split(line, ",")

			if isUpdateValid(pagesSplitAsString, rulesDict) {
				// do nothing
			} else {
				// reorder and append middle number
				validOrder := getValidOrder(pagesSplitAsString, rulesDict)
				middlePagesToSum = append(middlePagesToSum, validOrder[len(validOrder)/2])
			}
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %w", err)
	}

	sum := 0
	for _, middlePage := range middlePagesToSum {
		sum += middlePage
	}

	return sum, nil
}

func getValidOrder(pagesSplitAsString []string, rulesDict map[int][]int) []int {
	// create relevant rules dict
	pageNums := parsePageNums(pagesSplitAsString)
	for pageIndex, page := range pageNums {
		for prevPageIndex, prevPage := range pageNums[:pageIndex] {
			if slices.Contains(rulesDict[page], prevPage) {
				pageNums[pageIndex], pageNums[prevPageIndex] = pageNums[prevPageIndex], pageNums[pageIndex]
			}
		}
	}
	return pageNums
}

func day5Part2() (string, error) {
	sumOfMiddlePages, err := sumOfMiddlePagesOfOnlyFixedLines("days/day05/input")
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}

	// Output the results
	return strconv.Itoa(sumOfMiddlePages), nil
}
