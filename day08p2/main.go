package main

import (
	"fmt"
	// . "github.com/dsabdrashitov/adventofcode2024/pkg/boilerplate"
	"github.com/dsabdrashitov/adventofcode2024/pkg/fileread"
	ip "github.com/dsabdrashitov/adventofcode2024/pkg/intpoint"
)

const (
	inputFile = "input.txt"
	// inputFile = "sample.txt"
	// inputFile = "test.txt"
)

func solve(a []string) int {
	result := 0

	b := make(map[byte][]ip.Point)
	for i := range len(a) {
		for j := range len(a[i]) {
			if c := a[i][j]; c != '.' {
				if _, ok := b[c]; !ok {
					b[c] = make([]ip.Point, 0)
				}
				b[c] = append(b[c], ip.New(i, j))
			}
		}
	}

	u := make(map[ip.Point]bool)
	for _, l := range b {
		for _, p0 := range l {
			for _, p1 := range l {
				if p0 != p1 {
					d01 := p1.Sub(p0)
					for p := p0; p.InsideStrings(a); p = p.Add(d01) {
						u[p] = true
					}
					for p := p0; p.InsideStrings(a); p = p.Sub(d01) {
						u[p] = true
					}
				}
			}
		}
	}

	result = len(u)

	return result
}

func main() {
	inp := fileread.ReadLines(inputFile)

	answer := solve(inp)
	fmt.Println(answer)
}
