package main

import (
	"fmt"

	"github.com/dsabdrashitov/adventofcode2024/pkg/fileread"

	re "github.com/dsabdrashitov/adventofcode2024/pkg/regexp"
)

const (
	inputFile = "input.txt"
	// inputFile = "sample.txt"
)

func possible(t string, patterns []string) int {
	p := make([]int, len(t)+1)
	p[0] = 1
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			continue
		}
		for _, pattern := range patterns {
			if i+len(pattern) <= len(t) && t[i:i+len(pattern)] == pattern {
				p[i+len(pattern)] += p[i]
			}
		}
	}
	return p[len(t)]
}

func solve(patterns []string, task []string) int {
	answer := 0

	for _, t := range task {
		answer = answer + possible(t, patterns)
	}

	return answer
}

func main() {
	inp := fileread.ReadLines(inputFile)

	patre := re.List(re.Word(), re.Literal(", ")).Complie()
	patmatch := patre.Parse(inp[0])
	patterns := make([]string, len(patmatch.L))
	for i, m := range patmatch.L {
		patterns[i] = m.S
	}
	task := inp[2:]

	answer := solve(patterns, task)

	fmt.Println(answer)
}
