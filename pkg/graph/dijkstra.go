package graph

import (
	bp "github.com/dsabdrashitov/adventofcode2024/pkg/boilerplate"
	"github.com/dsabdrashitov/adventofcode2024/pkg/splaytree"
)

type Cost[N any, E any, C any] interface {
	Compare(other C) int
	Add(edge E) C
}

type NodeCost[K comparable, N Node[K, N, E], E any, C Cost[N, E, C]] struct {
	Node N
	Cost C
}

func (this NodeCost[K, N, E, C]) Compare(other NodeCost[K, N, E, C]) int {
	cc := this.Cost.Compare(other.Cost)
	if cc != 0 {
		return cc
	}
	return this.Node.Compare(other.Node)
}

func Dijkstra[K comparable, N Node[K, N, E], E Edge[N], C Cost[N, E, C]](starts []NodeCost[K, N, E, C]) map[K]C {
	heap := splaytree.NewWithComparator[NodeCost[K, N, E, C], struct{}](bp.ComparableComparator[NodeCost[K, N, E, C]])
	nearest := make(map[K]C)
	for _, start := range starts {
		update(start, heap, nearest)
	}
	for !heap.Empty() {
		best := heap.Min()
		heap.Delete(best)
		for _, e := range best.Node.Edges() {
			update(NodeCost[K, N, E, C]{e.To(), best.Cost.Add(e)}, heap, nearest)
		}
	}
	return nearest
}

func update[K comparable, N Node[K, N, E], E any, C Cost[N, E, C]](
	nc NodeCost[K, N, E, C],
	heap *splaytree.SplayTree[NodeCost[K, N, E, C], struct{}, struct{}],
	nearest map[K]C,
) {
	if ec, ok := nearest[nc.Node.Key()]; ok {
		if ec.Compare(nc.Cost) < 0 {
			return
		}
		heap.Delete(NodeCost[K, N, E, C]{nc.Node, ec})
	}
	heap.Set(nc, struct{}{})
	nearest[nc.Node.Key()] = nc.Cost
}
