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

func possible(t string, patterns []string) bool {
	p := make([]bool, len(t)+1)
	p[0] = true
	for i := 0; i < len(p); i++ {
		if !p[i] {
			continue
		}
		for _, pattern := range patterns {
			if i+len(pattern) <= len(t) && t[i:i+len(pattern)] == pattern {
				p[i+len(pattern)] = true
			}
		}
	}
	return p[len(t)]
}

func solve(patterns []string, task []string) int {
	answer := 0

	for _, t := range task {
		if possible(t, patterns) {
			answer = answer + 1
		}
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
