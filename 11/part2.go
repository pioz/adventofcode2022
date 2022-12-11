package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var regexp1 = regexp.MustCompile(`^  Starting items: (.*)$`)
var regexp2 = regexp.MustCompile(`^  Operation: (.*)$`)
var regexp3 = regexp.MustCompile(`^  Test: divisible by (\d+)$`)
var regexp4 = regexp.MustCompile(`^    If true: throw to monkey (\d+)$`)
var regexp5 = regexp.MustCompile(`^    If false: throw to monkey (\d+)$`)
var opRegexp = regexp.MustCompile(`^new = old ([\+\*]) (.*)$`)

type Monkey struct {
	Items      []int
	Op         string
	Divide     int
	True       int
	False      int
	Inspection int
}

func (m *Monkey) Exec(level int) int {
	tokens := opRegexp.FindAllStringSubmatch(m.Op, -1)
	if tokens == nil {
		panic("invalid operation")
	}
	var value int
	op := tokens[0][1]
	sValue := tokens[0][2]
	if sValue == "old" {
		value = level
	} else {
		value, _ = strconv.Atoi(sValue)
	}
	switch op {
	case "+":
		return level + value
	case "*":
		return level * value
	}
	panic("operation error")
}

var monkeys []Monkey

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for i := 0; scanner.Scan(); i++ {
		s := scanner.Text()
		idx := len(monkeys) - 1

		switch i % 7 {
		case 0:
			monkeys = append(monkeys, Monkey{})
		case 1:
			tokens := regexp1.FindAllStringSubmatch(s, -1)
			for _, v := range strings.Split(tokens[0][1], ", ") {
				level, _ := strconv.Atoi(v)
				monkeys[idx].Items = append(monkeys[idx].Items, level)
			}
		case 2:
			tokens := regexp2.FindAllStringSubmatch(s, -1)
			monkeys[idx].Op = tokens[0][1]
		case 3:
			tokens := regexp3.FindAllStringSubmatch(s, -1)
			v, _ := strconv.Atoi(tokens[0][1])
			monkeys[idx].Divide = v
		case 4:
			tokens := regexp4.FindAllStringSubmatch(s, -1)
			v, _ := strconv.Atoi(tokens[0][1])
			monkeys[idx].True = v
		case 5:
			tokens := regexp5.FindAllStringSubmatch(s, -1)
			v, _ := strconv.Atoi(tokens[0][1])
			monkeys[idx].False = v
		case 6:
		default:
			panic("parse error")
		}
	}

	lcm := 1
	for i := 0; i < len(monkeys); i++ {
		lcm *= monkeys[i].Divide
	}

	for i := 0; i < 10000; i++ {
		for j := 0; j < len(monkeys); j++ {
			for _, level := range monkeys[j].Items {
				new := monkeys[j].Exec(level) % lcm
				var to int
				if new%monkeys[j].Divide == 0 {
					to = monkeys[j].True
				} else {
					to = monkeys[j].False
				}
				monkeys[to].Items = append(monkeys[to].Items, new)
				monkeys[j].Inspection++
			}
			monkeys[j].Items = []int{}
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].Inspection > monkeys[j].Inspection
	})

	fmt.Println(monkeys[0].Inspection * monkeys[1].Inspection)
}
