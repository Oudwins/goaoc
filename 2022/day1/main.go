package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// open file
	file, err := os.Open("./day1/input.txt")
	if err != nil {
		fmt.Println("error:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	topThree := [3]int{0, 0, 0}
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		// base case
		if len(line) == 0 {
			isGreater := -1
			for idx, val := range topThree {
				if sum < val {
					break
				}
				isGreater = idx
			}
			updateTopThree(sum, isGreater, &topThree)
			sum = 0
			// fmt.Println(largest)
			continue
		}
		calories, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("ERROR:", err)
		}

		sum = sum + calories
	}
	// I can't find way to check for EOF so doing the last clean up here
	isGreater := -1
	for idx, val := range topThree {
		if sum < val {
			break
		}
		isGreater = idx
	}
	updateTopThree(sum, isGreater, &topThree)

	// print result
	fmt.Println("Array:", topThree)
	// Summing total
	total := 0
	for _, val := range topThree {
		total = total + val
	}
	fmt.Println(total)
}

func updateTopThree(val int, idx int, topThree *[3]int) {
	if idx < 0 || idx > len(topThree) {
		return
	}

	for i := 0; i < idx; i++ {
		topThree[i] = topThree[i+1]
	}
	topThree[idx] = val
}

func firstProblem() {

	// open file
	file, err := os.Open("./day1/input.txt")
	if err != nil {
		fmt.Println("error:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	largest := 0
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		// base case
		if len(line) == 0 {
			if sum > largest {
				largest = sum
			}
			sum = 0
			fmt.Println(largest)
			continue
		}
		calories, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("ERROR:", err)
		}

		sum = sum + calories
	}

	// Do while (there is file left)
	// greatest = 0;
	// sum until blank space
	// if sum > greatest;
	// greatest = sum;
	// when no file left compare last sum also!
	// return that

}
