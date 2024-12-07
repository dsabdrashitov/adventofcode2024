package fileread

import (
	"bufio"
	"os"

	. "github.com/dsabdrashitov/adventofcode2024/pkg/boilerplate"
)

func ReadWhole(filename string) string {
	return string(Must(os.ReadFile(filename)))
}

func ReadLines(filename string) []string {
	file := Must(os.Open(filename))
	defer func() { Must(false, file.Close()) }()
	result := make([]string, 0)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result
}
