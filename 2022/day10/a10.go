package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var cycle = 0
var cyclesToSum = []int{20, 60, 100, 140, 180, 220}
var registerX = 1
var signalSum = 0

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
		} else {
			newCycle()
		}

	}
	fmt.Println(signalSum)
}

func newCycle() {
	cycle++

	fmt.Println(registerX)
	if slices.Contains(cyclesToSum, cycle) {
		signalSum += cycle * registerX
	}
}
