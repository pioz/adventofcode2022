package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		s := scanner.Text()
		pair := strings.Split(s, ",")
		elf1 := strings.Split(pair[0], "-")
		elf2 := strings.Split(pair[1], "-")
		elf1From, _ := strconv.Atoi(elf1[0])
		elf1To, _ := strconv.Atoi(elf1[1])
		elf2From, _ := strconv.Atoi(elf2[0])
		elf2To, _ := strconv.Atoi(elf2[1])
		if elf1From >= elf2From && elf1To <= elf2To || elf1From <= elf2From && elf1To >= elf2To {
			count++
		}
	}
	fmt.Println(count)
}
