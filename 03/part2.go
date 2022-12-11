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

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var sum int32
	found := [52]byte{}
	for i := 0; scanner.Scan(); i++ {
		s := scanner.Text()
		for j := range s {
			p := priority(rune(s[j]))
			found[p] |= 1 << (i % 3) // 000 | 001 | 010 | 100 = 111 => 7
			if found[p] == 7 {
				sum += p + 1
				found = [52]byte{} // reset for next group
				break
			}
		}
	}
	fmt.Println(sum)
}
