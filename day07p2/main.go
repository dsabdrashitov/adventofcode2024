package main

import (
	"fmt"
	"strconv"

	. "github.com/dsabdrashitov/adventofcode2024/pkg/boilerplate"
	. "github.com/dsabdrashitov/adventofcode2024/pkg/fileread"
	P "github.com/dsabdrashitov/adventofcode2024/pkg/regexp"
)

const (
	inputFile = "input.txt"
	// inputFile = "sample.txt"
	// inputFile = "test.txt"
)

func solve(n int, b []int, a [][]int) int {
	result := 0
	for i := range n {
		if can(a[i], b[i]) {
			result += b[i]
		}
	}
	return result
}

func can(a []int, b int) bool {
	p0 := []int{a[0]}
	bs := fmt.Sprintf("%v", b)
	for i := 1; i < len(a); i++ {
		p1 := make([]int, 0)
		was := make(map[int]bool)
		for _, x := range p0 {
			var y int

			y = x + a[i]
			if y <= b && !was[y] {
				p1 = append(p1, y)
				was[y] = true
			}

			y = x * a[i]
			if y <= b && !was[y] {
				p1 = append(p1, y)
				was[y] = true
			}

			ys := fmt.Sprintf("%v%v", x, a[i])
			if len(ys) <= len(bs) {
				y = Must(strconv.Atoi(ys))
				if y <= b && !was[y] {
					p1 = append(p1, y)
					was[y] = true
				}
			}
		}
		p0 = p1
	}
	for _, x := range p0 {
		if x == b {
			return true
		}
	}
	return false
}

func main() {
	inp := ReadLines(inputFile)

	n := len(inp)
	a := make([][]int, n)
	b := make([]int, n)
	for i, line := range inp {
		p := P.ParseSliceInt(line)
		b[i] = p[0]
		a[i] = p[1:]
		fmt.Println(b[i], a[i])
	}
	answer := solve(n, b, a)
	fmt.Println(answer)
}
