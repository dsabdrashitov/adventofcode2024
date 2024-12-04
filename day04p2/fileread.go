package main

import (
	"bufio"
	"os"
)

func ReadWhole(filename string) string {
	return string(Must(os.ReadFile(inputFile)))
}

func ReadLines(filename string) []string {
	file := Must(os.Open(inputFile))
	defer func() { Must(false, file.Close()) }()
	result := make([]string, 0)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result
}
