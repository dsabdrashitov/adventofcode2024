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
}

type Graph struct {
	field [][]bool
}

func (g *Graph) Edges(node state) []graph.ArbitraryNodeCost[state, int] {
	result := make([]graph.ArbitraryNodeCost[state, int], 0)
	cpos := node.pos.Add(node.dir)
	if !g.field[cpos.X][cpos.Y] {
		result = append(result, graph.ArbitraryNodeCost[state, int]{Node: state{cpos, node.dir}, Cost: DIRECT})
	}
	for _, d := range ip.DIR4 {
		result = append(result, graph.ArbitraryNodeCost[state, int]{Node: state{node.pos, d}, Cost: ROTATE})
	}
	return result
}

type Dist int

func (d Dist) Compare(other Dist) int {
	return bp.OrderedComparator(d, other)
}

func (d Dist) Add(cost int) Dist {
	return Dist(int(d) + cost)
}

func exit(starts []state, field [][]bool) map[state]int {
	g := &Graph{field}
	eg := graph.NewEncodedGraph(g)
	sc := make([]int, len(starts))
	for i, start := range starts {
		sc[i] = eg.Encoder.Id(start)
	}
	dijkstra := graph.NewDijkstra[int, Dist](eg)
	dijkstra.SetZeroes(sc)
	nearest := dijkstra.Distances()
	result := make(map[state]int)
	for k, v := range nearest {
		result[eg.Encoder.Item(k)] = int(v)
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
