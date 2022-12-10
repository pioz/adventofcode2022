package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var commandRegexp = regexp.MustCompile(`^(\w+)\s?(-?\d*)$`)

func setPixel(crt *[240]bool, cycle, x int) {
	crt[cycle-1] = x-1 <= (cycle-1)%40 && (cycle-1)%40 <= x+1
}

func print(crt *[240]bool) {
	for i := 0; i < 6; i++ {
		for j := 0; j < 40; j++ {
			if crt[j+i*40] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	cycle := 1
	x := 1
	crt := [240]bool{}

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
			setPixel(&crt, cycle, x)
			cycle++
		case "addx":
			setPixel(&crt, cycle, x)
			cycle++
			setPixel(&crt, cycle, x)
			cycle++
			x += value
		default:
			panic("invalid command")
		}
	}
	print(&crt)
}
