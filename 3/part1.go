package main

import (
	"bufio"
	"fmt"
	"os"
)

func priority(r rune) int32 {
	if r > 90 {
		return r - 97
	}
	return r - 65 + 26
}

func foundPriority(s string) int32 {
	var found [52]bool
	for i := range s {
		p := priority(rune(s[i]))
		if i < len(s)/2 {
			found[p] = true
		} else if found[p] {
			return p + 1
		}
	}
	return 0
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var sum int32
	for scanner.Scan() {
		s := scanner.Text()
		sum += foundPriority(s)
	}
	fmt.Println(sum)
}
