package main

import (
	"fmt"
)

const (
	inputFile = "input.txt"
	// inputFile = "sample.txt"
	// inputFile = "test.txt"
)

func main() {
	inp := "zero := [1 -> (8, 1), -2 -> (), 3 -> (9, 10, 11), -4 -> (13)]"

	rp := Sequence{Number{}, Literal{` -> (`}, List{Number{}, `, `}, Literal{`)`}}
	p := Sequence{Word{}, Literal{` := [`}, List{rp, `, `}, Token{`]`}}
	pc := p.Complie()

	x := pc.Parse(inp)

	fmt.Println(x.list[0].str, x.list[2].str)
	for _, r := range x.list[1].list {
		a := Int(r.list[0].str)
		b := make([]int, len(r.list[1].list))
		for j, br := range r.list[1].list {
			b[j] = Int(br.str)
		}
		fmt.Println(a)
		fmt.Println(b)
	}
}
