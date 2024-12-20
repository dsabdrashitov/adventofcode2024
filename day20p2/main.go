package main

import (
	"fmt"

	bp "github.com/dsabdrashitov/adventofcode2024/pkg/boilerplate"
	"github.com/dsabdrashitov/adventofcode2024/pkg/fileread"
	"github.com/dsabdrashitov/adventofcode2024/pkg/graph"
	"github.com/dsabdrashitov/adventofcode2024/pkg/identificator"
	"github.com/dsabdrashitov/adventofcode2024/pkg/integer"
	ip "github.com/dsabdrashitov/adventofcode2024/pkg/intpoint"
)

const (
	inputFile = "input.txt"
	NEEDSAVE  = 100
	// inputFile = "sample.txt"
	// NEEDSAVE  = 75

	WALL   = '#'
	FREE   = '.'
	START  = 'S'
	FINISH = 'E'

	ALLOWED = 20
)

type Graph struct {
	free [][]bool
	enc  *identificator.ComparableEncoder[ip.Point]
}

func (g *Graph) Edges(node int) []graph.NodeCost[int] {
	p := g.enc.Item(node)
	result := make([]graph.NodeCost[int], 0)
	for _, d := range ip.DIR4 {
		np := p.Add(d)
		if np.X >= 0 && np.X < len(g.free) && np.Y >= 0 && np.Y < len(g.free[np.X]) && g.free[np.X][np.Y] {
			result = append(result, graph.NodeCost[int]{Node: g.Node(np), Cost: 1})
		}
	}
	return result
}

func (g *Graph) Node(p ip.Point) int {
	return g.enc.Id(p)
}

type Dist int

func (d Dist) Compare(other Dist) int {
	return bp.OrderedComparator(d, other)
}

func (d Dist) Add(to int, cost int) Dist {
	return Dist(int(d) + cost)
}

func solve(maze []string) int {
	answer := 0

	var start ip.Point
	var finish ip.Point
	free := make([][]bool, len(maze))
	for i, line := range maze {
		free[i] = make([]bool, len(maze[i]))
		for j := range maze[i] {
			switch line[j] {
			case WALL:
				free[i][j] = false
			case FREE:
				free[i][j] = true
			case START:
				free[i][j] = true
				start = ip.New(i, j)
			case FINISH:
				free[i][j] = true
				finish = ip.New(i, j)
			}
		}
	}

	g := &Graph{free: free, enc: identificator.New[ip.Point]()}
	fromStart := graph.Dijkstra[int, Dist](g, []graph.NodeCost[Dist]{{Node: g.Node(start), Cost: 0}})
	fromFinish := graph.Dijkstra[int, Dist](g, []graph.NodeCost[Dist]{{Node: g.Node(finish), Cost: 0}})
	baseline := fromStart[g.Node(finish)]

	fmt.Println(baseline, fromStart[g.Node(ip.New(finish.X-18, finish.Y))])

	for i0 := 0; i0 < len(free); i0++ {
		for j0 := 0; j0 < len(free[i0]); j0++ {
			if !free[i0][j0] {
				continue
			}
			p0 := ip.New(i0, j0)
			for i1 := max(0, i0-ALLOWED); i1 < min(len(free), i0+ALLOWED+1); i1++ {
				for j1 := max(0, j0-ALLOWED); j1 < min(len(free[i1]), j0+ALLOWED+1); j1++ {
					if !free[i1][j1] {
						continue
					}
					p1 := ip.New(i1, j1)
					dcheat := integer.Abs(i0-i1) + integer.Abs(j0-j1)
					if dcheat > ALLOWED {
						continue
					}
					// If cheat mode is active when the end position is reached, cheat mode ends automatically.
					// But the "right" answer doesn't account it.
					// if dcheat > ALLOWED-2 && ((i0 == finish.X && i1 == finish.X && integer.Abs(j0-finish.Y)+integer.Abs(j1-finish.Y) == dcheat && integer.Abs(j0-finish.Y) < dcheat) ||
					// 	(j0 == finish.Y && j1 == finish.Y && integer.Abs(i0-finish.X)+integer.Abs(i1-finish.X) == dcheat && integer.Abs(i0-finish.X) < dcheat)) {
					// 	continue
					// }
					ds, oks := fromStart[g.Node(p0)]
					df, okf := fromFinish[g.Node(p1)]
					if oks && okf {
						save := baseline - (ds + df + Dist(dcheat))
						if save >= NEEDSAVE {
							answer = answer + 1
						}
					}
				}
			}
		}
	}

	return answer
}

func main() {
	inp := fileread.ReadLines(inputFile)

	answer := solve(inp)

	fmt.Println(answer)
}
