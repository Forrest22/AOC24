package day14

import (
	"fmt"
)

func day14Part2() (string, error) {
	_, err := displayXmasTreeLike("days/day14/input", 101, 103, 1)
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}

	// Output the resultstytedssdddeett
	return "Search the above printed grids for xmas trees to find solution.", nil
}

func displayXmasTreeLike(filename string, w, h, sec int) (int, error) {
	robots := getRobots(filename)
	areaGrid := createGrid(w, h, robots)
	blinks := 0

	initSafetyScore := countSafetyScore(areaGrid)
	safetyScore := 0
	for safetyScore != initSafetyScore {
		areaGrid = processTime(areaGrid)
		blinks++

		// display xmas tree like areas, look manually good luck :P
		safetyScore = countSafetyScoreAndCheckXmasTree(areaGrid, blinks)

	}

	return safetyScore, nil
}

func countSafetyScoreAndCheckXmasTree(resultingGrid [][]Index, blinks int) int {
	northEastQuadrent, southEastQuadrent, southWestQuadrent, northWestQuadrent := 0, 0, 0, 0
	printGrid := true
	for x, column := range resultingGrid {
		for y, index := range column {
			if len(index.robots) > 0 {
				if len(index.robots) > 1 {
					// only print grids that don't overlap, a smaller set and more likely to be part of the sol'n
					printGrid = false
				}
				if y < len(resultingGrid[0])/2 && x > len(resultingGrid)/2 {
					// NE quadrant
					northEastQuadrent += len(index.robots)
				} else if y > len(resultingGrid[0])/2 && x > len(resultingGrid)/2 {
					// SE quadrant
					southEastQuadrent += len(index.robots)
				} else if y < len(resultingGrid[0])/2 && x < len(resultingGrid)/2 {
					// SW quadrant
					southWestQuadrent += len(index.robots)
				} else if y > len(resultingGrid[0])/2 && x < len(resultingGrid)/2 {
					// NW quadrant
					northWestQuadrent += len(index.robots)
				}
			}
		}
	}
	if printGrid {
		fmt.Println("blinks:", blinks)
		for _, column := range resultingGrid {
			for _, index := range column {
				fmt.Print(len(index.robots))
			}
			fmt.Println()
		}
		fmt.Println()
	}
	return northEastQuadrent * southEastQuadrent * southWestQuadrent * northWestQuadrent
}
