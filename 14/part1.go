package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	AIR = iota
	ROCK
	SAND
)

const (
	rows    = 172
	cols    = 80
	padding = 455
)

var ground [rows][cols]int

func print(x, y int) {
	for i := 0; i < len(ground); i++ {
		for j := 0; j < len(ground[i]); j++ {
			if j == 500-padding && i == 0 {
				fmt.Print("+")
				continue
			}
			if i == y && j == x {
				fmt.Print("o")
				continue
			}
			switch ground[i][j] {
			case ROCK:
				fmt.Print("#")
			case SAND:
				fmt.Print("O")
			default:
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func animate(x, y int) {
	print(x, y)
	time.Sleep(time.Millisecond * 50)
	fmt.Print("\033[H\033[2J")
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	height := 0
	for scanner.Scan() {
		s := scanner.Text()
		points := strings.Split(s, " -> ")
		var prevX, prevY int
		for _, point := range points {
			coords := strings.Split(point, ",")
			x, _ := strconv.Atoi(coords[0])
			y, _ := strconv.Atoi(coords[1])
			x -= padding
			if prevX == 0 && prevY == 0 {
				prevX = x
				prevY = y
				continue
			}
			if x < prevX {
				for i := x; i <= prevX; i++ {
					ground[y][i] = ROCK
				}
			} else if x > prevX {
				for i := prevX; i <= x; i++ {
					ground[y][i] = ROCK
				}
			} else if y < prevY {
				for i := y; i <= prevY; i++ {
					ground[i][x] = ROCK
				}
			} else if y > prevY {
				for i := prevY; i <= y; i++ {
					ground[i][x] = ROCK
				}
			}
			prevX = x
			prevY = y
			if y > height {
				height = y
			}
		}
	}

	count := 0
	abyss := false
	for {
		if abyss {
			break
		}
		x := 500 - padding
		y := 0
		for {
			// animate(x, y)
			if ground[y+1][x] == 0 {
				y++
				if y > height {
					abyss = true
					break
				}
				continue
			}
			if ground[y+1][x-1] == 0 {
				x--
				y++
				continue
			}
			if ground[y+1][x+1] == 0 {
				x++
				y++
				continue
			}
			ground[y][x] = SAND
			count++
			break
		}
	}
	fmt.Println(count)

}
