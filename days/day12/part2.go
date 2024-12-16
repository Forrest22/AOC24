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
