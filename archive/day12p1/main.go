package main

import (
	"fmt"

	// bp "github.com/dsabdrashitov/adventofcode2024/pkg/boilerplate"
	"github.com/dsabdrashitov/adventofcode2024/pkg/fileread"
	ip "github.com/dsabdrashitov/adventofcode2024/pkg/intpoint"
)

const (
	inputFile = "input.txt"
	// inputFile = "sample.txt"
	// inputFile = "test.txt"
)

func solve(ss []string) int {
	result := 0

	used := make([][]bool, len(ss))
	for i, line := range ss {
		used[i] = make([]bool, len(line))
	}

	for i, line := range ss {
		for j := range line {
			if !used[i][j] {
				s, p := measure(ip.New(i, j), ss, used)
				result += s * p
			}
		}
	}

	return result
}

func measure(start ip.Point, ss []string, used [][]bool) (s int, p int) {
	startc := ss[start.X][start.Y]
	stack := make([]ip.Point, 0)
	stack = append(stack, start)
	s += 1
	used[start.X][start.Y] = true
	for len(stack) > 0 {
		p0 := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, d := range ip.DIR4 {
			p1 := p0.Add(d)
			if !p1.InsideStrings(ss) {
				p += 1
				continue
			}
			if ss[p1.X][p1.Y] != startc {
				p += 1
				continue
			}
			if used[p1.X][p1.Y] {
				continue
			}
			stack = append(stack, p1)
			s += 1
			used[p1.X][p1.Y] = true
		}
	}
	return
}

func main() {
	inp := fileread.ReadLines(inputFile)
	answer := solve(inp)
	fmt.Println(answer)
}
