package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
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

	count := 0
	idx := 0
	var p1, p2 interface{}

	for i := 0; scanner.Scan(); i++ {
		s := scanner.Text()
		switch i % 3 {
		case 0:
			json.Unmarshal([]byte(s), &p1)
		case 1:
			json.Unmarshal([]byte(s), &p2)
			idx++
			r := compare(p1, p2)
			if r == 1 {
				count += idx
			}
		default:
			continue
		}
	}
	fmt.Println(count)
}
