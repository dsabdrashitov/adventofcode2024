package splaytree

import (
	bp "github.com/dsabdrashitov/adventofcode2024/pkg/boilerplate"
	"golang.org/x/exp/constraints"
)

type SplayTree[K any, V any, A any] struct {
	root      *node[K, V, A]
	compare   bp.Comparator[K]
	aggregate Aggregator[K, V, A]
}

func New[K constraints.Ordered, V any]() *SplayTree[K, V, struct{}] {
	return &SplayTree[K, V, struct{}]{nil, bp.OrderedComparator[K], EmptyAggregator[K, V, struct{}]}
}

func NewWithComparator[K any, V any](compare bp.Comparator[K]) *SplayTree[K, V, struct{}] {
	return &SplayTree[K, V, struct{}]{nil, compare, EmptyAggregator[K, V, struct{}]}
}

func NewWithAggregator[K constraints.Ordered, V any, A any](aggregate Aggregator[K, V, A]) *SplayTree[K, V, A] {
	return &SplayTree[K, V, A]{nil, bp.OrderedComparator[K], aggregate}
}

func NewWithSize[K constraints.Ordered, V any]() *SplayTree[K, V, TreeSize] {
	return &SplayTree[K, V, TreeSize]{nil, bp.OrderedComparator[K], SizeAggregator[K, V]}
}

func NewWithComparatorAndAggregator[K any, V any, A any](compare bp.Comparator[K], aggregate Aggregator[K, V, A]) *SplayTree[K, V, A] {
	return &SplayTree[K, V, A]{nil, compare, aggregate}
}

func (t *SplayTree[K, V, A]) assign(root *node[K, V, A]) *SplayTree[K, V, A] {
	return &SplayTree[K, V, A]{root, t.compare, t.aggregate}
}

func (t *SplayTree[K, V, A]) Clone() *SplayTree[K, V, A] {
	return &SplayTree[K, V, A]{t.root, t.compare, t.aggregate}
}

func (t *SplayTree[K, V, A]) Get(key K) (value V, ok bool) {
	if t.root == nil {
		return *new(V), false
	}

	t.root = t.root.splay(key, t.compare, t.aggregate)

	if t.compare(t.root.key, key) != 0 {
		return *new(V), false
	} else {
		return t.root.value, true
	}
}

func (t *SplayTree[K, V, A]) A() A {
	if t.root == nil {
		return *new(A)
	} else {
		return t.root.agg
	}
}

func (t *SplayTree[K, V, A]) Split(key K) (left *SplayTree[K, V, A], right *SplayTree[K, V, A]) {
	if t.root == nil {
		return t.assign(nil), t.assign(nil)
	}
	t.root = t.root.splay(key, t.compare, t.aggregate)
	c := t.compare(t.root.key, key)
	switch {
	case c > 0:
		return t.assign(t.root.left), t.assign(t.root.reassign(t.aggregate, nil, t.root.right))
	case c < 0:
		return t.assign(t.root.reassign(t.aggregate, t.root.left, nil)), t.assign(t.root.right)
	default:
		return t.assign(t.root.left), t.assign(t.root.right)
	}
}

func (t *SplayTree[K, V, A]) Set(key K, value V) {
	left, right := t.Split(key)
	t.root = makeNode(key, value, t.aggregate, left.root, right.root)
}

func (t *SplayTree[K, V, A]) Delete(key K) {
	if t.root == nil {
		return
	}
	tleft, tright := t.Split(key)
	t.root = merge(t.aggregate, tleft.root, tright.root)
}

func (t *SplayTree[K, V, A]) Min() K {
	if t.root == nil {
		return *new(K)
	}
	t.root = t.root.min(t.aggregate)
	return t.root.key
}

func (t *SplayTree[K, V, A]) Max() K {
	if t.root == nil {
		return *new(K)
	}
	t.root = t.root.max(t.aggregate)
	return t.root.key
}

func (t *SplayTree[K, V, A]) Empty() bool {
	return t.root == nil
}

func pushLeftToStack[K, V, A any](t *node[K, V, A], stack []*node[K, V, A]) []*node[K, V, A] {
	for n := t; n != nil; n = n.left {
		stack = append(stack, n)
	}
	return stack
}

func (t *SplayTree[K, V, A]) Keys() []K {
	result := make([]K, 0)
	stack := make([]*node[K, V, A], 0)
	stack = pushLeftToStack(t.root, stack)
	for len(stack) > 0 {
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, n.key)
		stack = pushLeftToStack(n.right, stack)
	}
	return result
}

func (t *SplayTree[K, V, A]) Values() []V {
	result := make([]V, 0)
	stack := make([]*node[K, V, A], 0)
	pushLeftToStack(t.root, stack)
	for len(stack) > 0 {
		n := stack[len(stack)-1]
		result = append(result, n.value)
		pushLeftToStack(n.right, stack)
	}
	return result
}
