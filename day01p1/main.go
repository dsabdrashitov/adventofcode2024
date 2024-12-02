package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"

	"golang.org/x/exp/constraints"
)

func abs[T constraints.Integer](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func solve(a1 []int, a2 []int) int {
	result := 0
	slices.Sort(a1)
	slices.Sort(a2)
	for i, v1 := range a1 {
		v2 := a2[i]
		result = result + abs(v1-v2)
	}
	return result
}

func check[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

func main() {
	file := check(os.Open("input.txt"))
	defer func() { check(0, file.Close()) }()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	a1 := make([]int, 0)
	a2 := make([]int, 0)
	for scanner.Scan() {
		n1 := check(strconv.Atoi(scanner.Text()))
		if !scanner.Scan() {
			panic("PE")
		}
		n2 := check(strconv.Atoi(scanner.Text()))
		a1 = append(a1, n1)
		a2 = append(a2, n2)
	}
	answer := solve(a1, a2)
	fmt.Println(answer)
}
