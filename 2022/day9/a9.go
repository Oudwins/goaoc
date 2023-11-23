package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

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
	head := coordinatePlane{0, 0}
	tail := coordinatePlane{0, 0}

	// planeSize := [2]coordinatePlane{{0,0}, {0,0}}
	// planeSize[1][0] > head[0] + amount * 2dDirection

	loc := fmt.Sprintf("%v:%v", tail[0], tail[1])
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
			movePoint(&head, direction)

			// debugPlane(8, 8, head, tail)
			if !areTouching(head, tail) {
				tailMove := moveTail(head, tail, direction)
				tail = tailMove
				// this is what runs on each loc
			}
			loc := fmt.Sprintf("%v:%v", tail[0], tail[1])
			traversed[loc] = true
			// if !areTouching(head, tail) {
			// fmt.Println("SOMETHING WENT WRONG. HEAD AND TAIL NOT TOUCHING")
			// }

			// debugPlane(8, 8, head, tail)
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

func moveTail(head coordinatePlane, tail coordinatePlane, move string) coordinatePlane {
	// 1. are we in the same col or row?
	// 2. if yes move as per move
	if head[0] == tail[0] || head[1] == tail[1] {
		movePoint(&tail, move)
		return tail
	}
	// 3. else  move to col or row & one as per move
	// first set the cur value that has a difference of 1 to the same then do the normal move
	// this is just head's previous location.
	if difference(head[0], tail[0]) == 1 {
		tail[0] = head[0]
		movePoint(&tail, move)
		return tail
	} else {
		tail[1] = head[1]
		movePoint(&tail, move)
		return tail
	}

}
func smoveTail(head coordinatePlane, tail coordinatePlane, move string) coordinatePlane {
	changeX := convertToSign(head[0] - tail[0])
	changeY := convertToSign(head[1] - tail[1])
	return coordinatePlane{tail[0] + changeX, tail[1] + changeY}
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
func difference(a int, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}

func debugPlane(rows int, cols int, head coordinatePlane, tail coordinatePlane) {

	for i := 0; i < rows; i++ {

		for j := 0; j < cols; j++ {
			if rows-head[0] == j && cols-head[1] == i {
				// print head
				fmt.Printf("H")
			} else if rows-tail[0] == j && cols-head[1] == i {
				// print tail
				fmt.Printf("T")
			} else {
				fmt.Printf(".")
			}

		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()

	fmt.Println(head, tail)
}
