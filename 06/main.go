package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	runTest()

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	line := scanner.Text()
	fmt.Println(markerSubroutine(line))
}

func markerSubroutine(input string) int {
	for i := 0; i+4 < len(input); i++ {
		currWindow := input[i : i+4]
		repeat := false
		for j, char := range currWindow {
			if strings.Index(currWindow, string(char)) != j {
				repeat = true
				break
			}
		}
		if !repeat {
			return i + 4
		}
	}
	return -1
}

func runTest() {
	for i, tc := range []struct {
		TestInput      string
		ExpectedMarker int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	} {
		if out := markerSubroutine(tc.TestInput); out != tc.ExpectedMarker {
			fmt.Printf("Expected %d but got %d for %s\n", tc.ExpectedMarker, out, tc.TestInput)
		} else {
			fmt.Printf("Pass test case %d\n", i)
		}
	}
}
