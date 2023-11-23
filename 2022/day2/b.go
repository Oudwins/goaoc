package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// Oponent
// A Rock
// B Paper
// C Scissors
//
//	Me
//
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
// PART 2
// X = lose
// Y = draw
// z = win

func main() {

	file, err := os.Open("./day2/input.txt")
	if err != nil {
		fmt.Println("error:", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	convertLetters := map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
		// lose
		"X": 0,
		// draw
		"Y": 1,
		// win
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
	myTotalScore := 0
	for scanner.Scan() {
		line := scanner.Text()
		curPlays := strings.Split(line, " ")
		if len(curPlays) == 2 {
			enemyPlay, epExists := convertLetters[curPlays[0]]
			desiredResult, drExists := convertLetters[curPlays[1]]
			if epExists && drExists {
				myPlay, _ := calcMyPlay(desiredResult, enemyPlay, battleResult)
				result := battleResult[enemyPlay][myPlay]
				shapeBonus := shapeScore[myPlay]
				myTotalScore += result + shapeBonus
			}
		}
	}
	fmt.Println(myTotalScore)
}

func calcMyPlay(desiredResult int, enemyPlay int, resultsArr [3][3]int) (int, error) {
	for i := 0; i < len(resultsArr[enemyPlay]); i++ {
		if resultsArr[enemyPlay][i] == desiredResult*3 {
			return i, nil
		}
	}
	return -1, errors.New("Couldn't find the play")
}
