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
	wall := ip.New(-1, -1)
	m := make([][]*ip.Point, len(maze))
	goods := make([]*ip.Point, 0)
	robot := ip.New(-1, -1)
	for i := range maze {
		m[i] = make([]*ip.Point, 2*len(maze[i]))
		for j := range len(maze[i]) {
			switch maze[i][j] {
			case ROBOT:
				robot = ip.New(i, 2*j)
				m[i][2*j+0] = &robot
			case MAZE:
				m[i][2*j+0] = &wall
				m[i][2*j+1] = &wall
			case GOOD:
				g := ip.New(i, 2*j)
				m[i][2*j+0] = &g
				m[i][2*j+1] = &g
				goods = append(goods, &g)
			case FREE:
				// nothing
			default:
				panic(fmt.Errorf("wrong char at %v, %v: %v", i, j, maze[i][j]))
			}
		}
	}

	for _, command := range commands {
		for j := range len(command) {
			switch command[j] {
			case UP:
				moveRobot(&robot, ip.UP, m, &wall)
			case DOWN:
				moveRobot(&robot, ip.DOWN, m, &wall)
			case LEFT:
				moveRobot(&robot, ip.LEFT, m, &wall)
			case RIGHT:
				moveRobot(&robot, ip.RIGHT, m, &wall)
			default:
				panic(fmt.Errorf("unknown command %v", command[j]))
			}
		}
	}

	printWorld(m, &wall, &robot)

	result := 0
	for _, g := range goods {
		result += g.X*GPS + g.Y
	}

	return result
}

func moveRobot(r *ip.Point, d ip.Point, m [][]*ip.Point, wall *ip.Point) {
	moved := make(map[*ip.Point]bool)
	if canPush(r.Add(d), d, m, wall, moved) {
		push(d, m, moved)
		m[r.X][r.Y] = nil
		*r = r.Add(d)
		m[r.X][r.Y] = r
	}
}

func push(d ip.Point, m [][]*ip.Point, moved map[*ip.Point]bool) {
	for k := range moved {
		m[k.X][k.Y] = nil
		m[k.X][k.Y+1] = nil
	}
	for k := range moved {
		k.X = k.X + d.X
		k.Y = k.Y + d.Y
		m[k.X][k.Y] = k
		m[k.X][k.Y+1] = k
	}
}

func canPush(p ip.Point, d ip.Point, m [][]*ip.Point, wall *ip.Point, done map[*ip.Point]bool) bool {
	mp := m[p.X][p.Y]
	if mp == nil {
		return true
	}
	if mp == wall {
		return false
	}
	if _, ok := done[mp]; ok {
		return true
	}
	done[mp] = false
	done[mp] = calcPush(mp, d, m, wall, done)
	return done[mp]
}

func calcPush(p *ip.Point, d ip.Point, m [][]*ip.Point, wall *ip.Point, done map[*ip.Point]bool) bool {
	n := p.Add(d)
	if !canPush(n, d, m, wall, done) {
		return false
	}
	n.Y = n.Y + 1
	if !canPush(n, d, m, wall, done) {
		return false
	}
	return true
}

func printWorld(w [][]*ip.Point, wall *ip.Point, robot *ip.Point) {
	for _, s := range w {
		for j := range s {
			switch s[j] {
			case nil:
				fmt.Print(".")
			case wall:
				fmt.Print("#")
			case robot:
				fmt.Print("@")
			default:
				fmt.Print("O")
			}
		}
		fmt.Println()
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
