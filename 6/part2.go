package main

import (
	"bufio"
	"fmt"
	"os"
)

const startOfPacketSize = 14

var freqTable [256]rune

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	var (
		i              = 0
		sameCharacters = 0
		mem            [startOfPacketSize]rune
	)

	for scanner.Scan() {
		c := rune(scanner.Text()[0])
		freqTable[c]++
		if freqTable[c] == 2 {
			sameCharacters++
		}
		if i >= startOfPacketSize-1 { // skip till window is full
			old := mem[i%startOfPacketSize] // first time is 0, but is ok
			freqTable[old]--
			if freqTable[old] == 1 {
				sameCharacters--
			}
			if sameCharacters == 0 {
				break
			}
		}
		mem[i%startOfPacketSize] = c
		i++
	}
	fmt.Println(i + 1)
}
