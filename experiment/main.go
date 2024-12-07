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
	inp := "zero := [1 -> (8, 1), 2 -> (), 3 -> (9, 10, 11), 4 -> (13)]"

	rpe := Sequence{IntNumber{}, Optional{`, `}}
	rp := Sequence{IntNumber{}, Literal{` -> (`}, Multiple{rpe, `*`}, Literal{`)`}, Optional{`, `}}
	p := Sequence{Regexp{`[a-z]+`}, Literal{` := [`}, Multiple{rp, `*`}, Literal{`]`}}
	pc := p.Complie()

	x := pc.Parse(inp)

	fmt.Println(x.values[0])
	for _, r := range x.values[2].values {
		a := r.values[0].intValue
		b := make([]int, len(r.values[2].values))
		for j, br := range r.values[2].values {
			b[j] = br.values[0].intValue
		}
		fmt.Println(a)
		fmt.Println(b)
	}
}
