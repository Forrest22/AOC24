package day04

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Reads each line of a input file and starts the wordsearch
func searchWordsearchForLiteralXMas(filename string) (int, error) {
	// Open the input file
	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Read and process the file line by line
	scanner := bufio.NewScanner(file)
	var allText []string
	for scanner.Scan() {
		allText = append(allText, scanner.Text())
	}

	var sum int
	for i, row := range allText {
		for j := range row {
			if isIndexPartOfXMas(allText, WordSearchIndex{row: i, column: j}) {
				sum++
			}
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %w", err)
	}

	return sum, nil
}

func getDiagonalDirectionFromIotaInt(i int) Direction {
	return Direction((i % 4) + 4)
}

/*
 * This word search allows words to be horizontal, vertical, diagonal, written backwards, or even overlapping other words.
 * Must find all instances of targetWord
 */
func isIndexPartOfXMas(wordsearch []string, index WordSearchIndex) bool {

	// check if the first letter matches the target, in this case we want the centerpeice 'A'
	if getRuneAtWordsearchIndex(wordsearch, index) != rune('A') {
		return false
	}

	sumOfFoundMatches := 0

	for i := range 4 {
		if isXMasMatchInDirection(wordsearch, index, getDiagonalDirectionFromIotaInt(i)) {
			sumOfFoundMatches++
		}
	}

	return sumOfFoundMatches == 2
}

// because we're statring at that 'A' we need to look diagnonally opposite for M and S
func isXMasMatchInDirection(wordsearch []string, startingIndex WordSearchIndex, direction Direction) bool {
	previousPoint := indexInDirection(startingIndex, getDiagonalDirectionFromIotaInt(int(direction)+2))
	previousTarget := 'M'
	nextPoint := indexInDirection(startingIndex, direction)
	nextTarget := 'S'

	if isWordSearchIndexOOB(wordsearch, previousPoint) {
		// OOB means no match
		return false
	}

	// check if rune matches in opposite directions
	if previousTarget != getRuneAtWordsearchIndex(wordsearch, previousPoint) {
		return false
	}

	if isWordSearchIndexOOB(wordsearch, nextPoint) {
		// OOB means no match
		return false
	}

	// check if rune matches in opposite directions
	if nextTarget != getRuneAtWordsearchIndex(wordsearch, nextPoint) {
		return false
	}

	return true
}

func day4Part2() (string, error) {
	sumOfWordsFound, err := searchWordsearchForLiteralXMas("days/day04/input")
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}

	// Output the results
	return strconv.Itoa(sumOfWordsFound), nil
}
