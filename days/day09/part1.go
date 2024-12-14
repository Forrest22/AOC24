package day09

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type FreeSpace bool

type ID int

type DiskBlock struct {
	isFreeSpace bool
	id          ID
}

func getFilesystemChecksum(filename string) (int, error) {
	fullDisk := getDiskMap(filename)
	sortedFullDisk := getSortedFullDisk(fullDisk)
	return getChecksum(sortedFullDisk), nil
}

func getChecksum(sortedFullDisk []DiskBlock) int {
	sum := 0
	for i := 0; i < len(sortedFullDisk); i++ {
		if !sortedFullDisk[i].isFreeSpace {
			sum += int(sortedFullDisk[i].id) * i
		}
	}
	return sum
}

// actually very cool dynamic programming?!
func getSortedFullDisk(fullDisk []DiskBlock) []DiskBlock {
	i, j := 0, len(fullDisk)-1
	for i < j {
		if fullDisk[i].isFreeSpace && !fullDisk[j].isFreeSpace {
			// swap
			fullDisk[i] = fullDisk[j]
			fullDisk[j] = DiskBlock{isFreeSpace: true, id: ID(-1)}
			i++
			j--
		} else if !fullDisk[i].isFreeSpace {
			i++
		} else {
			j--
		}
	}
	return fullDisk
}

func getDiskMap(filename string) []DiskBlock {
	// Open the input file
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Errorf("error opening file %v", err))
	}
	defer file.Close()

	// Read and process the file line by line
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	diskMap := scanner.Text()
	fullDisk := []DiskBlock{}

	for i, diskMapInstruction := range diskMap {
		d := string(diskMapInstruction)
		repeatXTimes, err := strconv.Atoi(d)
		if err != nil {
			panic(fmt.Errorf("error reading diskmap num %v, %v", d, err))
		}
		if isFileBlock(i) {
			// get the id
			id := i / 2
			for j := 0; j < repeatXTimes; j++ {
				fullDisk = append(fullDisk, DiskBlock{isFreeSpace: false, id: ID(id)})
			}
		} else {
			// is free blocks
			for j := 0; j < repeatXTimes; j++ {
				fullDisk = append(fullDisk, DiskBlock{isFreeSpace: true, id: ID(-1)})
			}
		}
	}
	return fullDisk
}

func isFileBlock(counter int) bool {
	return counter%2 == 0
}

func day9Part1() (string, error) {
	filesystemChecksum, err := getFilesystemChecksum("days/day09/input")
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}

	// Output the results
	return strconv.Itoa(filesystemChecksum), nil
}
