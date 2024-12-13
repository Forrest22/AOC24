package day08

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

type Slope struct {
	slope      float32
	yIntercept float32
}

// equivalent of checkin if y=mx+b
func isLocationOnSlope(loc Location, slope Slope) bool {
	return float32(loc.rowIndex) == (slope.slope*float32(loc.colIndex))-slope.yIntercept
}

func calculateAntinodeMapUsingSlopes(antennaMap []string, charMap map[rune][]Location, slopes []Slope) [][]rune {
	// make a blank antinode 2D array of runes filled with '.'
	antinodeMap := make([][]rune, len(antennaMap))
	for i := 0; i < len(antennaMap); i++ {
		// initialize each row with the specified number of columns
		antinodeMap[i] = make([]rune, len(antennaMap[0]))
		for j := 0; j < len(antennaMap[0]); j++ {
			// fill each cell with the provided rune
			currentLocation := Location{rowIndex: i, colIndex: j}
			// if isLocationAntenna(currentLocation, charMap) {
			// 	antinodeMap[i][j] = 'A'
			// } else
			if isAntinodeLocation(slopes, charMap, currentLocation) {
				antinodeMap[i][j] = '#'
			} else {
				antinodeMap[i][j] = '.'
			}
		}
	}
	return antinodeMap
}

func isAntinodeLocation(slopes []Slope, charMap map[rune][]Location, location Location) bool {
	// equivalent of checking y=mx+b for each slope
	for _, slope := range slopes {
		// if float32(location.rowIndex) == slope.slope*float32(location.colIndex)+slope.yIntercept {
		if isLocationOnSlope(location, slope) {
			fmt.Println(location, slopes, true)
			return true
		} else if isLocationAntenna(location, charMap) {
			return true
		}
	}
	// fmt.Println(location, slopes, false)
	return false
}

func isLocationAntenna(location Location, charMap map[rune][]Location) bool {
	for _, locationList := range charMap {
		if slices.Contains(locationList, location) {
			return true
		}
	}
	return false
}

func getSlopesFromCharMap(charMap map[rune][]Location) []Slope {
	// fmt.Println("charmap:", charMap)
	slopes := []Slope{}
	for k := range charMap {
		for l, charLocation := range charMap[k] {
			fmt.Println(k, charLocation)
			slopes = append(slopes, getSlopesFromLocations(charLocation, charMap[k][:l])...)
			fmt.Println(slopes)
		}
	}
	return slopes
}

func getSlopesFromLocations(targetLocation Location, locations []Location) []Slope {
	slopes := []Slope{}
	for _, location := range locations {
		if !locationsMatch(targetLocation, location) {
			// do some math
			slopes = append(slopes, getSlopeFromLocations(targetLocation, location))
		}
	}
	return slopes
}

func getSlopeFromLocations(targetLocation Location, location Location) Slope {
	rowDelta := float32(targetLocation.rowIndex - location.rowIndex)
	colDelta := float32(targetLocation.colIndex - location.colIndex)
	intercept := float32(float32(rowDelta/colDelta)*float32(targetLocation.rowIndex) - float32(targetLocation.colIndex))
	return Slope{
		slope:      rowDelta / colDelta,
		yIntercept: intercept,
	}
}

func getSumOfAntinodeAndTFrequencyLocations(filename string) (int, error) {
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

	slopes := getSlopesFromCharMap(charMap)

	fmt.Println("Antenna map:")
	for i := 0; i < len(antennaMap); i++ {
		fmt.Println(string(antennaMap[i]))
	}

	fmt.Println(charMap, slopes)
	antinodeMap := calculateAntinodeMapUsingSlopes(antennaMap, charMap, slopes)

	fmt.Println("Antinode map:")
	for i := 0; i < len(antinodeMap); i++ {
		fmt.Println(string(antinodeMap[i]))
	}
	fmt.Println()
	antinodeCount := countAntinodes(antinodeMap)

	return antinodeCount, nil
}

func day8Part2() (string, error) {
	sumOfAntinodeLocationsFound, err := getSumOfAntinodeAndTFrequencyLocations("days/day08/input")
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}

	// Output the results
	return strconv.Itoa(sumOfAntinodeLocationsFound), nil
}
