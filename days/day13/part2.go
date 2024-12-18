package day13

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day13Part2() (string, error) {
	fewestTokensCount, err := fewestTokensToWinAllPossiblePrizesAmended("days/day13/input")
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}

	// Output the results
	return strconv.Itoa(fewestTokensCount), nil
}

func fewestTokensToWinAllPossiblePrizesAmended(filename string) (int, error) {
	listOfClawMachines := getListOfClawMachinesAmended(filename)

	var aCount, bCount int
	aSum, bSum := 0, 0
	for i := range listOfClawMachines {
		aCount, bCount = getTokenCountOfPrizeEffiencient(listOfClawMachines[i])
		aSum += aCount
		bSum += bCount
	}

	return 3*aSum + bSum, nil
}

// Cramer's Rule
// A = (p_x*b_y - prize_y*b_x) / (a_x*b_y - a_y*b_x)
// B = (a_x*p_y - a_y*p_x) / (a_x*b_y - a_y*b_x)
func getTokenCountOfPrizeEffiencient(clawMachine map[string]int) (int, int) {
	a := float64(clawMachine["prizeX"]*clawMachine["bY"]-clawMachine["prizeY"]*clawMachine["bX"]) /
		float64(clawMachine["aX"]*clawMachine["bY"]-clawMachine["aY"]*clawMachine["bX"])
	if !isInt(a) {
		return 0, 0
	}
	b := float64(clawMachine["aX"]*clawMachine["prizeY"]-clawMachine["aY"]*clawMachine["prizeX"]) /
		float64(clawMachine["aX"]*clawMachine["bY"]-clawMachine["aY"]*clawMachine["bX"])
	if !isInt(b) {
		return 0, 0
	}
	return int(a), int(b)
}

func getListOfClawMachinesAmended(filename string) []map[string]int {
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
			"prizeX": prizeX + 10000000000000,
			"prizeY": prizeY + 10000000000000,
		}
		listOfClawMachines = append(listOfClawMachines, clawMachine)
	}
	return listOfClawMachines
}
