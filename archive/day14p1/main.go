package main

import (
	"fmt"

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

	steps = 100
)

func solve(robots [][]ip.Point) int {
	result := 1

	for range steps {
		doStep(robots)
	}

	var quad [3][3]int
	for _, robot := range robots {
		var i, j int
		if robot[0].X < width/2 {
			i = 0
		} else if robot[0].X >= width-width/2 {
			i = 1
		} else {
			i = 2
		}
		if robot[0].Y < height/2 {
			j = 0
		} else if robot[0].Y >= height-height/2 {
			j = 1
		} else {
			j = 2
		}
		quad[i][j] += 1
	}
	for _, d := range ip.DIRDR {
		result *= quad[d.X][d.Y]
	}

	return result
}

func doStep(robots [][]ip.Point) {
	for _, robot := range robots {
		robot[0] = robot[0].Add(robot[1])
		robot[0].X = ((robot[0].X % width) + width) % width
		robot[0].Y = ((robot[0].Y % height) + height) % height
	}
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
