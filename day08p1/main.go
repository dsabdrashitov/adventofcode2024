package main

import (
	"fmt"
	// . "github.com/dsabdrashitov/adventofcode2024/pkg/boilerplate"
	. "github.com/dsabdrashitov/adventofcode2024/pkg/fileread"
	. "github.com/dsabdrashitov/adventofcode2024/pkg/intpoint"
)

const (
	inputFile = "input.txt"
	// inputFile = "sample.txt"
	// inputFile = "test.txt"
)

func solve(a []string) int {
	result := 0

	b := make(map[byte][]Point)
	for i := range len(a) {
		for j := range len(a[i]) {
			if c := a[i][j]; c != '.' {
				if _, ok := b[c]; !ok {
					b[c] = make([]Point, 0)
				}
				b[c] = append(b[c], Point{i, j})
			}
		}
	}

	u := make(map[Point]bool)
	for _, l := range b {
		for _, p0 := range l {
			for _, p1 := range l {
				if p0 != p1 {
					i0, i1 := intersect(p0, p1)
					if PointInsideStrings(i0, a) {
						u[i0] = true
					}
					if PointInsideStrings(i1, a) {
						u[i1] = true
					}
				}
			}
		}
	}

	result = len(u)

	return result
}

func intersect(p0, p1 Point) (Point, Point) {
	d01 := p1.Sub(p0)
	return p1.Add(d01), p0.Sub(d01)
}

func main() {
	inp := ReadLines(inputFile)

	answer := solve(inp)
	fmt.Println(answer)
}
