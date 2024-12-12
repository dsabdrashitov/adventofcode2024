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
	reg := make(map[ip.Point]bool)
	stack = append(stack, start)
	s += 1
	used[start.X][start.Y] = true
	reg[start] = true
	for len(stack) > 0 {
		p0 := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, d := range ip.DIR4 {
			p1 := p0.Add(d)
			if !p1.InsideStrings(ss) {
				continue
			}
			if ss[p1.X][p1.Y] != startc {
				continue
			}
			if used[p1.X][p1.Y] {
				continue
			}
			stack = append(stack, p1)
			s += 1
			used[p1.X][p1.Y] = true
			reg[p1] = true
		}
	}
	p = perimeter(reg)
	// fmt.Println(s, p, s*p)
	return
}

func perimeter(reg map[ip.Point]bool) int {
	used := make(map[ip.Point]bool)
	result := 0
	for p, _ := range reg {
		for _, d1 := range ip.DIR9 {
			p1 := p.Add(d1)
			if used[p1] {
				continue
			} else {
				used[p1] = true
			}
			if reg[p1.Add(ip.New(0, 0))] && reg[p1.Add(ip.New(1, 1))] {
				if !reg[p1.Add(ip.New(1, 0))] {
					result += 1
				}
				if !reg[p1.Add(ip.New(0, 1))] {
					result += 1
				}
			}
			if reg[p1.Add(ip.New(0, 1))] && reg[p1.Add(ip.New(1, 0))] {
				if !reg[p1.Add(ip.New(0, 0))] {
					result += 1
				}
				if !reg[p1.Add(ip.New(1, 1))] {
					result += 1
				}
			}
			cnt := 0
			for i := 0; i <= 1; i++ {
				for j := 0; j <= 1; j++ {
					p2 := p1.Add(ip.New(i, j))
					if reg[p2] {
						cnt += 1
					}
				}
			}
			if cnt == 1 {
				result += 1
			}
		}
	}
	return result
}

func main() {
	inp := fileread.ReadLines(inputFile)
	answer := solve(inp)
	fmt.Println(answer)
}
