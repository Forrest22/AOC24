package day06

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// quick helper function for modifying strings
func modifyStringAt(s string, index int, newChar rune) string {
	runes := []rune(s)
	runes[index] = newChar
	return string(runes)
}

// Reads each line of a input file and starts the wordsearch
func getObstructionUniquePositions(filename string) (int, error) {
	// Open the input file
	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Read and process the file line by line
	scanner := bufio.NewScanner(file)
	var guardMapOriginal []string
	var guardStartingPoint MapIndex

	for scanner.Scan() {
		line := scanner.Text()
		guardMapOriginal = append(guardMapOriginal, line)

		// check for starting position
		if strings.Contains(line, "^") {
			guardStartingPoint = MapIndex{
				row:    len(guardMapOriginal) - 1,
				column: strings.Index(line, "^"),
			}

		}
	}

	guardMapTraced := make([]string, len(guardMapOriginal))
	copy(guardMapTraced, guardMapOriginal)

	// recursively move in direction until leaves
	recursivelyMoveGuard(&guardMapTraced, guardStartingPoint, North)

	sum := 0

	// now find where obstructions can go :(
	for rowIndex, rowOnMap := range guardMapTraced {
		for colIndex := range rowOnMap {
			// fmt.Println(rowOnMap, rowIndex, colIndex, guardStartingPoint)
			if !(guardStartingPoint.row == rowIndex && guardStartingPoint.column == colIndex) {
				fmt.Println("Adding an obsticle", rowIndex, colIndex)

				// check if adding an obsticle makes a loop
				guardMapWithObsticle := make([]string, len(guardMapOriginal))
				copy(guardMapWithObsticle, guardMapOriginal)

				guardMapWithObsticle[rowIndex] = modifyStringAt(guardMapWithObsticle[rowIndex], colIndex, '#')

				// for _, lineOnMap := range guardMapWithObsticle {
				// 	println(lineOnMap)
				// }

				locationCount := make([][]int, len(guardMapOriginal))
				for i := range locationCount {
					locationCount[i] = make([]int, len(guardMapOriginal[0]))
				}

				if rowIndex == 5 && colIndex == 2 {
					fmt.Println("Obstacle:", MapIndex{row: rowIndex, column: colIndex})
					for _, line := range guardMapWithObsticle {
						fmt.Println(line)
					}
					for _, line := range locationCount {
						stringLine := ""
						for _, count := range line {
							stringLine += strconv.Itoa(count)
						}
						fmt.Println(stringLine)
					}
				}

				if doesRecursivelyRunningGuardFindALoop(&guardMapWithObsticle, guardStartingPoint, North, locationCount) {
					fmt.Println("Obstacle creating a loop found:", MapIndex{row: rowIndex, column: colIndex})
					sum++
				}
				if rowIndex == 5 && colIndex == 2 {
					fmt.Println("Obstacle found:", MapIndex{row: rowIndex, column: colIndex})
					for _, line := range guardMapWithObsticle {
						fmt.Println(line)
					}
					for _, line := range locationCount {
						stringLine := ""
						for _, count := range line {
							stringLine += strconv.Itoa(count)
						}
						fmt.Println(stringLine)
					}
				}
			}
		}
	}

	for _, lineOnMap := range guardMapTraced {
		println(lineOnMap)
	}

	return sum, nil
}

func doesRecursivelyRunningGuardFindALoop(guardMap *[]string, guardStartingPoint MapIndex, direction Direction, locationCount [][]int) bool {
	if isMapIndexOOB(*guardMap, indexInDirection(guardStartingPoint, direction)) {
		// next step is oob
		// so need to manually check this?

		locationCount[guardStartingPoint.row][guardStartingPoint.column]++

		if locationCount[guardStartingPoint.row][guardStartingPoint.column] > 4 {
			// fmt.Println("Location hit 5 times!", MapIndex{row: guardStartingPoint.row, column: guardStartingPoint.column})
			// for _, line := range *guardMap {
			// 	fmt.Println(line)
			// }
			return true
		}

		(*guardMap)[guardStartingPoint.row] = modifyStringAt((*guardMap)[guardStartingPoint.row], guardStartingPoint.column, 'X')
		return false
	}
	newDirection := direction
	newGuardLocation := indexInDirection(guardStartingPoint, newDirection)
	if (*guardMap)[newGuardLocation.row][newGuardLocation.column] == '#' {
		fmt.Println("hit wall location", newGuardLocation, newDirection)
		// hit a wall, rotate 90*, try again
		newDirection = (newDirection + 1) % 4
		newGuardLocation = indexInDirection(guardStartingPoint, newDirection)
		if (*guardMap)[newGuardLocation.row][newGuardLocation.column] == '#' {
			fmt.Println("hit wall dep 2 location", newGuardLocation, newDirection)
			// hit a wall, rotate 90*, try again
			newDirection = (newDirection + 1) % 4
			newGuardLocation := indexInDirection(guardStartingPoint, newDirection)
			fmt.Println("new loc:", newGuardLocation, newDirection)

			if (*guardMap)[newGuardLocation.row][newGuardLocation.column] == '#' {
				fmt.Println("hit wall dep 3 location", newGuardLocation, newDirection)

				// hit a wall, rotate 90*, try again
				newDirection = (newDirection + 1) % 4
				newGuardLocation := indexInDirection(guardStartingPoint, newDirection)
				if (*guardMap)[newGuardLocation.row][newGuardLocation.column] == '#' {
					fmt.Println("hit wall dep 4 location", newGuardLocation, newDirection)
					panic("guard stuck")
				}
			}
		}
		fmt.Println("heree")
		for _, line := range *guardMap {
			fmt.Println(line)
		}
	}

	locationCount[guardStartingPoint.row][guardStartingPoint.column]++

	if locationCount[guardStartingPoint.row][guardStartingPoint.column] > 4 {
		// fmt.Println("Location hit 5 times!", MapIndex{row: guardStartingPoint.row, column: guardStartingPoint.column})
		// for _, line := range *guardMap {
		// 	fmt.Println(line)
		// }
		return true
	}

	(*guardMap)[guardStartingPoint.row] = modifyStringAt((*guardMap)[guardStartingPoint.row], guardStartingPoint.column, 'X')

	return doesRecursivelyRunningGuardFindALoop(guardMap, newGuardLocation, newDirection, locationCount)
}

func day6Part2() (string, error) {
	countOfLoopMakingObstructions, err := getObstructionUniquePositions("days/day06/input")
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}

	// Output the results
	return strconv.Itoa(countOfLoopMakingObstructions), nil
}
