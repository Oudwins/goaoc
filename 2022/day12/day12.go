package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
	"time"
)

const infinity = 100000000

// start = S -> problem 1
// start = a -> problem 2
const startingLocation = 'a'

func main() {

	file, err := os.ReadFile("./day12/input.txt")
	if err != nil {
		fmt.Println("error:", err)
	}
	fileContents := strings.Split(string(file), "\n")

	matrix := make([][]rune, len(fileContents))

	for i, line := range fileContents {
		runes := []rune(line)
		matrix[i] = runes
	}

	start := matrixIndexOf(&matrix, isChar('S'))[0]
	end := matrixIndexOf(&matrix, isChar('E'))[0]

	starts := [][2]int{start}

	if startingLocation != 'S' {
		as := matrixIndexOf(&matrix, isChar(startingLocation))
		starts = append(starts, as...)
	}
	matrix[start[0]][start[1]] = 'a'
	matrix[end[0]][end[1]] = 'z'

	allSteps := []int{}

	startTime := time.Now()
	for _, trailStart := range starts {
		allSteps = append(allSteps, dijkstrasShortestPath(trailStart, end, &matrix))
	}

	elapsedTime := time.Now().Sub(startTime)

	sort.Ints(allSteps)

	fmt.Println(allSteps[0])
	fmt.Println("IN ", elapsedTime)
}

func dijkstrasShortestPath(start [2]int, end [2]int, matrix *[][]rune) int {
	seen := make([][]bool, len(*matrix))
	distance := make([][]int, len(*matrix))

	for i := 0; i < len(*matrix); i++ {
		seen[i] = make([]bool, len((*matrix)[i]))
		distance[i] = make([]int, len((*matrix)[i]))
		for j := 0; j < len(distance[i]); j++ {
			distance[i][j] = infinity
		}
	}

	distance[start[0]][start[1]] = 0

	stepsInShortestestPath := 0

	// my distance + distance to that edge
	iterations := 0
	for {

		pos := getLowestUnvisited(&seen, &distance)
		seen[pos[0]][pos[1]] = true
		if pos[0] == -1 {
			log.Fatal("ERROR COULD NOT FIND NEXT LOWEST UNVISITED NODE")
		}

		if pos[0] == end[0] && pos[1] == end[1] {
			stepsInShortestestPath = distance[pos[0]][pos[1]]
			break
		}

		// edges
		newDistance := distance[pos[0]][pos[1]] + 1
		// top
		if canTravel(pos, [2]int{pos[0] - 1, pos[1]}, &seen, matrix) {
			distance[pos[0]-1][pos[1]] = updateDistance(distance[pos[0]-1][pos[1]], newDistance)
		}
		// bot
		if canTravel(pos, [2]int{pos[0] + 1, pos[1]}, &seen, matrix) {
			distance[pos[0]+1][pos[1]] = updateDistance(distance[pos[0]+1][pos[1]], newDistance)
		}

		// left
		if canTravel(pos, [2]int{pos[0], pos[1] - 1}, &seen, matrix) {
			distance[pos[0]][pos[1]-1] = updateDistance(distance[pos[0]][pos[1]-1], newDistance)
		}

		if canTravel(pos, [2]int{pos[0], pos[1] + 1}, &seen, matrix) {
			distance[pos[0]][pos[1]+1] = updateDistance(distance[pos[0]][pos[1]+1], newDistance)
		}
		iterations++

	}

	// get lowest unvisited
	// distance array
	// hasUnvisited

	// for each edge
	// if seen -> continue
	// if not seen -> add to list
	// calc shortest distance
	// do that step

	return stepsInShortestestPath
}

func updateDistance(old int, new int) int {
	if old > new {
		return new
	}
	return old
}

func canTravel(from [2]int, to [2]int, seen *[][]bool, matrix *[][]rune) bool {
	// is out of bounds
	if to[0] < 0 || to[1] < 0 || to[0] >= len(*matrix) || to[1] >= len((*matrix)[to[0]]) {
		return false
	}
	// has been visited already
	if (*seen)[to[0]][to[1]] {
		return false
	}
	// is larger > 1
	if (*matrix)[to[0]][to[1]] > (*matrix)[from[0]][from[1]]+1 {
		return false
	}

	return true
}

func getLowestUnvisited(seen *[][]bool, distance *[][]int) [2]int {
	lowest := [2]int{-1, -1}
	for i := 0; i < len(*seen); i++ {

		for j := 0; j < len((*seen)[i]); j++ {
			if (*seen)[i][j] {
				continue
			}

			if lowest[0] == -1 {
				lowest[0] = i
				lowest[1] = j
				continue
			}

			if (*distance)[i][j] < (*distance)[lowest[0]][lowest[1]] {
				lowest = [2]int{i, j}
			}

		}

	}

	return lowest
}

func calculateDistance(from [2]int, to [2]int) int {
	rows := math.Abs(float64(to[0] - from[0]))
	cols := math.Abs(float64(to[1] - from[1]))

	return int(rows + cols)
}

func matrixIndexOf(matrix *[][]rune, tester func(el rune) bool) (result [][2]int) {

	for i := 0; i < len(*matrix); i++ {

		for j, char := range (*matrix)[i] {
			if tester(char) {
				result = append(result, [2]int{i, j})
			}
		}
	}

	return result
}

func isChar(want rune) func(r rune) bool {

	return func(r rune) bool {
		if r == want {
			return true
		}
		return false
	}
}
