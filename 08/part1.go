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

	count := len(m)*2 + (len(m[0])-2)*2
	for i := 1; i < len(m)-1; i++ {
		for j := 1; j < len(m[i])-1; j++ {
			// search right
			visible := true
			for k := j + 1; k < len(m[i]); k++ {
				if m[i][j] <= m[i][k] {
					visible = false
					break
				}
			}
			if visible {
				count++
				continue
			}
			// search left
			visible = true
			for k := j - 1; k >= 0; k-- {
				if m[i][j] <= m[i][k] {
					visible = false
					break
				}
			}
			if visible {
				count++
				continue
			}
			// search bottom
			visible = true
			for k := i + 1; k < len(m); k++ {
				if m[i][j] <= m[k][j] {
					visible = false
					break
				}
			}
			if visible {
				count++
				continue
			}
			// search top
			visible = true
			for k := i - 1; k >= 0; k-- {
				if m[i][j] <= m[k][j] {
					visible = false
					break
				}
			}
			if visible {
				count++
				continue
			}
		}
	}
	fmt.Println(count)
}
