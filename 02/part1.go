package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	score := 0
	decryptElf := map[rune]rune{'A': 'R', 'B': 'P', 'C': 'S'}
	decryptMe := map[rune]rune{'X': 'R', 'Y': 'P', 'Z': 'S'}
	scores := map[rune]int{'R': 1, 'P': 2, 'S': 3}
	for scanner.Scan() {
		s := scanner.Text()
		elfMove := decryptElf[rune(s[0])]
		myMove := decryptMe[rune(s[2])]
		if elfMove == myMove {
			score += 3 + scores[myMove]
			continue
		}
		if (elfMove == 'R' && myMove == 'P') || (elfMove == 'P' && myMove == 'S') || (elfMove == 'S' && myMove == 'R') {
			score += 6 + scores[myMove]
			continue
		}
		score += 0 + scores[myMove]
	}
	fmt.Println(score)
}
