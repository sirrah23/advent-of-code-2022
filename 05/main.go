package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Stack struct {
	Items []string
}

func (s *Stack) Push(item string) {
	s.Items = append([]string{item}, s.Items...)
}

func (s *Stack) Pop() string {
	item := s.Items[0]
	s.Items = s.Items[1:]
	return item
}

func (s *Stack) Peek() string {
	return s.Items[0]
}

func NewStack(items ...string) *Stack {
	return &Stack{
		Items: items,
	}
}

func ExtractMove(text string) (numItems, fromIdx, toIdx int) {
	re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	moves := sliceAtoi(re.FindStringSubmatch(text)[1:])
	return moves[0], moves[1] - 1, moves[2] - 1
}

func sliceAtoi(in []string) []int {
	out := make([]int, len(in))
	for i, s := range in {
		pI, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		out[i] = pI
	}
	return out
}

func main() {
	stacks := []*Stack{
		NewStack("N", "Z"),
		NewStack("D", "C", "M"),
		NewStack("P"),
	}

	file, err := os.Open("moves.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		numItems, from, to := ExtractMove(scanner.Text())
		fmt.Println(numItems, from, to)
		for i := 0; i < numItems; i++ {
			item := stacks[from].Pop()
			stacks[to].Push(item)
		}
	}
	for _, stack := range stacks {
		fmt.Print(stack.Peek())
	}
}
