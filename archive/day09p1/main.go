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
	a := make([]int, n)
	n = 0
	for i, c := range line {
		cnt := int(c - '0')
		t := -1
		if i%2 == 0 {
			t = i / 2
		}
		for range cnt {
			a[n] = t
			n++
		}
	}

	for i, j := 0, n-1; i < j; {
		if a[i] != -1 {
			i++
			continue
		}
		if a[j] == -1 {
			j--
			continue
		}
		a[i] = a[j]
		a[j] = -1
		i++
		j--
	}

	for i, v := range a {
		if v == -1 {
			break
		}
		result += i * a[i]
	}

	return result
}

func main() {
	inp := fileread.ReadLines(inputFile)

	answer := solve(inp[0])
	fmt.Println(answer)
}
