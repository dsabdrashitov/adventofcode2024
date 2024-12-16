package main

import (
	"fmt"

	bp "github.com/dsabdrashitov/adventofcode2024/pkg/boilerplate"
	"github.com/dsabdrashitov/adventofcode2024/pkg/fileread"
	"github.com/dsabdrashitov/adventofcode2024/pkg/graph"
	ip "github.com/dsabdrashitov/adventofcode2024/pkg/intpoint"
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
	g   *[][]bool
}

func (ts state) Compare(os state) int {
	c := ts.pos.Compare(os.pos)
	if c != 0 {
		return c
	}
	c = ts.dir.Compare(os.dir)
	return c
}

func (ts state) Key() state {
	return ts
}

func (ts state) Edges() []edge {
	result := make([]edge, 0)
	npos := ts.pos.Add(ts.dir)
	if !(*ts.g)[npos.X][npos.Y] {
		result = append(result, edge{state{npos, ts.dir, ts.g}, DIRECT})
	}
	for _, d := range ip.DIR4 {
		result = append(result, edge{state{ts.pos, d, ts.g}, ROTATE})
	}
	return result
}

type edge struct {
	to   state
	cost int
}

func (this edge) To() state {
	return this.to
}

type dist int

func (this dist) Compare(other dist) int {
	return bp.OrderedComparator(this, other)
}

func (this dist) Add(e edge) dist {
	return this + dist(e.cost)
}

func exit(starts []state, g [][]bool) map[state]dist {
	sc := make([]graph.NodeCost[state, state, edge, dist], len(starts))
	for i, start := range starts {
		sc[i] = graph.NodeCost[state, state, edge, dist]{start, 0}
	}
	nearest := graph.Dijkstra(sc)
	return nearest
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

	bestDist := dist(-1)

	nearest := exit([]state{{start, START_DIR, &maze}}, maze)
	exits := make([]state, 0)
	for _, d := range ip.DIR4 {
		if c, ok := nearest[state{end, d, &maze}]; ok {
			if bestDist == -1 || bestDist > c {
				bestDist = c
				exits = make([]state, 0)
			}
			if bestDist == c {
				exits = append(exits, state{end, d.Mult(-1), &maze})
			}
		}
	}
	backnearest := exit(exits, maze)

	answer := make(map[ip.Point]bool)
	for direct, dirval := range nearest {
		if backval, ok := backnearest[state{direct.pos, direct.dir.Mult(-1), &maze}]; ok && dirval+backval == bestDist {
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
