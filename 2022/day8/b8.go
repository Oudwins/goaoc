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

	trees := loadMatrix(file)

	// fmt.Println("LENGTH OF TREES = ", len(trees), "|", len(trees[0]))
	// fmt.Println(tallestInCol, tallestInRow)
	// run algo
	maxScenicScore := 0
	for i := 0; i < len(trees); i++ {

		for j := 0; j < len(trees[i]); j++ {
			// css order, top, right, bottom, left
			// //1. checking to see if tall tree blocks that path
			// 2. Run algo -> css order
			// for ^ its i-- until i<0
			// for > its j++ until j < len-1
			// for ∨ its i++ until i < len-1
			// for < its j-- until j < 0

			upScore := calcScoreUp(trees, i, j)
			rightScore := calcScoreRight(trees, i, j)
			downScore := calcScoreDown(trees, i, j)
			leftScore := calcScoreLeft(trees, i, j)

			curTreeScore := upScore * rightScore * downScore * leftScore

			// fmt.Println("NUMBER AT: ", i, "|", j, " Score of: ", curTreeScore)
			// fmt.Printf("%v,%v,%v,%v\n", upScore, rightScore, downScore, leftScore)
			if curTreeScore > maxScenicScore {
				maxScenicScore = curTreeScore
			}
		}
	}
	fmt.Println(maxScenicScore)
}

func calcScoreUp(trees treeMatrix, row int, col int) int {
	for i := row - 1; i > -1; i-- {

		if trees[i][col] >= trees[row][col] {
			return row - i
		}
	}

	return row
}

func calcScoreRight(trees treeMatrix, row int, col int) int {
	for i := col + 1; i < len(trees[row]); i++ {
		if trees[row][i] >= trees[row][col] {
			return i - col
		}
	}
	return len(trees[row]) - col - 1
}

func calcScoreDown(trees treeMatrix, row int, col int) int {
	for i := row + 1; i < len(trees); i++ {
		if trees[i][col] >= trees[row][col] {
			return i - row
		}
	}
	return len(trees) - row - 1
}

func calcScoreLeft(trees treeMatrix, row int, col int) int {
	for i := col - 1; i > -1; i-- {
		if trees[row][i] >= trees[row][col] {
			return col - i
		}
	}
	return col
}

func loadMatrix(file *os.File) (matrix treeMatrix) {
	matrix = make(treeMatrix, 0)
	scanner := bufio.NewScanner(file)

	rowIdx := 0
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, make([]int8, 0))
		for _, char := range line {
			n, err := strconv.Atoi(string(char))
			if err != nil {
				fmt.Println("FAILED TO CONVERT MATRIX NUMBER. SOMETHING WENT WRONG")
			}

			matrix[rowIdx] = append(matrix[rowIdx], int8(n))
		}

		rowIdx++
	}

	return matrix
}
