package graph

type NodeCost[C any] struct {
	Node int
	Cost C
}

type Graph[C any] interface {
	Edges(node int) []NodeCost[C]
}
