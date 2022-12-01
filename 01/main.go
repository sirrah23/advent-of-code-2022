package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	currTotal := 0
	var elfTotals []int
	for scanner.Scan() {
		currLineRaw := scanner.Text()
		if currLineRaw == "" {
			elfTotals = append(elfTotals, currTotal)
			currTotal = 0
		}
		currLine, _ := strconv.Atoi(currLineRaw)
		currTotal += currLine
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elfTotals)))

	fmt.Printf("Top 3 calories total: %d\n", elfTotals[0]+elfTotals[1]+elfTotals[2])
}
