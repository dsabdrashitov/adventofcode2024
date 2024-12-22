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
	d := make([][]int, len(s))
	for i := range s {
		x := s[i]
		d[i] = make([]int, STEPS+1)
		d[i][0] = x % 10
		for j := range STEPS {
			x = step(x)
			d[i][j+1] = x % 10
		}
	}

	firstSum := make(map[[4]int]int)
	for _, dm := range d {
		first := make(map[[4]int]bool)
		for i := range len(dm) - 4 {
			k := [4]int{dm[i+1] - dm[i], dm[i+2] - dm[i+1], dm[i+3] - dm[i+2], dm[i+4] - dm[i+3]}
			if _, ok := first[k]; !ok {
				first[k] = true
				firstSum[k] = firstSum[k] + dm[i+4]
			}
		}
	}

	answer := 0

	for i0 := -9; i0 <= 9; i0++ {
		for i1 := -9; i1 <= 9; i1++ {
			for i2 := -9; i2 <= 9; i2++ {
				for i3 := -9; i3 <= 9; i3++ {
					k := [4]int{i0, i1, i2, i3}
					sum := firstSum[k]
					if answer < sum {
						fmt.Printf("%v, %v, %v, %v: %v\n", i0, i1, i2, i3, sum)
						answer = sum
					}
				}
			}
		}
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
