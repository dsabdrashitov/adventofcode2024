package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	inputFile = "input.txt"
	// inputFile = "sample.txt"
	// inputFile = "test.txt"
)

var MASES = [][]string{
	{
		"M.S",
		".A.",
		"M.S",
	},
	{
		"M.M",
		".A.",
		"S.S",
	},
	{
		"S.M",
		".A.",
		"S.M",
	},
	{
		"S.S",
		".A.",
		"M.M",
	},
}

func match(s string, p string) bool {
	if len(s) != len(p) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if p[i] != '.' && s[i] != p[i] {
			return false
		}
	}
	return true
}

func solve(s []string) int {
	result := 0
	for _, MAS := range MASES {
		for i := 0; i+len(MAS) <= len(s); i++ {
			for j := 0; j < len(s[i]); j++ {
				good := true
				for k := 0; k < len(MAS); k++ {
					if j+len(MAS[k]) > len(s[i+k]) || !match(s[i+k][j:j+len(MAS[k])], MAS[k]) {
						good = false
						break
					}
				}
				if good {
					result += 1
				}
			}
		}
	}
	return result
}

func main() {
	file := check(os.Open(inputFile))
	defer func() { check(0, file.Close()) }()
	a := make([]string, 0)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		a = append(a, scanner.Text())
	}
	answer := solve(a)
	fmt.Println(answer)
}
