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
}

var winAgainst = map[string]string{
	"A": "B",
	"B": "C",
	"C": "A",
}

var loseAgainst = map[string]string{
	"A": "C",
	"B": "A",
	"C": "B",
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
		oppMove, myPotentialMove := moves[0], moves[1]
		if myPotentialMove == "X" {
			score += 0 + moveValues[loseAgainst[oppMove]]
		} else if myPotentialMove == "Y" {
			score += 3 + moveValues[oppMove]
		} else {
			score += 6 + moveValues[winAgainst[oppMove]]
		}
	}
	fmt.Println(score)
}
