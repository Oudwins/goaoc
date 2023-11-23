package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cycle = 0
var registerX = 1

func main() {
	file, err := os.Open("./day10/input.txt")
	if err != nil {
		fmt.Println("error:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "add") {
			newCycle()
			newCycle()
			instruction := strings.Split(line, " ")
			n, _ := strconv.Atoi(instruction[1])
			registerX += n
		} else if strings.HasPrefix(line, "noop") {
			newCycle()
		}

	}
}

func newCycle() {
	// sprite = x
	// sprite is 3 pixels wide
	drawingAt := cycle % 40
	if cycle > 1 && drawingAt == 0 {
		fmt.Println()
	}
	// one of its pixels = cycle
	// fmt.Println(registerX, " | ", cycle, " | ", drawingAt)
	drawPixel(drawingAt, registerX)
	cycle++
}

func drawPixel(drawingAt int, pixelCenter int) {
	if difference(drawingAt, pixelCenter) <= 1 {
		fmt.Printf("#")
	} else {
		fmt.Printf(".")
	}
}

func difference(a int, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}
