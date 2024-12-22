package main

import (
	"fmt"

	// bp "github.com/dsabdrashitov/adventofcode2024/pkg/boilerplate"
	"github.com/dsabdrashitov/adventofcode2024/pkg/fileread"
	"github.com/dsabdrashitov/adventofcode2024/pkg/integer"
	// "github.com/dsabdrashitov/adventofcode2024/pkg/graph"
	// "github.com/dsabdrashitov/adventofcode2024/pkg/identificator"
	// ip "github.com/dsabdrashitov/adventofcode2024/pkg/intpoint"
)

const (
	inputFile = "input.txt"
	// inputFile = "sample.txt"

	STEPS = 2000
	MOD   = 16777216

	M1 = 64
	D2 = 32
	M3 = 2048
)

func step(x int) int {
	y := x * M1
	x = (x ^ y) % MOD

	// yr := x % D2
	y = x / D2
	// if yr >= 16 {
	// y = y + 1
	// }
	x = (x ^ y) % MOD

	y = x * M3
	x = (x ^ y) % MOD
	return x
}

func solve(s []int) int {
	answer := 0

	for i := range s {
		x := s[i]
		for range STEPS {
			x = step(x)
		}
		answer += x
		fmt.Printf("%v: %v\n", s[i], x)
	}

	return answer
}

func main() {
	inp := fileread.ReadLines(inputFile)
	s := make([]int, len(inp))
	for i := range len(inp) {
		s[i] = integer.Int(inp[i])
	}

	answer := solve(s)

	fmt.Println(answer)
}
