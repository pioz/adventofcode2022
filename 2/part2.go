package main

import (
	"bufio"
	"fmt"
	"os"
)

var scores = []int{1, 2, 3}

func myScore(elfMove int, result int) int {
	if result == 0 {
		return scores[elfMove] + 3
	}
	if result > 0 {
		index := elfMove + 1
		if index == 3 {
			index = 0
		}
		return scores[index] + 6
	}
	index := elfMove - 1
	if index == -1 {
		index = 2
	}
	return scores[index] + 0
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	score := 0
	decryptElf := map[rune]int{'A': 0, 'B': 1, 'C': 2}
	decryptResult := map[rune]int{'X': -1, 'Y': 0, 'Z': 1}
	for scanner.Scan() {
		s := scanner.Text()
		elfMove := decryptElf[rune(s[0])]
		result := decryptResult[rune(s[2])]
		score += myScore(elfMove, result)
	}
	fmt.Println(score)
}
