package day12

import (
	"fmt"
	"maps"
	"strconv"
)

func day12Part2() (string, error) {
	totalPrice, err := getTotalPriceOfFencingWithDiscount("days/day12/input")
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}

	// Output the results
	return strconv.Itoa(totalPrice), nil
}

func getTotalPriceOfFencingWithDiscount(filename string) (int, error) {
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
	gardenPlotSideCounts := map[string]int{}

	for i := range listOfGardenPlots {
		plotId := ""
		for j := range listOfGardenPlots[i] {
			plotId += strconv.Itoa(listOfGardenPlots[i][j].row) + "," + strconv.Itoa(listOfGardenPlots[i][j].column) + ", "
		}
		gardenPlotAreas[plotId] = len(listOfGardenPlots[i])
		gardenPlotSideCounts[plotId] = calculateNumberOfFencesToBuy(listOfGardenPlots[i], gardenMap)
	}

	totalPrice := 0
	for i := range maps.Keys(gardenPlotAreas) {
		totalPrice += gardenPlotAreas[i] * gardenPlotSideCounts[i]
	}

	return totalPrice, nil
}

// corner count = side count!
// both check two adjecnt sides and then if their rune matches, its a corner!
func calculateNumberOfFencesToBuy(gardenPlots []GardenPlotIndex, gardenMap []string) int {
	sideCount := 0
	for i := range gardenPlots {
		currentIndex := gardenPlots[i]
		for j := range 4 {
			adjacent := []GardenPlotIndex{indexInDirection(currentIndex, SearchDirection(j)), indexInDirection(currentIndex, SearchDirection((j+1)%4))}
			opposite := indexInDirection(adjacent[0], SearchDirection(j+1)%4)
			if indexIsACorner(gardenMap, rune(gardenMap[currentIndex.row][currentIndex.column]), adjacent[0], adjacent[1], opposite) {
				sideCount++
			}
		}
	}
	return sideCount
}

func indexIsACorner(gardenMap []string, currentPlant rune, gardenPlotIndex1, gardenPlotIndex2, opposite GardenPlotIndex) bool {
	if isIndexOOB(gardenMap, gardenPlotIndex1) && isIndexOOB(gardenMap, gardenPlotIndex2) {
		return true
	} else if !isIndexOOB(gardenMap, gardenPlotIndex1) && isIndexOOB(gardenMap, gardenPlotIndex2) {
		if rune(gardenMap[gardenPlotIndex1.row][gardenPlotIndex1.column]) != currentPlant {
			return true
		}
	} else if isIndexOOB(gardenMap, gardenPlotIndex1) && !isIndexOOB(gardenMap, gardenPlotIndex2) {
		if rune(gardenMap[gardenPlotIndex2.row][gardenPlotIndex2.column]) != currentPlant {
			return true
		}
	} else {
		plant1, plant2 := rune(gardenMap[gardenPlotIndex1.row][gardenPlotIndex1.column]), rune(gardenMap[gardenPlotIndex2.row][gardenPlotIndex2.column])
		oppositeCornerPlant := rune(gardenMap[opposite.row][opposite.column])
		if plant1 != currentPlant && plant2 != currentPlant {
			return true
		} else if plant1 == currentPlant && plant2 == currentPlant {
			if isIndexOOB(gardenMap, opposite) || oppositeCornerPlant != currentPlant {
				return true
			}
		}
	}
	return false
}
