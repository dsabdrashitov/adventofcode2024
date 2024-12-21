package graph

import (
	bp "github.com/dsabdrashitov/adventofcode2024/pkg/boilerplate"
	"github.com/dsabdrashitov/adventofcode2024/pkg/splaytree"
)

type Dist[C any, D any] interface {
	Compare(other D) int
	Add(cost C) D
}

type Start[C any, D Dist[C, D]] NodeCost[D]

func (this Start[C, D]) Compare(other Start[C, D]) int {
	cc := this.Cost.Compare(other.Cost)
	if cc != 0 {
		return cc
	}
	return bp.OrderedComparator(this.Node, other.Node)
}

type Dijkstra[C any, D Dist[C, D]] struct {
	g Graph[C]
	d map[int]D
}

func NewDijkstra[C any, D Dist[C, D]](g Graph[C]) *Dijkstra[C, D] {
	dij := &Dijkstra[C, D]{g, make(map[int]D)}
	return dij
}

func (dij *Dijkstra[C, D]) SetDists(starts []Start[C, D]) {
	heap := splaytree.NewWithComparable[Start[C, D], struct{}]()
	for _, start := range starts {
		dij.update(start, heap)
	}
	dij.processHeap(heap)
}

func (dij *Dijkstra[C, D]) Revisit(nodes []int) {
	heap := splaytree.NewWithComparable[Start[C, D], struct{}]()
	for _, n := range nodes {
		if existingDist, ok := dij.d[n]; ok {
			heap.Set(Start[C, D]{n, existingDist}, struct{}{})
		}
	}
	dij.processHeap(heap)
}

func (dij *Dijkstra[C, D]) SetZeroes(nodes []int) *Dijkstra[C, D] {
	starts := make([]Start[C, D], len(nodes))
	for i, n := range nodes {
		starts[i] = Start[C, D]{n, *new(D)}
	}
	dij.SetDists(starts)
	return dij
}

func (dij *Dijkstra[C, D]) Reachable(node int) bool {
	_, ok := dij.d[node]
	return ok
}

func (dij *Dijkstra[C, D]) Dist(node int) D {
	return dij.d[node]
}

func (dij *Dijkstra[C, D]) Distances() map[int]D {
	result := make(map[int]D)
	for k, v := range dij.d {
		result[k] = v
	}
	return result
}

func (dij *Dijkstra[C, D]) processHeap(heap *splaytree.SplayTree[Start[C, D], struct{}, struct{}]) {
	for !heap.Empty() {
		best := heap.Min()
		heap.Delete(best)
		for _, e := range dij.g.Edges(best.Node) {
			dij.update(Start[C, D]{e.Node, best.Cost.Add(e.Cost)}, heap)
		}
	}
}

func (dij *Dijkstra[C, D]) update(nodeDist Start[C, D], heap *splaytree.SplayTree[Start[C, D], struct{}, struct{}]) {
	if existingDist, ok := dij.d[nodeDist.Node]; ok {
		if existingDist.Compare(nodeDist.Cost) < 0 {
			return
		}
		heap.Delete(Start[C, D]{nodeDist.Node, existingDist})
	}
	heap.Set(nodeDist, struct{}{})
	dij.d[nodeDist.Node] = nodeDist.Cost
}
