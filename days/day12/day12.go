package day12

import "errors"

// Solve handles both parts for Day 1
func Solve(part int) (string, error) {
	switch part {
	case 1:
		return day12Part1()
	case 2:
		return day12Part2()
	default:
		return "", errors.New("invalid part number")
	}

}
