package day14

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Index struct {
	x      int
	y      int
	robots []Robot
}

type Velocity struct {
	vX int
	vY int
}

type Robot struct {
	startIndex Index
	velocity   Velocity
}

func day14Part1() (string, error) {
	safetyScore, err := getSafetyScore("days/day14/input", 101, 103, 100)
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}

	// Output the resultstytedssdddeett
	return strconv.Itoa(safetyScore), nil
}

func getSafetyScore(filename string, w, h, sec int) (int, error) {
	robots := getRobots(filename)
	areaGrid := createGrid(w, h, robots)

	safetyScore := 0
	for range sec {
		areaGrid = processTime(areaGrid)
		// count safety score
		safetyScore = countSafetyScore(areaGrid)
	}

	return safetyScore, nil
}

func processTime(areaGrid [][]Index) [][]Index {
	resultingGrid := createGrid(len(areaGrid), len(areaGrid[0]), []Robot{})
	var robotsToCopy []Robot
	for x, row := range areaGrid {
		for y, index := range row {
			robotsToCopy = index.robots
			// move the robots
			for _, robot := range robotsToCopy {
				resultingX, resultingY := (x+robot.velocity.vX)%len(resultingGrid), (y+robot.velocity.vY)%len(resultingGrid[0])
				for resultingX < 0 {
					resultingX = resultingX + len(resultingGrid)
				}
				for resultingY < 0 {
					resultingY = resultingY + len(resultingGrid[0])
				}
				// add robot to grid
				resultingGrid[resultingX][resultingY].robots = append(resultingGrid[resultingX][resultingY].robots, robot)
			}
		}
	}
	return resultingGrid
}

func countSafetyScore(resultingGrid [][]Index) int {
	northEastQuadrent, southEastQuadrent, southWestQuadrent, northWestQuadrent := 0, 0, 0, 0
	for x, column := range resultingGrid {
		for y, index := range column {
			if len(index.robots) > 0 {
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
	return northEastQuadrent * southEastQuadrent * southWestQuadrent * northWestQuadrent
}

func createGrid(w, h int, robots []Robot) [][]Index {
	grid := make([][]Index, w)
	for i := range grid {
		grid[i] = make([]Index, h)
	}
	for yIndex := range h {
		for xIndex := range w {
			grid[xIndex][yIndex] = Index{x: xIndex, y: yIndex}
			for _, robot := range robots {
				if robot.startIndex.x == xIndex && robot.startIndex.y == yIndex {
					grid[xIndex][yIndex].robots = append(grid[xIndex][yIndex].robots, robot)
				}
			}
		}
	}
	return grid
}

func getRobots(filename string) []Robot {
	robots := make([]Robot, 0)

	// Open the input file
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Errorf("error parsing claw machines: %v", err))
	}
	defer file.Close()

	// Read and process the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		robots = append(robots, getRobotFromLine(scanner.Text()))
	}
	return robots
}

func getRobotFromLine(line string) Robot {
	robot := Robot{}
	splitLine := strings.Split(line, " ")
	robot.startIndex = getPosition(splitLine[0])
	robot.velocity = getVelocity(splitLine[1])
	return robot
}

func getVelocity(s string) Velocity {
	velocity := strings.Split(s[2:], ",")
	x, err := strconv.Atoi(velocity[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(velocity[1])
	if err != nil {
		panic(err)
	}
	return Velocity{vX: x, vY: y}
}

func getPosition(s string) Index {
	position := strings.Split(s[2:], ",")
	x, err := strconv.Atoi(position[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(position[1])
	if err != nil {
		panic(err)
	}
	return Index{x: x, y: y}
}
