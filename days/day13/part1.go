package day13

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func day13Part1() (string, error) {
	fewestTokensCount, err := fewestTokensToWinAllPossiblePrizes("days/day13/input")
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}

	// Output the results
	return strconv.Itoa(fewestTokensCount), nil
}

func fewestTokensToWinAllPossiblePrizes(filename string) (int, error) {
	listOfClawMachines := getListOfClawMachines(filename)

	var aCount, bCount int
	aSum, bSum := 0, 0
	for i := range listOfClawMachines {
		aCount, bCount = getTokenCountOfPrize(listOfClawMachines[i])
		aSum += aCount
		bSum += bCount
	}

	return 3*aSum + bSum, nil
}

func getTokenCountOfPrize(clawMachine map[string]int) (int, int) {
	if isInt(float64(clawMachine["prizeX"])/float64(clawMachine["bX"])) &&
		isInt(float64(clawMachine["prizeY"])/float64(clawMachine["bY"])) {
		return 0, clawMachine["prizeX"] / clawMachine["bX"]
	}

	aTokens := 1
	bTokens := clawMachine["prizeX"] / clawMachine["bX"]
	for bTokens >= 0 {
		if clawMachine["aX"]*aTokens+clawMachine["bX"]*bTokens == clawMachine["prizeX"] &&
			clawMachine["aY"]*aTokens+clawMachine["bY"]*bTokens == clawMachine["prizeY"] {
			return aTokens, bTokens
		} else if clawMachine["aX"]*aTokens+clawMachine["bX"]*bTokens > clawMachine["prizeX"] &&
			clawMachine["aY"]*aTokens+clawMachine["bY"]*bTokens > clawMachine["prizeY"] {
			bTokens--
		} else {
			aTokens++
		}
	}
	return 0, 0
}

func isInt(num float64) bool {
	return num == math.Trunc(num)
}

func getListOfClawMachines(filename string) []map[string]int {
	listOfClawMachines := make([]map[string]int, 0)

	// Open the input file
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Errorf("error parsing claw machines: %v", err))
	}
	defer file.Close()

	// Read and process the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		buttonALine := strings.Split(scanner.Text(), "+")
		scanner.Scan() // goes to next line
		buttonBLine := strings.Split(scanner.Text(), "+")
		scanner.Scan()
		prizeLine := strings.Split(scanner.Text(), "=")
		scanner.Scan()

		aX, aY := getButtonValues(buttonALine)
		bX, bY := getButtonValues(buttonBLine)
		prizeX, prizeY := getPrizeValues(prizeLine)

		clawMachine := map[string]int{
			"aX":     aX,
			"aY":     aY,
			"bX":     bX,
			"bY":     bY,
			"prizeX": prizeX,
			"prizeY": prizeY,
		}
		listOfClawMachines = append(listOfClawMachines, clawMachine)
	}
	return listOfClawMachines
}

func getButtonValues(buttonLines []string) (int, int) {
	x, err := strconv.Atoi(strings.Split(buttonLines[1][:len(buttonLines[1])-3], ",")[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(buttonLines[2])
	if err != nil {
		panic(err)
	}
	return x, y
}

func getPrizeValues(prizeLines []string) (int, int) {
	x, err := strconv.Atoi(strings.Split(prizeLines[1][:len(prizeLines[1])-3], ",")[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(prizeLines[2])
	if err != nil {
		panic(err)
	}
	return x, y
}
