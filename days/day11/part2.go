package day11

import (
	"fmt"
	"maps"
	"strconv"
	"strings"
)

func day11Part2() (string, error) {
	totalCountOfStones, err := getTotalCountOfStonesMap("days/day11/input", 75)
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}

	// Output the results
	return strconv.Itoa(totalCountOfStones), nil
}

func getTotalCountOfStonesMap(filename string, blinks int) (int, error) {
	stones := getStonesFromInput(filename)

	stoneMap := make(map[string]int)
	for _, stone := range stones {
		stoneKey := strconv.Itoa(stone)
		stoneMap[stoneKey] = stoneMap[stoneKey] + 1
	}

	for i := 0; i < blinks; i++ {
		stoneMap = calculateStonesAfterBlinkUsingMap(stoneMap)
	}

	count := 0
	for stoneKey := range maps.Keys(stoneMap) {
		count += stoneMap[stoneKey]
	}
	return count, nil
}

func calculateStonesAfterBlinkUsingMap(stoneMap map[string]int) map[string]int {
	resultingStones := make(map[string]int)
	for stoneKey := range maps.Keys(stoneMap) {
		if truncateLeftZeroes(stoneKey) == "0" {
			resultingStones["1"] += stoneMap[stoneKey]
		} else if len(stoneKey)%2 == 0 {
			leftStoneString := stoneKey[:len(stoneKey)/2]
			rightStoneString := truncateLeftZeroes(stoneKey[len(stoneKey)/2:])
			resultingStones[leftStoneString] += stoneMap[stoneKey]
			resultingStones[rightStoneString] += stoneMap[stoneKey]
		} else {
			oldStoneKey, err := strconv.Atoi(stoneKey)
			if err != nil {
				panic(err)
			}
			newStoneKey := strconv.Itoa(2024 * oldStoneKey)
			resultingStones[newStoneKey] += stoneMap[stoneKey]
		}
	}
	return resultingStones
}

func truncateLeftZeroes(stoneKey string) string {
	truncatedKey := stoneKey
	for strings.IndexRune(truncatedKey, '0') == 0 && len(truncatedKey) > 1 {
		truncatedKey = truncatedKey[1:]
	}
	return truncatedKey
}
