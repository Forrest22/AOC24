package day04

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Define a struct to represent x, y coordinates
type WordSearchIndex struct {
	row    int
	column int
}

// directions to search
type Direction int

const (
	Horizontal Direction = iota
	HorizontalBackwards
	Vertical
	VerticalBackwards
	DiagonalNE
	DiagonalSE
	DiagonalSW
	DiagonalNW
)

func (d Direction) String() string {
	return [...]string{
		"Horizontal",
		"HorizontalBackwards",
		"Vertical",
		"VerticalBackwards",
		"DiagonalSE",
		"DiagonalNW",
		"DiagonalNE",
		"DiagonalSW",
	}[d]
}

func (d Direction) EnumIndex() int {
	return int(d)
}

func indexInDirection(index WordSearchIndex, direction Direction) WordSearchIndex {
	switch direction {
	case Horizontal:
		return WordSearchIndex{
			row:    index.row,
			column: index.column + 1,
		}
	case HorizontalBackwards:
		return WordSearchIndex{
			row:    index.row,
			column: index.column - 1,
		}
	case Vertical:
		return WordSearchIndex{
			row:    index.row + 1,
			column: index.column,
		}
	case VerticalBackwards:
		return WordSearchIndex{
			row:    index.row - 1,
			column: index.column,
		}
	case DiagonalNW:
		return WordSearchIndex{
			row:    index.row - 1,
			column: index.column - 1,
		}
	case DiagonalNE:
		return WordSearchIndex{
			row:    index.row - 1,
			column: index.column + 1,
		}
	case DiagonalSW:
		return WordSearchIndex{
			row:    index.row + 1,
			column: index.column - 1,
		}
	case DiagonalSE:
		return WordSearchIndex{
			row:    index.row + 1,
			column: index.column + 1,
		}
	default:
		fmt.Printf("Direction %d is not implemented yet.\n", direction)
		return index
	}
}

// Reads each line of a input file and starts the wordsearch
func searchWordsearchForXmas(filename string) (int, error) {
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
	targetWord := "XMAS"
	for i, row := range allText {
		for j := range row {
			sumOfMatchesAtIndex, err := getSumOfFoundMatchesFromIndex(allText, targetWord, WordSearchIndex{row: i, column: j})
			if err != nil {
				return 0, fmt.Errorf("error getting answer: %w", err)
			}
			sum += sumOfMatchesAtIndex
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %w", err)
	}

	return sum, nil
}

/*
 * This word search allows words to be horizontal, vertical, diagonal, written backwards, or even overlapping other words.
 * Must find all instances of targetWord
 */
func getSumOfFoundMatchesFromIndex(wordsearch []string, targetWord string, index WordSearchIndex) (int, error) {
	sumOfFoundMatches := 0

	// check if the first letter matches the target
	if getRuneAtWordsearchIndex(wordsearch, index) != rune(targetWord[0]) {
		return 0, nil
	}

	for dir := range DiagonalNW + 1 {
		if isMatchInDirection(wordsearch, index, targetWord, Direction(dir)) {
			sumOfFoundMatches++
		}
	}

	return sumOfFoundMatches, nil
}

func isMatchInDirection(wordsearch []string, startingIndex WordSearchIndex, targetWord string, direction Direction) bool {
	nextPoint := indexInDirection(startingIndex, direction)
	for i, targetLetter := range targetWord[1:] {
		if isWordSearchIndexOOB(wordsearch, nextPoint) {
			// OOB means no match
			return false
		}

		// check if rune matches
		if targetLetter != getRuneAtWordsearchIndex(wordsearch, nextPoint) {
			return false
		}

		nextPoint = indexInDirection(nextPoint, direction)

		// +2 looks weird here but its because i already chopped off the first index
		if len(targetWord) == i+2 {
			// found words
			return true
		}
	}
	return false
}

func isWordSearchIndexOOB(wordsearch []string, index WordSearchIndex) bool {
	// check i is inbounds
	if index.row < 0 || index.row >= len(wordsearch) {
		return true
	}
	// check j is inbounds
	if index.column < 0 || index.column >= len(wordsearch[0]) {
		return true
	}
	return false
}

func getRuneAtWordsearchIndex(wordsearch []string, index WordSearchIndex) rune {
	return []rune(wordsearch[index.row])[index.column]
}

func day4Part1() (string, error) {
	sumOfWordsFound, err := searchWordsearchForXmas("days/day04/input")
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}

	// Output the results
	return strconv.Itoa(sumOfWordsFound), nil
}
