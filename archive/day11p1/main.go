package main

import (
	"fmt"
	"math/big"
	"strconv"

	// bp "github.com/dsabdrashitov/adventofcode2024/pkg/boilerplate"
	"github.com/dsabdrashitov/adventofcode2024/pkg/fileread"
	re "github.com/dsabdrashitov/adventofcode2024/pkg/regexp"
)

const (
	inputFile = "input.txt"
	// inputFile = "sample.txt"
	// inputFile = "test.txt"

	CNT = 25
)

var M2024 *big.Int = big.NewInt(2024)

func solve(a []int) int {
	result := 0

	for _, v := range a {
		result += countStones(v, CNT)
	}

	return result
}

func add(m map[string]int, k string, v int) {
	m[k] = m[k] + v
}

func countStones(number int, cnt int) int {
	m0 := make(map[string]int)
	add(m0, strconv.Itoa(number), 1)
	for step := 0; step < cnt; step++ {
		m1 := make(map[string]int)
		for x, c := range m0 {
			ys := split(x)
			for _, y := range ys {
				add(m1, y, c)
			}
		}
		m0 = m1
	}
	result := 0
	for _, c := range m0 {
		result = result + c
	}
	return result
}

func split(x string) []string {
	if x == "0" {
		return []string{"1"}
	}
	if len(x)%2 == 0 {
		return []string{removeLeading(x[:len(x)/2]), removeLeading(x[len(x)/2:])}
	}
	y := new(big.Int)
	if _, ok := y.SetString(x, 10); !ok {
		panic(x)
	}
	y.Mul(y, M2024)
	return []string{y.String()}
}

func removeLeading(x string) string {
	i := 0
	for i+1 < len(x) && x[i] == '0' {
		i++
	}
	return x[i:]
}

func main() {
	inp := fileread.ReadLines(inputFile)
	a := re.ParseSliceInt(inp[0])
	answer := solve(a)
	fmt.Println(answer)
}
