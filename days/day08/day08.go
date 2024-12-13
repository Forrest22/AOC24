package day08

import "errors"

// Solve handles both parts for Day 1
func Solve(part int) (string, error) {
	switch part {
	case 1:
		return day8Part1()
	case 2:
		return day8Part2()
	default:
		return "", errors.New("invalid part number")
	}
}
