package splaytree

type Aggregator[K any, V any, A any] func(key K, value V, left *node[K, V, A], right *node[K, V, A]) A

func EmptyAggregator[K any, V any, A any](key K, value V, left *node[K, V, A], right *node[K, V, A]) A {
	return *new(A)
}

type TreeSize struct{ Size int }

func SizeAggregator[K any, V any](key K, value V, left *node[K, V, TreeSize], right *node[K, V, TreeSize]) TreeSize {
	result := TreeSize{Size: 1}
	if left != nil {
		result.Size += left.agg.Size
	}
	if right != nil {
		result.Size += right.agg.Size
	}
	return result
}

type TreeSizeDepth struct {
	Size  int
	Depth int
}

func SizeDepthAggregator[K any, V any](key K, value V, left *node[K, V, TreeSizeDepth], right *node[K, V, TreeSizeDepth]) TreeSizeDepth {
	result := TreeSizeDepth{Size: 1, Depth: 0}
	if left != nil {
		result.Size += left.agg.Size
		result.Depth = max(result.Depth, left.agg.Depth+1)
	}
	if right != nil {
		result.Size += right.agg.Size
		result.Depth = max(result.Depth, right.agg.Depth+1)
	}
	return result
}
