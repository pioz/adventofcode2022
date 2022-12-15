package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
)

const y = 2000000

// const y = 10

type Seg struct {
	Start  int
	End    int
	Merged bool
}

var regexpCoords = regexp.MustCompile(`^Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)$`)

func Abs(x int) int {
	return int(math.Abs(float64(x)))
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	segments := make([]Seg, 0)
	beacons := make(map[int]struct{})

	for scanner.Scan() {
		s := scanner.Text()
		tokens := regexpCoords.FindAllStringSubmatch(s, -1)
		sx, _ := strconv.Atoi(tokens[0][1])
		sy, _ := strconv.Atoi(tokens[0][2])
		bx, _ := strconv.Atoi(tokens[0][3])
		by, _ := strconv.Atoi(tokens[0][4])

		dist := Abs(sx-bx) + Abs(sy-by)
		intersection := dist - Abs(y-sy)
		if intersection >= 0 {
			startX := sx - intersection
			endX := sx + intersection
			segments = append(segments, Seg{startX, endX, false})
			if by == y && startX <= bx && bx <= endX {
				if _, found := beacons[bx]; !found {
					beacons[bx] = struct{}{}
				}
			}
		}
	}

	sort.Slice(segments, func(i, j int) bool {
		return segments[i].Start < segments[j].Start
	})

	// Merge segments
	prev := 0
	for i := 1; i < len(segments); i++ {
		if segments[prev].End >= segments[i].Start {
			if segments[i].End > segments[prev].End {
				segments[prev].End = segments[i].End
			}
			segments[i].Merged = true
		} else {
			prev = i
		}
	}

	// Count
	count := 0
	for i := 0; i < len(segments); i++ {
		if segments[i].Merged {
			continue
		}
		count += segments[i].End - segments[i].Start + 1
	}

	fmt.Println(count - len(beacons))
}
