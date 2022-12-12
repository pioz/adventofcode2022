package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

const (
	TOP = iota
	RIGHT
	BOTTOM
	LEFT
)

type Node struct {
	Height     float64
	Neighbours [4]*Node
	// Dijkstra stuffs
	Prev     *Node
	Distance float64
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	var nodes []*Node
	var end *Node
	set := make(map[*Node]struct{})

	i := 0
	cols := 0
	for scanner.Scan() {
		s := scanner.Text()
		cols = len(s)
		for _, c := range s {
			node := &Node{Distance: math.Inf(0)}
			set[node] = struct{}{}
			switch c {
			case 'S':
				c = 'a'
			case 'E':
				c = 'z'
				end = node
			}
			node.Height = float64(int(c) - 97)
			topIdx := i - cols
			if topIdx >= 0 {
				w := nodes[topIdx].Height - node.Height
				if w >= -1 {
					node.Neighbours[TOP] = nodes[topIdx]
				}
				w = node.Height - nodes[topIdx].Height
				if w >= -1 {
					nodes[topIdx].Neighbours[BOTTOM] = node
				}
			}
			leftIdx := i - 1
			if (i%cols)-1 >= 0 {
				w := nodes[leftIdx].Height - node.Height
				if w >= -1 {
					node.Neighbours[LEFT] = nodes[leftIdx]
				}
				w = node.Height - nodes[leftIdx].Height
				if w >= -1 {
					nodes[leftIdx].Neighbours[RIGHT] = node
				}
			}
			nodes = append(nodes, node)
			i++
		}
	}

	backupNodes := nodes

	// Perform Dijkstra
	end.Distance = 0
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Distance < nodes[j].Distance
	})
	for {
		if len(nodes) == 0 {
			break
		}
		current := nodes[0]
		nodes = nodes[1:] // remove current
		if current.Distance == math.Inf(0) {
			break
		}
		for i := 0; i < len(current.Neighbours); i++ {
			if current.Neighbours[i] != nil {
				alt := current.Distance + 1
				if alt < current.Neighbours[i].Distance {
					current.Neighbours[i].Distance = alt
					current.Neighbours[i].Prev = current
					sort.Slice(nodes, func(i, j int) bool {
						return nodes[i].Distance < nodes[j].Distance
					})
				}
			}
		}
	}
	// Find min path
	minCount := math.Inf(0)
	for i := 0; i < len(backupNodes); i++ {
		if backupNodes[i].Height == 0 {
			count := float64(0)
			current := backupNodes[i]
			for current.Prev != nil {
				current = current.Prev
				count++
			}
			if count > 0 && count < minCount {
				minCount = count
			}
		}

	}
	fmt.Println(minCount)
}
