package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	total := 0
	for i := 0; i < len(lines); i += 3 {
		res := computeCurr(lines[i : i+3])
		if 65 <= res && res <= 90 {
			total += int(res - 64 + 26)
		} else {
			total += int(res - 96)
		}
	}
	fmt.Println(total)
}

func computeCurr(inputs []string) rune {
	counts := map[rune]int{}
	for _, input := range inputs {
		uniqueChars := map[rune]struct{}{}
		for _, c := range input {
			uniqueChars[c] = struct{}{}
		}
		for u := range uniqueChars {
			counts[u]++
		}
	}
	for k, v := range counts {
		if v == 3 {
			return k
		}
	}
	return 0
}
