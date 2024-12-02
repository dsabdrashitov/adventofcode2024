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

	inf = 2147483647
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

func safe(a []int, d int) bool {
	pd := true
	ld := inf
	ps := false
	ls := 0
	pg := false
	lg := 0
	for i := 0; i < len(a); i++ {
		npd := pd && (ld == inf || (compare(a[i], ld) == d && abs(a[i]-ld) >= 1 && abs(a[i]-ld) <= 3))
		nld := a[i]
		nps := pd
		nls := ld
		npg := ((pg && (lg == inf || (compare(a[i], lg) == d && abs(a[i]-lg) >= 1 && abs(a[i]-lg) <= 3))) ||
			(ps && (ls == inf || (compare(a[i], ls) == d && abs(a[i]-ls) >= 1 && abs(a[i]-ls) <= 3))))
		nlg := a[i]
		pd = npd
		ld = nld
		ps = nps
		ls = nls
		pg = npg
		lg = nlg
	}
	return pd || ps || pg
}

func solve(a [][]int) int {
	result := 0
	for _, ai := range a {
		if safe(ai, 1) || safe(ai, -1) {
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
