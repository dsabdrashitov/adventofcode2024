package main

import (
	"fmt"
)

const (
	inputFile = "input.txt"
	// inputFile = "sample.txt"
	// inputFile = "test.txt"
)

func solve(n int, m int, field [][]bool, start Point) int {
	result := 0
	b := make([][]bool, n)
	for i := range b {
		b[i] = make([]bool, m)
	}
	p := start
	d := 0
	for {
		if !b[p.x][p.y] {
			b[p.x][p.y] = true
			result += 1
		}
		np := p.Add(DIR4[d])
		if !PointInside(np, field) {
			break
		}
		if field[np.x][np.y] {
			d = (d + 1) % len(DIR4)
		} else {
			p = np
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
