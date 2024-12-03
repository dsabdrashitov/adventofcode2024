package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	inputFile = "input.txt"
	// inputFile = "sample.txt"
	// inputFile = "test.txt"
)

func solve(input string) int {
	reCommand := regexp.MustCompile(`(?:mul\(([0-9]{1,3}),([0-9]{1,3})\))|(?:do\(\))|(?:don't\(\))`)
	allMatches := reCommand.FindAllStringSubmatch(input, -1)
	result := 0
	enabled := true
	for _, m := range allMatches {
		switch {
		case strings.HasPrefix(m[0], "mul"):
			if enabled {
				x := check(strconv.Atoi(m[1]))
				y := check(strconv.Atoi(m[2]))
				result = result + x*y
			}
		case strings.HasPrefix(m[0], "do("):
			enabled = true
		case strings.HasPrefix(m[0], "don't"):
			enabled = false
		default:
			panic("unknown command")
		}
		fmt.Println(m)
	}
	return result
}

func main() {
	input := string(check(os.ReadFile(inputFile)))
	answer := solve(input)
	fmt.Println(answer)
}
