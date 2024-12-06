package main

import (
	"fmt"
)

const (
	inputFile = "input.txt"
	// inputFile = "sample.txt"
	// inputFile = "test.txt"
)

func stuck(n int, m int, field [][]bool, start Point) bool {
	b := make([][][]bool, len(DIR4))
	for i := range b {
		b[i] = make([][]bool, n)
		for j := range b[i] {
			b[i][j] = make([]bool, m)
		}
	}
	p := start
	d := 0
	for {
		if !b[d][p.x][p.y] {
			b[d][p.x][p.y] = true
		} else {
			return true
		}
		np := p.Add(DIR4[d])
		if !PointInside(np, field) {
			return false
		}
		if field[np.x][np.y] {
			d = (d + 1) % len(DIR4)
		} else {
			p = np
		}
	}
}

func solve(n int, m int, field [][]bool, start Point) int {
	result := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if (Point{i, j}) == start {
				continue
			}
			if field[i][j] {
				continue
			}
			field[i][j] = true
			if stuck(n, m, field, start) {
				result += 1
			}
			field[i][j] = false
		}
	}
	return result
}

func main() {
	inp := ReadLines(inputFile)

	n := len(inp)
	m := len(inp[0])
	a := make([][]bool, n)
	var start Point
	for i := range a {
		a[i] = make([]bool, m)
		if len(inp[i]) != m {
			panic("")
		}
		for j := 0; j < m; j++ {
			switch inp[i][j] {
			case '.':
				// nothing
			case '#':
				a[i][j] = true
			case '^':
				start = Point{i, j}
			default:
				panic(string(inp[i][j]))
			}
		}
	}
	answer := solve(n, m, a, start)
	fmt.Println(answer)
}
