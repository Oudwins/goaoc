package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const matrixCubeSize = 99

type treeMatrix [][]int8

func main() {
	// each tree = single digit
	// count number of trees visible from outside
	// tree is visible if all other trees between it & the end of the grid are shorter than it (row or col)
	// All edge trees are visible

	// NOTES
	// 1. if I tree is visible, its visible so no need to continue trying to see if it is
	// 2. length_of_cube * 4 - 4 is the formula for the side trees
	// 3. We probably want to check the sides that are closer first
	// 4. there has to be a way do to this better by keeping track of highest tree in different directions
	// 5. start at matrix[1][1] until matrix[len-2][len-2]

	// WITHOUT THE POSSIBLE DIRECTIONS ITS TOO HIGH
	// WITH THE POSSIBLE DIRECTIONS ITS TOO LOW...
	// MEANS THERE IS A MISTAKE ON BOTH PARTS OF MY CODE...

	file, err := os.Open("./day8/input.txt")
	if err != nil {
		fmt.Println("error:", err)
	}
	defer file.Close()

	trees, tallestInCol, tallestInRow := loadMatrix(file)

	visibleTrees := (len(trees) * 4) - 4
	// fmt.Println("LENGTH OF TREES = ", len(trees), "|", len(trees[0]))
	fmt.Println(tallestInCol, tallestInRow)
	// run algo

	for i := 1; i < len(trees)-1; i++ {

		for j := 1; j < len(trees[i])-1; j++ {
			// css order, top, right, bottom, left
			possibleDirections := [4]bool{true, true, true, true}
			// //1. checking to see if tall tree blocks that path
			for _, val := range tallestInCol[j] {
				if val > i {
					possibleDirections[2] = false
				}
				if val < i {
					possibleDirections[0] = false
				}
			}
			for _, val := range tallestInRow[i] {
				if val > j {
					possibleDirections[1] = false
				}
				if val < j {
					possibleDirections[3] = false
				}
			}

			// 2. Run algo -> css order
			// for ^ its i-- until i<0
			// for > its j++ until j < len-1
			// for ∨ its i++ until i < len-1
			// for < its j-- until j < 0
			if (possibleDirections[0] && isTreeVisibleUp(trees, i, j)) || (possibleDirections[1] && isTreeVisibleRight(trees, i, j)) || (possibleDirections[2] && isTreeVisibleDown(trees, i, j)) || (possibleDirections[3] && isTreeVisibleLeft(trees, i, j)) {
				fmt.Println("TREE AT POS: ", i, " | ", j)
				visibleTrees++
			}

		}
	}
	// fmt.Println(trees)
	// for _, row := range trees {

	// 	for _, char := range row {
	// 		fmt.Printf("%v", char)
	// 	}

	// fmt.Println()
	// }
	fmt.Println(visibleTrees)
}

func isTreeVisibleUp(trees treeMatrix, row int, col int) bool {
	for i := row - 1; i > -1; i-- {

		if trees[i][col] >= trees[row][col] {
			return false
		}
	}

	return true
}

func isTreeVisibleRight(trees treeMatrix, row int, col int) bool {
	for i := col + 1; i < len(trees[row]); i++ {
		if trees[row][i] >= trees[row][col] {
			return false
		}
	}
	return true
}

func isTreeVisibleDown(trees treeMatrix, row int, col int) bool {
	for i := row + 1; i < len(trees); i++ {
		if trees[i][col] >= trees[row][col] {
			return false
		}
	}
	return true
}

func isTreeVisibleLeft(trees treeMatrix, row int, col int) bool {
	for i := col - 1; i > -1; i-- {
		if trees[row][i] >= trees[row][col] {
			return false
		}
	}
	return true
}

func loadMatrix(file *os.File) (matrix treeMatrix, tallestInCol [matrixCubeSize][]int, tallestInRow [matrixCubeSize][]int) {
	matrix = make(treeMatrix, 0)
	scanner := bufio.NewScanner(file)

	rowIdx := 0
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, make([]int8, 0))
		for colIdx, char := range line {
			n, err := strconv.Atoi(string(char))
			if err != nil {
				fmt.Println("FAILED TO CONVERT MATRIX NUMBER. SOMETHING WENT WRONG")
			}

			matrix[rowIdx] = append(matrix[rowIdx], int8(n))
			if n == 9 {
				tallestInCol[colIdx] = append(tallestInCol[colIdx], rowIdx)
				tallestInRow[rowIdx] = append(tallestInRow[rowIdx], colIdx)
			}
		}

		rowIdx++
	}

	return matrix, tallestInCol, tallestInRow
}
