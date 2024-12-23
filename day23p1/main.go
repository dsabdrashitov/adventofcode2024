package main

import (
	"fmt"
	"sort"

	// bp "github.com/dsabdrashitov/adventofcode2024/pkg/boilerplate"
	"github.com/dsabdrashitov/adventofcode2024/pkg/fileread"
	// "github.com/dsabdrashitov/adventofcode2024/pkg/graph"
	// "github.com/dsabdrashitov/adventofcode2024/pkg/identificator"
	re "github.com/dsabdrashitov/adventofcode2024/pkg/regexp"
	// ip "github.com/dsabdrashitov/adventofcode2024/pkg/intpoint"
)

const (
	inputFile = "input.txt"
	// inputFile = "sample.txt"
)

func solve(net [][2]string) int {
	answer := 0
	g := make(map[string]map[string]bool)
	for _, line := range net {
		add(g, line[0], line[1])
		add(g, line[1], line[0])
	}
	sets := make(map[[3]string]bool)
	for k1, gk1 := range g {
		if k1[0] != 't' {
			continue
		}
		for k2 := range gk1 {
			for k3 := range gk1 {
				if k3 <= k2 {
					continue
				}
				if g[k2][k3] {
					s := [3]string{k1, k2, k3}
					sort.Strings(s[:])
					sets[s] = true
				}
			}
		}
	}
	answer = len(sets)
	return answer
}

func add(g map[string]map[string]bool, s1, s2 string) {
	if m, ok := g[s1]; ok {
		m[s2] = true
	} else {
		m = make(map[string]bool)
		m[s2] = true
		g[s1] = m
	}
}

func main() {
	inp := fileread.ReadLines(inputFile)

	linere := re.Sequence(re.Word(), re.Literal(`-`), re.Word()).Complie()
	net := make([][2]string, len(inp))
	for i := range len(inp) {
		m := linere.Parse(inp[i])
		net[i] = [2]string{m.L[0].S, m.L[1].S}
	}

	answer := solve(net)

	fmt.Println(answer)
}
