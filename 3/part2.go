package main

import (
	"bufio"
	"fmt"
	"os"
)

func priority(r byte) int {
	if r > 90 {
		return int(r - 97)
	}
	return int(r - 65 + 26)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	sum := 0
	found := [52]byte{}
	for i := 0; scanner.Scan(); i++ {
		s := scanner.Text()
		for j := range s {
			p := priority(s[j])
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
