package day11

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day11Part1() (string, error) {
	totalCountOfStones, err := getTotalCountOfStones("days/day11/input", 25)
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}

	// Output the results
	return strconv.Itoa(totalCountOfStones), nil
}

// I rewrote this for part 2 hoping that ints would save me but it was not the problem of why my program would run out of memory and crash :(
func getTotalCountOfStones(filename string, blinks int) (int, error) {
	stones := getStonesFromInput(filename)
	for i := 0; i < blinks; i++ {
		stones = calculateStonesAfterBlink(stones)
	}
	return len(stones), nil
}

func calculateStonesAfterBlink(stones []int) []int {
	resultingStones := []int{}
	for _, stone := range stones {
		stringStone := strconv.Itoa(stone)
		if stringStone == "0" {
			resultingStones = append(resultingStones, 1)
		} else if len(stringStone)%2 == 0 {
			leftStone, err := strconv.Atoi(stringStone[:len(stringStone)/2])
			if err != nil {
				panic(err)
			}
			rightStone, err := strconv.Atoi(stringStone[len(stringStone)/2:])
			if err != nil {
				panic(err)
			}
			resultingStones = append(resultingStones, leftStone, rightStone)
		} else {
			resultingStones = append(resultingStones, 2024*stone)
		}
	}
	return resultingStones
}

func getStonesFromInput(filename string) []int {
	// Open the input file
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Errorf("error parsing trail map: %v", err))
	}
	defer file.Close()

	// Read and process the file line by line
	scanner := bufio.NewScanner(file)
	var stones []int
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		for i := range line {
			stone, err := strconv.Atoi(line[i])
			if err != nil {
				panic(err)
			}
			stones = append(stones, stone)
		}
	}
	return stones
}
