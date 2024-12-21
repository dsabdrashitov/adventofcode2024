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
	inputFile = "input.txt"
	// inputFile = "sample.txt"

	DOWNS     = "v"
	UPS       = "^"
	LEFTS     = "<"
	RIGHTS    = ">"
	ACTIVATES = "A"
	SPACES    = " "
)

var DIRK = []string{
	" ^A",
	"<v>",
}

var DIRKU, DIRKV = find(ACTIVATES[0], DIRK)

var NUMK = []string{
	"789",
	"456",
	"123",
	" 0A",
}

var NUMKU, NUMKV = find(ACTIVATES[0], NUMK)

func shortest(s string, kb []string, ku int, kv int) string {
	var b strings.Builder
	u, v := ku, kv
	for i := range s {
		nu, nv := find(s[i], kb)
		var cb, pb strings.Builder
		for step := v; step < nv; step++ {
			cb.WriteString(RIGHTS)
			pb.WriteByte(kb[u][step])
		}
		for step := v; step > nv; step-- {
			cb.WriteString(LEFTS)
			pb.WriteByte(kb[u][step])
		}
		for step := u; step < nu; step++ {
			cb.WriteString(DOWNS)
			pb.WriteByte(kb[step][nv])
		}
		for step := u; step > nu; step-- {
			cb.WriteString(UPS)
			pb.WriteByte(kb[step][nv])
		}
		if !strings.Contains(pb.String(), SPACES) {
			b.WriteString(cb.String())
		} else {
			cb.Reset()
			pb.Reset()
			for step := u; step < nu; step++ {
				cb.WriteString(DOWNS)
				pb.WriteByte(kb[step][v])
			}
			for step := u; step > nu; step-- {
				cb.WriteString(UPS)
				pb.WriteByte(kb[step][v])
			}
			for step := v; step < nv; step++ {
				cb.WriteString(RIGHTS)
				pb.WriteByte(kb[nu][step])
			}
			for step := v; step > nv; step-- {
				cb.WriteString(LEFTS)
				pb.WriteByte(kb[nu][step])
			}
			b.WriteString(cb.String())
		}
		u, v = nu, nv
		b.WriteString(ACTIVATES)
	}
	return b.String()
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

	for _, s := range inp {
		num := integer.Int(s[:len(s)-1])
		r1 := shortest(s, NUMK, NUMKU, NUMKV)
		r2 := shortest(r1, DIRK, DIRKU, DIRKV)
		r3 := shortest(r2, DIRK, DIRKU, DIRKV)
		fmt.Println(s)
		fmt.Println(r1)
		fmt.Println(r2)
		fmt.Println(r3)
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
