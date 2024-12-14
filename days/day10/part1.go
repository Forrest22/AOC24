package day10

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TrailMapIndex struct {
	row    int
	column int
}

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

func indexInDirection(index TrailMapIndex, direction Direction) TrailMapIndex {
	switch direction {
	case North:
		return TrailMapIndex{
			row:    index.row - 1,
			column: index.column,
		}
	case East:
		return TrailMapIndex{
			row:    index.row,
			column: index.column + 1,
		}
	case South:
		return TrailMapIndex{
			row:    index.row + 1,
			column: index.column,
		}
	case West:
		return TrailMapIndex{
			row:    index.row,
			column: index.column - 1,
		}
	default:
		fmt.Printf("Direction %d is not implemented yet.\n", direction)
		return index
	}
}

func isIndexOOB(trailMap []string, index TrailMapIndex) bool {
	// check row is inbounds
	if index.row < 0 || index.row >= len(trailMap) {
		return true
	}
	// check column is inbounds
	if index.column < 0 || index.column >= len(trailMap[0]) {
		return true
	}
	return false
}

func getSumOfTrailheadScores(filename string) (int, error) {
	trailMap := getTrailMap(filename)
	sum := 0
	for i := range trailMap {
		for j := range len(trailMap[0]) {
			if trailMap[i][j] == '0' {
				visitedPeaks := &[]TrailMapIndex{}
				getTrailheadScore(visitedPeaks, trailMap, TrailMapIndex{row: i, column: j}, 0)
				sum += len(*visitedPeaks)
			}
		}
	}
	return sum, nil
}

func getTrailheadScore(visitedPeaks *[]TrailMapIndex, trailMap []string, trailMapIndex TrailMapIndex, depth int) {
	if depth == 9 {
		*visitedPeaks = append(*visitedPeaks, trailMapIndex)
		return
	}

	for i := range 4 {
		nextIndex := indexInDirection(trailMapIndex, Direction(i))
		if !isIndexOOB(trailMap, nextIndex) {
			nextIndexVal, err := strconv.Atoi(string(trailMap[nextIndex.row][nextIndex.column]))
			if err != nil {
				panic(err)
			}
			if nextIndexVal == depth+1 {
				hasBeenVisited := false
				for i := range *visitedPeaks {
					if (*visitedPeaks)[i].row == nextIndex.row && (*visitedPeaks)[i].column == nextIndex.column {
						hasBeenVisited = true
					}
				}
				if !hasBeenVisited {
					getTrailheadScore(visitedPeaks, trailMap, nextIndex, depth+1)
				}
			}
		}
	}
}

func getTrailMap(filename string) []string {
	// Open the input file
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Errorf("error parsing trail map", err))
	}
	defer file.Close()

	// Read and process the file line by line
	scanner := bufio.NewScanner(file)
	var allText []string
	for scanner.Scan() {
		allText = append(allText, scanner.Text())
	}
	return allText
}

func day10Part1() (string, error) {
	sumOfAllTrailheads, err := getSumOfTrailheadScores("days/day10/input")
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}

	// Output the results
	return strconv.Itoa(sumOfAllTrailheads), nil
}
