package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// const linesInStack = 8
// const N_OF_STACKS = 9

const N_OF_STACKS = 9

func main() {
	file, err := os.Open("./day5/input.txt")
	if err != nil {
		fmt.Println("error:", err)
	}
	defer file.Close()
	// what ends up on top
	scanner := bufio.NewScanner(file)
	stacks := loadStacks(scanner)

	for scanner.Scan() {
		line := scanner.Text()
		instruction, err := parseInstruction(line)
		if err != nil {
			fmt.Println("ERROR PARSING LINE: ", err)
			fmt.Println("AT LINE: ", line)
			continue
		}
		// fmt.Println(instruction)
		// DO INSTRUCTION

		creates := stacks[instruction[1]-1][len(stacks[instruction[1]-1])-instruction[0]:]
		stacks[instruction[1]-1] = stacks[instruction[1]-1][:len(stacks[instruction[1]-1])-instruction[0]]
		stacks[instruction[2]-1] = append(stacks[instruction[2]-1], creates...)
		// for i := 0; i < instruction[0]; i++ {
		// 	box, hasItems := stacks[instruction[1]-1].Pop()
		// 	if !hasItems {
		// 		fmt.Println(stacks)
		// 		fmt.Println("problem stack = ", instruction[1])
		// 		fmt.Println(stacks[instruction[1]-1])
		// 		log.Fatal("Tried to take out boxes from stack that didn't have any")

		// 	}
		// }
		printStacks(stacks)
	}

	// prints first item in each stack
	for _, stack := range stacks {
		char, _ := stack.Pop()
		fmt.Printf("%c", char)
	}
	fmt.Println()

	// mys := []rune{
	// 	'a',
	// 	'b',
	// 	'c',
	// }

	// fmt.Println(mys)
	// reverseList(mys)

	// fmt.Println(mys)
}

func parseInstruction(s string) ([3]int, error) {

	var instruction [3]int
	commands := 0
	// could also check if it has prefix (move)
	sParts := strings.Split(s, " ")
	for _, part := range sParts {
		IPart, err := strconv.Atoi(part)
		if err != nil {
			continue
		}

		instruction[commands] = IPart
		commands += 1

		if commands >= 3 {
			break
		}
	}

	if commands < 3 {
		return instruction, errors.New("Less commands than expected, expected 3")
	}

	return instruction, nil
}

func loadStacks(scanner *bufio.Scanner) [N_OF_STACKS]Stack {

	// scanner := bufio.NewScanner(file)
	var stacks [N_OF_STACKS]Stack
	for scanner.Scan() {
		// read file line by line,
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		for idx, char := range line {
			if unicode.IsLetter(char) {
				stackIndex := int(math.Floor(float64(idx) / float64(4)))

				// fmt.Println(i, idx, stackIndex)

				stacks[stackIndex] = append(stacks[stackIndex], char)
			}
		}
		// I check if its alphanumeric character
		// if it is I check what multiple of 4 it is to see in which bucket it belongs! ciel(x/4)
		// then I reverse each list
		// return the stacks
	}

	for _, stack := range stacks {
		reverseList(stack)
	}

	return stacks
}

func reverseList(l []rune) {

	for i := 0; i < len(l)/2; i++ {
		j := len(l) - i - 1

		l[i], l[j] = l[j], l[i]
	}
}

type Stack []rune

// takes pointer to a stack & element to append
func (s *Stack) Push(val rune) {
	*s = append(*s, val)
}

func (s *Stack) Pop() (rune, bool) {
	length := len(*s)
	if length == 0 {
		return 0, false
	}

	idx := length - 1
	el := (*s)[idx]
	*s = (*s)[:idx]
	return el, true
}

func printStacks(stacks [N_OF_STACKS]Stack) {

	// for _, runesSlice := range stacks {
	// 	str := string(runesSlice)
	// 	fmt.Println(str)
	// }
}
