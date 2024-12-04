package main

import (
	"fmt"
)

const (
	inputFile = "input.txt"
	// inputFile = "sample.txt"
	// inputFile = "test.txt"

	XMAS = "XMAS"
)

func solve(s []string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			for _, d := range DIR8 {
				good := true
				for k := 0; k < len(XMAS); k++ {
					p := Point{i, j}.Add(d.Mult(k))
					if !PointInsideStrings(p, s) {
						good = false
						break
					}
					if s[p.x][p.y] != XMAS[k] {
						good = false
						break
					}
				}
				if good {
					result += 1
				}
			}
		}
	}
	return result
}

func main() {
	a := ReadLines(inputFile)
	answer := solve(a)
	fmt.Println(answer)
}
