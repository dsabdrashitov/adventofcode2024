package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	inputFile = "input.txt"
	// inputFile = "sample.txt"
	// inputFile = "test.txt"
)

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func compare(a int, b int) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}

func safeNormal(a []int, d int) bool {
	for i := 0; i+1 < len(a); i++ {
		diff := (a[i+1] - a[i]) * d
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func safeStupid(a []int, d int) bool {
	if safeNormal(a, d) {
		return true
	}
	reduced := make([]int, len(a)-1)
	for i := 0; i < len(a); i++ {
		copy(reduced[0:i], a[0:i])
		copy(reduced[i:], a[i+1:])
		if safeNormal(reduced, d) {
			return true
		}
	}
	return false
}

func safe(a []int, d int) bool {
	lowest := a[0]
	for _, ai := range a {
		if compare(lowest, ai) == d {
			lowest = ai
		}
	}
	pd := true
	ld := lowest - d
	ps := false
	ls := 0
	for i := 0; i < len(a); i++ {
		npd := pd && compare(a[i], ld) == d && abs(a[i]-ld) >= 1 && abs(a[i]-ld) <= 3
		nld := a[i]
		nps := pd
		nls := ld
		if ps && compare(a[i], ls) == d && abs(a[i]-ls) >= 1 && abs(a[i]-ls) <= 3 {
			if !nps || compare(nls, a[i]) == d {
				nps = true
				nls = a[i]
			}
		}
		pd = npd
		ld = nld
		ps = nps
		ls = nls
	}
	return pd || ps
}

func solve(a [][]int) int {
	result := 0
	for _, ai := range a {
		if safeStupid(ai, 1) || safeStupid(ai, -1) {
			if !(safe(ai, 1) || safe(ai, -1)) {
				fmt.Println(ai)
			}
			result += 1
		}
	}
	return result
}

func check[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

func main() {
	file := check(os.Open(inputFile))
	defer func() { check(0, file.Close()) }()
	a := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		ai := make([]int, 0)
		lineScanner := bufio.NewScanner(strings.NewReader(scanner.Text()))
		lineScanner.Split(bufio.ScanWords)
		for lineScanner.Scan() {
			x := check(strconv.Atoi(lineScanner.Text()))
			ai = append(ai, x)
		}
		a = append(a, ai)
	}
	answer := solve(a)
	fmt.Println(answer)
}
