package day15

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

type IndexPart2 struct {
	x               int
	y               int
	objectIndicator rune
}

func day15Part2() (string, error) {
	gpsSum, err := getGPSCoordinateSumWarehouse2("days/day15/input")
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}

	// Output the resultstytedssdddeett
	return strconv.Itoa(gpsSum), nil
}

func getGPSCoordinateSumWarehouse2(filename string) (int, error) {
	areaMap, moves, robotX, robotY := getStartingMapWarehouse2(filename)
	printAreaMapWarehouse2(areaMap)

	gpsSum := 0
	for _, move := range moves {
		areaMap, gpsSum, robotX, robotY = processMoveWarehouse2(areaMap, move, robotX, robotY)
		printAreaMapWarehouse2(areaMap)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("Final warehouse map:", moves, robotX, robotY)

	return gpsSum, nil
}

func printAreaMapWarehouse2(areaMap [][]IndexPart2) {
	for _, row := range areaMap {
		for _, index := range row {
			fmt.Print(string(index.objectIndicator))
		}
		fmt.Println()
	}
	fmt.Println()
}

// something has gone terribly wrong
func processMoveWarehouse2(areaMap [][]IndexPart2, move rune, robotX, robotY int) ([][]IndexPart2, int, int, int) {
	areaMap, robotX, robotY = moveRobotInDirectionWarehouse2(move, areaMap, robotX, robotY)

	gpsSum := 0
	for x, row := range areaMap {
		for y, index := range row {
			if index.objectIndicator == '[' {
				gpsSum += 100*x + y
			}
		}
	}
	return areaMap, gpsSum, robotX, robotY
}

func moveRobotInDirectionWarehouse2(move rune, areaMap [][]IndexPart2, robotX, robotY int) ([][]IndexPart2, int, int) {
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

	fmt.Println("move x,y", string(move), robotX, robotY, newRobotX, newRobotY)
	// printAreaMapWarehouse2(areaMap)
	locationsBeingPushed := []IndexPart2{areaMap[newRobotY][newRobotX]}
	fmt.Println("checkSpots", locationsBeingPushed, string(locationsBeingPushed[0].objectIndicator))
	if locationsBeingPushed[0].objectIndicator == '#' {
		// do nothing
		return areaMap, robotX, robotY
	} else if locationsBeingPushed[0].objectIndicator == ']' || locationsBeingPushed[0].objectIndicator == '[' {
		// if a box
		fmt.Println("box!")
		// get the full box
		for areAnyLocationsAreBoxes(locationsBeingPushed) {
			if move == '<' || move == '>' {
				// go to end of boxes
				locationsBeingPushed[0] = areaMap[locationsBeingPushed[0].y+yModifier][locationsBeingPushed[0].x+xModifier]
			} else {
				// get the end of the box being pushed
				if locationsBeingPushed[0].objectIndicator == '[' {
					// get right box bracket
					locationsBeingPushed = append(locationsBeingPushed, areaMap[locationsBeingPushed[0].y][locationsBeingPushed[0].x+1])
				} else {
					// get left box bracket
					locationsBeingPushed = append(locationsBeingPushed, areaMap[locationsBeingPushed[0].y][locationsBeingPushed[0].x-1])
				}
				// nextSet := getNextSetOfBoxes(areaMap, locationsBeingPushed, move)
				nextSet := locationsBeingPushed
				fmt.Println("locationsBeingPushed", locationsBeingPushed)
				fmt.Println("nextSetOfBoxes", nextSet)
				// get the rest of the boxes, one row at a time
				// for areAnyLocationsAreBoxes(nextSetOfBoxes) {
				for len(nextSet) > 0 {
					nextSet = getNextSetOfIndexes(areaMap, nextSet, move)
					i := 0
					for _, location := range nextSet {
						if location.objectIndicator == '#' {
							// if any are walls, exit
							fmt.Println("can't move bc of wall")
							return areaMap, robotX, robotY
						} else if location.objectIndicator == '.' {
							// remove it from the box list
							fmt.Println("i, nextSet[:i], nextSet[i+1:]", i, nextSet[:i], nextSet[i+1:])
							nextSet = append(nextSet[:i], nextSet[i+1:]...)
						} else {
							i++
						}
					}
					locationsBeingPushed = append(locationsBeingPushed, nextSet...)
					fmt.Println("locationsBeingPushed", locationsBeingPushed)
					fmt.Println("nextSetOfBoxes", nextSet)
				}
				locationsBeingPushed = append(locationsBeingPushed, nextSet...)
				fmt.Println("locationsBeingPushed", locationsBeingPushed)
				// move allBoxes up or down
				// not moving boxes properly
				// resultLocations := map[string]IndexPart2{}
				resultLocations := []IndexPart2{}
				for i := range locationsBeingPushed {
					location := locationsBeingPushed[i]
					resultLocations = append(resultLocations, IndexPart2{x: location.x, y: location.y + yModifier, objectIndicator: location.objectIndicator})
				}
				fmt.Println("resultLocations", resultLocations)
				// for i := range locationsBeingPushed {
				// 	location := locationsBeingPushed[i]
				// 	fmt.Println("location", location)
				// 	areaMap[location.y+yModifier][location.x] = IndexPart2{x: location.x, y: location.y + yModifier, objectIndicator: location.objectIndicator}
				// 	makeFreeSpace := false
				// 	var resultLocation IndexPart2
				// 	for j := range resultLocations {
				// 		resultLocation = resultLocations[j]
				// 		fmt.Println("resultLocation", resultLocation)
				// 		if location.x == resultLocation.x && location.y+yModifier == resultLocation.y {
				// 			// don't change
				// 			fmt.Println("")
				// 		} else {
				// 			makeFreeSpace = true
				// 		}
				// 	}
				// 	if makeFreeSpace {
				// 		areaMap[location.y][location.x].objectIndicator = '.'
				// 	} else {
				// 		areaMap[location.y][location.x].objectIndicator = resultLocation.objectIndicator

				// 	}
				// }

				// make the results
				for _, resultLocation := range resultLocations {
					makeFreeSpace := true
					for _, location := range locationsBeingPushed {
						if resultLocation.y == location.y+yModifier && resultLocation.x == location.x {
							areaMap[resultLocation.y][resultLocation.x].objectIndicator = location.objectIndicator
							makeFreeSpace = false
						}
					}
					if makeFreeSpace {
						areaMap[resultLocation.y][resultLocation.x].objectIndicator = '.'
					}
				}

				// clean up the boxes moves
				for _, location := range locationsBeingPushed {
					locationPartOfResults := false
					for _, resultLocation := range resultLocations {
						if resultLocation.y == location.y && resultLocation.x == location.x {
							fmt.Println("location part of results!", location)
							locationPartOfResults = true
						}
					}
					if !locationPartOfResults {
						areaMap[location.y][location.x].objectIndicator = '.'
					}
				}

				areaMap[robotY][robotX].objectIndicator = '.'
				newRobotY = robotY + yModifier
				areaMap[newRobotY][robotX].objectIndicator = '@'
				// printAreaMapWarehouse2(areaMap)
				return areaMap, robotX, newRobotY
			}
		}
		if locationsBeingPushed[0].objectIndicator != '#' {
			// not a wall
			// move the boxes and robots
			fmt.Println("not a wall", locationsBeingPushed)
			// areaMap[checkSpot.y][checkSpot.x].box = true
			if move == 'v' {
				// moving things down
				for i := range locationsBeingPushed[0].y - robotY {
					fmt.Println("i", i)
					areaMap[locationsBeingPushed[0].y-i][robotX].objectIndicator = areaMap[locationsBeingPushed[0].y-(i+1)][robotX].objectIndicator
				}
				areaMap[robotX][robotX].objectIndicator = '.'
			} else if move == '^' {
				// moving things up
				// for i := range robotY - checkSpot.y {
				// 	areaMap[checkSpot.y+i][robotX].objectIndicator = areaMap[checkSpot.y+(i+1)][robotX].objectIndicator
				// }
				// areaMap[robotX][robotX].objectIndicator = '.'
			} else if move == '>' {
				// moving things right
				for i := range locationsBeingPushed[0].x - robotX {
					areaMap[locationsBeingPushed[0].y][locationsBeingPushed[0].x-i].objectIndicator = areaMap[locationsBeingPushed[0].y][locationsBeingPushed[0].x-(i+1)].objectIndicator
				}
				areaMap[robotY][robotX].objectIndicator = '.'
			} else {
				// moving things left
				for i := range robotX - locationsBeingPushed[0].x {
					areaMap[locationsBeingPushed[0].y][locationsBeingPushed[0].x+i].objectIndicator = areaMap[locationsBeingPushed[0].y][locationsBeingPushed[0].x+(i+1)].objectIndicator
				}
				areaMap[robotY][robotX].objectIndicator = '.'
			}

			newRobotX = robotX + xModifier
			newRobotY = robotY + yModifier
			// areaMap[newRobotY][newRobotX].robot = true
			// areaMap[newRobotY][newRobotX].box = false
			// areaMap[robotY][robotX].robot = false
		} else {
			// can't move boxes, do nothing
			return areaMap, robotX, robotY
		}
	} else {
		// otherwise is free, move
		fmt.Println("free space", robotX, robotY, newRobotX, newRobotY, xModifier)
		areaMap[newRobotY][newRobotX].objectIndicator = '@'
		areaMap[robotY][robotX].objectIndicator = '.'
		// areaMap[newRobotY][newRobotX].robot = true
		// areaMap[robotY][robotX].robot = false
	}
	return areaMap, newRobotX, newRobotY
}

func getNextSetOfIndexes(areaMap [][]IndexPart2, boxes []IndexPart2, move rune) []IndexPart2 {
	xModifier, yModifier := 0, 0
	nextSetOfBoxes := []IndexPart2{}
	// fmt.Println("string(move),boxes", string(move), boxes)
	for _, box := range boxes {
		if move == '^' {
			yModifier = -1
		} else if move == '>' {
			xModifier = 1
		} else if move == 'v' {
			yModifier = 1
		} else if move == '<' {
			xModifier = -1
		}
		// fmt.Println("xModifier,yModifier", xModifier, yModifier)
		// fmt.Println("next box is", areaMap[box.y+yModifier][box.x+xModifier])
		possibleNextBox := areaMap[box.y+yModifier][box.x+xModifier]
		nextSetOfBoxes = append(nextSetOfBoxes, areaMap[box.y+yModifier][box.x+xModifier])
		if possibleNextBox.objectIndicator == '[' {
			// get right box bracket
			nextSetOfBoxes = append(nextSetOfBoxes, areaMap[possibleNextBox.y][possibleNextBox.x+1])
		} else {
			// get left box bracket
			nextSetOfBoxes = append(nextSetOfBoxes, areaMap[possibleNextBox.y][possibleNextBox.x-1])
		}

	}
	// fmt.Println("string(move),nextSetOfBoxes", string(move), nextSetOfBoxes)

	return nextSetOfBoxes
}

func areAnyLocationsAreBoxes(locationsBeingPushed []IndexPart2) bool {
	for i := range locationsBeingPushed {
		if locationsBeingPushed[i].objectIndicator == ']' || locationsBeingPushed[i].objectIndicator == '[' {
			return true
		}
	}
	return false
}

func getStartingMapWarehouse2(filename string) ([][]IndexPart2, string, int, int) {
	areaMap := make([][]IndexPart2, 0)

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
		areaMap = append(areaMap, []IndexPart2{})
		line := scanner.Text()
		for col := range line {
			if line[col] == byte('#') {
				areaMap[row] = append(areaMap[row], IndexPart2{x: 2 * col, y: row, objectIndicator: '#'})
				areaMap[row] = append(areaMap[row], IndexPart2{x: 2*col + 1, y: row, objectIndicator: '#'})
			} else if line[col] == byte('O') {
				areaMap[row] = append(areaMap[row], IndexPart2{x: 2 * col, y: row, objectIndicator: '['})
				areaMap[row] = append(areaMap[row], IndexPart2{x: 2*col + 1, y: row, objectIndicator: ']'})
			} else if line[col] == byte('@') {
				areaMap[row] = append(areaMap[row], IndexPart2{x: 2 * col, y: row, objectIndicator: '@'})
				robotStartX = 2 * col
				robotStartY = row
				areaMap[row] = append(areaMap[row], IndexPart2{x: 2*col + 1, y: row, objectIndicator: '.'})
			} else {
				areaMap[row] = append(areaMap[row], IndexPart2{x: 2 * col, y: row, objectIndicator: '.'})
				areaMap[row] = append(areaMap[row], IndexPart2{x: 2*col + 1, y: row, objectIndicator: '.'})
			}
			col++
		}
		row++
	}
	moves := ""
	for scanner.Scan() {
		moves += scanner.Text()
	}
	return areaMap, moves, robotStartX, robotStartY
}
