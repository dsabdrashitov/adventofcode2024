package main

import (
	"fmt"

	"github.com/dsabdrashitov/adventofcode2024/pkg/fileread"
	re "github.com/dsabdrashitov/adventofcode2024/pkg/regexp"
)

const (
	inputFile = "input.txt"
	// inputFile = "sample.txt"
	// inputFile = "test.txt"
)

func simOne(a int) int {
	if a == 0 {
		return -1
	}
	b := a % 8
	b = b ^ 6
	c := a >> b
	b = b ^ c
	b = b ^ 4
	return b % 8
}

func solve(a int, instructions []int, last int) int {
	fmt.Printf("%v: %v\n", a, instructions[last+1:])
	if last == -1 {
		return a
	}
	result := -1
	for i := 0; i < 8; i++ {
		d := simOne((a << 3) | i)
		if d == instructions[last] {
			s := solve((a<<3)|i, instructions, last-1)
			if s != -1 {
				if result == -1 || result > s {
					result = s
				}
			}
		}
	}
	return result
}

func main() {
	inp := fileread.ReadLines(inputFile)

	instre := re.Sequence(re.Literal(`Program: `), re.List(re.Number(), re.Literal(`,`))).Complie()
	instructions := re.DecodeSliceInt(instre.Parse(inp[4]).L[0])
	answer := solve(0, instructions, len(instructions)-1)

	fmt.Println(answer)
}
