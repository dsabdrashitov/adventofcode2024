package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const (
	inputFile = "input.txt"
	// inputFile = "sample.txt"
	// inputFile = "test.txt"
)

func solve(input string) int {
	reMul := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	allMatches := reMul.FindAllStringSubmatch(input, -1)
	result := 0
	for _, m := range allMatches {
		x := check(strconv.Atoi(m[1]))
		y := check(strconv.Atoi(m[2]))
		result = result + x*y
		fmt.Println(m)
	}
	return result
}

func main() {
	input := string(check(os.ReadFile(inputFile)))
	answer := solve(input)
	fmt.Println(answer)
}
