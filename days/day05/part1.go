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
func sumOfCorrectMiddlePages(filename string) (int, error) {
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
				pageNums := parsePageNums(pagesSplitAsString)
				middlePagesToSum = append(middlePagesToSum, pageNums[len(pageNums)/2])
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

func parsePageNums(pages []string) []int {
	var pageNums []int
	for _, p := range pages {
		pageNum, err := strconv.Atoi(string(p))
		if err != nil {
			panic(fmt.Errorf("error parsing 3 int: %v", err))
		}
		pageNums = append(pageNums, pageNum)
	}
	return pageNums
}

func isUpdateValid(pagesSplitAsString []string, rulesDict map[int][]int) bool {
	pageNums := parsePageNums(pagesSplitAsString)
	for i, page := range pageNums {
		if !isValidSoFar(page, pageNums[:i], rulesDict) {
			return false
		}
	}
	return true
}

// checks the previously scanned pages to check if they comply with the rules
func isValidSoFar(pageToCheck int, previousPagesInUpdate []int, rulesDict map[int][]int) bool {
	for _, prevPage := range previousPagesInUpdate {
		if slices.Contains(rulesDict[pageToCheck], prevPage) {
			return false
		}
	}
	return true
}

func day5Part1() (string, error) {
	sumOfWordsFound, err := sumOfCorrectMiddlePages("days/day05/input")
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}

	// Output the results
	return strconv.Itoa(sumOfWordsFound), nil
}
