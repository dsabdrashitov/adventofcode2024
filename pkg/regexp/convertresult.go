package regexp

import (
	. "github.com/dsabdrashitov/adventofcode2024/pkg/integer"
)

func ParseSliceInt(s string) []int {
	p := List(Number(), Regexp(`[^-+.e\d]+`))
	pr := p.Complie().Parse(s)
	return DecodeSliceInt(pr)
}

func DecodeSliceInt(pr ParsingResult) []int {
	result := make([]int, len(pr.L))
	for i, cr := range pr.L {
		result[i] = Int(cr.S)
	}
	return result
}

func DecodeSliceSliceInt(pr ParsingResult) [][]int {
	result := make([][]int, len(pr.L))
	for i, cr := range pr.L {
		result[i] = make([]int, len(cr.L))
		for j, ccr := range cr.L {
			result[i][j] = Int(ccr.S)
		}
	}
	return result
}
