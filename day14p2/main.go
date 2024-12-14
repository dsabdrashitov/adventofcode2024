package main

import (
	"fmt"
	"strings"

	// bp "github.com/dsabdrashitov/adventofcode2024/pkg/boilerplate"
	"github.com/dsabdrashitov/adventofcode2024/pkg/fileread"
	"github.com/dsabdrashitov/adventofcode2024/pkg/integer"
	ip "github.com/dsabdrashitov/adventofcode2024/pkg/intpoint"
	re "github.com/dsabdrashitov/adventofcode2024/pkg/regexp"
)

const (
	inputFile = "input.txt"
	width     = 101
	height    = 103
	// inputFile = "sample.txt"
	// width     = 11
	// height    = 7
	// inputFile = "test.txt"

	steps = 8168
)

func good(robots [][]ip.Point) bool {
	cnt0 := 0
	cnt1 := 0
	for _, robot := range robots {
		if robot[0].Y < height/2 && robot[0].X < width/4-robot[0].Y {
			cnt0++
		}
		if robot[0].Y < height/2 && robot[0].X > 3*width/4+robot[0].Y {
			cnt1++
		}
	}
	return cnt0 < 10 && cnt1 < 10
}

func solve(robots [][]ip.Point) []int {
	result := make([]int, 0)
	for step := range steps {
		doStep(robots)
		if good(robots) {
			// fmt.Println(step)
			// printWorld(paintWorld(robots))
			result = append(result, step)
		}
	}

	printWorld(paintWorld(robots))

	return result
}

func doStep(robots [][]ip.Point) {
	for _, robot := range robots {
		robot[0] = robot[0].Add(robot[1])
		robot[0].X = ((robot[0].X % width) + width) % width
		robot[0].Y = ((robot[0].Y % height) + height) % height
	}
}

func paintWorld(robots [][]ip.Point) []string {
	var a [height][width]int
	for _, robot := range robots {
		a[robot[0].Y][robot[0].X] = 1
	}
	result := make([]string, height)
	for i := range height {
		buf := strings.Builder{}
		for j := range width {
			buf.WriteString(fmt.Sprintf("%v", a[i][j]))
		}
		result[i] = buf.String()
	}
	return result
}

func printWorld(w []string) {
	for _, s := range w {
		fmt.Println(s)
	}
	fmt.Println()
}

func main() {
	inp := fileread.ReadLines(inputFile)

	robotRe := re.Sequence(re.Literal("p="), re.Number(), re.Literal(","), re.Number(), re.Literal(" v="), re.Number(), re.Literal(","), re.Number()).Complie()
	robots := make([][]ip.Point, len(inp))
	for i, line := range inp {
		match := robotRe.Parse(line)
		robots[i] = []ip.Point{
			ip.New(integer.Int(match.L[0].S), integer.Int(match.L[1].S)),
			ip.New(integer.Int(match.L[2].S), integer.Int(match.L[3].S)),
		}
	}
	answer := solve(robots)
	fmt.Println(answer)
}
