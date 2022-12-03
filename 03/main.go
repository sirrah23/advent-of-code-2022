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
	total := 0
	for scanner.Scan() {
		rucksack := scanner.Text()
		compOne, compTwo := rucksack[:len(rucksack)/2], rucksack[len(rucksack)/2:]
		match := 0
		for _, c1 := range compOne {
			for _, c2 := range compTwo {
				if c1 == c2 {
					val := int(c1)
					if 65 <= val && val <= 90 {
						match = val - 64 + 26
					} else {
						match = val - 96
					}
				}
			}
		}
		total += match
	}
	fmt.Println(total)
}
