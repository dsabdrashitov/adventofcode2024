package main

import (
	"fmt"
	"sort"
	"strings"

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

func solve(net [][2]string) string {
	g := make(map[string]map[string]bool)
	for _, line := range net {
		add(g, line[0], line[1])
		add(g, line[1], line[0])
	}
	keys := make([]string, 0, len(g))
	possible := make(map[string]bool)
	for k := range g {
		keys = append(keys, k)
		possible[k] = true
	}
	sort.Strings(keys)
	largest := largest(0, keys, g, possible, 0)

	var buf strings.Builder
	for _, s := range largest {
		buf.WriteString(s)
		buf.WriteString(",")
	}
	answer := buf.String()
	if len(answer) > 0 {
		answer = answer[:len(answer)-1]
	}
	return answer
}

func largest(from int, keys []string, g map[string]map[string]bool, possible map[string]bool, current int) []string {
	result := make([]string, 0)
	for i := from; i < len(keys); i++ {
		k := keys[i]
		if !possible[k] {
			continue
		}
		gk := g[k]
		cp := make(map[string]bool)
		for child := range possible {
			if child <= k {
				continue
			}
			if !gk[child] {
				continue
			}
			cp[child] = true
		}
		ca := largest(i+1, keys, g, cp, current-1)
		if current < 1+len(ca) {
			current = 1 + len(ca)
			result = result[:0]
			result = append(result, k)
			result = append(result, ca...)
			fmt.Println(result)
		}
	}
	return result
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
