package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./day6/input.txt")
	if err != nil {
		fmt.Println("error:", err)
	}
	defer file.Close()
	// what ends up on top
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		startPaket := -1
		line := scanner.Text()
		// line := "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"

		nCharsInPacketPrefix := 4
		idx := 0
		for idx < len(line)-nCharsInPacketPrefix {
			chars := []rune{}
			for i := 0; i < nCharsInPacketPrefix; i++ {
				charIdx := idx + i
				// fmt.Println(chars)
				chars = append(chars, rune(line[charIdx]))
			}
			// fmt.Println(chars)
			if areAllUnique(chars) {
				// fmt.Println("INDEX IS: ", idx)
				startPaket = idx + nCharsInPacketPrefix
				break
			}
			idx += 1

		}
		fmt.Println(startPaket)
	}

}

func areAllUnique(charList []rune) bool {
	uniqueChars := make(map[rune]bool)
	for _, char := range charList {
		if uniqueChars[char] {
			return false
		}
		uniqueChars[char] = true
	}
	return true
}

// 4 unique characters in a row represents end of subroutine. Return n of characters processed before start of sequence
