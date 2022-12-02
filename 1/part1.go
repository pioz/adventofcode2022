package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	max, current := 0, 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			if current > max {
				max = current
			}
			current = 0
		} else {
			kal, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			current += kal
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println(max)
}
