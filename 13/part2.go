package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"sort"
)

func compare(p1, p2 interface{}) int {
	p1Value := reflect.ValueOf(p1)
	p2Value := reflect.ValueOf(p2)
	if p1Value.Kind() == reflect.Float64 && p2Value.Kind() == reflect.Float64 {
		if p1.(float64) < p2.(float64) {
			return 1
		}
		if p1.(float64) > p2.(float64) {
			return -1
		}
		return 0
	}
	if p1Value.Kind() == reflect.Float64 {
		return compare([]float64{p1.(float64)}, p2)
	}
	if p2Value.Kind() == reflect.Float64 {
		return compare(p1, []float64{p2.(float64)})
	}
	p1Len := p1Value.Len()
	p2Len := p2Value.Len()
	for i := 0; i < p1Len && i < p2Len; i++ {
		v := compare(p1Value.Index(i).Interface(), p2Value.Index(i).Interface())
		if v != 0 {
			return v
		}
	}
	if p1Len < p2Len {
		return 1
	}
	if p1Len > p2Len {
		return -1
	}
	return 0
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	var divider1, divider2 interface{}
	json.Unmarshal([]byte("[[2]]"), &divider1)
	json.Unmarshal([]byte("[[6]]"), &divider2)
	packets := []*interface{}{&divider1, &divider2}

	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			continue
		}
		var p interface{}
		json.Unmarshal([]byte(s), &p)
		packets = append(packets, &p)
	}

	sort.Slice(packets, func(i, j int) bool {
		return compare(*packets[i], *packets[j]) > 0
	})

	key := 1
	for i := 0; i < len(packets); i++ {
		if packets[i] == &divider1 || packets[i] == &divider2 {
			key *= i + 1
		}
	}
	fmt.Println(key)
}
