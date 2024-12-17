package main

import (
	"fmt"

	// bp "github.com/dsabdrashitov/adventofcode2024/pkg/boilerplate"
	"github.com/dsabdrashitov/adventofcode2024/pkg/fileread"
	"github.com/dsabdrashitov/adventofcode2024/pkg/integer"
	ip "github.com/dsabdrashitov/adventofcode2024/pkg/intpoint"
	re "github.com/dsabdrashitov/adventofcode2024/pkg/regexp"
)

const (
	inputFile = "input.txt"
	// inputFile = "sample.txt"
	// inputFile = "test.txt"

	ACOST = 3
	BCOST = 1
)

var DELTA ip.Point = ip.New(10000000000000, 10000000000000)

func solve(a, b, prize []ip.Point) int {
	result := 0
	n := len(a)
	if n != len(b) || n != len(prize) {
		panic("wrong length")
	}

	for i := range n {
		result = result + cost(a[i], b[i], prize[i])
	}

	return result
}

func cost(a, b, p ip.Point) int {
	p = p.Add(DELTA)
	d := a.X*b.Y - a.Y*b.X
	da := p.X*b.Y - p.Y*b.X
	db := a.X*p.Y - a.Y*p.X
	if d < 0 {
		d = -d
		da = -da
		db = -db
	}
	if d == 0 {
		result := 0

		for i := range b.X {
			r := p.Sub(a.Mult(i))
			if r.X < 0 || r.Y < 0 {
				break
			}
			if r.X%b.X != 0 || r.Y%b.Y != 0 {
				continue
			}
			if r.X/b.X != r.Y/b.Y {
				break
			}
			j := r.X / b.X
			pc := ACOST*i + BCOST*j
			if result == 0 || result > pc {
				result = pc
			}
		}

		for i := range a.X {
			r := p.Sub(b.Mult(i))
			if r.X < 0 || r.Y < 0 {
				break
			}
			if r.X%a.X != 0 || r.Y%a.Y != 0 {
				continue
			}
			if r.X/a.X != r.Y/a.Y {
				break
			}
			j := r.X / a.X
			pc := ACOST*j + BCOST*i
			if result == 0 || result > pc {
				result = pc
			}
		}

		return result
	} else {
		if da < 0 || db < 0 {
			return 0
		}
		if da%d != 0 {
			return 0
		}
		if db%d != 0 {
			return 0
		}
		ia := da / d
		ib := db / d
		return ACOST*ia + BCOST*ib
	}
}

func main() {
	inp := fileread.ReadLines(inputFile)

	buttonRe := re.Sequence(re.Literal("Button "), re.Regexp("[AB]"), re.Literal(": X+"), re.Number(), re.Literal(", Y+"), re.Number()).Complie()
	prizeRe := re.Sequence(re.Literal("Prize: X="), re.Number(), re.Literal(", Y="), re.Number()).Complie()
	a := make([]ip.Point, 0)
	b := make([]ip.Point, 0)
	prize := make([]ip.Point, 0)
	for i := 0; i < len(inp); i += 4 {
		apr := buttonRe.Parse(inp[i+0])
		if apr.L[0].S != "A" {
			panic("wrong letter")
		}
		bpr := buttonRe.Parse(inp[i+1])
		if bpr.L[0].S != "B" {
			panic("wrong letter")
		}
		ppr := prizeRe.Parse(inp[i+2])
		if i+3 < len(inp) {
			if inp[i+3] != "" {
				panic("not empty")
			}
		}
		a = append(a, ip.New(integer.Int(apr.L[1].S), integer.Int(apr.L[2].S)))
		b = append(b, ip.New(integer.Int(bpr.L[1].S), integer.Int(bpr.L[2].S)))
		prize = append(prize, ip.New(integer.Int(ppr.L[0].S), integer.Int(ppr.L[1].S)))
	}
	answer := solve(a, b, prize)
	fmt.Println(answer)
}
