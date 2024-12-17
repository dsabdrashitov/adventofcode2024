package main

import (
	"fmt"

	"github.com/dsabdrashitov/adventofcode2024/pkg/fileread"
	"github.com/dsabdrashitov/adventofcode2024/pkg/integer"
	re "github.com/dsabdrashitov/adventofcode2024/pkg/regexp"
)

const (
	// inputFile = "input.txt"
	inputFile = "sample.txt"
	// inputFile = "test.txt"
)

func solve(a, b, c int, instructions []int) []int {
	result := make([]int, 0)

	p := 0
	for ; p+1 < len(instructions); p += 2 {
		com := instructions[p]
		op := instructions[p+1]
		switch com {
		case 0:
			cmb, err := combo(op, a, b, c)
			if err != nil {
				fmt.Println(err)
				return result
			}
			a = a >> cmb
		case 1:
			b = b ^ op
		case 2:
			cmb, err := combo(op, a, b, c)
			if err != nil {
				fmt.Println(err)
				return result
			}
			b = cmb % 8
		case 3:
			if a != 0 {
				p = op - 2
			}
		case 4:
			b = b ^ c
		case 5:
			cmb, err := combo(op, a, b, c)
			if err != nil {
				fmt.Println(err)
				return result
			}
			result = append(result, cmb%8)
		case 6:
			cmb, err := combo(op, a, b, c)
			if err != nil {
				fmt.Println(err)
				return result
			}
			b = a >> cmb
		case 7:
			cmb, err := combo(op, a, b, c)
			if err != nil {
				fmt.Println(err)
				return result
			}
			c = a >> cmb
		default:
			fmt.Printf("Illegal instruction %v at %v\n", com, p)
			return result
		}
	}

	fmt.Printf("Pointer at the end %v\n", p)
	return result
}

func combo(op, a, b, c int) (val int, err error) {
	switch op {
	case 0:
		return op, nil
	case 1:
		return op, nil
	case 2:
		return op, nil
	case 3:
		return op, nil
	case 4:
		return a, nil
	case 5:
		return b, nil
	case 6:
		return c, nil
	case 7:
		return 0, fmt.Errorf("illegal combo operand %v\n", op)
	default:
		return 0, fmt.Errorf("large combo operand %v\n", op)
	}
}

func main() {
	inp := fileread.ReadLines(inputFile)

	regare := re.Sequence(re.Literal(`Register A: `), re.Number()).Complie()
	a := integer.Int(regare.Parse(inp[0]).L[0].S)
	regbre := re.Sequence(re.Literal(`Register B: `), re.Number()).Complie()
	b := integer.Int(regbre.Parse(inp[1]).L[0].S)
	regcre := re.Sequence(re.Literal(`Register C: `), re.Number()).Complie()
	c := integer.Int(regcre.Parse(inp[2]).L[0].S)

	instre := re.Sequence(re.Literal(`Program: `), re.List(re.Number(), re.Literal(`,`))).Complie()
	instructions := re.DecodeSliceInt(instre.Parse(inp[4]).L[0])
	answer := solve(a, b, c, instructions)

	for i, a := range answer {
		if i > 0 {
			fmt.Print(",")
		}
		fmt.Print(a)
	}
	fmt.Println()
}
