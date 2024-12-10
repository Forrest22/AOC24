package day06

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Define a struct to represent x, y coordinates
type MapIndex struct {
	row    int
	column int
}

// directions to search
type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{
		"North",
		"East",
		"South",
		"West",
	}[d]
}

func (d Direction) EnumIndex() int {
	return int(d)
}

func indexInDirection(index MapIndex, direction Direction) MapIndex {
	switch direction {
	case North:
		return MapIndex{
			row:    index.row - 1,
			column: index.column,
		}
	case East:
		return MapIndex{
			row:    index.row,
			column: index.column + 1,
		}
	case South:
		return MapIndex{
			row:    index.row + 1,
			column: index.column,
		}
	case West:
		return MapIndex{
			row:    index.row,
			column: index.column - 1,
		}
	default:
		fmt.Printf("Direction %d is not implemented yet.\n", direction)
		return index
	}
}

func isMapIndexOOB(mapOfRoom []string, index MapIndex) bool {
	// check i is inbounds
	if index.row < 0 || index.row >= len(mapOfRoom) {
		return true
	}
	// check j is inbounds
	if index.column < 0 || index.column >= len(mapOfRoom[0]) {
		return true
	}
	return false
}

// Reads each line of a input file and starts the wordsearch
func getGuardDistictPositions(filename string) (int, error) {
	// Open the input file
	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Read and process the file line by line
	scanner := bufio.NewScanner(file)
	var guardMap []string
	var guardStartingPoint MapIndex

	for scanner.Scan() {
		line := scanner.Text()
		guardMap = append(guardMap, line)

		// check for starting position
		if strings.Contains(line, "^") {
			guardStartingPoint = MapIndex{
				row:    len(guardMap) - 1,
				column: strings.Index(line, "^"),
			}

		}
	}

	// recursively move in direction until leaves
	// Convert the string to a slice of runes to allow modification
	rowAsRunes := []rune(guardMap[guardStartingPoint.row])
	rowAsRunes[guardStartingPoint.column] = 'X'           // Modify the character
	guardMap[guardStartingPoint.row] = string(rowAsRunes) // Convert back to string
	recursivelyMoveGuard(&guardMap, guardStartingPoint, North)

	sum := 0
	for _, lineOnMap := range guardMap {
		sum += strings.Count(lineOnMap, "X")
	}
	return sum, nil
}

func recursivelyMoveGuard(guardMap *[]string, guardStartingPoint MapIndex, direction Direction) {
	if isMapIndexOOB(*guardMap, indexInDirection(guardStartingPoint, direction)) {
		return
	}
	newGuardLocation := indexInDirection(guardStartingPoint, direction)
	if (*guardMap)[newGuardLocation.row][newGuardLocation.column] == '#' {
		// hit a wall, rotate 90*, try again
		direction = (direction + 1) % 4
		newGuardLocation = indexInDirection(guardStartingPoint, direction)
		if (*guardMap)[newGuardLocation.row][newGuardLocation.column] == '#' {
			fmt.Println("hit wall depth 2")
			// hit a wall, rotate 90*, try again
			direction = (direction + 1) % 4
			newGuardLocation := indexInDirection(guardStartingPoint, direction)
			if (*guardMap)[newGuardLocation.row][newGuardLocation.column] == '#' {
				fmt.Println("hit wall depth 3")

				// hit a wall, rotate 90*, try again
				direction = (direction + 1) % 4
				newGuardLocation := indexInDirection(guardStartingPoint, direction)
				if (*guardMap)[newGuardLocation.row][newGuardLocation.column] == '#' {
					fmt.Println("hit wall depth 4")
					panic("guard stuck")
				}
			}
		}
	}

	// Convert the string to a slice of runes to allow modification
	rowAsRunes := []rune((*guardMap)[newGuardLocation.row])
	rowAsRunes[newGuardLocation.column] = 'X'              // Modify the character
	(*guardMap)[newGuardLocation.row] = string(rowAsRunes) // Convert back to string

	recursivelyMoveGuard(guardMap, newGuardLocation, direction)
}

func day6Part1() (string, error) {
	sumOfWordsFound, err := getGuardDistictPositions("days/day06/input")
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}

	// Output the results
	return strconv.Itoa(sumOfWordsFound), nil
}
