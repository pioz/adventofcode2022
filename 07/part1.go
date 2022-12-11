package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Node struct {
	Name     string
	Size     int
	Dir      bool
	Parent   *Node
	Children []*Node
}

func (n *Node) Add(child *Node) {
	child.Parent = n
	n.Children = append(n.Children, child)
	p := child.Parent
	for p != nil {
		p.Size += child.Size
		p = p.Parent
	}
}

func (n *Node) TotSizeAtMost(atMost int) int {
	s := 0
	if n.Size < atMost {
		s += n.Size
	}
	for i := 0; i < len(n.Children); i++ {
		if n.Children[i].Dir {
			s += n.Children[i].TotSizeAtMost(atMost)
		}
	}
	return s
}

var cdRegexp = regexp.MustCompile(`^\$ cd (.*)`)
var lsRegexp = regexp.MustCompile(`^\$ ls`)
var dirRegexp = regexp.MustCompile(`^dir (.*)`)
var fileRegexp = regexp.MustCompile(`^(\d+) (.*)`)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	var tree, current *Node

	for scanner.Scan() {
		s := scanner.Text()
		if m := cdRegexp.FindAllStringSubmatch(s, -1); m != nil {
			path := m[0][1]
			switch path {
			case "/":
				tree = &Node{Name: "/", Dir: true}
				current = tree
			case "..":
				current = current.Parent
			default:
				n := &Node{Name: path, Dir: true}
				if current != nil {
					current.Add(n)
				}
				current = n
			}
			continue
		}
		if m := fileRegexp.FindAllStringSubmatch(s, -1); m != nil {
			size, _ := strconv.Atoi(m[0][1])
			filename := m[0][2]
			n := &Node{Name: filename, Size: size}
			if current != nil {
				current.Add(n)
			}
			continue
		}
	}
	fmt.Println(tree.TotSizeAtMost(100000))
}
