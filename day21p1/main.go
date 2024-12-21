package main

import (
	"fmt"
	"strings"

	// bp "github.com/dsabdrashitov/adventofcode2024/pkg/boilerplate"
	"github.com/dsabdrashitov/adventofcode2024/pkg/fileread"
	// "github.com/dsabdrashitov/adventofcode2024/pkg/graph"
	// "github.com/dsabdrashitov/adventofcode2024/pkg/identificator"
	"github.com/dsabdrashitov/adventofcode2024/pkg/integer"
	// ip "github.com/dsabdrashitov/adventofcode2024/pkg/intpoint"
)

const (
	// inputFile = "input.txt"
	inputFile = "sample.txt"

	DOWN     = "v"
	UP       = "^"
	LEFT     = "<"
	RIGHT    = ">"
	ACTIVATE = "A"
)

var DIRK = []string{
	" ^A",
	"<v>",
}

var NUMK = []string{
	"789",
	"456",
	"123",
	" 0A",
}

func shortest(s string, kb []string, u *int, v *int) string {
	var buf strings.Builder
	su, sv := find(' ', kb)
	for cur := 0; cur < len(s); {
		nu, nv := find(s[cur], kb)
		for *u != nu || *v != nv {
			// Avoid space
			if *v == sv && nu != su {
				if *v < nv {
					buf.WriteString(RIGHT)
					*v += 1
				} else {
					buf.WriteString(LEFT)
					*v -= 1
				}
				continue
			}

			switch {
			case *u < nu:
				buf.WriteString(DOWN)
				*u += 1
			case *u > nu:
				buf.WriteString(UP)
				*u -= 1
			default:
				if *v < nv {
					buf.WriteString(RIGHT)
					*v += 1
				} else {
					buf.WriteString(LEFT)
					*v -= 1
				}
			}
		}
		buf.WriteString(ACTIVATE)
		cur += 1
	}
	return buf.String()
}

func find(c byte, kb []string) (int, int) {
	for u, line := range kb {
		for v := range line {
			if line[v] == c {
				return u, v
			}
		}
	}
	return -1, -1
}

func solve(inp []string) int {
	answer := 0

	r1u, r1v := find(ACTIVATE[0], NUMK)
	r2u, r2v := find(ACTIVATE[0], DIRK)
	r3u, r3v := find(ACTIVATE[0], DIRK)

	for _, s := range inp {
		num := integer.Int(s[:len(s)-1])
		r1 := shortest(s, NUMK, &r1u, &r1v)
		r2 := shortest(r1, DIRK, &r2u, &r2v)
		r3 := shortest(r2, DIRK, &r3u, &r3v)
		fmt.Println(s)
		fmt.Println(r1, r1u, r1v)
		fmt.Println(r2, r2u, r2v)
		fmt.Println(r3, r3u, r3v)
		fmt.Println(num, len(r3), num*len(r3))
		answer += num * len(r3)
	}

	return answer
}

func main() {
	inp := fileread.ReadLines(inputFile)

	answer := solve(inp)

	fmt.Println(answer)
}
