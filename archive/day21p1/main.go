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

func shortest(s string, kb []string, ku int, kv int, depth int) string {
	if depth == 0 {
		return s
	}
	if s == "" {
		return ""
	}
	var b strings.Builder
	u, v := ku, kv
	for i := range s {
		nu, nv := find(s[i], kb)
		var cb1, pb1 strings.Builder
		for step := v; step < nv; step++ {
			cb1.WriteString(RIGHTS)
			pb1.WriteByte(kb[u][step])
		}
		for step := v; step > nv; step-- {
			cb1.WriteString(LEFTS)
			pb1.WriteByte(kb[u][step])
		}
		for step := u; step < nu; step++ {
			cb1.WriteString(DOWNS)
			pb1.WriteByte(kb[step][nv])
		}
		for step := u; step > nu; step-- {
			cb1.WriteString(UPS)
			pb1.WriteByte(kb[step][nv])
		}
		var cb2, pb2 strings.Builder
		for step := u; step < nu; step++ {
			cb2.WriteString(DOWNS)
			pb2.WriteByte(kb[step][v])
		}
		for step := u; step > nu; step-- {
			cb2.WriteString(UPS)
			pb2.WriteByte(kb[step][v])
		}
		for step := v; step < nv; step++ {
			cb2.WriteString(RIGHTS)
			pb2.WriteByte(kb[nu][step])
		}
		for step := v; step > nv; step-- {
			cb2.WriteString(LEFTS)
			pb2.WriteByte(kb[nu][step])
		}
		if strings.Contains(pb1.String(), SPACES) && strings.Contains(pb2.String(), SPACES) {
			panic("spaces everywhere!!!")
		}
		if strings.Contains(pb1.String(), SPACES) {
			cb1 = cb2
		}
		if strings.Contains(pb2.String(), SPACES) {
			cb2 = cb1
		}
		s1 := shortest(cb1.String()+ACTIVATES, DIRK, DIRKU, DIRKV, depth-1)
		s2 := shortest(cb2.String()+ACTIVATES, DIRK, DIRKU, DIRKV, depth-1)
		if len(s1) <= len(s2) {
			b.WriteString(s1)
		} else {
			b.WriteString(s2)
		}
		u, v = nu, nv
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
		c := shortest(s, NUMK, NUMKU, NUMKV, 3)
		fmt.Println(c)
		fmt.Println(num, len(c), num*len(c))
		answer += num * len(c)
	}

	return answer
}

func main() {
	inp := fileread.ReadLines(inputFile)

	answer := solve(inp)

	fmt.Println(answer)
}
