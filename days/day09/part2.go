package day09

import (
	"fmt"
	"strconv"
)

func getFilesystemChecksumNoFileFrag(filename string) (int, error) {
	fullDisk := getDiskMap(filename)
	sortedFullDisk := getSortedFullDiskNoFileFrag(fullDisk)
	return getChecksum(sortedFullDisk), nil
}

func getSortedFullDiskNoFileFrag(fullDisk []DiskBlock) []DiskBlock {
	backIndexStart, backIndexEnd := findBackIndexBounds(fullDisk, len(fullDisk))
	fileLength := backIndexEnd - backIndexStart
	found, frontIndexStart, frontIndexEnd := findLeftMostFreeBoundsOfSize(fullDisk, 0, fileLength, backIndexStart)

	for backIndexStart >= 0 && frontIndexStart < len(fullDisk) {
		if found {
			// swap found
			fullDisk = swap(fullDisk, frontIndexStart, frontIndexEnd, backIndexStart, backIndexEnd)

			backIndexStart, backIndexEnd = findBackIndexBounds(fullDisk, backIndexStart)
			fileLength = backIndexEnd - backIndexStart
			found, frontIndexStart, frontIndexEnd = findLeftMostFreeBoundsOfSize(fullDisk, 0, fileLength, backIndexStart)
		} else {
			// no swap match found
			backIndexStart, backIndexEnd = findBackIndexBounds(fullDisk, backIndexStart)
			fileLength = backIndexEnd - backIndexStart
			found, frontIndexStart, frontIndexEnd = findLeftMostFreeBoundsOfSize(fullDisk, 0, fileLength, backIndexStart)

		}
	}

	return fullDisk
}

func swap(fullDisk []DiskBlock, frontIndexStart, frontIndexEnd, backIndexStart, backIndexEnd int) []DiskBlock {
	result, freeSpaces := []DiskBlock{}, []DiskBlock{}

	// swap full words
	for i := 0; i < frontIndexEnd-frontIndexStart; i++ {
		freeSpaces = append(freeSpaces, DiskBlock{id: ID(-1), isFreeSpace: true})
	}
	file := fullDisk[backIndexStart:backIndexEnd]

	result = append(fullDisk[:frontIndexStart], file...)
	result = append(result, fullDisk[frontIndexEnd:backIndexStart]...)
	result = append(result, freeSpaces...)
	result = append(result, fullDisk[backIndexEnd:]...)

	return result
}

func findLeftMostFreeBoundsOfSize(fullDisk []DiskBlock, startIndex, fileLength, backIndexStart int) (bool, int, int) {
	i := startIndex
	for i < backIndexStart {
		if !fullDisk[i].isFreeSpace {
			for !fullDisk[i].isFreeSpace && i < backIndexStart {
				i++
			}
		} else {
			newStartIndex := i
			for fullDisk[i].isFreeSpace && i < backIndexStart {
				if 1+i-newStartIndex == fileLength {
					return true, newStartIndex, i + 1
				}
				i++
			}
		}
	}
	return false, -1, -1
}
func findBackIndexBounds(fullDisk []DiskBlock, backIndexEnd int) (int, int) {
	if backIndexEnd == 0 {
		return -1, -1
	}

	if fullDisk[backIndexEnd-1].isFreeSpace {
		for fullDisk[backIndexEnd-1].isFreeSpace {
			backIndexEnd--
		}
	}

	backIndexStart := backIndexEnd
	for backIndexStart > 0 && !fullDisk[backIndexStart-1].isFreeSpace && fullDisk[backIndexStart-1].id == fullDisk[backIndexEnd-1].id {
		backIndexStart--
	}
	return backIndexStart, backIndexEnd
}

func day9Part2() (string, error) {
	sumOfAntinodeLocationsFound, err := getFilesystemChecksumNoFileFrag("days/day09/input")
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}

	// Output the results
	return strconv.Itoa(sumOfAntinodeLocationsFound), nil
}
