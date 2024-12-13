package day08

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Location struct {
	rowIndex int
	colIndex int
}

func getSumOfAntinodeLocations(filename string) (int, error) {
	// Open the input file
	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Read and process the file line by line
	scanner := bufio.NewScanner(file)
	var antennaMap []string
	charMap := make(map[rune][]Location)

	for scanner.Scan() {
		antennaMap = append(antennaMap, scanner.Text())
		// doing some preprocessing of frequencies
		for i, char := range antennaMap[len(antennaMap)-1] {
			if char != '.' {
				// locations of each frequency in a map
				charMap[char] = append(charMap[char], Location{rowIndex: len(antennaMap) - 1, colIndex: i})
			}
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %w", err)
	}

	antinodeMap := calculateAntinodeMap(antennaMap, charMap)

	fmt.Println("Antinode map:")
	for i := 0; i < len(antinodeMap); i++ {
		fmt.Println(string(antinodeMap[i]))
	}
	fmt.Println()
	antinodeCount := countAntinodes(antinodeMap)

	return antinodeCount, nil
}

func calculateAntinodeMap(antennaMap []string, charMap map[rune][]Location) [][]rune {
	// make a blank antinode 2D array of runes filled with '.'
	antinodeMap := make([][]rune, len(antennaMap))
	for i := 0; i < len(antennaMap); i++ {
		// initialize each row with the specified number of columns
		antinodeMap[i] = make([]rune, len(antennaMap[0]))
		for j := 0; j < len(antennaMap[0]); j++ {
			// fill each cell with the provided rune
			antinodeMap[i][j] = '.'
		}
	}

	// do some math on locations
	// for each point in []location, make antinode locations
	keys := make([]rune, 0, len(charMap))
	for k := range charMap {
		keys = append(keys, k)
	}
	antinodeLocations := make([]Location, 0)
	for _, key := range keys {
		for _, location := range charMap[key] {
			antinodeLocations = append(antinodeLocations, calculateAntinodeLocations(location, charMap[key])...)
		}
	}

	// for each location, if it matches mark it as an antinode
	for i := 0; i < len(antennaMap); i++ {
		for j := 0; j < len(antennaMap[0]); j++ {
			if locationListContainsLocation(antinodeLocations, Location{rowIndex: i, colIndex: j}) {
				antinodeMap[i][j] = '#'
			}
		}
	}

	return antinodeMap
}

func locationListContainsLocation(locationList []Location, location Location) bool {
	for i := 0; i < len(locationList); i++ {
		if locationsMatch(location, locationList[i]) {
			return true
		}
	}
	return false
}

func calculateAntinodeLocations(targetLocation Location, allMatchingFrequencyLocations []Location) []Location {
	antinodeLocations := []Location{}
	for _, location := range allMatchingFrequencyLocations {
		if locationsMatch(location, targetLocation) {
			// don't calculate
		} else {
			// do some math
			antinodeLocations = append(antinodeLocations, calculateSingleAntinode(targetLocation, location))
		}
	}
	return antinodeLocations
}

func calculateSingleAntinode(targetLocation Location, location Location) Location {
	return Location{
		rowIndex: targetLocation.rowIndex + (targetLocation.rowIndex - location.rowIndex),
		colIndex: targetLocation.colIndex + (targetLocation.colIndex - location.colIndex),
	}
}

func locationsMatch(location Location, targetLocation Location) bool {
	if location.rowIndex == targetLocation.rowIndex && location.colIndex == targetLocation.colIndex {
		fmt.Println("true match loc", location, targetLocation)
		return true
	}
	fmt.Println("false match loc", location, targetLocation)
	return false
}

func countAntinodes(antinodeMap [][]rune) int {
	sum := 0
	for i := 0; i < len(antinodeMap); i++ {
		for j := 0; j < len(antinodeMap[0]); j++ {
			if antinodeMap[i][j] != '.' {
				sum++
			}
		}
	}
	return sum
}

func day8Part1() (string, error) {
	sumOfAntinodeLocationsFound, err := getSumOfAntinodeLocations("days/day08/input")
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}

	// Output the results
	return strconv.Itoa(sumOfAntinodeLocationsFound), nil
}
