package main

import (
	"fmt"

	"github.com/dsabdrashitov/adventofcode2024/pkg/fileread"
	"github.com/dsabdrashitov/adventofcode2024/pkg/graph"
	"github.com/dsabdrashitov/adventofcode2024/pkg/identificator"

	bp "github.com/dsabdrashitov/adventofcode2024/pkg/boilerplate"
	ip "github.com/dsabdrashitov/adventofcode2024/pkg/intpoint"
	re "github.com/dsabdrashitov/adventofcode2024/pkg/regexp"
)

const (
	inputFile = "input.txt"
	WSIZE     = 70 + 1
	STEPS     = 1024
	// inputFile = "sample.txt"
	// WSIZE     = 6 + 1
	// STEPS     = 12
)

type Graph struct {
	field [][]bool
	enc   *identificator.ComparableEncoder[ip.Point]
}

func (g *Graph) Edges(node int) []graph.NodeCost[int] {
	p := g.enc.Item(node)
	result := make([]graph.NodeCost[int], 0)
	for _, d := range ip.DIR4 {
		np := p.Add(d)
		if np.X >= 0 && np.X < WSIZE && np.Y >= 0 && np.Y < WSIZE && !g.field[np.X][np.Y] {
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

func solve(bytes []ip.Point) int {
	result := 0

	occupied := make([][]bool, WSIZE)
	for i := range WSIZE {
		occupied[i] = make([]bool, WSIZE)
	}

	for i := 0; i < STEPS; i++ {
		occupied[bytes[i].X][bytes[i].Y] = true
	}

	g := &Graph{occupied, identificator.New[ip.Point]()}

	start := ip.New(0, 0)
	finish := ip.New(WSIZE-1, WSIZE-1)
	dist := graph.Dijkstra[int, Dist](g, []graph.NodeCost[Dist]{{Node: g.Node(start), Cost: 0}})
	result = int(dist[g.Node(finish)])

	return result
}

func main() {
	inp := fileread.ReadLines(inputFile)

	bytes := make([]ip.Point, len(inp))
	linere := re.List(re.Number(), re.Literal(",")).Complie()
	for i, line := range inp {
		c := re.DecodeSliceInt(linere.Parse(line))
		bytes[i] = ip.New(c[1], c[0])
	}

	answer := solve(bytes)

	fmt.Println(answer)
}
