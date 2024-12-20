package boilerplate

import "golang.org/x/exp/constraints"

const (
	LESS    = -1
	EQUAL   = 0
	GREATER = 1
)

type Comparator[K any] func(k1 K, k2 K) int

func OrderedComparator[K constraints.Ordered](k1 K, k2 K) int {
	switch {
	case k1 < k2:
		return LESS
	case k1 > k2:
		return GREATER
	default:
		return EQUAL
	}
}

type Comparable[T any] interface {
	Compare(other T) int
}

func ComparableComparator[T Comparable[T]](k1 T, k2 T) int {
	return k1.Compare(k2)
}
