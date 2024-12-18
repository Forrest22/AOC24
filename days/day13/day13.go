package day13

import "errors"

// Solve handles both parts for Day 1
func Solve(part int) (string, error) {
	switch part {
	case 1:
		return day13Part1()
	case 2:
		return day13Part2()
	default:
		return "", errors.New("invalid part number")
	}

}
