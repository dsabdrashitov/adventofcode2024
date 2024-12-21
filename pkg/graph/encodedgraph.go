package graph

import "github.com/dsabdrashitov/adventofcode2024/pkg/identificator"

type ArbitraryNodeCost[N any, C any] struct {
	Node N
	Cost C
}

type ArbitraryGraph[N any, C any] interface {
	Edges(node N) []ArbitraryNodeCost[N, C]
}

type EncodedGraph[N comparable, C any] struct {
	g       ArbitraryGraph[N, C]
	Encoder *identificator.ComparableEncoder[N]
}

func NewEncodedGraph[N comparable, C any](g ArbitraryGraph[N, C]) *EncodedGraph[N, C] {
	return &EncodedGraph[N, C]{g, identificator.New[N]()}
}

func (eg *EncodedGraph[N, C]) Edges(nodeId int) []NodeCost[C] {
	node := eg.Encoder.Item(nodeId)
	edges := eg.g.Edges(node)
	result := make([]NodeCost[C], len(edges))
	for i, e := range edges {
		child := eg.Encoder.Id(e.Node)
		result[i] = NodeCost[C]{child, e.Cost}
	}
	return result
}
