package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	ans := 0
	for scanner.Scan() {
		pair1, pair2 := extractPairs(scanner.Text())
		for i := pair1.Low; i <= pair1.High; i++ {
			if i >= pair2.Low && i <= pair2.High {
				ans++
				break
			}
		}
	}
	fmt.Println(ans)
}

type pair struct {
	Low  int
	High int
}

func extractPairs(line string) (pair, pair) {
	rawPairs := strings.Split(line, ",")
	rawPair1, rawPair2 := rawPairs[0], rawPairs[1]
	rawPair1Comps := strings.Split(rawPair1, "-")
	rawPair2Comps := strings.Split(rawPair2, "-")
	rawPair1Low, err := strconv.Atoi(rawPair1Comps[0])
	if err != nil {
		panic(err)
	}
	rawPair1High, err := strconv.Atoi(rawPair1Comps[1])
	if err != nil {
		panic(err)
	}
	rawPair2Low, err := strconv.Atoi(rawPair2Comps[0])
	if err != nil {
		panic(err)
	}
	rawPair2High, err := strconv.Atoi(rawPair2Comps[1])
	if err != nil {
		panic(err)
	}
	return pair{rawPair1Low, rawPair1High}, pair{rawPair2Low, rawPair2High}
}
