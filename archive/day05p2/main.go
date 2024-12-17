package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

const (
	inputFile = "input.txt"
	// inputFile = "sample.txt"
	// inputFile = "test.txt"
)

func solve(rulesList [][]int, pages [][]int) int {
	result := 0
	maxpage := 0
	for _, rule := range rulesList {
		maxpage = max(maxpage, rule[0], rule[1])
	}
	for _, page := range pages {
		for _, p := range page {
			maxpage = max(maxpage, p)
		}
	}
	rules := make([][]int, maxpage+1)
	for i := range rules {
		rules[i] = make([]int, 0)
	}
	a := make([][]bool, len(rules))
	for i := range a {
		a[i] = make([]bool, len(rules))
	}
	for _, rule := range rulesList {
		rules[rule[1]] = append(rules[rule[1]], rule[0])
		a[rule[0]][rule[1]] = true
	}

	for _, page := range pages {
		if !good(rules, page) {
			sort(a, page)
			result += page[len(page)/2]
		}
	}
	return result
}

func sort(a [][]bool, page []int) {
	slices.SortFunc(page, func(p1 int, p2 int) int {
		switch {
		case a[p1][p2]:
			return -1
		case a[p2][p1]:
			return 1
		case p1 < p2:
			return -1
		case p1 > p2:
			return 1
		default:
			return 0
		}
	})
}

func good(rules [][]int, page []int) bool {
	has := make(map[int]bool)
	for _, p := range page {
		has[p] = true
	}
	hasLeft := make(map[int]bool)
	for i := 0; i < len(page); i++ {
		for _, rule := range rules[page[i]] {
			if has[rule] && !hasLeft[rule] {
				return false
			}
		}
		hasLeft[page[i]] = true
	}
	return true
}

func main() {
	a := ReadLines(inputFile)

	splitIndex := slices.Index(a, "")
	rules := make([][]int, 0)
	ruleRE := regexp.MustCompile(`([^|]+)\|([^|]+)`)
	for _, line := range a[:splitIndex] {
		match := ruleRE.FindStringSubmatch(line)
		u1 := Must(strconv.Atoi(match[1]))
		u2 := Must(strconv.Atoi(match[2]))
		rules = append(rules, []int{u1, u2})
	}
	pages := make([][]int, 0)
	for _, line := range a[splitIndex+1:] {
		s := strings.Split(line, ",")
		p := make([]int, len(s))
		for i := range s {
			p[i] = Must(strconv.Atoi(s[i]))
		}
		pages = append(pages, p)
	}

	answer := solve(rules, pages)
	fmt.Println(answer)
}
