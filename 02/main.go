package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var moveValues = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var moveMapping = map[string]string{
	"A": "rock",
	"B": "paper",
	"C": "scissors",
	"X": "rock",
	"Y": "paper",
	"Z": "scissors",
}

func playGame(move1, move2 string) int {
	if move1 == move2 {
		return 0
	}
	if move1 == "rock" {
		if move2 == "paper" {
			return -1
		} else if move2 == "scissors" {
			return 1
		}
	}
	if move1 == "paper" {
		if move2 == "rock" {
			return 1
		} else if move2 == "scissors" {
			return -1
		}
	}
	if move1 == "scissors" {
		if move2 == "rock" {
			return -1
		} else if move2 == "paper" {
			return 1
		}
	}

	return -42
}

func main() {
	_ = moveValues
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	score := 0
	for scanner.Scan() {
		round := scanner.Text()
		moves := strings.Split(round, " ")
		oppMove, myMove := moves[0], moves[1]
		resultScore := 0
		gameResult := playGame(moveMapping[oppMove], moveMapping[myMove])
		if gameResult == 1 {
			resultScore = 0
		} else if gameResult == -1 {
			resultScore = 6
		} else if gameResult == 0 {
			resultScore = 3
		} else {
			panic("okay")
		}
		score += moveValues[myMove] + resultScore
		fmt.Println(score)
	}
}
