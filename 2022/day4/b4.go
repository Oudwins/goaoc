package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("./day4/input.txt")
	if err != nil {
		fmt.Println("error:", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// section has unique id number
	// each elf is assigned range of ids
	// asignments overlap
	// elves pair up & list list of assignments
	// count number of assignment pair where one range fully contains the others

	collisions := 0
	elvesPerGroup := 2
	for scanner.Scan() {
		// get the string split by ,
		// split again by -
		// compare numbers
		line := scanner.Text()
		splitLine := strings.Split(line, ",")
		elfs := [2][2]int{}
		// check nothing is wrong
		if len(elfs) != elvesPerGroup {
			fmt.Println("Couldn't parse both elves")
			continue
		}

		for idx, elf := range splitLine {
			vals := strings.Split(elf, "-")
			low, _ := strconv.Atoi(vals[0])
			high, _ := strconv.Atoi(vals[1])
			elfs[idx] = [2]int{low, high}
		}
		if pairHasColision(elfs) {
			collisions++
		} else {
			fmt.Println(elfs)
		}
	}
	fmt.Println(collisions)
}

func pairHasColision(elfs [2][2]int) bool {
	return rangeIsInside(elfs[0], elfs[1]) || rangeIsInside(elfs[1], elfs[0])
}

func rangeIsInside(r1 [2]int, r2 [2]int) bool {
	return (r1[0] >= r2[0] && r1[0] <= r2[1]) || (r1[1] >= r2[0] && r1[1] <= r2[1])
}
