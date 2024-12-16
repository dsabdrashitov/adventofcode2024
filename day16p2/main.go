package main

import (
	"fmt"

	// bp "github.com/dsabdrashitov/adventofcode2024/pkg/boilerplate"
	"github.com/dsabdrashitov/adventofcode2024/pkg/fileread"
	ip "github.com/dsabdrashitov/adventofcode2024/pkg/intpoint"
	"github.com/dsabdrashitov/adventofcode2024/pkg/splaytree"
)

const (
	inputFile = "input.txt"
	// inputFile = "sample.txt"
	// inputFile = "test.txt"

	START  = 'S'
	FINISH = 'E'
	WALL   = '#'
	FREE   = '.'

	DIRECT = 1
	ROTATE = 1000
)

var START_DIR = ip.RIGHT

type state struct {
	pos ip.Point
	dir ip.Point
}

type dist struct {
	state state
	val   int
}

func compare(d1 dist, d2 dist) int {
	if d1.val < d2.val {
		return splaytree.LESS
	}
	if d1.val > d2.val {
		return splaytree.GREATER
	}
	if d1.state.pos.X < d2.state.pos.X {
		return splaytree.LESS
	}
	if d1.state.pos.X > d2.state.pos.X {
		return splaytree.GREATER
	}
	if d1.state.pos.Y < d2.state.pos.Y {
		return splaytree.LESS
	}
	if d1.state.pos.Y > d2.state.pos.Y {
		return splaytree.GREATER
	}
	if d1.state.dir.X < d2.state.dir.X {
		return splaytree.LESS
	}
	if d1.state.dir.X > d2.state.dir.X {
		return splaytree.GREATER
	}
	if d1.state.dir.Y < d2.state.dir.Y {
		return splaytree.LESS
	}
	if d1.state.dir.Y > d2.state.dir.Y {
		return splaytree.GREATER
	}
	return splaytree.EQUAL
}

func setDist(d dist, heap *splaytree.SplayTree[dist, struct{}, struct{}], nearest map[state]int) {
	if n, ok := nearest[d.state]; ok {
		heap.Delete(dist{d.state, n})
	}
	heap.Set(d, struct{}{})
	nearest[d.state] = d.val
}

func exit(starts []state, graph [][]bool) map[state]int {
	heap := splaytree.NewWithComparator[dist, struct{}](compare)
	nearest := make(map[state]int)
	for _, start := range starts {
		setDist(dist{start, 0}, heap, nearest)
	}
	for !heap.Empty() {
		best := heap.Min()
		heap.Delete(best)
		for _, c := range children(best, graph) {
			if e, ok := nearest[c.state]; ok && e <= c.val {
				continue
			}
			setDist(c, heap, nearest)
		}
	}
	return nearest
}

func children(from dist, g [][]bool) []dist {
	result := make([]dist, 0)
	step := from.state.pos.Add(from.state.dir)
	if !g[step.X][step.Y] {
		result = append(result, dist{state{step, from.state.dir}, from.val + DIRECT})
	}
	for _, d := range ip.DIR4 {
		result = append(result, dist{state{from.state.pos, d}, from.val + ROTATE})
	}
	return result
}

func solve(inp []string) int {
	var start ip.Point
	var end ip.Point
	maze := make([][]bool, len(inp))

	for i := 0; i < len(inp); i++ {
		maze[i] = make([]bool, len(inp[i]))
		for j := 0; j < len(inp[i]); j++ {
			switch inp[i][j] {
			case START:
				start = ip.New(i, j)
			case FINISH:
				end = ip.New(i, j)
			case WALL:
				maze[i][j] = true
			case FREE:
				maze[i][j] = false
			default:
				panic(fmt.Errorf("unknown char %v", inp[i][j]))
			}
		}
	}

	bestDist := -1

	nearest := exit([]state{{start, START_DIR}}, maze)
	exits := make([]state, 0)
	for _, d := range ip.DIR4 {
		if c, ok := nearest[state{end, d}]; ok {
			if bestDist == -1 || bestDist > c {
				bestDist = c
				exits = make([]state, 0)
			}
			if bestDist == c {
				exits = append(exits, state{end, d.Mult(-1)})
			}
		}
	}
	backnearest := exit(exits, maze)

	answer := make(map[ip.Point]bool)
	for direct, dirval := range nearest {
		if backval, ok := backnearest[state{direct.pos, direct.dir.Mult(-1)}]; ok && dirval+backval == bestDist {
			answer[direct.pos] = true
		}
	}

	printWorld(maze, answer)

	return len(answer)
}

func printWorld(maze [][]bool, answer map[ip.Point]bool) {
	for i := range len(maze) {
		for j := range len(maze[i]) {
			if answer[ip.New(i, j)] {
				fmt.Print("O")
			} else if maze[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	inp := fileread.ReadLines(inputFile)

	answer := solve(inp)
	fmt.Println(answer)
}
