package main

import (
	"fmt"

	bp "github.com/dsabdrashitov/adventofcode2024/pkg/boilerplate"
	"github.com/dsabdrashitov/adventofcode2024/pkg/integer"
	re "github.com/dsabdrashitov/adventofcode2024/pkg/regexp"
)

func DoRegexpExperiments() {
	inp := "zero := [1 -> (8, 1), -2 -> (), 3 -> (9, 10, 11), -4 -> (13)]"

	rp := re.Sequence(re.Number(), re.Literal(` -> (`), re.List(re.Number(), re.Literal(`, `)), re.Literal(`)`))
	p := re.Sequence(re.Word(), re.Literal(` := [`), re.List(rp, re.Literal(`, `)), re.Token(`]`))
	pc := p.Complie()

	x := pc.Parse(inp)

	fmt.Println(x.L[0].S, x.L[2].S)
	for _, r := range x.L[1].L {
		a := integer.Int(r.L[0].S)
		b := make([]int, len(r.L[1].L))
		for j, br := range r.L[1].L {
			b[j] = integer.Int(br.S)
		}
		fmt.Println(a)
		fmt.Println(b)
	}

	bp.Must("", nil)
}
