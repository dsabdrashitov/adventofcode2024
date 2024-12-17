package main

import (
	"fmt"
	// . "github.com/dsabdrashitov/adventofcode2024/pkg/boilerplate"
	"github.com/dsabdrashitov/adventofcode2024/pkg/fileread"
	ip "github.com/dsabdrashitov/adventofcode2024/pkg/intpoint"
)

const (
	inputFile = "input.txt"
	// inputFile = "sample.txt"
	// inputFile = "test.txt"
)

func add(to map[int]bool, from map[int]bool) {
	for k, v := range from {
		to[k] = v
	}
}

func solve(s []string) int {
	result := 0

	a := make([][]map[int]bool, len(s))
	for i, si := range s {
		a[i] = make([]map[int]bool, len(si))
		for j, c := range si {
			a[i][j] = make(map[int]bool)
			if c == '9' {
				a[i][j][i+j*len(s)] = true
			}
		}
	}
	for cc := '8'; cc >= '0'; cc-- {
		for i, si := range s {
			for j, c := range si {
				if c == cc {
					ij := ip.New(i, j)
					for _, d := range ip.DIR4 {
						p := ij.Add(d)
						if p.InsideStrings(s) {
							if s[p.X][p.Y] == byte(cc)+1 {
								add(a[i][j], a[p.X][p.Y])
							}
						}
					}
					if cc == '0' {
						result += len(a[i][j])
						// fmt.Println(a[i][j])
					}
				}
			}
		}
	}

	return result
}

func main() {
	inp := fileread.ReadLines(inputFile)

	answer := solve(inp)
	fmt.Println(answer)
}
