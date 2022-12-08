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
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	var i, j int
	var m [][]int

	for scanner.Scan() {
		c := scanner.Text()
		if i == 0 {
			m = append(m, []int{})
		}
		if c == "\n" {
			i = 0
			j++
			continue
		}
		tree, _ := strconv.Atoi(c)
		m[j] = append(m[j], tree)
		i++
	}

	maxScore := 0
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			scoreRight := 0
			for k := j + 1; k < len(m[i]); k++ {
				scoreRight++
				if m[i][j] <= m[i][k] {
					break
				}
			}
			scoreLeft := 0
			for k := j - 1; k >= 0; k-- {
				scoreLeft++
				if m[i][j] <= m[i][k] {
					break
				}
			}
			scoreBottom := 0
			for k := i + 1; k < len(m); k++ {
				scoreBottom++
				if m[i][j] <= m[k][j] {
					break
				}
			}
			scoreTop := 0
			for k := i - 1; k >= 0; k-- {
				scoreTop++
				if m[i][j] <= m[k][j] {
					break
				}
			}
			score := scoreTop * scoreRight * scoreBottom * scoreLeft
			if score > maxScore {
				maxScore = score
			}
		}
	}
	fmt.Println(maxScore)
}
