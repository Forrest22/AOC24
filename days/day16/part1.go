package day16

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"slices"
	"strconv"
)

type Index struct {
	x               int
	y               int
	objectIndicator rune
	score           int
	direction       SearchDirection
	index           int
}

type PriorityQueue []*Index

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].score < pq[j].score
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Index)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

type SearchDirection int

const (
	North SearchDirection = iota
	East
	South
	West
)

func (d SearchDirection) String() string {
	return [...]string{
		"North",
		"East",
		"South",
		"West",
	}[d]
}

func (d SearchDirection) EnumIndex() int {
	return int(d)
}

func indexInDirection(mazeMap [][]Index, index Index, direction SearchDirection) Index {
	switch direction {
	case North:
		return mazeMap[index.y-1][index.x]
	case East:
		return mazeMap[index.y][index.x+1]
	case South:
		return mazeMap[index.y+1][index.x]
	case West:
		return mazeMap[index.y][index.x-1]
	default:
		fmt.Printf("Direction %d is not implemented yet.\n", direction)
		return index
	}
}

func day16Part1() (string, error) {
	lowestScore, err := getReindeerMazeLowestScore("days/day16/input")
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}
	// Output the resultstytedssdddeett
	return strconv.Itoa(lowestScore), nil
}

func getReindeerMazeLowestScore(filename string) (int, error) {
	mazeMap, startIndex, endIndex := getStartingMap(filename)
	lowestScore := findLowestScoreBFS(mazeMap, startIndex, endIndex, East)
	return lowestScore, nil
}

func printMaze(mazeMap [][]Index) {
	for y, row := range mazeMap {
		for x := range row {
			fmt.Print(string(mazeMap[y][x].objectIndicator))
		}
		fmt.Println(strconv.Itoa(y))
	}
	for x := range len(mazeMap[0]) {
		fmt.Print(strconv.Itoa(x % 10))
	}
	fmt.Println()
}

func findLowestScoreBFS(mazeMap [][]Index, startIndex, endIndex Index, currentDirection SearchDirection) int {
	pathScores := map[string]int{}
	currentScore := 0
	startIndex.score = 0
	startIndex.direction = currentDirection
	startIndex.index = 0
	// make the priority queue
	pq := make(PriorityQueue, 1)
	// Push to the queue
	pq[0] = &startIndex
	heap.Init(&pq)
	for pq.Len() > 0 {
		currentIndex := heap.Pop(&pq).(*Index)
		currentScore = currentIndex.score
		currentKey := fmt.Sprintf("(%v, %v, %v)", currentIndex.y, currentIndex.x, currentDirection)
		currentDirection = currentIndex.direction

		// if the current score is greater, don't update anything
		val, exists := pathScores[currentKey]
		if exists && val < currentScore {
			continue
		}

		// for each change in direction from this one
		for i := range 3 {
			newSearchDirectionInt := (int(currentDirection) + i + 1) % 4
			newSearchDirection := SearchDirection(newSearchDirectionInt)
			newDirectionKey := fmt.Sprintf("(%v, %v, %v)", currentIndex.y, currentIndex.x, newSearchDirectionInt)
			val, exists := pathScores[newDirectionKey]
			if !exists || val > currentScore+1001 {
				pathScores[newDirectionKey] = currentScore + 1001
				nextIndex := indexInDirection(mazeMap, *currentIndex, newSearchDirection)
				nextIndex.direction = newSearchDirection
				nextIndex.score = currentScore + 1001
				if nextIndex.objectIndicator != '#' {
					pq = append(pq, &nextIndex)
				}
			}

		}

		nextIndexInCurrentDirection := indexInDirection(mazeMap, *currentIndex, currentDirection)
		val, exists = pathScores[fmt.Sprintf("(%v, %v, %v)", nextIndexInCurrentDirection.y, nextIndexInCurrentDirection.x, currentDirection.EnumIndex())]
		if 0 <= nextIndexInCurrentDirection.y && 0 <= nextIndexInCurrentDirection.x && nextIndexInCurrentDirection.y < len(mazeMap) &&
			nextIndexInCurrentDirection.x < len(mazeMap[0]) && nextIndexInCurrentDirection.objectIndicator != '#' && (!exists || val > currentScore+1) {
			nextIndexKey := fmt.Sprintf("(%v, %v, %v)", nextIndexInCurrentDirection.y, nextIndexInCurrentDirection.x, currentDirection.EnumIndex())
			pathScores[nextIndexKey] = currentScore + 1
			nextIndex := indexInDirection(mazeMap, *currentIndex, currentDirection)
			nextIndex.direction = currentDirection
			nextIndex.score = currentScore + 1
			pq = append(pq, &nextIndex)
		}
	}
	endingScores := []int{}
	for i := range 4 {
		endingScores = append(endingScores, pathScores[fmt.Sprintf("(%v, %v, %v)", endIndex.y, endIndex.x, i)])
	}
	lowestScore := slices.Min(endingScores)

	return lowestScore
}

func getStartingMap(filename string) ([][]Index, Index, Index) {
	maze := make([][]Index, 0)

	// Open the input file
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Errorf("error parsing claw machines: %v", err))
	}
	defer file.Close()

	// Read and process the file line by line
	scanner := bufio.NewScanner(file)

	// get the wall
	startX, startY := 0, 0
	endX, endY := 0, 0
	y := 0
	for scanner.Scan() {
		maze = append(maze, []Index{})
		line := scanner.Text()
		for x, col := range line {
			if col == 'S' {
				startX = x
				startY = y
			} else if col == 'E' {
				endX = x
				endY = y
			}
			maze[y] = append(maze[y], Index{x: x, y: y, objectIndicator: col})
		}
		y++
	}
	return maze, Index{x: startX, y: startY}, Index{x: endX, y: endY}
}
