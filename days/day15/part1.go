package day15

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Index struct {
	x     int
	y     int
	robot Robot
	wall  Wall
	box   Box
}

type Robot bool

type Wall bool

type Box bool

func day15Part1() (string, error) {
	gpsSum, err := getGPSCoordinateSum("days/day15/input")
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}

	// Output the resultstytedssdddeett
	return strconv.Itoa(gpsSum), nil
}

func getGPSCoordinateSum(filename string) (int, error) {
	areaMap, moves, robotX, robotY := getStartingMap(filename)

	gpsSum := 0
	for _, move := range moves {
		areaMap, gpsSum, robotX, robotY = processMove(areaMap, move, robotX, robotY)
	}
	fmt.Println("Final warehouse map:")
	printAreaMap(areaMap)

	return gpsSum, nil
}

func printAreaMap(areaMap [][]Index) {
	for _, row := range areaMap {
		for _, index := range row {
			if index.wall {
				fmt.Print("#")
			} else if index.box {
				fmt.Print("O")
			} else if index.robot {
				fmt.Print("@")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func processMove(areaMap [][]Index, move rune, robotX, robotY int) ([][]Index, int, int, int) {
	areaMap, robotX, robotY = moveRobotInDirection(move, areaMap, robotX, robotY)

	gpsSum := 0
	for x, row := range areaMap {
		for y, index := range row {
			if index.box {
				gpsSum += 100*x + y
			}
		}
	}
	return areaMap, gpsSum, robotX, robotY
}

func moveRobotInDirection(move rune, areaMap [][]Index, robotX, robotY int) ([][]Index, int, int) {
	newRobotX, newRobotY := robotX, robotY
	xModifier, yModifier := 0, 0 // only one of these should ever be non-zero
	if move == '^' {
		newRobotY--
		yModifier = -1
	} else if move == '>' {
		newRobotX++
		xModifier = 1
	} else if move == 'v' {
		newRobotY++
		yModifier = 1
	} else if move == '<' {
		newRobotX--
		xModifier = -1
	} else {
		panic("unexpected move in move list")
	}

	checkSpot := areaMap[newRobotY][newRobotX]
	if checkSpot.wall {
		// do nothing
		return areaMap, robotX, robotY
	} else if checkSpot.box {
		// if a box
		for checkSpot.box {
			// go to end of boxes
			checkSpot = areaMap[checkSpot.y+yModifier][checkSpot.x+xModifier]
		}
		if !checkSpot.wall {
			// move the boxes and robots
			areaMap[checkSpot.y][checkSpot.x].box = true
			newRobotX = robotX + xModifier
			newRobotY = robotY + yModifier
			areaMap[newRobotY][newRobotX].robot = true
			areaMap[newRobotY][newRobotX].box = false
			areaMap[robotY][robotX].robot = false
		} else {
			// can't move boxes, do nothing
			return areaMap, robotX, robotY
		}
	} else {
		// otherwise is free
		areaMap[newRobotY][newRobotX].robot = true
		areaMap[robotY][robotX].robot = false
	}
	return areaMap, newRobotX, newRobotY
}

func getStartingMap(filename string) ([][]Index, string, int, int) {
	areaMap := make([][]Index, 0)

	// Open the input file
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Errorf("error parsing claw machines: %v", err))
	}
	defer file.Close()

	// Read and process the file line by line
	scanner := bufio.NewScanner(file)

	// get the wall
	row := 0
	robotStartX := 0
	robotStartY := 0
	for scanner.Scan() && scanner.Text() != "" {
		areaMap = append(areaMap, []Index{})
		line := scanner.Text()
		for col := range line {
			if line[col] == byte('#') {
				areaMap[row] = append(areaMap[row], Index{x: col, y: row, wall: true})
			} else if line[col] == byte('O') {
				areaMap[row] = append(areaMap[row], Index{x: col, y: row, box: true})
			} else if line[col] == byte('@') {
				areaMap[row] = append(areaMap[row], Index{x: col, y: row, robot: true})
				robotStartX = col
				robotStartY = row
			} else {
				areaMap[row] = append(areaMap[row], Index{x: col, y: row})
			}
		}
		row++
	}
	moves := ""
	for scanner.Scan() {
		moves += scanner.Text()
	}
	return areaMap, moves, robotStartX, robotStartY
}
