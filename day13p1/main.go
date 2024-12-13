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

	MAXPRESS = 100
	ACOST    = 3
	BCOST    = 1
)

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
	result := 0
	for i := range MAXPRESS + 1 {
		am := a.Mult(i)
		rem := p.Sub(am)
		if rem.X < 0 || rem.Y < 0 {
			continue
		}
		if (rem.X%b.X != 0) || (rem.Y%b.Y != 0) {
			continue
		}
		if rem.X/b.X != rem.Y/b.Y {
			continue
		}
		pc := ACOST*i + BCOST*(rem.X/b.X)
		if result == 0 || result > pc {
			result = pc
		}
	}
	return result
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
