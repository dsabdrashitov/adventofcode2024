package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	inputFile = "input.txt"
	// inputFile = "sample.txt"
	// inputFile = "test.txt"

	XMAS = "XMAS"
)

var D = [][]int{
	{-1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
	{1, 0},
	{1, -1},
	{0, -1},
	{-1, -1},
}

func add(p1 []int, p2 []int) []int {
	return []int{p1[0] + p2[0], p1[1] + p2[1]}
}

func mult(m int, p []int) []int {
	return []int{m * p[0], m * p[1]}
}

func point(x int, y int) []int {
	return []int{x, y}
}

func isIn(i int, j int, s []string) bool {
	if i < 0 {
		return false
	}
	if i >= len(s) {
		return false
	}
	if j < 0 {
		return false
	}
	if j >= len(s[i]) {
		return false
	}
	return true
}

func solve(s []string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			for _, d := range D {
				good := true
				for k := 0; k < len(XMAS); k++ {
					p := add(point(i, j), mult(k, d))
					if !isIn(p[0], p[1], s) {
						good = false
						break
					}
					if s[p[0]][p[1]] != XMAS[k] {
						good = false
						break
					}
				}
				if good {
					result += 1
				}
			}
		}
	}
	return result
}

func main() {
	file := check(os.Open(inputFile))
	defer func() { check(0, file.Close()) }()
	a := make([]string, 0)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		a = append(a, scanner.Text())
	}
	answer := solve(a)
	fmt.Println(answer)
}
