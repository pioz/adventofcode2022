package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var commandRegexp = regexp.MustCompile(`^(\w+)\s?(-?\d*)$`)

func checkSignal(cycle, x int) int {
	if cycle%40 == 20 {
		return cycle * x
	}
	return 0
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	cycle := 1
	x := 1
	signal := 0

	for scanner.Scan() {
		s := scanner.Text()
		tokens := commandRegexp.FindAllStringSubmatch(s, -1)
		cmd := tokens[0][1]
		value := 0
		if len(tokens[0]) == 3 {
			value, _ = strconv.Atoi(tokens[0][2])
		}
		switch cmd {
		case "noop":
			cycle++
			signal += checkSignal(cycle, x)
		case "addx":
			cycle++
			signal += checkSignal(cycle, x)
			x += value
			cycle++
			signal += checkSignal(cycle, x)
		default:
			panic("invalid command")
		}
	}
	fmt.Println(signal)
}
