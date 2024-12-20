package main

import (
	"fmt"

	bp "github.com/dsabdrashitov/adventofcode2024/pkg/boilerplate"
	"github.com/dsabdrashitov/adventofcode2024/pkg/fileread"
	"github.com/dsabdrashitov/adventofcode2024/pkg/graph"
	"github.com/dsabdrashitov/adventofcode2024/pkg/identificator"
	ip "github.com/dsabdrashitov/adventofcode2024/pkg/intpoint"
)

const (
	inputFile = "input.txt"
	NEEDSAVE  = 100
	// inputFile = "sample.txt"
	// NEEDSAVE  = 41

	WALL   = '#'
	FREE   = '.'
	START  = 'S'
	FINISH = 'E'
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

	for i := 0; i < len(free); i++ {
		for j := 0; j < len(free); j++ {
			if !free[i][j] {
				continue
			}
			for _, d := range ip.DIR4 {
				p0 := ip.New(i, j)
				p1 := p0.Add(d.Mult(2))
				if p1.InsideStrings(maze) && free[p1.X][p1.Y] {
					ds, oks := fromStart[g.Node(p0)]
					df, okf := fromFinish[g.Node(p1)]
					if oks && okf {
						save := baseline - (ds + df + 2)
						if save >= NEEDSAVE {
							answer = answer + 1
						}
					}
				}
			}
			for _, d := range ip.DIRDIAG {
				p0 := ip.New(i, j)
				p1 := p0.Add(d)
				if p1.InsideStrings(maze) && free[p1.X][p1.Y] {
					ds, oks := fromStart[g.Node(p0)]
					df, okf := fromFinish[g.Node(p1)]
					if oks && okf {
						save := baseline - (ds + df + 2)
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
