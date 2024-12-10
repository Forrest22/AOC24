package day07

import "errors"

// Solve handles both parts for Day 1
func Solve(part int) (string, error) {
	switch part {
	case 1:
		return day7Part1()
	case 2:
		return day7Part2()
	default:
		return "", errors.New("invalid part number")
	}
}
