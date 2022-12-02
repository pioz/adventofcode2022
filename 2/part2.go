package main

import (
	"bufio"
	"fmt"
	"os"
)

var scores = [3]int{1, 2, 3}

func myScore(elfMove int, result int) int {
	if result == 1 { // draw
		return scores[elfMove] + 3
	}
	if result == 2 { // win
		return scores[(elfMove+1)%3] + 6 // rotate right
	}
	// loss
	return scores[(elfMove-1+3)%3] + 0 // rotate left
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
		elfMove := int(s[0] - 65)
		result := int(s[2] - 88)
		score += myScore(elfMove, result)
	}
	fmt.Println(score)
}
