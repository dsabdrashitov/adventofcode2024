package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func safe(a []int) bool {
	for i := 0; i+1 < len(a); i++ {
		if compare(a[i], a[i+1]) != compare(a[0], a[1]) {
			return false
		}
		if c := abs(a[i] - a[i+1]); c < 1 || c > 3 {
			return false
		}
	}
	return true
}

func solve(a [][]int) int {
	result := 0
	for _, ai := range a {
		if safe(ai) {
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
	file := check(os.Open("input.txt"))
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
