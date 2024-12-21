package main

import (
	"fmt"

	"github.com/dsabdrashitov/adventofcode2024/pkg/fileread"
	"github.com/dsabdrashitov/adventofcode2024/pkg/graph"

	bp "github.com/dsabdrashitov/adventofcode2024/pkg/boilerplate"
	ip "github.com/dsabdrashitov/adventofcode2024/pkg/intpoint"
	re "github.com/dsabdrashitov/adventofcode2024/pkg/regexp"
)

const (
	inputFile = "input.txt"
	WSIZE     = 70 + 1
	// inputFile = "sample.txt"
	// WSIZE     = 6 + 1
)

type Graph struct {
	field [][]bool
}

func (g *Graph) Edges(p ip.Point) []graph.ArbitraryNodeCost[ip.Point, int] {
	result := make([]graph.ArbitraryNodeCost[ip.Point, int], 0)
	for _, d := range ip.DIR4 {
		np := p.Add(d)
		if np.X >= 0 && np.X < WSIZE && np.Y >= 0 && np.Y < WSIZE && !g.field[np.X][np.Y] {
			result = append(result, graph.ArbitraryNodeCost[ip.Point, int]{Node: np, Cost: 1})
		}
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

func solve(bytes []ip.Point) ip.Point {
	occupied := make([][]bool, WSIZE)
	for i := range WSIZE {
		occupied[i] = make([]bool, WSIZE)
	}
	for _, b := range bytes {
		occupied[b.X][b.Y] = true
	}
	g := &Graph{occupied}
	eg := graph.NewEncodedGraph(g)

	start := ip.New(0, 0)
	finish := ip.New(WSIZE-1, WSIZE-1)

	dijkstra := graph.NewDijkstra[int, Dist](eg)
	dijkstra.SetZeroes([]int{eg.NodeId(start)})
	if dijkstra.Reachable(eg.NodeId(finish)) {
		return ip.Point{X: -1, Y: -1}
	}
	for i := len(bytes) - 1; i >= 0; i-- {
		b := bytes[i]
		occupied[b.X][b.Y] = false
		neighbours := make([]int, 0)
		for _, d := range ip.DIR4 {
			neigh := b.Add(d)
			if neigh.X >= 0 && neigh.X < WSIZE && neigh.Y >= 0 && neigh.Y < WSIZE {
				neighbours = append(neighbours, eg.NodeId(neigh))
			}
		}
		dijkstra.Revisit(neighbours)
		if dijkstra.Reachable(eg.NodeId(finish)) {
			return bytes[i]
		}
	}
	return ip.Point{X: -2, Y: -2}
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

	fmt.Printf("%v,%v\n", answer.Y, answer.X)
}
