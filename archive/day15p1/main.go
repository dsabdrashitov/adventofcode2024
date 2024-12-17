package main

import (
	"fmt"

	// bp "github.com/dsabdrashitov/adventofcode2024/pkg/boilerplate"
	"github.com/dsabdrashitov/adventofcode2024/pkg/fileread"
	ip "github.com/dsabdrashitov/adventofcode2024/pkg/intpoint"
)

const (
	inputFile = "input.txt"
	// inputFile = "sample.txt"
	// inputFile = "test.txt"

	ROBOT = '@'
	MAZE  = '#'
	GOOD  = 'O'
	FREE  = '.'

	UP    = '^'
	DOWN  = 'v'
	LEFT  = '<'
	RIGHT = '>'
	GPS   = 100
)

func solve(maze []string, commands []string) int {
	m := make([][]byte, len(maze))
	r := ip.New(-1, -1)
	for i := range maze {
		m[i] = make([]byte, len(maze[i]))
		for j := range len(maze[i]) {
			switch maze[i][j] {
			case ROBOT:
				r.X = i
				r.Y = j
				m[i][j] = FREE
			case MAZE:
				m[i][j] = maze[i][j]
			case GOOD:
				m[i][j] = maze[i][j]
			case FREE:
				m[i][j] = maze[i][j]
			default:
				panic(fmt.Errorf("wrong char at %v, %v: %v", i, j, maze[i][j]))
			}
		}
	}

	for _, command := range commands {
		for j := range len(command) {
			switch command[j] {
			case UP:
				moveRobot(&r, ip.UP, m)
			case DOWN:
				moveRobot(&r, ip.DOWN, m)
			case LEFT:
				moveRobot(&r, ip.LEFT, m)
			case RIGHT:
				moveRobot(&r, ip.RIGHT, m)
			default:
				panic(fmt.Errorf("unknown command %v", command[j]))
			}
		}
	}

	printWorld(m)

	result := 0
	for i := range m {
		for j := range m[i] {
			if m[i][j] == GOOD {
				result += i*GPS + j
			}
		}
	}

	return result
}

func moveRobot(r *ip.Point, d ip.Point, m [][]byte) {
	if push(r.Add(d), d, m) {
		*r = r.Add(d)
	}
}

func push(p, d ip.Point, m [][]byte) bool {
	switch m[p.X][p.Y] {
	case FREE:
		return true
	case MAZE:
		return false
	case GOOD:
		if push(p.Add(d), d, m) {
			m[p.X][p.Y] = FREE
			p = p.Add(d)
			m[p.X][p.Y] = GOOD
			return true
		} else {
			return false
		}
	default:
		panic("botva")
	}
}

func printWorld(w [][]byte) {
	for _, s := range w {
		fmt.Println(string(s))
	}
	fmt.Println()
}

func main() {
	inp := fileread.ReadLines(inputFile)

	mazel := 0
	for inp[mazel] != "" {
		mazel++
	}

	answer := solve(inp[:mazel], inp[mazel+1:])
	fmt.Println(answer)
}
