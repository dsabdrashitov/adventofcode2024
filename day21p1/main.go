package main

import (
	"fmt"

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

	DOWNS     = "v"
	UPS       = "^"
	LEFTS     = "<"
	RIGHTS    = ">"
	ACTIVATES = "A"
	SPACE     = ' '
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

func dm(depth int, zero int, other int) int {
	if depth == 0 {
		return zero
	} else {
		return other
	}
}

func shortest(s string, depth int, kb []string, u int, v int, lastsu int, lastsv int) (result string, lasteu int, lastev int) {
	if depth == 0 {
		return s, lastsu, lastsv
	}
	if s == "" {
		return "", lastsu, lastsv
	}
	if kb[u][v] == SPACE {
		panic("SPACE")
	}
	if s[0] == kb[u][v] {
		sr, seu, sev := shortest(ACTIVATES, depth-1, DIRK, dm(depth-2, lastsu, DIRKU), dm(depth-2, lastsv, DIRKV), lastsu, lastsv)
		cr, ceu, cev := shortest(s[1:], depth, kb, u, v, dm(depth-1, u, seu), dm(depth-1, v, sev))
		return sr + cr, ceu, cev
	}
	ok := false
	nu, nv := find(s[0], kb)
	if u < nu && kb[u+1][v] != SPACE {
		sr, seu, sev := shortest(DOWNS, depth-1, DIRK, dm(depth-2, lastsu, DIRKU), dm(depth-2, lastsv, DIRKV), lastsu, lastsv)
		cr, ceu, cev := shortest(s, depth, kb, u+1, v, dm(depth-1, u+1, seu), dm(depth-1, v, sev))
		if !ok || len(result) > len(sr)+len(cr) {
			result, lasteu, lastev = sr+cr, ceu, cev
			ok = true
		}
	}
	if u > nu && kb[u-1][v] != SPACE {
		sr, seu, sev := shortest(UPS, depth-1, DIRK, dm(depth-2, lastsu, DIRKU), dm(depth-2, lastsv, DIRKV), lastsu, lastsv)
		cr, ceu, cev := shortest(s, depth, kb, u-1, v, dm(depth-1, u-1, seu), dm(depth-1, v, sev))
		if !ok || len(result) > len(sr)+len(cr) {
			result, lasteu, lastev = sr+cr, ceu, cev
			ok = true
		}
	}
	if v < nv && kb[u][v+1] != SPACE {
		sr, seu, sev := shortest(RIGHTS, depth-1, DIRK, dm(depth-2, lastsu, DIRKU), dm(depth-2, lastsv, DIRKV), lastsu, lastsv)
		cr, ceu, cev := shortest(s, depth, kb, u, v+1, dm(depth-1, u, seu), dm(depth-1, v+1, sev))
		if !ok || len(result) > len(sr)+len(cr) {
			result, lasteu, lastev = sr+cr, ceu, cev
			ok = true
		}
	}
	if v > nv && kb[u][v-1] != SPACE {
		sr, seu, sev := shortest(LEFTS, depth-1, DIRK, dm(depth-2, lastsu, DIRKU), dm(depth-2, lastsv, DIRKV), lastsu, lastsv)
		cr, ceu, cev := shortest(s, depth, kb, u, v-1, dm(depth-1, u, seu), dm(depth-1, v-1, sev))
		if !ok || len(result) > len(sr)+len(cr) {
			result, lasteu, lastev = sr+cr, ceu, cev
			ok = true
		}
	}
	if !ok {
		panic("!ok")
	}
	return result, lasteu, lastev
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
		commands, _, _ := shortest(s, 3, NUMK, NUMKU, NUMKV, DIRKU, DIRKV)
		fmt.Println(commands)
		fmt.Println(num, len(commands), num*len(commands))
		answer += num * len(commands)
	}

	return answer
}

func main() {
	inp := fileread.ReadLines(inputFile)

	answer := solve(inp)

	fmt.Println(answer)
}
