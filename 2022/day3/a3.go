package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {

	file, err := os.Open("./day3/input.txt")
	if err != nil {
		fmt.Println("error:", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// input is always an even number
	// split in two to get the two sides of the backpack
	// find the letter (case sensative) that appears in both parts
	sumPriorities := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineLen := len(line)

		// make dictionary
		items := map[rune]bool{}

		var duplicateItem rune
		// loop over first half then second half until rune key is already true in which case stop
		for i := 0; i < (lineLen / 2); i++ {
			char := []rune(line)[i]
			items[char] = true
		}
		for i := lineLen / 2; i < lineLen; i++ {
			char := []rune(line)[i]

			if items[char] {
				duplicateItem = char
				break
			}
		}
		// something went wrong
		if !unicode.IsLetter(duplicateItem) {
			fmt.Printf("Invalid input %c\n", duplicateItem)
			continue
		}
		// priority -> a-z 1-26 | A-Z 27-52
		// find priority of missing item
		// sum of priorities
		// sumPriorities := 0
		var itemPriority int
		if unicode.IsUpper(duplicateItem) {
			itemPriority = int(duplicateItem) - 38
		} else {
			itemPriority += int(duplicateItem) - 96
		}
		sumPriorities += itemPriority
	} // add priority to sum
	fmt.Println(sumPriorities)
}
