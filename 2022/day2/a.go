package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("./day2/input.txt")
	if err != nil {
		fmt.Println("error:", err)
	}
	defer file.Close()
	playPostion := map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
		"X": 0,
		"Y": 1,
		"Z": 2,
	}

	battleResult := [3][3]int{
		//[0][x]
		{3, 6, 0},
		//[1][x]
		{0, 3, 6},
		//[2][x]
		{6, 0, 3},
	}

	shapeScore := [3]int{1, 2, 3}
	// Oponent
	// A Rock
	// B Paper
	// C Scissors
	//  Me
	// x Rock
	// y paper
	// c scissors
	// SCORES
	// shape
	// x = 1
	// y = 2
	// c = 3tring
	// outcome
	// 0 if lost
	// 3 if draw
	// 6 if won
	myTotalScore := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		curPlays := strings.Split(line, " ")
		if len(curPlays) == 2 {
			enemyPlay, epExists := playPostion[curPlays[0]]
			myPlay, mpExists := playPostion[curPlays[1]]

			if epExists && mpExists {
				result := battleResult[enemyPlay][myPlay]
				shapeBonus := shapeScore[myPlay]
				myTotalScore += result + shapeBonus
			}
		}
	}
	fmt.Println(myTotalScore)
}
