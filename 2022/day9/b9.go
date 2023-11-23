package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// note to self. seems to work correctly. But seems to work for test case. So i'm fucked...

type coordinatePlane [2]int

func main() {

	file, err := os.Open("./day9/input.txt")
	if err != nil {
		fmt.Println("error:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// x=0
	// y=1
	traversed := make(map[string]bool)
	rope := [10]coordinatePlane{}
	tailIdx := len(rope) - 1
	headIdx := 0
	for i := 0; i < len(rope); i++ {
		rope[i] = coordinatePlane{0, 0}
	}

	// planeSize := [2]coordinatePlane{{0,0}, {0,0}}
	// planeSize[1][0] > head[0] + amount * 2dDirection

	loc := fmt.Sprintf("%v:%v", rope[tailIdx][0], rope[tailIdx][1])
	traversed[loc] = true
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		direction := parts[0]
		amount, _ := strconv.Atoi(parts[1])

		// fmt.Println("MOVE IS: ", direction, amount)
		for i := 1; i <= amount; i++ {
			// adding cur location to traversed
			// fmt.Println("LOOOP START")

			// move head
			movePoint(&rope[headIdx], direction)

			// debugPlane(8, 8, head, tail)

			for i := headIdx + 1; i < len(rope); i++ {
				if !areTouching(rope[i-1], rope[i]) {
					newCoordenates := moveFollower(rope[i-1], rope[i], direction)
					rope[i] = newCoordenates
				}
			}
			loc := fmt.Sprintf("%v:%v", rope[tailIdx][0], rope[tailIdx][1])
			traversed[loc] = true
			// if !areTouching(head, tail) {
			// fmt.Println("SOMETHING WENT WRONG. HEAD AND TAIL NOT TOUCHING")
			// }

			debugPlane(6, 6, rope)
			// fmt.Println(areTouching(head, tail))

		}

	}

	fmt.Println(len(traversed))
}

func areTouching(head coordinatePlane, tail coordinatePlane) bool {

	// calculating the distance using distance formula
	diviation := math.Pow(float64(tail[0]-head[0]), 2) + math.Pow(float64(tail[1]-head[1]), 2)
	// fmt.Println("DIVIATION: ", diviation)
	if diviation < 4 {
		return true
	}
	// (x1,y1) = (x2,y2)
	// x1 = x2 + - 1
	// y1 = y2 +-1

	// 00 = 11
	// 00 = 1-1
	// 00 = -11
	// 00 = -1-1name
	return false
}

func movePoint(point *coordinatePlane, direction string) {
	switch direction {
	case "U":
		point[1]++
	case "D":
		point[1]--
	case "R":
		point[0]++
	case "L":
		point[0]--
	default:
		fmt.Println("INVALID DIRECTION")
	}
}

func moveFollower(head coordinatePlane, tail coordinatePlane, move string) coordinatePlane {
	changeX := convertToSign(head[0] - tail[0])
	changeY := convertToSign(head[1] - tail[1])
	return coordinatePlane{tail[0] + changeX, tail[1] + changeY}
}

func difference(a int, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}

func convertToSign(number int) int {
	if number > 0 {
		return 1
	} else if number < 0 {
		return -1
	} else {
		return 0
	}
}
func debugPlane(rows int, cols int, rope [10]coordinatePlane) {

	fmt.Println(rope)
	for i := 0; i < cols; i++ {

		for j := 0; j < rows; j++ {

			found := false

			for idx := 0; idx < len(rope); idx++ {
				if rows-1-rope[idx][0] == i && rope[idx][1] == j {
					fmt.Printf("%v", idx)
					found = true
					break
				}
			}
			if !found {
				fmt.Printf(".")
			}

		}
		fmt.Println()
	}
	fmt.Println()
}
