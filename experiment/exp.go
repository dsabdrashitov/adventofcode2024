package main

import (
	"fmt"

	. "github.com/dsabdrashitov/adventofcode2024/pkg/boilerplate"
	. "github.com/dsabdrashitov/adventofcode2024/pkg/integer"
	P "github.com/dsabdrashitov/adventofcode2024/pkg/regexp"
)

func doExp() {
	inp := "zero := [1 -> (8, 1), -2 -> (), 3 -> (9, 10, 11), -4 -> (13)]"

	rp := P.Sequence(P.Number(), P.Literal(` -> (`), P.List(P.Number(), P.Literal(`, `)), P.Literal(`)`))
	p := P.Sequence(P.Word(), P.Literal(` := [`), P.List(rp, P.Literal(`, `)), P.Token(`]`))
	pc := p.Complie()

	x := pc.Parse(inp)

	fmt.Println(x.L[0].S, x.L[2].S)
	for _, r := range x.L[1].L {
		a := Int(r.L[0].S)
		b := make([]int, len(r.L[1].L))
		for j, br := range r.L[1].L {
			b[j] = Int(br.S)
		}
		fmt.Println(a)
		fmt.Println(b)
	}

	Must("", nil)
}
