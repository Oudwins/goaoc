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

	fileContents := strings.Split(string(file), "\n")
	divPkg1 := "[[2]]"
	divPkg2 := "[[6]]"
	fileContents = append(fileContents, divPkg1, divPkg2)
	answer := 1
	var arr [][]interface{}
	for _, line := range fileContents {
		if line == "" {
			continue
		}
		arr = append(arr, parseLine(line))
	}

	quickSort(&arr)
	for i := 0; i < len(arr); i++ {
		if len(arr[i]) == 1 {
			s, ok := arr[i][0].([]interface{})
			if ok && len(s) == 1 {
				n, ok := s[0].(int)
				if ok && (n == 2 || n == 6) {
					// fmt.Println(arr[i], i)
					answer *= (i + 1)
				}
			}
		}
	}

	// fmt.Println(arr)
	fmt.Println(answer)

	// fmt.Println("TOTAL IN ORDER: ", sum)

	// p2 is a sorting algorith
	// Quick sort -> How to do it? -> less memory but might be slower
	// merge sort -> I can do it! -> more memory but always fast
	// keep track of position of my two additional elements!
	// [[2]]
	// [[6]]

}

func quickSort(arr *[][]interface{}) {
	qs(arr, 0, len(*arr)-1)
}

func qs_partition(arr *[][]interface{}, low int, high int) int {
	pivot := (*arr)[high]

	idx := low - 1

	for i := low; i < high; i++ {
		if inOrder((*arr)[i], pivot) {
			idx++
			tmp := (*arr)[i]
			(*arr)[i] = (*arr)[idx]
			(*arr)[idx] = tmp
		}
	}
	idx++
	(*arr)[high] = (*arr)[idx]
	(*arr)[idx] = pivot

	return idx
}

func qs(arr *[][]interface{}, low int, high int) {
	// base case
	if low >= high {
		return
	}

	pivotIdx := qs_partition(arr, low, high)

	qs(arr, low, pivotIdx-1)
	qs(arr, pivotIdx+1, high)
}

func inOrder(pkg1 []interface{}, pkg2 []interface{}) bool {
	res, _ := areInOrder(pkg1, pkg2)
	return res
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
