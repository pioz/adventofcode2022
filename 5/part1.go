package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var stack []string = []string{}

func fillStack(s string) bool {
	empty := true
	i := 0
	for j := 0; j < len(s); j += 4 {
		if len(stack) <= i {
			stack = append(stack, "")
		}
		if s[j] == '[' {
			stack[i] = s[j+1:j+2] + stack[i]
			empty = false
		}
		i++
	}
	return empty
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	moveRegexp := regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`)
	readingCurrentStack := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		if readingCurrentStack {
			if fillStack(s) {
				readingCurrentStack = false
			}
		}
		values := moveRegexp.FindAllStringSubmatch(s, -1)
		if values == nil {
			continue
		}
		amount, _ := strconv.Atoi(values[0][1])
		from, _ := strconv.Atoi(values[0][2])
		to, _ := strconv.Atoi(values[0][3])
		for i := 0; i < amount; i++ {
			l := len(stack[from-1])
			crate := stack[from-1][l-1 : l]
			stack[from-1] = stack[from-1][:l-1]
			stack[to-1] += crate
		}
	}
	for i := range stack {
		l := len(stack[i])
		fmt.Print(stack[i][l-1 : l])
	}
	fmt.Println()
}
