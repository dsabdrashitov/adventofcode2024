package main

import (
	"fmt"
	// . "github.com/dsabdrashitov/adventofcode2024/pkg/boilerplate"
	"github.com/dsabdrashitov/adventofcode2024/pkg/fileread"
)

const (
	inputFile = "input.txt"
	// inputFile = "sample.txt"
	// inputFile = "test.txt"
)

func solve(line string) int {
	result := 0

	n := 0
	for _, c := range line {
		n += int(c - '0')
	}
	a := make([][]int, n)
	b := make([][]int, 0)
	n = 0
	for i, c := range line {
		cnt := int(c - '0')
		var p []int
		if i%2 == 0 {
			p = []int{i / 2, n, n + cnt}
			b = append(b, p)
		} else {
			p = []int{-1, n, n + cnt}
		}
		for range cnt {
			a[n] = p
			n++
		}
	}

	for i := len(b) - 1; i >= 0; i-- {
		size := b[i][2] - b[i][1]
		for j := 0; j < b[i][1]; j = a[j][2] {
			if a[j][0] != -1 {
				continue
			}
			if a[j][2]-a[j][1] < size {
				continue
			}
			free(a, b[i][1], b[i][2])
			wasfree := a[j]
			b[i][1] = wasfree[1]
			b[i][2] = wasfree[1] + size
			wasfree[1] = b[i][2]
			for u := b[i][1]; u < b[i][2]; u++ {
				a[u] = b[i]
			}
			break
		}
	}

	for i := 0; i < n; i++ {
		if a[i][0] != -1 {
			result += a[i][0] * i
		}
	}

	return result
}

func free(a [][]int, from int, to int) {
	if from == to {
		return
	}
	if from > 0 && a[from-1][0] == -1 {
		a[from-1][2] = to
		for i := from; i < to; i++ {
			a[i] = a[from-1]
		}
	} else {
		p := []int{-1, from, to}
		for i := from; i < to; i++ {
			a[i] = p
		}
	}
	if to < len(a) && a[to][0] == -1 {
		if a[to][2]-a[to][1] > a[from][2]-a[from][1] {
			a[to][1] = from
			for i := from; i < to; i++ {
				a[i] = a[to]
			}
		} else {
			a[from][2] = a[to][2]
			for i := to; i < a[from][2]; i++ {
				a[i] = a[from]
			}
		}
	}
}

func main() {
	inp := fileread.ReadLines(inputFile)

	answer := solve(inp[0])
	fmt.Println(answer)
}
