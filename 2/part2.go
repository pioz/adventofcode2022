package main

import (
	"bufio"
	"fmt"
	"os"
)

func myScore(elfMove int, result int) int {
	switch result {
	case 0: // loss
		return 1 + (elfMove-1+3)%3 + 0 // rotate left
	case 1: // draw
		return 1 + elfMove + 3
	default: // win
		return 1 + (elfMove+1)%3 + 6 // rotate right
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	score := 0
	for scanner.Scan() {
		s := scanner.Text()
		elfMove := int(s[0]) - 65
		result := int(s[2]) - 88
		score += myScore(elfMove, result)
	}
	fmt.Println(score)
}
