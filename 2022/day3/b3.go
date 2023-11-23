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
	i := 0
	items := map[rune]int{}
	for scanner.Scan() {
		line := scanner.Text()
		i += 1
		elfChars := []rune(line)
		uniqueElfChars := removeDuplicateStr(elfChars)

		for _, val := range uniqueElfChars {
			items[val] += 1
		}
		// 3r is last elf in team
		if i%3 == 0 {
			var teamLetter rune
			for key := range items {
				if items[key] == 3 {
					teamLetter = key
				}
				items[key] = 0
			}
			if !unicode.IsLetter(teamLetter) {
				fmt.Printf("Invalid input %c\n", teamLetter)
				continue
			}
			// priority -> a-z 1-26 | A-Z 27-52
			// find priority of missing item
			// sum of priorities
			// sumPriorities := 0
			var itemPriority int
			if unicode.IsUpper(teamLetter) {
				itemPriority = int(teamLetter) - 38
			} else {
				itemPriority += int(teamLetter) - 96
			}
			sumPriorities += itemPriority
		}
	} // add priority to sum
	fmt.Println(sumPriorities)
}

func removeDuplicateStr(strSlice []rune) []rune {
	allKeys := make(map[rune]bool)
	list := []rune{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
