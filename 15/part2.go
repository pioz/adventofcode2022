package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Point struct {
	X int
	Y int
}

type Sensor struct {
	Coord       Point
	BeaconCoord Point
	Dist        int
}

var regexpCoords = regexp.MustCompile(`^Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)$`)

func Abs(x int) int {
	return int(math.Abs(float64(x)))
}

func Dist(p1, p2 Point) int {
	return Abs(p1.X-p2.X) + Abs(p1.Y-p2.Y) + 1
}

var dir [4][2]int = [4][2]int{{-1, -1}, {-1, 1}, {1, 1}, {1, -1}}

// Returns the number of sensors they would find a beacon at x,y
func Find(sensors []Sensor, p Point) int {
	count := 0
	for i := 0; i < len(sensors); i++ {
		if sensors[i].BeaconCoord.X == p.X && sensors[i].BeaconCoord.Y == p.Y {
			return 0
		}
		dist := Dist(p, sensors[i].Coord)
		if dist <= sensors[i].Dist {
			count++
		}
	}
	return count
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	sensors := make([]Sensor, 0)

	for scanner.Scan() {
		s := scanner.Text()
		tokens := regexpCoords.FindAllStringSubmatch(s, -1)
		sx, _ := strconv.Atoi(tokens[0][1])
		sy, _ := strconv.Atoi(tokens[0][2])
		bx, _ := strconv.Atoi(tokens[0][3])
		by, _ := strconv.Atoi(tokens[0][4])
		p := Point{sx, sy}
		b := Point{bx, by}
		sensors = append(sensors, Sensor{p, b, Dist(p, b)})
	}

	for i := 0; i < len(sensors); i++ {
		x := sensors[i].Coord.X
		y := sensors[i].Coord.Y
		dist := sensors[i].Dist
		x += dist

		for d := 0; d < 4; d++ {
			for j := 0; j < dist; j++ {
				if Find(sensors, Point{x, y}) == 0 && 0 <= x && x <= 4000000 && 0 <= y && y <= 4000000 {
					fmt.Println(4000000*x + y)
					os.Exit(0)
				}
				if j != dist-1 {
					x += dir[d][0]
					y += dir[d][1]
				}
			}
		}
	}
}
