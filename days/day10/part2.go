package day10

import (
	"fmt"
	"strconv"
)

func getSumOfTrailheadScoresRating(filename string) (int, error) {
	trailMap := getTrailMap(filename)
	sum := 0
	for i := range trailMap {
		for j := range len(trailMap[0]) {
			if trailMap[i][j] == '0' {
				visitedPeaks := &[]TrailMapIndex{}
				getTrailheadScoreRating(visitedPeaks, trailMap, TrailMapIndex{row: i, column: j}, 0)
				sum += len(*visitedPeaks)
			}
		}
	}
	return sum, nil
}

func getTrailheadScoreRating(visitedPeaks *[]TrailMapIndex, trailMap []string, trailMapIndex TrailMapIndex, depth int) {
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
				getTrailheadScoreRating(visitedPeaks, trailMap, nextIndex, depth+1)
			}
		}
	}
}

func day10Part2() (string, error) {
	sumOfAllTrailheads, err := getSumOfTrailheadScoresRating("days/day10/input")
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}

	// Output the results
	return strconv.Itoa(sumOfAllTrailheads), nil
}
