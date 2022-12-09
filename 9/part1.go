package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const rowSize = 1000
const bufferSize = rowSize * rowSize

/*
Shift by direction

	[7]      [0]      [1]
	    +7   +8    +9
	       \  |  /
	[6] -1 <- X -> +1 [2]
	       /  |  \
	    -9   -8    -7
	[5]      [4]      [3]
*/
var dir [8]int = [8]int{
	rowSize,      // 0
	rowSize + 1,  // 1
	1,            // 2
	-rowSize + 1, // 3
	-rowSize,     // 4
	-rowSize - 1, // 5
	-1,           // 6
	rowSize - 1,  // 7
}

func moveToIndex(move string) int {
	switch move {
	case "R":
		return 2
	case "L":
		return 6
	case "U":
		return 4
	case "D":
		return 0
	}
	panic("invalid move")
}

func tailMove(h, t int) bool {
	if h == t {
		return false
	}
	for i := 0; i < 8; i++ {
		if h+dir[i] == t {
			return false
		}
	}
	return true
}

func followTrail(h, t int) int {
	col := (h % rowSize) - (t % rowSize)
	row := (h / rowSize) - (t / rowSize)
	if col > 0 && row > 0 {
		return 1
	}
	if col > 0 && row < 0 {
		return 3
	}
	if col < 0 && row < 0 {
		return 5
	}
	if col < 0 && row > 0 {
		return 7
	}
	return 0
}

var moveRegexp = regexp.MustCompile(`^([UDLR]) (\d+)$`)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	p := make([]int, bufferSize, bufferSize)
	h := bufferSize/2 + rowSize/2
	t := bufferSize/2 + rowSize/2
	p[t]++

	for scanner.Scan() {
		s := scanner.Text()
		tokens := moveRegexp.FindAllStringSubmatch(s, -1)
		move := moveToIndex(tokens[0][1])
		count, _ := strconv.Atoi(tokens[0][2])

		for i := 0; i < count; i++ {
			h += dir[move]
			if tailMove(h, t) {
				idx := followTrail(h, t)
				if idx == 0 {
					t += dir[move]
				} else {
					t += dir[idx]
				}
				p[t]++
			}
		}
	}
	count := 0
	for i := range p {
		if p[i] > 0 {
			count++
		}
	}
	fmt.Println(count)
}
