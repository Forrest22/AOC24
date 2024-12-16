package day12

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"strconv"
)

// Define a struct to represent i, j indexes for the []strings garden map
type GardenPlotIndex struct {
	row    int
	column int
}

type SearchDirection int

const (
	North SearchDirection = iota
	East
	South
	West
)

func (d SearchDirection) String() string {
	return [...]string{
		"North",
		"East",
		"South",
		"West",
	}[d]
}

func (d SearchDirection) EnumIndex() int {
	return int(d)
}

func indexInDirection(index GardenPlotIndex, direction SearchDirection) GardenPlotIndex {
	switch direction {
	case North:
		return GardenPlotIndex{
			row:    index.row - 1,
			column: index.column,
		}
	case East:
		return GardenPlotIndex{
			row:    index.row,
			column: index.column + 1,
		}
	case South:
		return GardenPlotIndex{
			row:    index.row + 1,
			column: index.column,
		}
	case West:
		return GardenPlotIndex{
			row:    index.row,
			column: index.column - 1,
		}
	default:
		fmt.Printf("Direction %d is not implemented yet.\n", direction)
		return index
	}
}

func isIndexOOB(gardenMap []string, index GardenPlotIndex) bool {
	if index.row < 0 || index.row >= len(gardenMap) {
		return true
	}
	if index.column < 0 || index.column >= len(gardenMap[0]) {
		return true
	}
	return false
}

func day12Part1() (string, error) {
	totalPrice, err := getTotalPriceOfFencing("days/day12/input")
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}

	// Output the results
	return strconv.Itoa(totalPrice), nil
}

func getTotalPriceOfFencing(filename string) (int, error) {
	gardenMap := getGardenMap(filename)

	listOfGardenPlots := [][]GardenPlotIndex{}
	for i := range gardenMap {
		for j := range gardenMap[0] {
			if !isGardenPlotIndexInList(listOfGardenPlots, i, j) {
				listOfGardenPlots = append(listOfGardenPlots, getIndexListForPlotStartingAt(&[]GardenPlotIndex{}, gardenMap, rune(gardenMap[i][j]), GardenPlotIndex{row: i, column: j}))
			}
		}
	}

	gardenPlotAreas := map[string]int{}
	gardenPlotPerimeters := map[string]int{}

	for i := range listOfGardenPlots {
		plotId := ""
		for j := range listOfGardenPlots[i] {
			plotId += strconv.Itoa(listOfGardenPlots[i][j].row) + "," + strconv.Itoa(listOfGardenPlots[i][j].column)
		}
		gardenPlotAreas[plotId] = len(listOfGardenPlots[i])
		gardenPlotPerimeters[plotId] = calculateGardenPlotPerimeter(listOfGardenPlots[i], gardenMap)
	}

	totalPrice := 0
	for i := range maps.Keys(gardenPlotAreas) {
		totalPrice += gardenPlotAreas[i] * gardenPlotPerimeters[i]
	}

	return totalPrice, nil
}

func calculateGardenPlotPerimeter(gardenPlots []GardenPlotIndex, gardenMap []string) int {
	perimeter := 0
	for i := range gardenPlots {
		currentIndex := gardenPlots[i]
		for j := range 4 {
			borderIndex := indexInDirection(gardenPlots[i], SearchDirection(j))
			if isIndexOOB(gardenMap, borderIndex) || gardenMap[currentIndex.row][currentIndex.column] != gardenMap[borderIndex.row][borderIndex.column] {
				perimeter++
			}
		}
	}
	return perimeter
}

func isGardenPlotIndexInList(listOfGardenPlots [][]GardenPlotIndex, targetRow, targetColumn int) bool {
	if len(listOfGardenPlots) == 0 || len(listOfGardenPlots[0]) == 0 {
		return false
	}
	for i := range listOfGardenPlots {
		for j := range listOfGardenPlots[i] {
			if listOfGardenPlots[i][j].row == targetRow && listOfGardenPlots[i][j].column == targetColumn {
				return true
			}
		}
	}
	return false
}

func getIndexListForPlotStartingAt(searchedIndexes *[]GardenPlotIndex, gardenMap []string, startingPlantRune rune, gardenPlotIndex GardenPlotIndex) []GardenPlotIndex {
	if rune(gardenMap[gardenPlotIndex.row][gardenPlotIndex.column]) != startingPlantRune {
		*searchedIndexes = append(*searchedIndexes, GardenPlotIndex{row: gardenPlotIndex.row, column: gardenPlotIndex.column})
		return nil
	}

	// for each search direction
	var listOfGardenPlots []GardenPlotIndex
	*searchedIndexes = append(*searchedIndexes, gardenPlotIndex)
	for i := range 4 {
		nextIndex := indexInDirection(gardenPlotIndex, SearchDirection(i))
		if !isIndexOOB(gardenMap, nextIndex) {
			if !isGardenPlotIndexInList([][]GardenPlotIndex{*searchedIndexes}, nextIndex.row, nextIndex.column) {
				plots := getIndexListForPlotStartingAt(searchedIndexes, gardenMap, startingPlantRune, indexInDirection(gardenPlotIndex, SearchDirection(i)))
				listOfGardenPlots = append(listOfGardenPlots, plots...)
			}
		}
	}
	return append(listOfGardenPlots, gardenPlotIndex)
}

func getGardenMap(filename string) []string {
	// Open the input file
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Errorf("error parsing trail map: %v", err))
	}
	defer file.Close()

	// Read and process the file line by line
	scanner := bufio.NewScanner(file)
	gardenMap := []string{}
	// gardenPlotAreaMapping := map[rune]int{}
	for scanner.Scan() {
		rowInGarden := scanner.Text()
		gardenMap = append(gardenMap, rowInGarden)
		// for i := range rowInGarden {
		// 	gardenPlotAreaMapping[rune(rowInGarden[i])] += 1
		// }
	}
	return gardenMap
}
