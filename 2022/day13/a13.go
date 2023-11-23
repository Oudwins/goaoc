package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	file, err := os.ReadFile("./day13/input.txt")
	if err != nil {
		log.Fatal("FAILED TO READ FILE")
	}

	fileContents := strings.Split(string(file), "\n\n")

	sum := 0
	for i, msgStr := range fileContents {
		msg := strings.Split(msgStr, "\n")
		p1 := parseLine(msg[0])
		p2 := parseLine(msg[1])
		inOrder, _ := areInOrder(p1, p2)

		if inOrder {
			parIdx := i + 1
			// fmt.Println("PAIR IN ORDER: ", parIdx)
			sum += parIdx
		}
	}

	fmt.Println("TOTAL IN ORDER: ", sum)

	// p2 is a sorting algorith
	// Quick sort -> How to do it? -> less memory but might be slower
	// merge sort -> I can do it! -> more memory but always fast
	// keep track of position of my two additional elements!
	// [[2]]
	// [[6]]

}

func areInOrder(pkg1 []interface{}, pkg2 []interface{}) (inOrder bool, exit bool) {
	maxIterations := int(math.Max(float64(len(pkg1)), float64(len(pkg2))))

loop:
	for i := 0; i < maxIterations; i++ {
		if i >= len(pkg1) || i >= len(pkg2) {
			inOrder = len(pkg2) > len(pkg1)
			exit = true
			break loop
		}

		switch v1 := pkg1[i].(type) {
		case int:
			switch v2 := pkg2[i].(type) {
			case int:
				// both ints
				if v1 == v2 {
					continue
				}
				inOrder = v2 > v1
				exit = true
				break loop

			case []interface{}:
				newSlice := []interface{}{v1}
				inOrder, exit = areInOrder(newSlice, v2)

				if exit {
					break loop
				}

			default:
				fmt.Println(v1)
				fmt.Println("SOMETHING WENT WRONG. INVALID TYPE")
			}

		case []interface{}:
			switch v2 := pkg2[i].(type) {
			case int:
				newSlice := []interface{}{v2}
				inOrder, exit = areInOrder(v1, newSlice)

				if exit {
					break loop
				}
			case []interface{}:
				// both arrays
				inOrder, exit = areInOrder(v1, v2)

				if exit {
					break loop
				}

			default:
				fmt.Println(v1)
				fmt.Println("SOMETHING WENT WRONG. INVALID TYPE")
			}
		default:
			fmt.Println(v1)
			fmt.Println("SOMETHING WENT WRONG. INVALID TYPE")
		}

	}

	// both ints =>
	// left < right => true
	// left > right => false
	// left == right => continue

	// both lists
	// compare first value of each list
	// left.legth < right.length => true

	// one int & 1 list
	// convert the int to list & compare lists
	return inOrder, exit
}

func parseLine(input string) []interface{} {
	idx := 1
	return recParseLine(input, &idx)
}

func recParseLine(input string, i *int) []interface{} {

	var res []interface{}

	// list := parseLine("[[[]]]")
parseLoop:
	for ; *i < len(input); *i++ {
		switch {
		case unicode.IsDigit(rune(input[*i])):
			for j := *i + 1; j < len(input); j++ {
				if !unicode.IsDigit(rune(input[j])) {
					n, _ := strconv.Atoi(string(input[*i:j]))
					res = append(res, n)

					*i = j
					if input[j] == ']' {
						break parseLoop
					}
					break
				}
			}

			// iterar hasta , o ]
			// if ] return
		case input[*i] == ',':
			continue
		case input[*i] == '[':
			// new list
			*i++
			child := recParseLine(input, i)
			res = append(res, child)
			// fmt.Println(res, i)

		case input[*i] == ']':
			break parseLoop
			// return res
		// number
		default:
			fmt.Println("INVALID INPUT COULD NOT PARSE")
		}
	}
	return res
}
