package graph

import (
	bp "github.com/dsabdrashitov/adventofcode2024/pkg/boilerplate"
	"github.com/dsabdrashitov/adventofcode2024/pkg/splaytree"
)

type Dist[C any, D any] interface {
	Compare(other D) int
	Add(to int, cost C) D
}

type dijkstraDist[C any, D Dist[C, D]] NodeCost[D]

func (this dijkstraDist[C, D]) Compare(other dijkstraDist[C, D]) int {
	cc := this.Cost.Compare(other.Cost)
	if cc != 0 {
		return cc
	}
	return bp.OrderedComparator(this.Node, other.Node)
}

func Dijkstra[C any, D Dist[C, D]](g Graph[C], starts []NodeCost[D]) map[int]D {
	heap := splaytree.NewWithComparator[dijkstraDist[C, D], struct{}](bp.ComparableComparator[dijkstraDist[C, D]])
	nearest := make(map[int]D)
	for _, start := range starts {
		update(dijkstraDist[C, D](start), heap, nearest)
	}
	for !heap.Empty() {
		best := heap.Min()
		heap.Delete(best)
		for _, e := range g.Edges(best.Node) {
			update(dijkstraDist[C, D]{e.Node, best.Cost.Add(e.Node, e.Cost)}, heap, nearest)
		}
	}
	return nearest
}

func update[C any, D Dist[C, D]](
	nodeDist dijkstraDist[C, D],
	heap *splaytree.SplayTree[dijkstraDist[C, D], struct{}, struct{}],
	nearest map[int]D,
) {
	if existingDist, ok := nearest[nodeDist.Node]; ok {
		if existingDist.Compare(nodeDist.Cost) < 0 {
			return
		}
		heap.Delete(dijkstraDist[C, D]{nodeDist.Node, existingDist})
	}
	heap.Set(nodeDist, struct{}{})
	nearest[nodeDist.Node] = nodeDist.Cost
}
