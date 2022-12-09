package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"os"
	"regexp"
	"strconv"
)

const rowSize = 1000
const bufferSize = rowSize * rowSize
const knotsSize = 10
const simulation = false

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
		return 0
	case "D":
		return 4
	}
	panic("invalid move")
}

func followTrail(h, t int) (int, bool) {
	col := (h % rowSize) - (t % rowSize)
	row := (h / rowSize) - (t / rowSize)

	if col == 0 && row == 2 {
		return 0, true
	}
	if col == 0 && row == -2 {
		return 4, true
	}
	if row == 0 && col == 2 {
		return 2, true
	}
	if row == 0 && col == -2 {
		return 6, true
	}
	if (col == 1 && row == 2) || (col == 2 && row == 1) || (col == 2 && row == 2) {
		return 1, true
	}
	if (col == 1 && row == -2) || (col == 2 && row == -1) || (col == 2 && row == -2) {
		return 3, true
	}
	if (col == -1 && row == -2) || (col == -2 && row == -1) || (col == -2 && row == -2) {
		return 5, true
	}
	if (col == -1 && row == 2) || (col == -2 && row == 1) || (col == -2 && row == 2) {
		return 7, true
	}
	return 0, false
}

func newImage(knots []int, p []int) *image.Paletted {
	palette := color.Palette{color.Black, color.White, color.RGBA{0x33, 0x33, 0x33, 0xff}}
	img := image.NewPaletted(image.Rect(0, 0, rowSize, rowSize), palette)
	for i := rowSize - 1; i >= 0; i-- {
		for j := 0; j < rowSize; j++ {
			img.Set(i, j, color.RGBA{0x00, 0x00, 0x00, 0xff})
			if p[j+i*rowSize] > 0 {
				img.Set(i, j, color.RGBA{0x33, 0x33, 0x33, 0xff})
			}
			for k := 0; k < knotsSize; k++ {
				if j+i*rowSize == knots[k] {
					img.Set(i, j, color.RGBA{0xff, 0xff, 0xff, 0xff})
					break
				}
			}
		}
	}
	return img
}

var moveRegexp = regexp.MustCompile(`^([UDLR]) (\d+)$`)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	var images []*image.Paletted
	var delays []int

	p := make([]int, bufferSize, bufferSize)
	knots := make([]int, knotsSize, knotsSize)
	for i := 0; i < knotsSize; i++ {
		knots[i] = bufferSize/2 + rowSize/2
	}
	p[knots[knotsSize-1]]++

	for scanner.Scan() {
		s := scanner.Text()
		tokens := moveRegexp.FindAllStringSubmatch(s, -1)
		move := moveToIndex(tokens[0][1])
		count, _ := strconv.Atoi(tokens[0][2])

		for i := 0; i < count; i++ {
			knots[0] += dir[move]
			for j := 1; j < knotsSize; j++ {
				idx, m := followTrail(knots[j-1], knots[j])
				if m {
					knots[j] += dir[idx]
				}
			}
			p[knots[knotsSize-1]]++
			if simulation {
				images = append(images, newImage(knots, p))
				delays = append(delays, 1)
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

	if simulation {
		f, err := os.OpenFile("simulation3.gif", os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		gif.EncodeAll(f, &gif.GIF{Image: images})
	}
}
