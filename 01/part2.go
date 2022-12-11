package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func findMinIndex(a []int) int {
	index := 0
	for i, value := range a {
		if value < a[index] {
			index = i
		}
	}
	return index
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	top3 := []int{0, 0, 0}
	minIndex := 0
	current := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			if current > top3[minIndex] {
				top3[minIndex] = current
				minIndex = findMinIndex(top3)
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

	fmt.Println(top3[0] + top3[1] + top3[2])
}
