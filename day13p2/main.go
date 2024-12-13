package main

import (
	"fmt"
	"math/big"

	// bp "github.com/dsabdrashitov/adventofcode2024/pkg/boilerplate"
	"github.com/dsabdrashitov/adventofcode2024/pkg/fileread"
	"github.com/dsabdrashitov/adventofcode2024/pkg/integer"
	ip "github.com/dsabdrashitov/adventofcode2024/pkg/intpoint"
	re "github.com/dsabdrashitov/adventofcode2024/pkg/regexp"
)

const (
	// inputFile = "input.txt"
	inputFile = "sample.txt"
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
	result := 0
	var pc int

	pc = costDirected(a, b, p, ACOST, BCOST)
	if result == 0 || result > pc {
		result = pc
	}

	pc = costDirected(b, a, p, BCOST, ACOST)
	if result == 0 || result > pc {
		result = pc
	}

	return result
}

func costDirected(a, b, p ip.Point, acost, bcost int) int {
	ZERO := big.NewInt(0)
	px := big.NewInt(int64(p.X) + int64(DELTA.X))
	py := big.NewInt(int64(p.Y) + int64(DELTA.Y))
	bx := big.NewInt(int64(b.X))
	by := big.NewInt(int64(b.Y))
	result := 0

	fmt.Println(a, b, p)
	for i := range b.X * b.Y {
		i64 := int64(i)
		ax := big.NewInt(int64(a.X))
		ay := big.NewInt(int64(a.Y))
		ax.Mul(ax, big.NewInt(i64))
		ay.Mul(ay, big.NewInt(i64))
		ax.Sub(px, ax)
		ay.Sub(py, ay)
		if ax.Cmp(ZERO) < 0 || ay.Cmp(ZERO) < 0 {
			break
		}
		var rx, ry big.Int
		rx.Mod(ax, bx)
		ry.Mod(ay, by)
		fmt.Println("  ", rx.String(), ry.String())
		if rx.Cmp(ZERO) != 0 || ry.Cmp(ZERO) != 0 {
			continue
		}
		var dx, dy big.Int
		dx.Div(ax, bx)
		dy.Div(ay, by)
		if dx.Cmp(&dy) != 0 {
			continue
		}
		pc := acost*i + bcost*int(dx.Int64())
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
